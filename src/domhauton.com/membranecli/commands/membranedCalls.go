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
	TRACKED_FILES   string = "watched-files"
	TRACKED_FOLDERS string = "watched-folders"
	WATCH_LIST      string = "watch-list"
	WATCH_ADD       string = "watch-add"
	WATCH_REMOVE    string = "watch-remove"
	DAEMON_STATUS   string = "status"
	STORAGE_STATUS  string = "status-storage"
	NETWORK_STATUS  string = "status-network"
	CONTRACT_STATUS string = "status-contract"
	SHOW_PEERS      string = "peers"
	ALL_FILES       string = "files"
	FILE_HISTORY    string = "history"
	FILE_RECOVER    string = "recover"
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

func PrintStorageStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Storage Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", STORAGE_STATUS)
	} else {
		status, err := daemon.GetStorageStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			mb := 1024 * 1024
			fmt.Printf("%d current files stored.\n%d total files.\nLocal Shard Storage. %dMB of %dMB used (max %dMB)\nPeer Block Storage. %dMB of %dMB used (max %dMB)\n",
				len(status.CurrentFiles), len(status.ReferencedFiles),
				status.LocalShardStorageSize/mb, status.TargetLocalShardStorageSize/mb, status.MaxLocalShardStorageSize/mb,
				status.PeerBlockStorageSize/mb, status.TargetPeerBlockStorageSize/mb, status.MaxPeerBlockStorageSize/mb)
		}
	}
}

func PrintFiles(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Files at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", ALL_FILES)
	} else {
		status, err := daemon.GetStorageStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			if len(status.ReferencedFiles) == 0 {
				fmt.Print("No files have been backed up yet.")
			} else {
				fmt.Printf("File(s):\n\t%s\n", strings.Join(daemon.GetFiles(status), "\n\t"))
			}
		}
	}
}

func PrintNetworkStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Network Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", NETWORK_STATUS)
	} else {
		status, err := daemon.GetNetworkStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			if status.Enabled {
				fmt.Printf("User ID:\t%s\nPeers:\t\t%d (max. %d)\nPeer Port:\t%d\nUPnP Address:\t%s\n",
					status.NetworkUID,
					status.ConnectedPeers, status.MaxConnectionCount,
					status.PeerListeningPort,
					status.UpnpAddress)
			} else {
				fmt.Print("Networking Disabled. Enable Contracts for Networking.")
			}

		}
	}
}

func PrintContractStatus(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting Membrane Contract Status at %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", CONTRACT_STATUS)
	} else {
		status, err := daemon.GetContractStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			if status.Enabled {
				fmt.Printf("Peers Contracted:\t%d (max. %d)\nUndeployed Shards:\t%d\nPartially Deployed:\t%d\nFully Deployed:\t\t%d\n",
					len(status.ContractedPeers), status.ContractTarget,
					len(status.UndeployedShards),
					len(status.PartiallyDeployedShards),
					len(status.FullyDeployedShards))
			} else {
				fmt.Print("Peer Contracts Disabled.")
			}

		}
	}
}

func PrintContractedPeers(ip string, port int, verbose bool, help bool) {
	if verbose {
		fmt.Printf("Getting contracted peers from %s:%d\n", ip, port)
	}
	if help {
		fmt.Printf("Usage: membrane %s\n", SHOW_PEERS)
	} else {
		status, err := daemon.GetContractStatus(ip, port)
		if err != nil {
			fmt.Printf("Status:\tOFFLINE\nHost:\t%s:%d\n", ip, port)
			if verbose {
				log.Fatal(err)
			}
		} else {
			if status.Enabled {
				if len(status.ContractedPeers) == 0 {
					fmt.Print("No peers have been contracted.")
				} else {
					fmt.Printf("Contracted Peer(s):\n\t%s\n", strings.Join(status.ContractedPeers, "\n\t"))
				}
			} else {
				fmt.Print("Peer Contracts Disabled.")
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
		status, err := daemon.GetWatchFolders(ip, port)
		if err != nil {
			fmt.Println("Membrane Offline. Status unknown.")
			if verbose {
				log.Fatal(err)
			}
		} else {
			watchFolders := daemon.GetWatchFoldersAsString(status.WatchFolders)
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

		if len(directory) == 0 || directory[0] != '/' {
			fmt.Printf("You must specify the full directory.\nExample: /tmp/dir1\nYou provided: %s\n", directory)
			return
		}

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
			fmt.Fprintf(os.Stderr, "Error %s watch folder. %s\n", printableOp, err)
			if verbose {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("Successfully %s watch folder [%s].\n", printableOp, directory)
		}
	}
}

func GetFileHistory(ip string, port int, verbose bool, help bool, args []string) {
	if help {
		fmt.Printf("Usage: membrane %s <file>\n", FILE_HISTORY)
	} else {
		if len(args) < 1 {
			fmt.Fprint(os.Stderr, "Invalid arguments supplied. Check usage.\n")
			return
		}
		filePath := args[0]

		if len(filePath) == 0 || filePath[0] != '/' {
			fmt.Printf("You must specify the full file.\nExample: /tmp/dir1/file.txt\nYou provided: %s\n", filePath)
			return
		}

		if verbose {
			fmt.Printf("Getting file history for [%s] at %s:%d\n", filePath, ip, port)
		}
		fileId := daemon.FileID{
			FilePath:       filePath,
			DateTime:       "",
			TargetFilePath: "/tmp/mbrn/" + filePath,
		}

		history, err := daemon.GetFileHistory(ip, port, fileId)
		if err != nil {
			fmt.Printf("Failed to retrieve file history. Error: %s\n", err)
		} else {
			historyStrings := daemon.History2Strings(history)
			fmt.Printf("File History for [%s]:\n\t%s\n", history.FilePath, strings.Join(historyStrings, "\n\t"))
		}
	}
}

func RecoverFile(ip string, port int, help bool, args []string) {
	if help {
		fmt.Printf("Usage: membrane %s <file> <destination> <OPTIONAL Time: 2017-04-24T03:10:13.000>\n", FILE_RECOVER)
	} else {
		if len(args) < 2 {
			fmt.Fprint(os.Stderr, "Invalid arguments supplied. Check usage.\n")
			return
		}
		filePath := args[0]
		targetFilePath := args[1]

		if len(filePath) == 0 || filePath[0] != '/' {
			fmt.Printf("You must specify the full filePath.\nExample: /tmp/dir1/file.txt\nYou provided: %s\n", filePath)
			return
		}

		if len(targetFilePath) == 0 || targetFilePath[0] != '/' {
			fmt.Printf("You must specify the full filePath.\nExample: /tmp/dir1/file.txt\nYou provided: %s\n", targetFilePath)
			return
		}

		var timeString string

		if len(args) > 2 {
			timeString = args[2]
			_, err := time.Parse("2006-01-02T15:04:05.999", timeString)
			if err != nil {
				fmt.Printf("Given date must follow format YYYY-MM-DDTHH:MM:ss.SSS.\nExample: 2006-01-02T15:04:05.999\nYou provided: %s\n", timeString)
				return
			}
		} else {
			timeString = ""
		}

		fileId := daemon.FileID{
			FilePath:       targetFilePath,
			DateTime:       timeString,
			TargetFilePath: targetFilePath,
		}

		if err := daemon.RecoverFile(ip, port, fileId); err != nil {
			fmt.Fprintf(os.Stderr, "Error recovering file. %s\n", err)
		} else {
			fmt.Printf("Successfully recovered file. [%s]\n", filePath)
		}
	}
}
