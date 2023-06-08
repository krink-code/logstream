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
		case "--help":
			printUsage()
		case "--version":
			fmt.Println("Version ", version)
		default:
			fmt.Println("Invalid command. Usage: program grep <pattern>")
		}
	} else {
		err := logstream.Stream()
		if err != nil {
			panic(err)
		}
	}

}

func printUsage() {
	usage := `Usage: program [command] [options]

Commands:
  grep    Filter log output based on a pattern

Options:
  --help  Display this help message
  --version  Display version
  `

	fmt.Println(usage)
}

