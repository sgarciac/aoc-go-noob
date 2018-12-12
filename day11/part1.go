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

func threeByThreeVal(x, y int) int {
	total := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += values[x + i][y + j]
		}
	}
	return total
}

func maxThreeByThree() (int, int){
	max := math.MinInt32
	topx := -1
	topy := -1
	for x := 0; x < 298; x++ {
		for y := 0; y < 298; y++ {
			currentVal := threeByThreeVal(x, y)
			if currentVal > max {
				max = currentVal
				topx = x
				topy = y
			}
		}
	}
	return topx, topy
}

func main(){
	loadValues()
	x,y := maxThreeByThree()
	fmt.Printf("%d,%d\n", x, y)
}
