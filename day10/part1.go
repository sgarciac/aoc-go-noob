package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"math"
//	"fmt"
	"io"
	"image"
	"image/color"
	"image/gif"
)

var validEntry = regexp.MustCompile(`^position=< *(-?\d+), *(-?\d+)> *velocity=< *(-?\d+), *(-?\d+)>$`)

type point struct {
	posx int
	posy int
	velx int
	vely int
}

var points [1000]point
var pointCount int

func printSky(out io.Writer) {
	topleft, bottomright := calculateCorners();
	width := (bottomright.posx - topleft.posx) + 1
	height := (bottomright.posy - topleft.posy) + 1

	palette := []color.Color{color.White, color.Black}
	rect := image.Rect(0, 0,
		width,
		height)
	img := image.NewPaletted(rect, palette)
	for i := 0; i < pointCount; i++ {
		img.SetColorIndex(points[i].posx - topleft.posx, points[i].posy - topleft.posy, 1)
	}
	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}
	gif.EncodeAll(out, &anim)
}

func calculateArea() int {
	topleft, bottomright := calculateCorners();
	width := (bottomright.posx - topleft.posx) + 1
	height := (bottomright.posy - topleft.posy) + 1
	return width * height
}

func tic() {
	for i := 0; i < pointCount; i++ {
		points[i].posx += points[i].velx
		points[i].posy += points[i].vely
	}
}

func backtic() {
	for i := 0; i < pointCount; i++ {
		points[i].posx -= points[i].velx
		points[i].posy -= points[i].vely
	}
}


func calculateCorners() (point, point){
	var topleft, bottomright point
	topleft.posx = math.MaxInt32
	topleft.posy = math.MaxInt32
	bottomright.posx = math.MinInt32
	bottomright.posy = math.MinInt32

	for i := 0; i < pointCount; i++ {
		p := points[i]
		if p.posx < topleft.posx {
			topleft.posx = p.posx
		}
		if p.posy < topleft.posy {
			topleft.posy = p.posy
		}
		if p.posx > bottomright.posx {
			bottomright.posx = p.posx
		}
		if p.posy > bottomright.posy {
			bottomright.posy = p.posy
		}
	}
	return topleft, bottomright
}



func readPoints(){
	scanner := bufio.NewScanner(os.Stdin)
	pointCount = 0
	for scanner.Scan() {
		str := scanner.Text();
		if(len(str) > 0){
			match := validEntry.FindStringSubmatch(str);
			posx, _ := strconv.Atoi(match[1])
			posy, _ := strconv.Atoi(match[2])
			velx, _ := strconv.Atoi(match[3])
			vely, _ := strconv.Atoi(match[4])
			points[pointCount] = point{posx,posy,velx,vely}
			pointCount++
		}
	}
}

func main(){
	readPoints()
	currentArea := calculateArea()
	for true {
		tic()
		newArea := calculateArea();
		if (newArea >= currentArea){
			break
		}
		currentArea = newArea
	}
	backtic()
	printSky(os.Stdout)
}
