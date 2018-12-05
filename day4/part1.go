package main

import (
	"fmt"
)

func main(){
	log := ReadLog();
	times := make(map[int]int);
	currentGuard := -1;
	currentMinute := -1
	currentState := -1
	maxGuard := -1;
	maxTime := -1
	for _, logentry := range log {
		switch logentry.logtype {
		case awake:
			newMinute := logentry.date.Minute()
			times[currentGuard] += (newMinute - currentMinute)
			currentMinute = newMinute
			currentState = awake
			if times[currentGuard] > maxTime {
				maxGuard = currentGuard
				maxTime = times[currentGuard]
			}
		case asleep:
			currentMinute = logentry.date.Minute()
			currentState = asleep
		case shift:
			if (currentGuard != -1){
				if currentState == asleep {
					times[currentGuard] += (60 - currentMinute)

				}
			}
			currentGuard = logentry.guardId
			if(logentry.date.Hour() != 0) {
				currentMinute = 0;
			} else {
				currentMinute = logentry.date.Minute()
			}
		}
	}

	fmt.Printf("max guard: %d hours = %d\n",maxGuard, maxTime);
}
