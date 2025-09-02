package utilities

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// DaemonAlreadyRunning checks if a daemon process with the given appName is already running on the system.
// Returns true if the process exists, otherwise false.
func DaemonAlreadyRunning(appName string) bool {
	_, err := FindDaemonProcessPID(appName)
	return err == nil
}

// MustFindDaemonProcessPID retrieves the PID of a running daemon process for the given appName or returns 0 if not found.
// It panics if an error occurs during the search, enforcing that the process must be found.
func MustFindDaemonProcessPID(appName string) int {
	pid, err := FindDaemonProcessPID(appName)
	if err != nil {
		return 0
	}
	return pid
}

// FindDaemonProcessPID searches for the PID of a running daemon process matching the given appName.
// It checks the local process table on Linux and utilizes platform-specific logic for macOS.
// Returns the PID of the found process or an error if the process is not found or on failure.
func FindDaemonProcessPID(appName string) (int, error) {
	return FindDaemonProcessPIDWithArg(appName, "run")
}

// FindDaemonProcessPIDWithArg searches for the PID of a running daemon process matching the given appName and given argument.
// It checks the local process table on Linux and utilizes platform-specific logic for macOS.
// Returns the PID of the found process or an error if the process is not found or on failure.
func FindDaemonProcessPIDWithArg(appName string, argName string) (int, error) {
	if runtime.GOOS == "darwin" {
		return FindProcessPIDMAC(appName)
	}
	ownPID := os.Getpid()
	const procDir = "/proc"
	var processPID int

	err := filepath.Walk(procDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Ignore permission errors.
			if os.IsPermission(err) || errors.Is(err, os.ErrNotExist) {
				return nil
			}
			return err
		}

		// Skip non-directories, the /proc directory itself, and dont recurse.
		if !info.IsDir() || path == procDir || len(strings.Split(path, "/")) > 3 {
			return nil
		}

		// Skip non-numeric directories inside /proc.
		if _, err := strconv.Atoi(info.Name()); err != nil {
			return nil
		}

		// Read the cmdline file to get the process name.
		cmdlinePath := filepath.Join(path, "cmdline")
		cmdlineBytes, err := os.ReadFile(cmdlinePath)
		if err != nil {
			// Ignore permission errors and the case where the cmdline file might have been removed after opening.
			if os.IsPermission(err) || errors.Is(err, os.ErrNotExist) {
				return nil
			}
			return err
		}

		cmdlineFields := bytes.SplitN(cmdlineBytes, []byte{0}, 2)
		if len(cmdlineFields) == 0 {
			return nil
		}
		cmdline := string(cmdlineFields[0])
		if filepath.Base(cmdline) != appName {
			return nil
		}
		if argName != "" {
			if len(cmdlineFields) != 2 {
				return nil
			}
			arg := string(cmdlineFields[1])
			if !strings.Contains(arg, argName) {
				return nil
			}
		}
		pid, err := strconv.Atoi(info.Name())
		if err != nil {
			return err
		}
		if pid == ownPID {
			return nil
		}
		processPID = pid
		return syscall.EEXIST // Returning a specific error to stop walking
	})

	if err == syscall.EEXIST {
		return processPID, nil
	}
	if err != nil {
		return 0, err
	}
	return 0, errors.New("process not found")
}

// FindProcessPIDMAC searches for a specific process by name on macOS and returns its PID.
// It filters processes using the `ps` command, ignoring the caller's own PID, and matches the process name and argument.
// Returns the PID of the found process or an error if the process is not found.
func FindProcessPIDMAC(appName string) (int, error) {
	argName := "run"
	ownPID := os.Getpid()
	// List all processes using the `ps` command
	out, err := exec.Command("ps", "-A", "-o", "pid,args").Output()
	if err != nil {
		return 0, err
	}

	// Parse the output
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		// Check if the process name matches
		if strings.Contains(fields[1], appName) && strings.Contains(line, argName) {
			pid, err := strconv.Atoi(fields[0])
			if err != nil {
				return 0, err
			}
			if pid == ownPID {
				continue
			}
			return pid, nil
		}
	}

	return 0, errors.New("process not found")
}
