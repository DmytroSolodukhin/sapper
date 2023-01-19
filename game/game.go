package game

import (
	"fmt"

	"github.com/DmytroSolodukhin/sapper/model"
	out "github.com/DmytroSolodukhin/sapper/output"
	"github.com/DmytroSolodukhin/sapper/validators"
)

const (
	minParameterSize = 5
	maxParameterSize = 100
)

// Output Arena for player visualization.
type Output interface {
	Show()
	ShowFull()
	AskCoordinate()
	Loose()
	Win()
	ErrorMsg()
}

// Start game.
func Start() {
	validator := validators.New(minParameterSize, maxParameterSize)
	x, y, mines := out.InitializePlayingField(validator)
	arena := model.New(x, y, mines)
	run(arena)
}

func run(arena model.Arena) {
	out.Show(arena)

	x, y := out.AskCoordinate()

	err := checkCell(arena, x, y)
	if err != nil {
		out.ErrorMsg(err.Error())
		run(arena)

		return
	}

	if !arena.Open(x, y) {
		out.ShowFull(arena)
		out.Lost()

		return
	}

	if arena.Win() {
		out.ShowFull(arena)
		out.Win()

		return
	}

	run(arena)
}

func checkCell(arena model.Arena, x, y int) error {
	if x < 0 || x > len(arena)-1 || y < 0 || y > len(arena[0])-1 {
		return fmt.Errorf("cell %v, %v does not exist", x, y)
	}

	if arena[x][y].IsOpen() {
		return fmt.Errorf("cell %v, %v allready opened", x, y)
	}

	return nil
}
