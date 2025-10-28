package utilities

import (
	"fmt"
	"github.com/mostlygeek/arp"
	"net"
)

// GetMacAddressFromIp retrieves the MAC address associated with the given IP address by parsing the ARP table.
// Returns the MAC address as a string if found, otherwise returns an error if the IP is invalid or no matching MAC address is found.
//
// revive:disable-next-line var-naming // keep exported name for API compatibility
func GetMacAddressFromIp(ipAddress string) (string, error) {
	tvip := net.ParseIP(ipAddress)
	if tvip == nil {
		return "", fmt.Errorf("invalid TV IP address: %s", ipAddress)
	}
	table := arp.Table()
	for ipStr, mac := range table {
		ip := net.ParseIP(ipStr)
		if ip.Equal(tvip) {
			return mac, nil
		}
	}
	return "", fmt.Errorf("mac address not found for TV ip %s", ipAddress)
}
