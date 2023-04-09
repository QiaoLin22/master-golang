package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {
		// block here until someone is sending a message to the channel
		case msg := <-s.msgch:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a graceful shutdown")
			// logic for graceful shutdown
			break free
		default:
		}
	}

	fmt.Println("the server is shut down!")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "qiao",
		Payload: payload,
	}

	msgch <- msg
}

func gracefulQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello qiaolin")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		gracefulQuitServer(s.quitch)
	}()

	select {}
}
