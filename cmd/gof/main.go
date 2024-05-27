package main

import (
	"github.com/danblok/gameoflife/internal/game"
)

func main() {
	g := game.New(32, 128)
	g.Start()
}
