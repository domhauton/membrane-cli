package daemon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type NetworkStatus struct {
	Enabled            bool   `json:"enabled"`
	ConnectedPeers     int    `json:"connectedPeers"`
	NetworkUID         string `json:"networkUID"`
	MaxConnectionCount int    `json:"maxConnectionCount"`
	PeerListeningPort  int    `json:"peerListeningPort"`
	UpnpAddress        string `json:"upnpAddress"`
}

func GetNetworkStatus(ip string, port int) (response NetworkStatus, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/network")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = NetworkStatus{}
	json.Unmarshal(body, &response)
	return
}
