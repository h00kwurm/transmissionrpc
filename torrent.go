package transmissionrpc

import ()

type Torrent struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Size      int    `json:"sizeWhenDone,omitempty"`
	AddedTime int    `json:"addedDate",omitempty"`
	Files     []File `json:"files,omitempty"`
	Location  string `json:"downloadDir,omitempty"`
	TorrentError
}

type TorrentError struct {
	Error       int    `json:"error,omitempty"`
	ErrorString string `json:"errorString,omitempty"`
}

func (client *Client) GetTorrents() ([]Torrent, error) {
	return client.GetTorrentsWithFields([]string{"name", "id", "size", "addedDate", "downloadDir"})
}

func (client *Client) GetTorrentsWithFields(fields []string) ([]Torrent, error) {

	request := Request{
		Method: "torrent-get",
		Args: RequestArguments{
			Fields: fields,
		},
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return []Torrent{}, err
	}
	return resp.Args.Torrents, nil

}
