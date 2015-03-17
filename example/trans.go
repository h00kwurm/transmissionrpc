package main

import (
	"bufio"
	"fmt"
	"github.com/h00kwurm/transmissionrpc"
	"os"
	// "time"
)

func main() {
	transmission := transmissionrpc.New("http://192.168.0.105", "9091")

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("trans$ ")
		reader.Scan()
		text := reader.Text()
		if text == "list" {
			getTorrents(transmission)
		} else if text == "stats" {
			getSessionStats(transmission)
		} else if text == "help" {
			printHelp()
		} else if text == "q" || text == "quit" {
			break
		}
	}

}

func printHelp() {
	fmt.Println("Usage: super-basic example of a client\nCommands\n\tlist\n\tstats\n\thelp\n\tquit")
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

func getSessionStats(client *transmissionrpc.Client) {
	fmt.Println("getting session stats")
	stats, err := client.GetSessionStats()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stats)
}
