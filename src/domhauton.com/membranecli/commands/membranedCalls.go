package commands

import (
	"domhauton.com/membranecli/daemon"
	"fmt"
	"log"
	"time"
)

func PrintStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Print("Status Help Placeholder")
	} else {
		status, err := daemon.Status(ip, port)
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
					int(duration.Hours()), int(duration.Minutes())%60)
			}

		}
	}
}
