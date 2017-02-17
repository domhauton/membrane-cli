package commands

import (
	"domhauton.com/membranecli/daemon"
	"fmt"
	"log"
	"time"
	"strings"
)

const (
	TRACKED_FILES string = "tracked-files"
	TRACKED_FOLDERS string = "tracked-folders"
	DAEMON_STATUS string = "status"
)

func PrintStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Print("Status Help Placeholder")
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
				fmt.Printf("Status:\t\t%s\nHost:\t\t%s:%d\nVersion:\t%s\nUptime:\t\t%02dH%02d\n",
					status.Status, status.Hostname, status.Port, status.Version,
					int(duration.Hours()), int(duration.Minutes()) % 60)
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

			if (trackingType == TRACKED_FILES) {
				watchList = watcherStatus.TrackedFile
				name = "file/s"
			} else if (trackingType == TRACKED_FOLDERS) {
				watchList = watcherStatus.TrackedFolders
				name = "folder/s"
			}

			if watchList == nil || len(watchList) == 0 {
				watchList = []string{"None"}
			}

			fmt.Printf("Tracked %s:\n\t%s\n", name, strings.Join(watchList, "\n\t"))
		}
	}
}
