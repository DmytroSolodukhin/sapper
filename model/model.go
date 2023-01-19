package model

import (
	"math/rand"
	"strconv"
	"time"
)

// Arena minefield - is a slice with length x as coordinates, each element of which is a slice with length y containing a cell.
type Arena [][]Cell

// Cell that can contain a mine or information about neighboring cells.
type Cell struct {
	open   bool
	hflag  bool
	marker int
}

// New initialize playing field.
func New(x, y, mines int) Arena {
	arena := make(Arena, x)

	for i := range arena {
		arena[i] = make([]Cell, y)
	}

	for i := 0; i <= mines; i++ {
		arena.randomSetMine()
	}

	return arena
}

// Open cell.
func (p Arena) Open(x, y int) bool {
	if p[x][y].IsMine() {
		return false
	}

	p[x][y].open = true

	p.visiting(p.open, x, y)

	return true
}

func (p Arena) Win() bool {
	for _, x := range p {
		for _, y := range x {
			if !y.open && !y.hflag {
				return false
			}
		}
	}

	return true
}

func (p Arena) randomSetMine() {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(len(p))    //nolint
	y := rand.Intn(len(p[0])) //nolint

	if p[x][y].IsMine() {
		p.randomSetMine()

		return //nolint
	}

	p[x][y].hflag = true

	p.visiting(p.marked, x, y)
}

func (p Arena) marked(x, y int) {
	p[x][y].marker++
}

func (p Arena) open(x, y int) {
	if p[x][y].open || p[x][y].IsMine() {
		return
	}

	p[x][y].open = true

	if p[x][y].marker == 0 {
		p.Open(x, y)
	}
}

func (p Arena) visiting(f func(x, y int), x, y int) {
	if x > 0 {
		f(x-1, y)

		if y > 0 {
			f(x-1, y-1)
		}

		if y < len(p[0])-1 {
			f(x-1, y+1)
		}
	}

	if y > 0 {
		f(x, y-1)
	}

	if y < len(p[0])-1 {
		f(x, y+1)
	}

	if x < len(p)-1 {
		f(x+1, y)

		if y > 0 {
			f(x+1, y-1)
		}

		if y < len(p)-1 {
			f(x+1, y+1)
		}
	}
}

// IsOpen check cell.
func (c Cell) IsOpen() bool {
	return c.open
}

// IsMine check mine.
func (c Cell) IsMine() bool {
	return c.hflag
}

// GetMarker show information about neighboring cells.
func (c Cell) GetMarker() string {
	return strconv.FormatInt(int64(c.marker), 10)
}
