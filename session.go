package transmissionrpc

import (
	"errors"
	"fmt"
)

type Session struct {
}

type SessionStats struct {
	NumTorrents int     `json:"torrentCount,omitempty"`
	DownSpeed   float64 `json:"downloadSpeed,omitempty"`
	UpSpeed     float64 `json:"uploadSpeed,omitempty"`
}

func (client *Client) GetSessionStats() error {
	request := Request{
		Method: "session-stats",
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return err
	}
	fmt.Println(resp)

	if resp.Result != "success" {
		return errors.New("something totally busted because transmission doesnt care about unfound ids")
	} else {
		return nil
	}
}
