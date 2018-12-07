package main

import (
	"os"
	"regexp"
	"bufio"
	"strconv"
	"math"
)

var validPoint = regexp.MustCompile(`^(\d+), (\d+)$`);

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
    if x < y {
	return y
    }
    return x
}

func min(x, y int) int {
    if x > y {
	return y
    }
    return x
}

// Global vars
var pointsCount = 0;
var points ([100]point)

var top = math.MaxInt32
var bottom = math.MinInt32
var left = math.MaxInt32
var right = math.MinInt32

func manhattan(p point, x, y int) int{
	return abs(x - p.x) + abs(y - p.y)
}

func readPoints() {
	scanner := bufio.NewScanner(os.Stdin)
	index := 0
	for scanner.Scan() {
		str := scanner.Text()
		if(len(str) > 0){
			match := validPoint.FindStringSubmatch(str)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			left = min(left, x)
			right = max(right, x)
			top = min(top, y)
			bottom = max(bottom, y)
			points[index] = point{x,y}
			index++
		}
	}
	pointsCount = index;
}
