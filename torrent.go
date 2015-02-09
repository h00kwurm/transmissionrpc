package transmissionrpc

import ()

type Torrent struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Size      int    `json:"sizeWhenDone,omitempty"`
	AddedTime int    `json:"addedDate",omitempty"`
	Files     []File `json:"files,omitempty"`
	Location  string `json:"downloadDir,omitempty"`
	Hash      string `json:"hashString,omitEmpty"`
	TorrentError
}

type TorrentError struct {
	Error       int    `json:"error,omitempty"`
	ErrorString string `json:"errorString,omitempty"`
}

type GetTorrentArguments struct {
	Ids    []int    `json:"ids,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

func (client *Client) GetTorrents() ([]Torrent, error) {
	return client.GetTorrentsWithFields([]string{"name", "id", "size", "addedDate", "downloadDir"})
}

func (client *Client) GetTorrentsWithFields(fields []string) ([]Torrent, error) {

	request := Request{
		Method: "torrent-get",
		Args: GetTorrentArguments{
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

type AddTorrentArguments struct {
	URI      string `json:"filename,omitempty"`
	Location string `json:"download-dir,omitempty"`
}

//TODO: decide whether we should get the url ourselves
//  and then base64 encode the .torrent contents
//  to shove it in `metadata` instead of passing
//  the pain down to transmission
func (client *Client) AddTorrent(url, location string) (Torrent, error) {

	request := Request{
		Method: "torrent-add",
		Args: AddTorrentArguments{
			URI:      url,
			Location: location,
		},
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return Torrent{}, err
	}

	return resp.Args.Added, nil
}
