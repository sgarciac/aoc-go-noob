package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"fmt"
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

func printPoint(){
	
}

func readPoints(){
	scanner := bufio.NewScanner(os.Stdin)
	pointsCount := 0
	for scanner.Scan() {
		str := scanner.Text();
		if(len(str) > 0){
			fmt.Println(str)
			match := validEntry.FindStringSubmatch(str);
			posx, _ := strconv.Atoi(match[1])
			posy, _ := strconv.Atoi(match[2])
			velx, _ := strconv.Atoi(match[3])
			vely, _ := strconv.Atoi(match[4])
			points[pointsCount] = point{posx,posy,velx,vely}
			pointsCount++
		}
	}
}

func main(){
	readPoints();
}
