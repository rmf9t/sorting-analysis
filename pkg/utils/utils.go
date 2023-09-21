package utils

import (
	"os"
	"os/exec"
)

// ClearScreen clears the terminal screen for the linux terminal
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
