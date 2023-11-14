package main

type Game struct {
	score int
	snake *Snake
	food  position
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

func newSnake() *Snake {
	//TODO: Remove hardcode
	maxX, maxY := 10, 10
	pos := position{maxX / 2, maxY / 2}

	return &Snake{
		body:      []position{pos},
		direction: up,
	}
}

func newGame() *Game {
	newGameSnake := newSnake()

	game := &Game{
		score: 0,
		snake: newGameSnake,
		food:  position{},
	}
	return game
}
