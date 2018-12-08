package main

import (
	"fmt"
	"math"
)

var times [26]int

// get the minimum (other than zero)
func getMinTime() int {
	minTime := math.MaxInt32
	for i:= 0; i <= maxNode; i++ {
		if (times[i] < minTime && times[i] != 0) {
			minTime = times[i]
		}
	}
	return minTime
}

func workersStillBusy() bool {
	for i:= 0; i <= maxNode; i++ {
		if times[i] > 0 {
			return true
		}
	}
	return false
}

func main() {
	readInput();
	makeRootsAvailable();

	totaltime := 0
	freeworkers := 5
	extratime := 60

	for len(available) > 0  || workersStillBusy() {
		// first, find all availables and start them
		for freeworkers > 0 && len(available) > 0{
			head := popAvailable()
			times[head] = head + 1 + extratime
			freeworkers--
		}

		// --- finish the next batch
		mintime := getMinTime()

		// mark as done
		for i := 0; i <= maxNode; i++ { 
			if (times[i] == mintime) {
				done[i] = true
				fmt.Printf("%c",i + offset)
				freeworkers++
			}
		}

		// add new availables
		for i := 0; i <= maxNode; i++ { 
			if (times[i] == mintime) {
				for j := 0; j <= maxNode; j++ {
					if !done[j] && goesTo(i,j) && allParentsDone(j) {
						makeAvailable(j)
					}
				}
			}
		}
		// advance clock...
		for i := 0; i <= maxNode; i++ { 
			if (times[i] > 0) {
				times[i] -= mintime
			}
		}
		totaltime += mintime
	}
	fmt.Printf(" %d\n",totaltime)
}
