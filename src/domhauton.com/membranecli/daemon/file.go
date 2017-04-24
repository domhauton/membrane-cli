package daemon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type FileID struct {
	FilePath       string `json:"filepath"`
	TargetFilePath string `json:"targetFilePath"`
	DateTime       string `json:"dateTime"`
}

type FileHistory struct {
	FilePath        string          `json:"filePath"`
	FileHistEntries []FileHistEntry `json:"fileHistory"`
}

type FileHistEntry struct {
	DateTime string   `json:"dateTime"`
	Shards   []string `json:"hashes"`
	Size     int      `json:"size"`
	RemoveOp bool     `json:"remove"`
}

func GetFileHistory(ip string, port int, fileId FileID) (response FileHistory, err error) {

	data, err := json.Marshal(&fileId)
	if err != nil {
		return
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + "/request/history"
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
	response = FileHistory{}
	json.Unmarshal(body, &response)
	return
}

func History2Strings(history FileHistory) []string {
	historyStrings := make([]string, len(history.FileHistEntries), len(history.FileHistEntries))
	for i, historyEntry := range history.FileHistEntries {
		historyStrings[i] = historyEntry2String(historyEntry)
	}
	return historyStrings
}

func historyEntry2String(entry FileHistEntry) string {
	var printableOp string
	if entry.RemoveOp {
		printableOp = "DEL"
	} else {
		printableOp = "ADD"
	}

	size := entry.Size
	sizeEnding := "B"

	if entry.Size/(1024*1024) != 0 {
		size = entry.Size / (1024 * 1024)
		sizeEnding = "MB"
	} else if entry.Size/1024 != 0 {
		size = entry.Size / 1024
		sizeEnding = "KB"
	}

	return fmt.Sprintf("%s %s %d%s", printableOp, entry.DateTime, size, sizeEnding)
}

func RecoverFile(ip string, port int, fileId FileID) (err error) {

	data, err := json.Marshal(&fileId)
	if err != nil {
		return
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + "/request/reconstruct"
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
