package commands

import (
	"domhauton.com/membranecli/daemon"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	TRACKED_FILES   string = "tracked-files"
	TRACKED_FOLDERS string = "tracked-folders"
	WATCH_LIST      string = "watch-list"
	WATCH_ADD       string = "watch-add"
	WATCH_REMOVE    string = "watch-remove"
	DAEMON_STATUS   string = "status"
)

func PrintStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", DAEMON_STATUS)
	} else {
		status, err := daemon.GetDaemonStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			startingTime, err := daemon.StatusTime(&status)
			if err != nil {
				fmt.Printf("Status:\t\t%s\nHost:\t\t%s:%d\nVersion:\t%s\n",
					status.Status, status.Hostname, status.Port, status.Version)
				if verbose {
					fmt.Printf("Unknown uptime. Given startTime is %s", status.StartTime)
				}
			} else {
				var duration time.Duration = time.Since(startingTime)
				fmt.Printf("Status:\t\t%s\nHost:\t\t%s:%d\nVersion:\t%s\nUptime:\t\t%02d:%02d:%02d\n",
					status.Status, status.Hostname, status.Port, status.Version,
					int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
			}

		}
	}
}

func PrintTrackingInfo(ip string, port int, verbose bool, help bool, trackingType string) {
	if verbose {
		fmt.Printf("Getting %s for %s:%d\n", trackingType, ip, port)
	}

	if help {
		fmt.Printf("Usage: membrane %s\n", trackingType)
	} else {
		watcherStatus, err := daemon.GetWatcherStatus(ip, port)
		if err != nil {
			fmt.Print("Membrane Offline. Status unknown.\n")
			if verbose {
				log.Fatal(err)
			}
		} else {
			var watchList []string
			var name string

			if trackingType == TRACKED_FILES {
				watchList = watcherStatus.TrackedFile
				name = "file(s)"
			} else if trackingType == TRACKED_FOLDERS {
				watchList = watcherStatus.TrackedFolders
				name = "folder(s)"
			}

			if watchList == nil || len(watchList) == 0 {
				watchList = []string{"None"}
			}

			fmt.Printf("Tracked %s:\n\t%s\n", name, strings.Join(watchList, "\n\t"))
		}
	}
}

func PrintWatchedFolders(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Watched Folders at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", WATCH_LIST)
	} else {
		status, err := daemon.GetDaemonSettings(ip, port)
		if err != nil {
			fmt.Println("Membrane Offline. Status unknown.")
			if verbose {
				log.Fatal(err)
			}
		} else {
			watchFolders := daemon.GetWatchFoldersAsString(status.Watcher.WatchFolders)
			if len(watchFolders) == 0 {
				watchFolders = []string{"None"}
			}
			fmt.Printf("Watch Folder(s):\n\t%s\n", strings.Join(watchFolders, "\n\t"))
		}
	}
}

func ModifyWatchedFolders(ip string, port int, verbose bool, help bool, opType string, recursive bool, args []string) {
	if help {
		fmt.Printf("Usage: membrane %s <folder>\nOptions:\n\t-r\trecursive watch folder", opType)
	} else {
		if len(args) < 1 {
			fmt.Fprint(os.Stderr, "Invalid arguments supplied. Check usage.\n")
			return
		}
		directory := args[0]
		var isAdd bool
		var printableOp string

		if opType == WATCH_ADD {
			printableOp = "Adding"
			isAdd = true
		} else if opType == WATCH_REMOVE {
			printableOp = "Removing"
			isAdd = false
		}

		if verbose {
			fmt.Printf("%s Watched Folder [%s] at %s:%d\n", printableOp, directory, ip, port)
		}
		watchFolder := daemon.WatchFolder{
			Recursive: recursive,
			Directory: directory,
		}

		if err := daemon.ConfigureWatchFolder(ip, port, watchFolder, isAdd); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding watch folder. %s\n", err)
			if verbose {
				log.Fatal(err)
			}
		} else {
			// TODO Fix this printout
			fmt.Printf("Successfully added watch folder [%s].\nRun %s to force load.", directory, "n/a")
		}
	}
}
