package commands

import (
	"fmt"
	"strings"
)

const Version = "1.0.0-alpha.1"

func PrintHelp() {
	commands := []string{DAEMON_STATUS, TRACKED_FOLDERS, TRACKED_FILES, "watch-add", "watch-list", "watch-remove", "file-recover"}
	fmt.Printf(`Usage: membrane <command>

	where <command> is one of: [%s]

Options:
	-h	print this message
	-v	enable verbose mode`, strings.Join(commands, ", "))
}

func NoArgs() {
	fmt.Printf("Membrane CLI version %s\nUse flag '-h' for more options", Version)
}
