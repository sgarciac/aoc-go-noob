package main
import (
	"os"
	"regexp"
	"bufio"
	"strconv"
	"fmt"
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

// areas[i][j] = index of point closer to this, if any that is unique
// or -1 otherwise
var areas [1000][1000] int

// sizes[i] = size of area (inside the inner rectangle) of points[i]
var sizes [100] int;

var top = math.MaxInt32
var bottom = math.MinInt32
var left = math.MaxInt32
var right = math.MinInt32


func manhattan(p point, x, y int) int{
	return abs(x - p.x) + abs(y - p.y)
}

func calculateClosest(i,j int) int {
	closestPoint := -1
	closestDistance := math.MaxInt32
	unique := true
	for index := 0; index < pointsCount; index++ {
		d := manhattan(points[index], i, j)
		if d < closestDistance {
			closestDistance = d
			closestPoint = index
			unique = true
		} else if d == closestDistance {
			unique = false
		}
	}
	if (unique) {
		return closestPoint
	} else {
		return -1
	}
}

// does a point have an infinite area?
// (it touches the border of the inner rectangle)
func isInfinite(pindex int) bool {
	for i := left; i <= right; i++ {
		if areas[i][top] == pindex || areas[i][bottom] == pindex {
			return true;
		}
	}

	for i := top; i<= bottom; i++ {
		if areas[left][i] == pindex || areas[right][i] == pindex {
			return true;
		}
	}

	return false
}


// calculate areas only for the inner rectangle
func calculateAreas(){
	for i := top; i <= bottom; i++ {
		for j :=  left; j <= right; j++ {
			areas[i][j] = calculateClosest(i,j);
		}
	}
}

func calculateSizes(){
	for i := top; i <= bottom; i++ {
		for j :=  left; j <= right; j++ {
			if areas[i][j] != -1{
				sizes[areas[i][j]]++
			}
		}
	}
}

func printAreas(){
	for i := top; i <= bottom; i++ {
		for j :=  left; j <= right; j++ {
			fmt.Printf("[%d]",areas[j][i])
		}
		fmt.Printf("\n")
	}
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

func main(){
	readPoints()
	calculateAreas()
	calculateSizes()
	maxAreaSize := -1
	for index := 0; index < pointsCount; index++ {
		if(!isInfinite(index)){
			if(sizes[index] > maxAreaSize){
				maxAreaSize = sizes[index]
			}
		}
	}
	fmt.Printf("%d\n", maxAreaSize)
}
