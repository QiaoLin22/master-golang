package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("player is taking damage")
	player.health -= dmg
}

func takeDamageFromExplosion(player *Player, dmg int) {
	fmt.Println("player is taking damage")
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	fmt.Printf("before explosion %+v\n", player)
	player.takeDamageFromExplosion(50)
	fmt.Printf("after explosion %+v\n", player)
}
