package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initArena() Arena {
	return Arena{
		[]Cell{
			{}, {}, {},
		},
		[]Cell{
			{}, {hflag: true}, {hflag: true},
		},
		[]Cell{
			{}, {hflag: true}, {},
		},
	}
}

func TestMarkedAndOpenCell(t *testing.T) {
	arena := initArena()

	for x, xVal := range arena {
		for y, cell := range xVal {
			if cell.IsMine() {
				arena.visiting(arena.marked, x, y)
			}
		}
	}

	assert.Equal(t, 1, arena[0][0].marker)
	assert.Equal(t, 2, arena[0][1].marker)
	assert.Equal(t, 2, arena[0][2].marker)
	assert.Equal(t, 2, arena[1][0].marker)
	assert.Equal(t, 2, arena[2][0].marker)
	assert.Equal(t, 3, arena[2][2].marker)

	arena.Open(0, 0)

	assert.Equal(t, true, arena[0][0].IsOpen())
	assert.Equal(t, true, arena[0][1].IsOpen())
	assert.Equal(t, false, arena[0][2].IsOpen())
	assert.Equal(t, true, arena[1][0].IsOpen())
	assert.Equal(t, false, arena[1][1].IsOpen())
	assert.Equal(t, false, arena[1][2].IsOpen())
	assert.Equal(t, false, arena[2][0].IsOpen())
	assert.Equal(t, false, arena[2][1].IsOpen())
	assert.Equal(t, false, arena[2][2].IsOpen())

	arena.Open(2, 2)

	assert.Equal(t, true, arena[0][0].IsOpen())
	assert.Equal(t, true, arena[0][1].IsOpen())
	assert.Equal(t, false, arena[0][2].IsOpen())
	assert.Equal(t, true, arena[1][0].IsOpen())
	assert.Equal(t, false, arena[1][1].IsOpen())
	assert.Equal(t, false, arena[1][2].IsOpen())
	assert.Equal(t, false, arena[2][0].IsOpen())
	assert.Equal(t, false, arena[2][1].IsOpen())
	assert.Equal(t, true, arena[2][2].IsOpen())

	arena.Open(0, 2)

	assert.Equal(t, true, arena[0][0].IsOpen())
	assert.Equal(t, true, arena[0][1].IsOpen())
	assert.Equal(t, true, arena[0][2].IsOpen())
	assert.Equal(t, true, arena[1][0].IsOpen())
	assert.Equal(t, false, arena[1][1].IsOpen())
	assert.Equal(t, false, arena[1][2].IsOpen())
	assert.Equal(t, false, arena[2][0].IsOpen())
	assert.Equal(t, false, arena[2][1].IsOpen())
	assert.Equal(t, true, arena[2][2].IsOpen())
}

func TestWin(t *testing.T) {
	arena := initArena()

	for x, xVal := range arena {
		for y, cell := range xVal {
			if cell.IsMine() {
				arena.visiting(arena.marked, x, y)
			}
		}
	}

	assert.Equal(t, false, arena.Win())

	for x, xVal := range arena {
		for y, cell := range xVal {
			if !cell.IsMine() {
				arena[x][y].open = true
			}
		}
	}

	assert.Equal(t, true, arena.Win())
}
