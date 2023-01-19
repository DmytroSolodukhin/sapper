package output

import (
	"fmt"

	"github.com/DmytroSolodukhin/sapper/model"
)

// Validator usage for getting parameters.
type Validator interface {
	ValidSizeParameter(a int) error
	ValidateGameParameters(c, mines int) error
	GetMinParameters() int
	GetMaxParameters() int
}

// Show game arena.
func Show(p model.Arena) {
	for _, x := range p {
		for _, y := range x {
			i := " "
			if y.IsOpen() {
				i = y.GetMarker()
			}
			fmt.Printf("|%v", i)
		}
		fmt.Print("|\n")
	}
}

// ShowFull showing opened game arena with all information.
func ShowFull(p model.Arena) {
	for _, x := range p {
		for _, y := range x {
			i := y.GetMarker()
			if y.IsMine() {
				i = "X"
			}
			fmt.Printf("|%v", i)
		}
		fmt.Print("|\n")
	}
}

func AskCoordinate() (x, y int) {
	fmt.Println("please set number of column what you want to open.")
	fmt.Scan(&y)

	fmt.Println("please set number of string what you want to open.")
	fmt.Scan(&x)

	return
}

func Lost() {
	fmt.Print("You lost!")
}

func Win() {
	fmt.Print("You Win!")
}

func ErrorMsg(msg string) {
	fmt.Println(msg)
}

func InitializePlayingField(v Validator) (x, y, mines int) {
	fmt.Printf("please specify the length of the field (min: %v, max: %v)\n", v.GetMinParameters(), v.GetMaxParameters())
INIT_LEN:
	fmt.Scan(&x)

	err := v.ValidSizeParameter(x)
	if err != nil {
		fmt.Printf("%v. Please try again.\n", err)

		goto INIT_LEN
	}

	fmt.Printf("please specify the height of the field (min: %v, max: %v)\n", v.GetMinParameters(), v.GetMaxParameters())
INIT_HEIGHT:
	fmt.Scan(&y)

	err = v.ValidSizeParameter(y)
	if err != nil {
		fmt.Printf("%v. Please try again.\n", err)

		goto INIT_HEIGHT
	}

	maxMines := x*y - 10

	fmt.Printf("please specify the count of mines (min: %v, max: %v)\n", v.GetMinParameters(), maxMines)

INIT_MINES:
	fmt.Scan(&mines)

	err = v.ValidateGameParameters(maxMines, mines)
	if err != nil {
		fmt.Printf("%v. Please try again.\n", err)

		goto INIT_MINES
	}

	return
}
