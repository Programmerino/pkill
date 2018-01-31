// +build windows

package pkill

import (
	"os/exec"
	"strings"
)

// Pkill will kill a process based on it's name, and will output a string output from the windows taskkill command, and also return any errors that occur
func Pkill(name string) (output string, err error) {
	var nameAppend string
	if !strings.HasSuffix(name, ".exe") {
		nameAppend = nameAppend + ".exe"
	}
	bytes, err := exec.Command("taskkill", "/f", "/im", name+nameAppend).Output()
	output = string(bytes)
	return
}
