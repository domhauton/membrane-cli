package daemon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type StatusResponse struct {
	Hostname  string `json:"hostname"`
	StartTime string `json:"startTime"`
	Port      int    `json:"port"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Tagline   string `json:"tagline"`
}

func Status(ip string, port int) (response StatusResponse, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = StatusResponse{}
	json.Unmarshal(body, &response)
	return
}

func StatusTime(status *StatusResponse) (startTime time.Time, err error) {
	startTime, err = time.Parse("2006-01-02T15:04:05.999", status.StartTime)
	return
}
