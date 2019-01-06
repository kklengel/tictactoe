package main

import "fmt"

type OTurn struct {
	game *Game
	name string
}

func NewOTurn(g *Game) *OTurn {
	return &OTurn{
		game: g,
		name: "O",
	}
}

func (o *OTurn) PressButton(number int) error {
	placeStoneErr := o.game.world.PlaceStone(o.name, number)

	if placeStoneErr != nil {
		return placeStoneErr
	}

	o.game.world.printWorld()

	//check if o won after placing that stone
	won, draw := o.game.world.checkForWinOrDraw(o.name)
	if won {
		fmt.Println("Player ", o.name, " won the game!  Press a button to go to the splash screen again")
		o.game.SetNewState(o.game.Ende)
		return nil
	}

	if draw {
		fmt.Println("we have a draw!  Press a button to go to the splash screen again")
		o.game.SetNewState(o.game.Ende)
		return nil
	}

	o.game.SetNewState(o.game.xTurn)

	return nil
}
