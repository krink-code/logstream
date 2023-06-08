package main

import (
	"fmt"
	"os"

	"gitlab.com/krink/logstream/golang/logstream"
)

var version = "1.0.0"

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "grep":
			if len(os.Args) > 2 {
				pattern := os.Args[2]
				err := logstream.Grep(pattern)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println("Pattern missing. Usage: program grep <pattern>")
			}
		case "--output":
			runOutPut()
		case "--help":
			printUsage()
		case "--version":
			fmt.Println("Version ", version)
		default:
			fmt.Println("Invalid argument ", os.Args[1])
		}
	} else {
		err := logstream.Stream()
		if err != nil {
			panic(err)
		}
	}

}

func runOutPut() {

	output, err := logstream.OutPut()
	if err != nil {
		panic(err)
	}

	// Process the captured output
	for line := range output {
		fmt.Println("Received output:", line)

		// Add any desired logic or break condition here
	}

	fmt.Println("Finished capturing and printing output.")
}

func printUsage() {
	usage := `Usage: logstream [command|options]

Commands:
  grep    Filter log output based on a pattern

Options:
  --help  Display this help message
  --version  Display version
  --output Iterator function
  `

	fmt.Println(usage)
}

