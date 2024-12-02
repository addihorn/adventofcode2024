package aocutils

import "math"

type Roomsize struct {
	MinX int
	MinY int
	MinZ int
	MaxX int
	MaxY int
	MaxZ int
}

func NewRoomSize() *Roomsize {
	return &Roomsize{math.MaxInt, math.MaxInt, math.MaxInt, math.MinInt, math.MinInt, math.MinInt}
}

func (this *Roomsize) RecalibrateTo(point [3]int) {

	this.MinX = Min([]int{this.MinX, point[0]})
	this.MinY = Min([]int{this.MinY, point[1]})
	this.MinZ = Min([]int{this.MinZ, point[2]})
	this.MaxX = Max([]int{this.MaxX, point[0]})
	this.MaxY = Max([]int{this.MaxY, point[1]})
	this.MaxZ = Max([]int{this.MaxZ, point[2]})

}
