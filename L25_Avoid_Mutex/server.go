package server

import "fmt"

type Player struct {
	Name string
}

type SetFooMsg struct {
	value int
}

type GameState struct {
	players []*Player
	msgch   chan any
	foo     int
}

func (g *GameState) Receive(msg any) {
	g.msgch <- msg
}

func (g *GameState) loop() {
	for msg := range g.msgch {
		g.handleMessage(msg)
	}
}

func (g *GameState) handleMessage(message any) {
	switch msg := message.(type) {
	case *Player:
		g.addPlayer(msg)
	case *SetFooMsg:
		g.handleSetFoo(msg)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) handleSetFoo(foo *SetFooMsg) {
	g.foo = foo.value
	fmt.Println("setting foo", foo.value)
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)
	fmt.Println("adding player:", p.Name)
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch:   make(chan any, 10),
	}

	go g.loop()

	return g
}

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

func (s *Server) handleSetFoo(val int) error {
	s.gameState.Receive(&SetFooMsg{value: val})
	return nil
}

func (s *Server) handleNewPlayer(player *Player) error {
	s.gameState.Receive(player)
	return nil
}
