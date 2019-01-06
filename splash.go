package main

import "fmt"

type Splash struct {
	game *Game
}

func NewSplash(g *Game) *Splash {
	return &Splash{
		game: g,
	}
}

func (s *Splash) PressButton(number int) error {
	fmt.Println("Button pressed on splash screen, starting game")

	s.game.world.Reset()

	s.game.SetNewState(s.game.xTurn)

	return nil
}
