package main

import (
	"bufio"
	"os"
	"regexp"
	"time"
	"sort"
	"strconv"
)

const (
	asleep = iota
	awake
	shift
)

var validLogEntry = regexp.MustCompile(`^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] (.*)$`)
var validBeginShift = regexp.MustCompile(`^Guard #(\d+) begins shift$`);


type logentry struct {
	date time.Time
	logtype int
	guardId int
}

// returns the log, sorted
func ReadLog() ([]logentry){
	scanner := bufio.NewScanner(os.Stdin)
	var log []logentry
	for scanner.Scan() {
		str := scanner.Text()
		if(len(str) > 0){
			match := validLogEntry.FindStringSubmatch(str)
			date, _ := time.Parse("2006-01-02 15:04", match[1])
			// parse logtype
			guardId := 0
			logtype := -1
			if match[2] == "wakes up" {
				logtype = awake
			} else if match[2] == "falls asleep" {
				logtype = asleep
			} else {
				logtype = shift
				matchShift := validBeginShift.FindStringSubmatch(match[2]);
				guardId, _ = strconv.Atoi(matchShift[1]);
			}

			log = append(log, logentry{
				date,
				logtype,
				guardId})
		}
	}
	sort.Slice(log,func(i, j int) bool {
		return log[i].date.Before(log[j].date);
	});
	return log;
}
