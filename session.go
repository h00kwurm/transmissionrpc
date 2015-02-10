package transmissionrpc

import (
	"encoding/json"
)

type Session struct {
}

type SessionStats struct {
	NumTorrents int     `json:"torrentCount,omitempty"`
	DownSpeed   float64 `json:"downloadSpeed,omitempty"`
	UpSpeed     float64 `json:"uploadSpeed,omitempty"`
}

func (client *Client) GetSessionStats() (SessionStats, error) {
	request := Request{
		Method: "session-stats",
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return SessionStats{}, err
	}

	var response SessionStats
	err = json.Unmarshal(resp.Args, &response)
	if err != nil {
		return SessionStats{}, err
	}

	return response, nil
}
