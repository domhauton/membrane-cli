package main

import (
	"domhauton.com/membranecli/commands"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	help := false
	verbose := true
	commandStart := -1

	// Sift input for actual start and flags

	for argIdx := 0; argIdx < len(args); argIdx++ {
		if currentArg := args[0]; strings.HasPrefix(currentArg, "-") {
			for flagIdx := 1; flagIdx < len(currentArg); flagIdx++ {
				switch flag := currentArg[flagIdx]; flag {
				case 'v':
					verbose = true
				case 'h':
					commands.PrintHelp()
					os.Exit(0)
				default:
					fmt.Fprintf(os.Stderr, "Invalid argument. %q", flag)
					os.Exit(1)
				}
			}
		} else if commandStart != -1 {
			commandStart = argIdx
		}
	}

	// Execute input

	if commandStart == -1 {
		if help {
			commands.PrintHelp()
		} else {
			commands.NoArgs()
		}
	} else {
		switch command := args[commandStart]; command {
		case "status":
			commands.PrintStatus(verbose, help)
		case "watch-add":
		case "watch-list":
		case "watch-remove":
			fmt.Print("Command unimplemented.")
		default:
			fmt.Fprintf(os.Stderr, "Invalid command. %s", command)
			os.Exit(1)
		}
	}
}
