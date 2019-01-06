package main

import "fmt"

type XTurn struct {
	game *Game
	name string
}

func NewXTurn(g *Game) *XTurn {
	return &XTurn{
		game: g,
		name: "X",
	}
}

func (x *XTurn) PressButton(number int) error {
	//x holt sich die world und führt dann eine aktion aus

	placeStoneErr := x.game.world.PlaceStone(x.name, number)

	if placeStoneErr != nil {
		return placeStoneErr
	}

	x.game.world.printWorld()

	won, draw := x.game.world.checkForWinOrDraw(x.name)
	//check if x won after placing that stone
	if won {
		fmt.Println("Player ", x.name, " won the game!  Press a button to go to the splash screen again")
		x.game.SetNewState(x.game.Ende)
		return nil
	}

	if draw {
		fmt.Println("we have a draw! Press a button to go to the splash screen again")
		x.game.SetNewState(x.game.Ende)
		return nil
	}

	x.game.SetNewState(x.game.oTurn)
	return nil
}
