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
	recursive := false
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
				case 'r':
					recursive = true
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
		case commands.STORAGE_STATUS:
			commands.PrintStorageStatus(ip, port, verbose, help)
		case commands.NETWORK_STATUS:
			commands.PrintNetworkStatus(ip, port, verbose, help)
		case commands.CONTRACT_STATUS:
			commands.PrintContractStatus(ip, port, verbose, help)
		case commands.SHOW_PEERS:
			commands.PrintContractedPeers(ip, port, verbose, help)
		case commands.ALL_FILES:
			commands.PrintFiles(ip, port, verbose, help)
		case commands.TRACKED_FILES:
			commands.PrintTrackingInfo(ip, port, verbose, help, command)
		case commands.TRACKED_FOLDERS:
			commands.PrintTrackingInfo(ip, port, verbose, help, command)
		case commands.WATCH_LIST:
			commands.PrintWatchedFolders(ip, port, verbose, help)
		case commands.WATCH_ADD:
			commands.ModifyWatchedFolders(ip, port, verbose, help, command, recursive, args[commandStart+1:])
		case commands.WATCH_REMOVE:
			commands.ModifyWatchedFolders(ip, port, verbose, help, command, recursive, args[commandStart+1:])
		case commands.FILE_HISTORY:
			commands.GetFileHistory(ip, port, verbose, help, args[commandStart+1:])
		case commands.FILE_RECOVER:
			commands.RecoverFile(ip, port, help, args[commandStart+1:])
		default:
			fmt.Fprintf(os.Stderr, "Invalid command. %s", command)
			os.Exit(1)
		}
	}
}
