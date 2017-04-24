package daemon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ContractStatus struct {
	Enabled                 bool     `json:"contractManagerActive"`
	ContractTarget          int      `json:"contractTarget"`
	ContractedPeers         []string `json:"contractedPeers"`
	UndeployedShards        []string `json:"undeployedShards"`
	PartiallyDeployedShards []string `json:"partiallyDistributedShards"`
	FullyDeployedShards     []string `json:"fullyDistributedShards"`
}

func GetContractStatus(ip string, port int) (response ContractStatus, err error) {
	resp, err := http.Get("http://" + ip + ":" + strconv.Itoa(port) + "/status/contract")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = ContractStatus{}
	json.Unmarshal(body, &response)
	return
}
