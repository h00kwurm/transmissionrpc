package transmissionrpc

import (
	"encoding/json"
	"errors"
)

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

type GetTorrentsResponse struct {
	Torrents []Torrent `json:"torrents,omitempty"`
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

	response := GetTorrentsResponse{}
	err = json.Unmarshal(resp.Args, &response)
	if err != nil {
		return []Torrent{}, err
	}

	return response.Torrents, nil
}

type AddTorrentArguments struct {
	URI      string `json:"filename,omitempty"`
	Location string `json:"download-dir,omitempty"`
}

type AddTorrentResponse struct {
	Added Torrent `json:"torrent-added,omitempty"`
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

	response := AddTorrentResponse{}
	err = json.Unmarshal(resp.Args, &response)
	if err != nil {
		return Torrent{}, err
	}

	return response.Added, nil
}

type RemoveTorrentArguments struct {
	Ids          []int `json:"ids,omitempty"`
	ShouldDelete bool  `json:"delete-local-data,omitempty"`
}

// this is misleading because while it takes an array it only takes one
// "should i delete it" boolean. its a vestige of the rpc call. i'd
// assume that in practice it just distributes the bool over all ids
// I hate it. i might make it an official single id and wrap this
// for any multiplicity.
func (client *Client) RemoveTorrent(ids []int, clean bool) error {

	request := Request{
		Method: "torrent-remove",
		Args: RemoveTorrentArguments{
			Ids:          ids,
			ShouldDelete: clean,
		},
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return err
	}

	if resp.Result != "success" {
		return errors.New("something totally busted because transmission doesnt care about unfound ids")
	} else {
		return nil
	}
}

type MoveTorrentArguments struct {
	Ids        []int  `json:"ids,omitempty"`
	Location   string `json:"location,omitempty"`
	ShouldMove bool   `json:"move,omitempty"`
}

// I feel the same way about this as i do the remove function
// the rpc is irritating. maybe that's the point of writing something like this
// cleaning up the underlying stuff. i think with that thought expect soon
// me to make these only take single ids and have helper functions for many
// with some structure like [{id, location, move}] as args
func (client *Client) MoveTorrent(ids []int, location string, move bool) error {

	request := Request{
		Method: "torrent-set-location",
		Args: MoveTorrentArguments{
			Ids:        ids,
			Location:   location,
			ShouldMove: move,
		},
	}

	resp, err := client.makeRequest(request)
	if err != nil {
		dealWithIt(err.Error())
		return err
	}

	if resp.Result != "success" {
		return errors.New("something totally busted because transmission doesnt care about unfound ids")
	} else {
		return nil
	}
}
