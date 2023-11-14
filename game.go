package main

type Game struct {
	score  int
	snake  *Snake
	food   position
	drawer *TermDrawer
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
	//g.drawer.HideCursor()
	g.drawer.MoveCursor([2]int{1, 10})
	g.drawer.ShowCursor()
}
