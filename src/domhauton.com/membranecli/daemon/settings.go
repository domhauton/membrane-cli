package daemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type WatchFolder struct {
	Directory string `json:"directory"`
	Recursive bool   `json:"recursive"`
}

type DaemonWatcherSettings struct {
	FileRescan   int           `json:"fileRescan"`
	FolderRescan int           `json:"folderRescan"`
	ChunkSize    int           `json:"chunkSize"`
	WatchFolders []WatchFolder `json:"watchFolders"`
}

type DaemonStorageSettings struct {
	SoftStorageCap   int `json:"softStorageCap"`
	HardStorageCap   int `json:"hardStorageCap"`
	StorageDirectory int `json:"directory"`
	TrimFrequency    int `json:"trimFrequency"`
}

type RestAPIConfig struct {
	Port int `json:"port"`
}

type DaemonSettings struct {
	Watcher            DaemonWatcherSettings `json:"watcher"`
	LocalStorage       DaemonStorageSettings `json:"localStorage"`
	DistributedStorage DaemonStorageSettings `json:"distributedStorage"`
	RestAPI            RestAPIConfig         `json:"restAPI"`
}

func GetDaemonSettings(ip string, port int) (response DaemonSettings, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/config")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = DaemonSettings{}
	json.Unmarshal(body, &response)
	return
}

func GetWatchFoldersAsString(watchFolderInfo []WatchFolder) []string {
	watchFolderStrings := make([]string, len(watchFolderInfo), len(watchFolderInfo))
	for i := 0; i < len(watchFolderInfo); i++ {
		var preamble string
		if watchFolderInfo[i].Recursive {
			preamble = "[Recursive]"
		} else {
			preamble = "[Non-Recur]"
		}
		watchFolderStrings[i] = fmt.Sprintf("%s %s", preamble, watchFolderInfo[i].Directory)
	}
	return watchFolderStrings
}
