package main

import "fmt"

type Ende struct {
	game *Game
}

func NewEnde(g *Game) *Ende {
	return &Ende{
		game: g,
	}
}

func (e *Ende) PressButton(number int) error {
	fmt.Println("a button was pressed in end state, going to splash screen")

	e.game.SetNewState(e.game.splash)

	return nil
}
