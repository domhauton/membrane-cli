package daemon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type WatcherStatus struct {
	TrackedFolders []string `json:"trackedFolders"`
	TrackedFile    []string `json:"trackedFiles"`
}

func GetWatcherStatus(ip string, port int) (response WatcherStatus, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/watcher")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = WatcherStatus{}
	json.Unmarshal(body, &response)
	return
}
