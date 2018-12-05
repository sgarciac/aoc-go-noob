package main

import (
	"fmt"
)

func main(){
	log := ReadLog();
	guardsIntervals := calculateIntervals(log);
	maxGuard := -1;
	maxTime := -1;
	for guard, intervals := range guardsIntervals {
		total := 0
		for _, interval := range intervals {
			total += interval[1] - interval[0]
		}
		if total > maxTime {
			maxTime = total;
			maxGuard = guard;
		}
	}
	fmt.Printf("Most lazy guard: %d (%dm)\n", maxGuard, maxTime);
	maxMinute := -1
	maxMinuteHits := -1;
	for i := 0; i < 59; i++ {
		hits := 0
		for _, interval := range guardsIntervals[maxGuard] {
			if(i >= interval[0] && i < interval[1]){
				hits++;
			}
		}
		if hits > maxMinuteHits {
			maxMinuteHits = hits;
			maxMinute = i
		}
	}
	fmt.Printf("Most lazy minute: %d (%d hits)\n", maxMinute, maxMinuteHits);
	fmt.Printf("%d x %d = %d\n",maxGuard, maxMinute, maxGuard * maxMinute);
}
