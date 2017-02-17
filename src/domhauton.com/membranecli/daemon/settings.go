package daemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DaemonWatchFolderInfo struct {
	Directory string `json:"directory"`
	Recursive bool   `json:"recursive"`
}

type DaemonWatcherSettings struct {
	FileRescan   int                     `json:"fileRescan"`
	FolderRescan int                     `json:"folderRescan"`
	ChunkSize    int                     `json:"chunkSize"`
	WatchFolders []DaemonWatchFolderInfo `json:"watchFolders"`
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
	Watcher DaemonWatcherSettings `json:"watcherConfig"`
	Storage DaemonStorageSettings `json:"storageConfig"`
	RestAPI RestAPIConfig         `json:"restAPIConfig"`
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

func GetWatchFoldersAsString(watchFolderInfo []DaemonWatchFolderInfo) []string {
	watchFolderStrings := make([]string, len(watchFolderInfo), len(watchFolderInfo))
	for i := 0; i < len(watchFolderInfo); i++ {
		start := "\t\t"
		if watchFolderInfo[i].Recursive {
			start = "recursive"
		}
		watchFolderStrings[i] = fmt.Sprintf("%s %s", start, watchFolderInfo[i].Directory)
	}
	return watchFolderStrings
}
