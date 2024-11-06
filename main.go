package main

import (
	"time"

	"github.com/Daniel-Rammirez/blockchain/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)
	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Trum won"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
