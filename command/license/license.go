// MIT License
// 
// Copyright (c) 2023 Alyx Holms
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package license

import (
	"log/slog"
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
	} else if files, err := license.GetAllSupportedFiles("."); err != nil {
		return err
	} else {
		slog.Info("Writing license to files", "files", files, "license", formatted)
		for _, file := range files {
			if err := license.WriteLicenseToFile(file, formatted); err != nil {
				return err
			}
		}
		return nil
	}
}
