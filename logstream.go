package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

func logstream() error {

    switch runtime.GOOS {
    case "windows":
        fmt.Println("TODO: Windows")
        os.Exit(1)
    case "darwin":
        cmd := exec.Command("log", "stream", "--style", "ndjson")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd.Run()
    case "linux":
        cmd := exec.Command("journalctl", "-f", "-o", "json")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd.Run()
    }

    return nil
}

func main() {

    err := logstream()

    if err != nil {
        panic(err)
    }

}

