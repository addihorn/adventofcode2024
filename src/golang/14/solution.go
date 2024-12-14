package main

import (
	"bytes"
	"example/hello/src/golang/aocutils"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"regexp"
	"strconv"

	"github.com/gosuri/uilive"
	"github.com/icza/mjpeg"
)

type robot struct {
	x, y             int
	dx, dy           int
	x_final, y_final int
	R, G, B, A       uint8
}

type gameData struct {
	width, height int
	rounds        int
}

const debug = false
const printPictures = false
const inputFile = "input.txt"

var game = gameData{width: 101, height: 103, rounds: 100}

var inputData []string
var allRobots []*robot

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	solve()

	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	allRobots = make([]*robot, len(inputData))
	r, _ := regexp.Compile(`p=(.+?),(.+?) v=(.+?),(.+)`)
	for i, robotData := range inputData {
		matches := r.FindAllStringSubmatch(robotData, -1)[0]
		robotEntity := robot{}
		robotEntity.x, _ = strconv.Atoi(matches[1])
		robotEntity.y, _ = strconv.Atoi(matches[2])
		robotEntity.dx, _ = strconv.Atoi(matches[3])
		robotEntity.dy, _ = strconv.Atoi(matches[4])

		robotEntity.R = uint8(rand.Intn(255))
		robotEntity.G = uint8(rand.Intn(255))
		robotEntity.B = uint8(rand.Intn(255))
		robotEntity.B = uint8(rand.Intn(255))

		allRobots[i] = &robotEntity
	}
}

func solve() {
	for _, robotEntity := range allRobots {
		newX := (robotEntity.x + (robotEntity.dx * game.rounds)) % game.width
		newY := (robotEntity.y + (robotEntity.dy * game.rounds)) % game.height

		switch true {
		case newX < 0:
			robotEntity.x_final = game.width + newX
		default:
			robotEntity.x_final = newX
		}

		switch true {
		case newY < 0:
			robotEntity.y_final = game.height + newY
		default:
			robotEntity.y_final = newY
		}

	}
}

/* Solve here */

func part1() {
	var q1, q2, q3, q4 int
	for _, robotEntity := range allRobots {

		x := robotEntity.x_final
		y := robotEntity.y_final

		switch true {
		case x < game.width/2 && y < game.height/2:
			q1++
		case x > game.width/2 && y < game.height/2:
			q2++
		case x < game.width/2 && y > game.height/2:
			q3++
		case x > game.width/2 && y > game.height/2:
			q4++
		}

	}

	_ = inputData
	fmt.Printf("Solution for part 1: %d\n", q1*q2*q3*q4)
}

func part2() {
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	var canvas *image.Paletted
	images := []*image.Paletted{}
	delays := []int{}
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	writer := uilive.New()
	framerate := int32(24)
	aw, err := mjpeg.New("test.avi", int32(game.width), int32(game.height), framerate)
	checkErr(err)
	writer.Start()

	for i := 7500; i < 8000; i++ {
		game.rounds = i
		solve()

		canvas = image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{game.width, game.height}}, palette)
		for _, robotEntity := range allRobots {
			canvas.Set(
				robotEntity.x_final, robotEntity.y_final,
				color.RGBA{R: robotEntity.R, G: robotEntity.G, B: robotEntity.B, A: robotEntity.A},
			)
		}
		images = append(images, canvas)
		delays = append(delays, 0)
		fmt.Fprintf(writer, "Generated images : %d/%d \n", i, 9999)
		buf := new(bytes.Buffer)

		checkErr(jpeg.Encode(buf, canvas, nil))
		checkErr(aw.AddFrame(buf.Bytes()))

		if printPictures {
			file, _ := os.Create(fmt.Sprintf("images/%d.png", i))
			png.Encode(file, canvas)
		}

		if i == 7774 {
			for i := 0; i < int(framerate*3); i++ {
				images = append(images, canvas)
				delays = append(delays, 1)

				buf := new(bytes.Buffer)
				checkErr(jpeg.Encode(buf, canvas, nil))
				checkErr(aw.AddFrame(buf.Bytes()))
			}

		}

	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays, LoopCount: 3,
	})

	checkErr(aw.Close())
	writer.Stop()
	//solved by manually looking through all 9999 pictures
	fmt.Printf("Solution for part 2: %d\n", 0)
}
