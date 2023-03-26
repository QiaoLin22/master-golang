package main

import (
	"flag"
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

type store struct{}

func newStore() actor.Receiver {
	return &store{}
}

func (s *store) Receive(c *actor.Context) {
	switch msg := c.Message.(type) {
	case *types.CatFact:
		fmt.Println("stored fact into db", msg.Fact)
	case actor.Started:
		fmt.Println("store is started")
	case actor.Stopped:
	}
}

func main() {
	listenaddr := flag.String("listenaddr", "127.0.0.1:4000", "todo")
	flag.Parse()

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{ListenAddr: *listenaddr})
	e.WithRemote(r)
	e.Spawn(newStore, "store")

	select {}
}
