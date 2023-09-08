package command

import (
	"github.com/superlinkx/go-runner/command/license"
)

type subCmd int

const (
	InvalidSubCmd subCmd = iota - 1
	LicenseSubCmd
)

// String implements Stringer for the Command enum
func (s subCmd) String() string {
	switch s {
	case LicenseSubCmd:
		return license.Name
	default:
		return "invalid command"
	}
}

// Commands returns our valid set of Command options
func Commands() []subCmd {
	return []subCmd{LicenseSubCmd}
}

// Commands usage returns a slice of Command usage statements indexed by their enum
func CommandsUsage() []string {
	var usage = make([]string, len(Commands()))

	usage[LicenseSubCmd] = license.Usage

	return usage
}
