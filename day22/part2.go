package main

import (
	"fmt"
	"github.com/yourbasic/graph"
)

var depth = 10689
var targetX = 11
var targetY = 722

//var depth = 510
//var targetX = 10
//var targetY = 10

var maxX = targetX * 3
var maxY = targetY * 3

// Z = equipment, 0 = torch, 1 = climbing, 2 = nothing


var gindexCache = make(map[int]int)
var erosionCache = make(map[int]int)

var paths = graph.New(maxX * maxY * 3)

func pointId(x, y, z int) int {
	return (z * maxX * maxY) + (y * maxX) + x
}

func idToPoint(id int) (int,int,int){
	z := id / (maxX * maxY)
	y := (id - (z * maxX * maxY)) / maxX
	x := id % maxX
	return x,y,z
}

func gIndex(x, y int) int{
	pointId := pointId(x, y, 0)

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
	pointId := pointId(x, y, 0)
	val, ok := erosionCache[pointId]

	if ok {
		return val
	}

	val = ((gIndex(x, y) + depth) % 20183)
	erosionCache[pointId] = val
	return val
}

// 0 rocky, 1 wet, 2 narrow
func regType(x, y int) int {
	return erosionLevel(x,y) % 3
}

func printAddCost(col, row, eq, col2, row2, eq2 int) {
	if false {
		fmt.Printf("(%v)<%v> %v,%v,%v -> (%v)<%v> %v,%v,%v\n",
			pointId(col, row, eq), regType(col,row), row, col, eq,
			pointId(col2, row2, eq2), regType(col2,row2), row2, col2, eq2)

	}
}

func validEquipment(col, row, eq int) bool {
	valid := (regType(col, row) == 0 && eq != 2) ||
		(regType(col, row) == 1 && eq != 0) ||
		(regType(col, row) == 2 && eq != 1)
	return valid
}

func createPaths(col, row, eq int) {
	// if not a valid node, return immediately
	if !validEquipment(col, row, eq) {
		return
	}

	// move down?
	targetRow := row + 1
	targetCol := col
	if (validEquipment(targetCol, targetRow, eq)) {
		printAddCost(col,row,eq,targetCol,targetRow,eq)
		paths.AddBothCost(pointId(col, row, eq), pointId(targetCol, targetRow, eq), 1)
	}

	// move right?
	targetRow = row
	targetCol = col + 1
	if (validEquipment(targetCol, targetRow, eq)) {
		printAddCost(col,row,eq,targetCol,targetRow, eq)
		paths.AddBothCost(pointId(col, row, eq), pointId(targetCol, targetRow, eq), 1)
	}

	// change equipment
	for i := 0; i < 3; i++ {
		if (i != eq && validEquipment(col, row, i)) {
			printAddCost(col,row,eq,col,row, i)
			paths.AddBothCost(pointId(col, row, eq), pointId(col, row, i), 7)
		}
	}
}

func main (){
	for equipment := 0; equipment < 3; equipment++ {
		for row := 0; row <= maxY - 2; row++ {
			for col := 0; col <= maxX - 2; col++ {
				createPaths(col, row, equipment)
			}
		}
	}

	_, distance := graph.ShortestPath(
		paths,
		pointId(0, 0, 0),
		pointId(targetX, targetY, 0))

	fmt.Println(distance)
}
