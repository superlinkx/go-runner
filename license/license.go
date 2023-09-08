package license

import (
	"bufio"
	"fmt"
	"io"
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
	bytes := []byte(license)
	if _, licenseFirstLine, err := bufio.ScanLines(bytes, true); err != nil {
		return fmt.Errorf("error getting first line of license: %s", err)
	} else if file, err := os.OpenFile(path, os.O_RDWR, 0644); err != nil {
		return fmt.Errorf("error opening file %s: %s", path, err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()

		if string(licenseFirstLine) == scanner.Text() {
			// We found the start of the license, so do nothing
			return nil
		} else if err := scanner.Err(); err != nil {
			return fmt.Errorf("scanning for license header failed with: %s", err)
		} else if _, err := file.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek to beginning of file %s: %s", path, err)
		} else if written, err := file.Write([]byte(license)); err != nil {
			return fmt.Errorf("error writing to file %s: %s", path, err)
		} else if written != len(license) {
			return fmt.Errorf("wrote %d of %d bytes to file %s", written, len(license), path)
		} else {
			return nil
		}
	}
}
