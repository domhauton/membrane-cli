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

type WatchFolders struct {
	WatchFolders []WatchFolder `json:"watchFolders"`
}

func GetWatchFolders(ip string, port int) (response WatchFolders, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/watch_folder")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = WatchFolders{}
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
