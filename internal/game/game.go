package game

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	dead  rune = ' '
	alive rune = '\u25A0'
)

type Game struct {
	rows, cols  int
	front, back [][]rune

	updateTime time.Duration
}

func New(rows, cols int) *Game {
	if rows < 2 || cols < 2 {
		panic("rows and cols must be greater than 2")
	}
	back := make([][]rune, rows)
	front := make([][]rune, rows)

	g := &Game{
		rows:       rows,
		cols:       cols,
		back:       back,
		front:      front,
		updateTime: 16 * time.Millisecond,
	}

	for y := range rows {
		back[y] = make([]rune, cols)
		front[y] = make([]rune, cols)
		for x := range back[y] {
			if rand.Int()%2 == 0 {
				back[y][x] = dead
			} else {
				back[y][x] = alive
			}
		}
	}

	g.updateFront()

	return g
}

func (g *Game) SetUpdateTime(t time.Duration) {
	g.updateTime = t
}

func (g *Game) Start() {
	if g.back == nil || g.front == nil {
		panic("no initial state")
	}

	for {
		g.Print()
		g.Next()
		time.Sleep(g.updateTime)
	}
}

func (g *Game) Next() {
	for y := range g.back {
		for x := range g.back[y] {
			ns := g.countNeighbours(x, y)
			if g.front[y][x] == alive && (ns == 2 || ns == 3) {
				g.back[y][x] = alive
			} else if g.front[y][x] == dead && ns == 3 {
				g.back[y][x] = alive
			} else {
				g.back[y][x] = dead
			}
		}
	}

	g.updateFront()
}

func (g *Game) Restart() {
	for y := range g.rows {
		for x := range g.cols {
			g.back[y][x] = dead
			if rand.Int()%2 == 0 {
				g.back[y][x] = alive
			}
		}
	}

	g.updateFront()
}

func (g *Game) updateFront() {
	for y := range g.rows {
		copy(g.front[y], g.back[y])
	}
}

func (g *Game) countNeighbours(cx, cy int) int {
	var count int
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			yy := mod(cy+y, g.rows)
			xx := mod(cx+x, g.cols)

			if yy == cy && xx == cx {
				continue
			}

			if g.front[yy][xx] == alive {
				count++
			}
		}
	}

	return count
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (g *Game) Print() {
	clean()
	g.Display(os.Stdout)
}

func clean() {
	fmt.Print("\033[H\033[2J")
}

func (g *Game) Display(w io.Writer) {
	for y := range g.front {
		for x := range g.front[y] {
			fmt.Fprintf(w, "%c", g.front[y][x])
		}
		fmt.Fprintln(w)
	}
}
