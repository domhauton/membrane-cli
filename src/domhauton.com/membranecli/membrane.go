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
	verbose := false
	commandStart := -1

	ip := "127.0.0.1"
	port := 13200

	// Sift input for actual start and flags

	for argIdx := 0; argIdx < len(args); argIdx++ {
		if currentArg := args[argIdx]; strings.HasPrefix(currentArg, "-") {
			for flagIdx := 1; flagIdx < len(currentArg); flagIdx++ {
				switch flag := currentArg[flagIdx]; flag {
				case 'v':
					verbose = true
				case 'h':
					help = true
				default:
					fmt.Fprintf(os.Stderr, "Invalid argument. %q", flag)
					os.Exit(1)
				}
			}
		} else if commandStart == -1 {
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
		case commands.DAEMON_STATUS:
			commands.PrintStatus(ip, port, verbose, help)
		case commands.TRACKED_FILES:
			commands.PrintTrackingInfo(ip, port, verbose, help, command)
		case commands.TRACKED_FOLDERS:
			commands.PrintTrackingInfo(ip, port, verbose, help, command)
		case commands.WATCH_LIST:
			commands.PrintWatchedFolders(ip, port, verbose, help)
		case "watch-add":
			fmt.Print("Command unimplemented.")
		case "watch-remove":
			fmt.Print("Command unimplemented.")
		default:
			fmt.Fprintf(os.Stderr, "Invalid command. %s", command)
			os.Exit(1)
		}
	}
}
