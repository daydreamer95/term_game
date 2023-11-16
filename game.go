package main

import (
	"strconv"
	"time"
)

type Game struct {
	score      int
	termHeight int
	termWidth  int
	snake      *Snake
	wall       []position
	food       position
	drawer     *TermDrawer
}

type Snake struct {
	body      []position
	direction direction
}

type position [2]int

type direction int

const (
	up direction = iota
	down
	left
	right
)

func NewSnake() *Snake {
	//TODO: Remove hardcode
	maxX, maxY := 10, 10
	pos := position{maxX / 2, maxY / 2}

	return &Snake{
		body:      []position{pos},
		direction: up,
	}
}

func NewGame() *Game {
	newGameSnake := NewSnake()

	game := &Game{
		score:  0,
		snake:  newGameSnake,
		food:   position{},
		drawer: NewTermDrawer(),
	}
	return game
}

func (g *Game) mustBeforeGame() {
	g.drawer.HideCursor()
	w, h := getTermSize()
	g.termWidth = w
	g.termHeight = h
	for i := 2; i <= h; i++ {
		for j := 0; j <= w; j++ {
			//Top and bottom wall
			if i == 2 || i == h {
				g.wall = append(g.wall, position{j, i})
				continue
			}

			// Left and Right border
			if j == 1 || j == w {
				g.wall = append(g.wall, position{j, i})
			}
		}
	}

	//fmt.Printf("Terminal width %v height %v\n", w, h)
	//fmt.Println("Wall: ", g.wall)
}

func (g *Game) draw() {
	g.drawer.Clear()
	g.drawer.MoveCursor(position{0, 0})

	//draw status bar
	status := "score: " + strconv.Itoa(g.score)
	statusXPos := g.termWidth/2 - len(status)/2

	g.drawer.MoveCursor(position{statusXPos, 1})
	g.drawer.Write(status)

	//draw wall
	for _, wallPos := range g.wall {
		g.drawer.MoveCursor(wallPos)
		g.drawer.Write("#")
	}
	g.drawer.Render()
}

func (g *Game) Start() {
	//fmt.Println("Start game loop")
	g.draw()

	for {
		time.Sleep(time.Millisecond * 100)
	}
}
