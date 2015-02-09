package main

import (
	"fmt"
	"github.com/h00kwurm/transmissionrpc"
	"time"
)

func main() {
	transmission := transmissionrpc.New("http://192.168.0.106", "9091")
	getTorrents(transmission)

	torrentId, err := addTorrent(transmission)
	if err != nil {
		fmt.Println("failed adding torrent")
		return
	}

	// this is so you can watch it work from another UI (web ui in my case)
	time.Sleep(10 * time.Second)
	moveTorrent(transmission, []int{torrentId})

	// this is so you can watch it work from another UI (web ui in my case)
	time.Sleep(10 * time.Second)
	removeTorrent(transmission, []int{torrentId})

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

func addTorrent(client *transmissionrpc.Client) (int, error) {
	torrent, err := client.AddTorrent("http://sample-file.bazadanni.com/download/applications/torrent/sample.torrent", "/home/anatraj/Downloads")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	fmt.Println("added torrent: ", torrent)
	return torrent.Id, nil
}

func removeTorrent(client *transmissionrpc.Client, ids []int) {
	err := client.RemoveTorrent(ids, true)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func moveTorrent(client *transmissionrpc.Client, ids []int) {
	err := client.MoveTorrent(ids, "/home/anatraj/Music", true)
	if err != nil {
		fmt.Println(err)
		return
	}
}
