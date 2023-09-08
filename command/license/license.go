package license

import (
	"fmt"
	"os"

	"github.com/superlinkx/go-runner/environment"
	"github.com/superlinkx/go-runner/license"
)

const (
	Name  = "license"
	Usage = "Ensure licenses are present in current repository"
)

type cmd struct {
	LicenseFile string
}

func Create() cmd {
	var (
		filePath = os.Getenv(environment.LicenseFile.Env())
	)

	if filePath == "" {
		filePath = "LICENSE"
	}

	return cmd{
		LicenseFile: filePath,
	}
}

func (c cmd) Name() string {
	return Name
}

func (c cmd) Usage() string {
	return Usage
}

func (c cmd) Run() error {
	if formatted, err := license.ReadLicenseFile(c.LicenseFile); err != nil {
		return err
	} else {
		fmt.Print(formatted)
		return nil
	}
}
