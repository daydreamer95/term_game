package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
)

type Writer struct {
}

func (*Writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type TermDrawer struct {
	buf *bufio.Writer
}

func NewTermDrawer() *TermDrawer {
	return &TermDrawer{
		buf: bufio.NewWriter(os.Stdout),
	}
}

func (t *TermDrawer) HideCursor() {
	fmt.Fprintf(t.buf, "\u001b[?25l")
}

func (t *TermDrawer) ShowCursor() {
	fmt.Fprintf(t.buf, "\u001b[?25h")
}

func (t *TermDrawer) Clear() {
	fmt.Fprintf(t.buf, "\u001b[2J")
}

func (t *TermDrawer) MoveCursor(pos [2]int) {
	fmt.Fprintf(t.buf, "\u001b[%d;%dH", pos[1], pos[0])
}

func (t *TermDrawer) Write(s string) {
	fmt.Fprintf(t.buf, s)
}

func (t *TermDrawer) Render() {
	t.buf.Flush()
}

func getTermSize() (int, int) {
	w, h, err := term.GetSize(0)
	if err != nil {
		return 0, 0
	}
	return w, h
}
