package main

import (
	"fmt"
	"github.com/h00kwurm/transmissionrpc"
)

func main() {
	transmission := transmissionrpc.New("http://192.168.0.106", "9091")
	getTorrents(transmission)
	addTorrent(transmission)
}

func getTorrents(client *transmissionrpc.Client) {
	torrents, err := client.GetTorrents()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(torrents[i])
	}
}

func addTorrent(client *transmissionrpc.Client) {
	torrent, err := client.AddTorrent("http://sample-file.bazadanni.com/download/applications/torrent/sample.torrent", "/home/anatraj/Downloads")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("added torrent: ", torrent)
}
