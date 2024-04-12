package windows

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunPowershellOneLine executes a given `command` and returns the string result,
// plus any errors. This is ideal for commands that are meant for single-line
// results.
func RunPowershellOneLine(command string) (string, error) {
	cmd := exec.Command("powershell.exe", "-Command", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.ReplaceAll(string(output), "\n", "")), nil
}

// RunPowershell executes a given `command` and returns the string result,
// plus any errors.
func RunPowershell(command string) (string, error) {
	cmd := exec.Command("powershell.exe", "-Command", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func GetProperty(path, property string) (string, error) {
	return RunPowershellOneLine(fmt.Sprintf("(Get-ItemProperty '%s').%s", path, property))
}

func GetProperties(path, property string) (string, error) {
	return RunPowershell(fmt.Sprintf("(Get-ItemProperty '%s').%s | Format-List *", path, property))
}
