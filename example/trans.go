package main

import (
	"fmt"
	"github.com/h00kwurm/transmissionrpc"
)

func main() {
	getTorrents()
}

func getTorrents() {
	transmission := transmissionrpc.New("http://192.168.0.106", "9091")
	torrents, err := transmission.GetTorrents()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(torrents[i])
	}
}
