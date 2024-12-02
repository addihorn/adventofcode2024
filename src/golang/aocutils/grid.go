package aocutils

import (
	"fmt"
	"math"
	"time"

	"github.com/inancgumus/screen"
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
	screen.Clear()
	output := ""
	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				output = output + string(val)
			} else {
				output = output + " "
			}

		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 10)

}
