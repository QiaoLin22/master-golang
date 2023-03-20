package server

import (
	"testing"
)

func TestServer(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		// player := &Player{
		// 	Name: fmt.Sprintf("player_%d", i),
		// }
		// go server.handleNewPlayer(player)
		go server.handleSetFoo(i)
	}
}
