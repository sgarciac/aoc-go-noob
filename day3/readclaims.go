package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var validClaim = regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")


type claim struct {
	id int
	topx int
	topy int
	width int
	height int
}

func ReadClaims() ([]claim){
	scanner := bufio.NewScanner(os.Stdin)
	var claims []claim;
	for scanner.Scan() {
		str := scanner.Text();
		if(len(str) > 0){
			match := validClaim.FindStringSubmatch(str);
			id, _ := strconv.Atoi(match[1])
			topx, _ := strconv.Atoi(match[2])
			topy, _ := strconv.Atoi(match[3])
			width, _ := strconv.Atoi(match[4])
			height, _ := strconv.Atoi(match[5]) 
			claims = append(claims, claim{
				id,
				topx,
				topy,
				width,
				height})
		}
	}
	return claims;
}

