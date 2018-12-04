package main

import (
	"bufio"
	"os"
	"regexp"
)

var validLogEntry = regexp.MustCompile(`^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] (.*)$`)

type logentry struct {
	date string
	content string
}

func ReadLog() ([]logentry){
	scanner := bufio.NewScanner(os.Stdin)
	var log []logentry
	for scanner.Scan() {
		str := scanner.Text()
		if(len(str) > 0){
			match := validLogEntry.FindStringSubmatch(str)
			date := match[1]
			content := match[2]
			log = append(log, logentry{
				date,
				content})
		}
	}
	return log
}
