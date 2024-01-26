package gtp

import (
    "os/exec"
)

func Run() {
    cmd := exec.Command("notepad")
    cmd.Run()
}