package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type SetState struct {
	value uint
}

type ResetState struct{}

type Handler struct {
	state uint
}

func newHandler() actor.Receiver {
	return &Handler{}
}

func (h *Handler) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case SetState:
		h.state = msg.value
		fmt.Println("handler received new state", h.state)
	case actor.Initialized:
		h.state = 10
		fmt.Println("handler initialized, my state", h.state)
	case actor.Started:
		fmt.Println("handler started")
	case actor.Stopped:
		_ = msg
	}
}

func (h *Handler) handleMessage(msg uint) {
	h.state = msg
}

func main() {
	e := actor.NewEngine()
	pid := e.Spawn(newHandler, "handler")
	for x := 0; x < 10; x++ {
		e.Send(pid, SetState{value: uint(x)})
	}
	e.Send(pid, ResetState{})
}
