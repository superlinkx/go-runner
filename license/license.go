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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadLicenseFile(path string) (string, error) {
	var (
		eof     bool
		idx     int
		license strings.Builder
	)

	if file, err := os.ReadFile(path); err != nil {
		return "", fmt.Errorf("error reading license file: %s", err)
	} else {
		for !eof {
			if advance, line, err := bufio.ScanLines(file[idx:], true); err != nil {
				return "", fmt.Errorf("error scanning license file: %s", err)
			} else {
				fmt.Fprintf(&license, "// %s\n", string(line))
				idx += advance
				if idx >= len(file) {
					fmt.Fprintf(&license, "\n")
					eof = true
				}
			}
		}

		return license.String(), nil
	}
}

func GetAllSupportedFiles(startDir string) ([]string, error) {
	var files []string
	if err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		} else {
			files = append(files, path)
			return nil
		}
	}); err != nil {
		return files, fmt.Errorf("failed to get all supported files: %s", err)
	} else {
		return files, nil
	}
}

func WriteLicenseToFile(path string, license string) error {
	licenseBytes := []byte(license)
	if _, licenseFirstLine, err := bufio.ScanLines(licenseBytes, true); err != nil {
		return fmt.Errorf("error getting first line of license: %s", err)
	} else if file, err := os.ReadFile(path); err != nil {
		return fmt.Errorf("error opening file %s: %s", path, err)
	} else if _, line, err := bufio.ScanLines(file, true); err != nil {
		return fmt.Errorf("error scanning first line of file %s: %s", path, err)
	} else if string(licenseFirstLine) == string(line) {
		// We found the start of the license, so do nothing
		return nil
	} else if info, err := os.Stat(path); err != nil {
		return fmt.Errorf("error getting file info for %s: %s", path, err)
	} else if err := os.WriteFile(path, append(licenseBytes, file...), info.Mode().Perm()); err != nil {
		return fmt.Errorf("error writing file %s: %s", path, err)
	} else {
		return nil
	}
}
