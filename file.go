package gtp

import (
    "os/exec"
)

func f() {
    cmd := exec.Command("notepad")
    cmd.Run()
}