package aocutils

import (
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/inancgumus/screen"
	"github.com/nsf/termbox-go"
)

type Gridsize struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func NewGridSize() *Gridsize {
	return &Gridsize{math.MaxInt, math.MaxInt, math.MinInt, math.MinInt}
}

func (this *Gridsize) RecalibrateTo(point [2]int) {

	this.MinX = Min([]int{this.MinX, point[0]})
	this.MinY = Min([]int{this.MinY, point[1]})
	this.MaxX = Max([]int{this.MaxX, point[0]})
	this.MaxY = Max([]int{this.MaxY, point[1]})
}

func (this *Gridsize) PaintGrid(grid map[[2]int]rune) {

	fmt.Println(reflect.TypeOf(grid).String())
	screen.Clear()
	output := ""
	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				output = output + string(val)
			} else {
				output = output + "."
			}

		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 10)

}

func Paint[K ~[2]int, V ~rune](grid map[K]V, this *Gridsize) {

	if !termbox.IsInit {
		screen.Clear()
		output := ""
		for y := this.MinY; y <= this.MaxY; y++ {
			for x := this.MinX; x <= this.MaxX; x++ {
				if val, ok := grid[[2]int{x, y}]; ok {
					output = output + string(val)
				} else {
					output = output + "."
				}

			}
			output = output + "\n"
		}
		fmt.Println(output)
		time.Sleep(time.Millisecond * 10)
		return
	}

	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {

				fg_color := termbox.ColorLightGreen

				switch val {
				case '.':
					val = ' '
				case '#':
					fg_color = termbox.ColorLightRed
				case '@':
					fg_color = termbox.ColorDefault
				}
				termbox.SetCell(x, y, rune(val), fg_color, termbox.ColorDefault)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorLightGreen, termbox.ColorDefault)
			}

		}
	}
	termbox.Flush()
	time.Sleep(time.Millisecond)

}
