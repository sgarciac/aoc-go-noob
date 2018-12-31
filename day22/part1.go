package main

import (
	"fmt"
)

var depth = 10689
var targetX = 11
var targetY = 722

//var depth = 510
//var targetX = 10
//var targetY = 10

var gindexCache = make(map[int]int)
var erosionCache = make(map[int]int)

func pointId(x, y int) int {
	return (y * (targetX + 1)) + x
}

func gIndex(x, y int) int{
	pointId := pointId(x, y)

	val, ok := gindexCache[pointId]

	if ok {
		return val
	}

	// mouth or target
	if (x == 0 && y == 0) || (x == targetX && y == targetY) {
		val = 0
		gindexCache[pointId] = val
		return val
	}

	if y == 0 {
		val = x * 16807
		gindexCache[pointId] = val
		return val
	}

	if x == 0 {
		val = y * 48271
		gindexCache[pointId] = val
		return val
	}

	val = erosionLevel(x-1,y) * erosionLevel(x,y-1)
	gindexCache[pointId] = val
	return val
}

func erosionLevel(x, y int) int {
	pointId := pointId(x, y)
	val, ok := erosionCache[pointId]

	if ok {
		return val
	}

	val = ((gIndex(x, y) + depth) % 20183)
	erosionCache[pointId] = val
	return val
}

func regType(x, y int) int {
	return erosionLevel(x,y) % 3
}

func main (){
	total := 0
	for row := 0; row <= targetY; row++ {
		for col := 0; col <= targetX; col++ {
			total += regType(col, row)
		}
	}
	fmt.Println(total)
}
