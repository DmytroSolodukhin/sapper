package sapper

import (
	"fmt"
	"github.com/DmytroSolodukhin/sapper/game"
	"log"
	"runtime/debug"
)

func main()  {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic: %s: %s", r, string(debug.Stack()))
		}
	}()

	game.New()
}
