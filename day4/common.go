package main

func calculateIntervals(log []logentry) (map[int][][]int){
	times := make(map[int][][]int);
	currentGuard := -1;
	currentMinute := -1
	currentState := -1
	for _, logentry := range log {
		switch logentry.logtype {
		case awake:
			newMinute := logentry.date.Minute()
			times[currentGuard] = append(times[currentGuard],[]int{currentMinute, newMinute});
			currentMinute = newMinute
			currentState = awake
		case asleep:
			currentMinute = logentry.date.Minute()
			currentState = asleep
		case shift:
			// not sure if this is necessary, but
			// better be safe than sorry.
			if (currentGuard != -1){
				if currentState == asleep {
					newMinute := logentry.date.Minute()
					times[currentGuard] = append(times[currentGuard],[]int{currentMinute, newMinute});

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
	return times;
}
