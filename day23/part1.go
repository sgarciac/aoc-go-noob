package main

import (
	"fmt"
	"regexp"
	"bufio"
	"strconv"
	"os"
)

type nanobot struct {
	x int
	y int
	z int
	r int
}

func intAbs(x, y int) int {
	diff := x - y
	if diff < 0 {
		return diff * -1
	} else {
		return diff
	}
}

func distance(n1, n2 nanobot) int {
	return intAbs(n1.x, n2.x) +
		intAbs(n1.y, n2.y) +
		intAbs(n1.z, n2.z)
}

func inRange(origin, sat nanobot) bool {
	inRange := distance(origin,sat) <= origin.r
	//fmt.Printf("%v <- %v : %v, %v\n", origin,sat, distance(origin,sat), inRange)
	return inRange
}

func nanobotEq(n1, n2 nanobot) bool {
	return n1.x == n2.x && n1.y == n2.y && n1.z == n1.z && n1.r == n2.r
}

func (this nanobot) String() string { return fmt.Sprintf("<%v,%v,%v> (%v)",this.x, this.y, this.z, this.r) }

var validEntry = regexp.MustCompile(`^pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)$`)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	maxRad := -1
	var maxNanobot nanobot
	var nanobots []nanobot
	// read nanobots, find max
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) > 0 {
			match := validEntry.FindStringSubmatch(str)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			z, _ := strconv.Atoi(match[3])
			r, _ := strconv.Atoi(match[4])
			n := nanobot{x,y,z,r}
			if r > maxRad {
				maxRad = r
				maxNanobot = n
			}
			nanobots = append(nanobots, n)
		}
	}
	//
	total := 0
	for _, n := range nanobots {
		if inRange(maxNanobot, n) {
			total++
		}
	}
	fmt.Println(total)
}
