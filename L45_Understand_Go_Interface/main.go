package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f CR7) KickBall() {
	shot := f.stamina + f.power*f.SUI
	fmt.Println("CR7 is kicking the ball", shot)
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("I am kicking the ball", shot)
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			stamina: int(rand.Intn(10)),
			power:   int(rand.Intn(10)),
		}
	}
	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}

}
