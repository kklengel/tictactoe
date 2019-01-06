package main

import "fmt"

type World struct {
	game   *Game
	fields map[int]string
}

func NewWorld(g *Game) *World {
	return &World{
		game:   g,
		fields: map[int]string{},
	}
}

//reset the map of fields on world
func (w *World) Reset() {
	fmt.Println("resetting world")
	w.fields = map[int]string{}
	w.printWorld()
}

//PlaceStone returns no error if the player could place the stone on given number
func (w *World) PlaceStone(player string, number int) error {

	fmt.Println("player (", player, ") wants to place stone at ", number)

	return w.placeStone(player, number)
}

func (w *World) placeStone(player string, number int) error {

	if number < 0 || number > 9 {
		return fmt.Errorf("%v is not a valid number to place", number)
	}

	if _, exists := w.fields[number]; !exists {
		w.fields[number] = player
		return nil
	}

	return fmt.Errorf("player %s could not place the stone at %v because it already belongs to %s", player, number, w.fields[number])

}

//checkForWin returns (true, false) if a player won the  game. returns (false, true) on a draw and (false,false) on neither of them.
func (w *World) checkForWinOrDraw(player string) (bool, bool) {
	//1, 2, 3
	//4, 5, 6
	//7, 8, 9
	owner1, exists1 := w.fields[1]
	owner2, exists2 := w.fields[2]
	owner3, exists3 := w.fields[3]
	owner4, exists4 := w.fields[4]
	owner5, exists5 := w.fields[5]
	owner6, exists6 := w.fields[6]
	owner7, exists7 := w.fields[7]
	owner8, exists8 := w.fields[8]
	owner9, exists9 := w.fields[9]

	//check horizontal
	if exists1 && exists2 && exists3 {
		if owner1 == owner2 && owner2 == owner3 && owner1 == player {
			fmt.Println("player ", player, " owns  horizontal row 1")
			return true, false
		}
	}

	if exists4 && exists5 && exists6 {
		if owner4 == owner5 && owner5 == owner6 && owner4 == player {
			fmt.Println("player ", player, " owns  horizontal row 2")
			return true, false
		}
	}

	if exists7 && exists8 && exists9 {
		if owner7 == owner8 && owner8 == owner9 && owner7 == player {
			fmt.Println("player ", player, " owns  horizontal row 3")
			return true, false
		}
	}
	//check vertical

	if exists1 && exists4 && exists7 {
		if owner1 == owner4 && owner4 == owner7 && owner1 == player {
			fmt.Println("player ", player, " owns  vertical column 1")
			return true, false
		}
	}

	if exists2 && exists5 && exists8 {
		if owner2 == owner5 && owner5 == owner8 && owner2 == player {
			fmt.Println("player ", player, " owns  vertical column 2")
			return true, false
		}
	}

	if exists3 && exists6 && exists9 {
		if owner3 == owner6 && owner6 == owner9 && owner3 == player {
			fmt.Println("player ", player, " owns  vertical column 3")
			return true, false
		}
	}
	//check diagonal
	if exists1 && exists5 && exists9 {
		if owner1 == owner5 && owner5 == owner9 && owner1 == player {
			fmt.Println("player ", player, " owns  diagonal")
			return true, false
		}
	}

	if exists3 && exists5 && exists7 {
		if owner3 == owner5 && owner5 == owner7 && owner3 == player {
			fmt.Println("player ", player, " owns reverse diagonal")
			return true, false
		}
	}

	//if one of the fields is not set, we dont have to check for draw
	for i := 1; i < 11; i++ {

		// if we would check for the 10th field and still did not exit, all fields are set and there  is no winner, so we assume a draw and return false, true
		if i == 10 {
			return false, true
		}

		if _, exists := w.fields[i]; !exists {
			break
		}
	}

	return false, false
}

func (w *World) printWorld() {
	for i := 1; i < 10; i++ {

		if i == 4 || i == 7 {
			fmt.Println()
		}
		if _, exists := w.fields[i]; exists {
			fmt.Printf(" %s ", w.fields[i])
		} else {
			fmt.Printf(" %s ", "_")
		}
	}
	fmt.Println()
}
