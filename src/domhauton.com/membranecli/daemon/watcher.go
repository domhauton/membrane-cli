package daemon

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	ADD_CHANGE    = "ADD"
	REMOVE_CHANGE = "REMOVE"
)

type WatcherStatus struct {
	TrackedFolders []string `json:"trackedFolders"`
	TrackedFile    []string `json:"trackedFiles"`
}

type watcherModifier struct {
	ChangeType  string      `json:"type"`
	WatchFolder WatchFolder `json:"watchFolder"`
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

func ConfigureWatchFolder(ip string, port int, watchFolder WatchFolder, isAddition bool) (err error) {
	request := watcherModifier{}
	request.WatchFolder = watchFolder
	if isAddition {
		request.ChangeType = ADD_CHANGE
	} else {
		request.ChangeType = REMOVE_CHANGE
	}
	data, err := json.Marshal(&request)
	if err != nil {
		return
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + "/configure/watch_folder"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(string(body[:]))
	}
	return
}
