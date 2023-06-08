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


func OutPut() (<-chan string, error) {

	// Command to execute
	var cmd *exec.Cmd

        switch runtime.GOOS {
        case "windows":
		return nil, fmt.Errorf("TODO: windows unsupported operating system: %s", runtime.GOOS)
        case "darwin":
                cmd = exec.Command("log", "stream", "--style", "ndjson")
        case "linux":
                cmd = exec.Command("journalctl", "-f", "-o", "json")
        default:
                return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
        }

	// Create a pipe to capture the output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// Create a channel to receive the captured output
	outputCh := make(chan string)

	// Start a goroutine to capture the output
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			outputCh <- line
		}
		close(outputCh)
	}()

	return outputCh, nil
}


func Grep(word string) error {

        var cmd *exec.Cmd

        switch runtime.GOOS {
        case "windows":
		return fmt.Errorf("TODO: windows unsupported operating system: %s", runtime.GOOS)
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


