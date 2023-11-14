package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	game := NewGame()
	game.mustBeforeGame()
	go game.listenForKeyPress()

	//listen for terminate
	c := make(chan os.Signal)
	done := make(chan bool, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
	<-done
	fmt.Println("exiting")

}

func (g *Game) listenForKeyPress() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			panic(err)
			return
		}
		switch string(b) {
		case "A":
			g.snake.direction = up
		case "B":
			g.snake.direction = down
		case "C":
			g.snake.direction = right
		case "D":
			g.snake.direction = left
		}
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}
}
