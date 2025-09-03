package utilities

import (
	"fmt"
	"github.com/sanity-io/litter"

	"log/slog"
)

// init initializes the default configuration for the litter package with custom options and preserves existing field exclusions.
func init() {
	fe := litter.Config.FieldExclusions
	litter.Config = litter.Options{
		Compact:                   false,
		StripPackageNames:         false,
		HidePrivateFields:         false,
		HideZeroValues:            false,
		FieldExclusions:           fe,
		FieldFilter:               nil,
		HomePackage:               "",
		Separator:                 " ",
		StrictGo:                  false,
		DumpFunc:                  nil,
		DisablePointerReplacement: false,
	}
}

// LitterCheckErr logs an error if present, outputs a debug dump of the provided value using the litter library, and then returns the value.
func LitterCheckErr[T any](out T, err error) T {
	if err != nil {
		slog.Error(fmt.Sprintf("*****DEBUG****** error encountered: %v", err))
	}
	litter.Dump(out)
	return out
}
