package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	game := NewGame()

	for {
		err := game.ProcessInput()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

type GameState interface {
	PressButton(number int) error
}

type Game struct {
	splash       GameState
	xTurn        GameState
	oTurn        GameState
	Ende         GameState
	world        *World
	currentState GameState
}

func NewGame() *Game {

	g := &Game{}

	splash := NewSplash(g)
	xTurn := NewXTurn(g)
	oTurn := NewOTurn(g)
	ende := NewEnde(g)
	world := NewWorld(g)

	g.currentState = splash

	fmt.Println("new game created. State: splash")

	g.splash = splash
	g.xTurn = xTurn
	g.oTurn = oTurn
	g.Ende = ende
	g.world = world
	return g
}

func (w *Game) ProcessInput() error {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)

	text = strings.TrimSuffix(text, "\n")

	if text == "end" {
		os.Exit(0)
		return nil
	}

	num, err := strconv.Atoi(text)

	if err != nil {
		return fmt.Errorf("not a valid number. please enter a number of 1 to 9")
	}

	pressButtonErr := w.currentState.PressButton(num)

	if pressButtonErr != nil {
		fmt.Println(pressButtonErr.Error())
		return err
	}

	return nil

}

func (w *Game) SetNewState(newState GameState) {
	//w.world.printWorld()

	w.currentState = newState
}
