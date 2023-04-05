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
	msgch chan Message
}

func (s *Server) StartAndListen() {
	for {
		select {
		// block here until someone is sending a message to the channel
		case msg := <-s.msgch:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
		default:
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "qiao",
		Payload: payload,
	}

	msgch <- msg
}

func main() {
	s := &Server{
		msgch: make(chan Message),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(time.Second)
		sendMessageToServer(s.msgch, "Hello qiaolin")
	}()

	select {}
}
