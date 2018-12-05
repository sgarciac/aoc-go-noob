package main

func calculateIntervals(log []logentry) (map[int][][]int){
	times := make(map[int][][]int);
	currentGuard := -1;
	currentMinute := -1
	for _, logentry := range log {
		switch logentry.logtype {
		case awake:
			newMinute := logentry.date.Minute()
			times[currentGuard] = append(times[currentGuard],[]int{currentMinute, newMinute});
			currentMinute = newMinute
		case asleep:
			currentMinute = logentry.date.Minute()
		case shift:
			currentGuard = logentry.guardId
			if(logentry.date.Hour() != 0) {
				currentMinute = 0;
			} else {
				currentMinute = logentry.date.Minute()
			}
		}
	}
	return times;
}
