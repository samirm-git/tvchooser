//go:build linux

package mounted

import (
	"os/exec"
	"strings"
)

// Get a list of mounted devices on Linux.
func GetWindowsDriveLetters() ([]string, error) {
	out, err := exec.Command("df", "--output=target").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")[1:] // skip header
	var mounts []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			mounts = append(mounts, line)
		}
	}
	return mounts, nil
}
