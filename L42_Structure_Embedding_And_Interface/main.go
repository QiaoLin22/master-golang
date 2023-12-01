package main

import "fmt"

type Tile struct{}

type TileWalker interface {
	WalkTile(Tile)
}

type Updater interface {
	Update()
}

type Transform struct {
	position int
}

type Enemy struct {
	Transform
	tileWalker TileWalker
}

func (e *Enemy) checkTilesCollided() {
	fmt.Println("enemy walking on the tile", e.position)
	e.tileWalker.WalkTile(Tile{})
}

func (e *Enemy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

type FireEnemy struct {
	*Enemy
}

func (e *FireEnemy) WalkTile(tile Tile) {
	fmt.Println("fire enemy is walking on tile")
}

func main() {
	e := &FireEnemy{}
	e.Enemy = &Enemy{
		tileWalker: e,
	}
	for i := 0; i < 100; i++ {
		Update(e)
	}
}

func Update(u Updater) {
	u.Update()
}
