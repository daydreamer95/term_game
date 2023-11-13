package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	game := newGame()
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	//scanner := bufio.NewScanner(reader)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
