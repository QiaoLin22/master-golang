package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

type Position struct {
	x float64
	y float64
}

type Player struct {
	*Position
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {
	player := NewPlayer()
	fmt.Println(player.Position)
	boss := NewEnemy()
	boss.Move(1, 2)
	fmt.Println(boss.Position)
	boss.MoveSpecial(1, 2)
	fmt.Println(boss.Position)
}
