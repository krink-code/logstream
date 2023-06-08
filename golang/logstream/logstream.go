package logstream

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "bufio"
    "regexp"
)


func Stream() error {

    switch runtime.GOOS {
    case "windows":
        fmt.Println("TODO: Windows")
        return nil
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


func Grep(word string) error {

        var cmd *exec.Cmd

        switch runtime.GOOS {
        case "windows":
                fmt.Println("TODO: Windows")
                return nil
        case "darwin":
                cmd = exec.Command("log", "stream", "--style", "ndjson")
        case "linux":
                cmd = exec.Command("journalctl", "-f", "-o", "json")
        default:
                return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
        }

        stdoutPipe, err := cmd.StdoutPipe()
        if err != nil {
                return err
        }

        err = cmd.Start()
        if err != nil {
                return err
        }

        go func() {
                defer cmd.Wait()
                scanner := bufio.NewScanner(stdoutPipe)
                for scanner.Scan() {
                        line := scanner.Text()
                        if grep(word, line) {
                                fmt.Println(line)
                        }
                }
        }()

        err = cmd.Wait()
        if err != nil {
                return err
        }

        return nil
}

func grep(pattern string, text string) bool {
        regex := regexp.MustCompile(pattern)
        return regex.MatchString(text)
}


