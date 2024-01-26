package go-test-package

import (
    "os/exec"
)

func run() {
    cmd := exec.Command("notepad")
    cmd.Run()
}