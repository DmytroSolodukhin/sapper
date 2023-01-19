package main

import (
	"log"
	"runtime/debug"

	"github.com/DmytroSolodukhin/sapper/game"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic: %s: %s", r, string(debug.Stack()))
		}
	}()

	game.Start()
}
