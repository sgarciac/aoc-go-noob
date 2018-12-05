package main

import (
	"fmt"
)


func main(){
	log := ReadLog();
	guardsIntervals := calculateIntervals(log);
	maxGuard := -1;
	absoluteMaxMinuteHits := -1;
	absoluteMaxMinute := -1;
	for guard, intervals := range guardsIntervals {
		maxMinuteHits := -1
		maxMinute := -1
		for i := 0; i < 59; i++ {
			hits := 0
			for _, interval := range intervals {
				if(i >= interval[0] && i < interval[1]){
					hits++;
				}
			}
			if hits > maxMinuteHits {
				maxMinuteHits = hits;
				maxMinute = i
			}
		}
		if (maxMinuteHits > absoluteMaxMinuteHits){
			absoluteMaxMinuteHits = maxMinuteHits
			absoluteMaxMinute = maxMinute
			maxGuard = guard
		}
	}
	fmt.Printf("Guard: %d minute:%d hits:%d. Result = %d\n", maxGuard, absoluteMaxMinute, absoluteMaxMinuteHits, maxGuard * absoluteMaxMinute);

}
