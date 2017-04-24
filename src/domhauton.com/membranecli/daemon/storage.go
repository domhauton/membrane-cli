package daemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type StorageStatus struct {
	CurrentFiles                []string `json:"currentFiles"`
	ReferencedFiles             []string `json:"referencedFiles"`
	LocalShardStorageSize       int      `json:"localShardStorageSize"`
	TargetLocalShardStorageSize int      `json:"targetLocalShardStorageSize"`
	MaxLocalShardStorageSize    int      `json:"maxLocalShardStorageSize"`
	PeerBlockStorageSize        int      `json:"peerBlockStorageSize"`
	TargetPeerBlockStorageSize  int      `json:"targetPeerBlockStorageSize"`
	MaxPeerBlockStorageSize     int      `json:"maxPeerBlockStorageSize"`
}

func GetStorageStatus(ip string, port int) (response StorageStatus, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/storage")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = StorageStatus{}
	json.Unmarshal(body, &response)
	return
}

func GetFiles(status StorageStatus) []string {
	allFiles := status.ReferencedFiles

	currentFilesMap := map[string]bool{}
	for _, file := range status.CurrentFiles {
		currentFilesMap[file] = true
	}

	fileStrings := make([]string, len(allFiles), len(allFiles))
	for i, file := range status.ReferencedFiles {
		var preamble string
		if currentFilesMap[file] {
			preamble = "* "
		} else {
			preamble = "  "
		}
		fileStrings[i] = fmt.Sprintf("%s %s", preamble, file)
	}
	return fileStrings
}
