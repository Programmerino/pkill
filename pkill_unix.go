// +build !windows

package pkill

import (
	"os/exec"
	"strings"
)

// Pkill will kill a process based on it's name, and will output a string output from the linux pkill command, and also return any errors that occur
func Pkill(name string) (output string, err error) {
	name = strings.TrimSuffix(name, ".exe")
	bytes, err := exec.Command("pkill", name).Output()
	output = string(bytes)
	return
}
