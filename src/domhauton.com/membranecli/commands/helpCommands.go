package commands

import (
	"fmt"
	"strings"
)

const Version = "1.0.0-alpha.4"

func PrintHelp() {
	available_commands := []string{DAEMON_STATUS, ALL_FILES, SHOW_PEERS, FILE_HISTORY, STORAGE_STATUS, NETWORK_STATUS, CONTRACT_STATUS, TRACKED_FOLDERS, TRACKED_FILES, WATCH_ADD, WATCH_LIST, WATCH_REMOVE, "file-recover"}
	fmt.Printf(`Usage: membrane <command>

	where <command> is one of: [%s]

Options:
	-h	print this message
	-v	enable verbose mode
`, strings.Join(available_commands, ", "))
}

func NoArgs() {
	fmt.Printf("Membrane CLI version %s\nUse flag '-h' for more options", Version)
}
