package aocutils

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func TBprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
