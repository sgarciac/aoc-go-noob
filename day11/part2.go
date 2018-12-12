package main

import (
	"fmt"
	"math"
)

var values [300][300]int;

var testInput = 6878

func powerLevel(x, y, i int) int {
	rackId := x + 10
	pl := rackId * y
	pl += i
	pl *= rackId
	pl = (pl / 100) % 10
	pl -= 5
	return pl
}

func loadValues(){
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			values[x][y] = powerLevel(x+1, y+1, testInput)
		}
	}
}

func lByLVal(x, y, l int) int {
	total := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			total += values[x + i][y + j]
		}
	}
	return total
}

func maxLByL(l int) (int, int, int){
	max := math.MinInt32
	topx := -1
	topy := -1
	for x := 0; x < 300 - (l - 1); x++ {
		for y := 0; y < 300 - (l - 1); y++ {
			currentVal := lByLVal(x, y, l)
			if currentVal > max {
				max = currentVal
				topx = x
				topy = y
			}
		}
	}
	return topx, topy, max
}

func realMax() (int, int, int){
	bigMax := math.MinInt32
	xMax := -1
	yMax := -1
	lMax := -1
	for l := 1; l <= 300; l++ {
		cx, cy, cmax := maxLByL(l)
		if cmax > bigMax {
			bigMax = cmax
			xMax = cx
			yMax = cy
			lMax = l
		}
	}
	return xMax, yMax, lMax
}

func main(){
	loadValues()
	x,y, l := realMax()
	fmt.Printf("%d,%d,%d\n", x, y, l)
}
