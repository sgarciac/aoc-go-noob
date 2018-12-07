package main

import (
	"fmt"
)

func main(){
	readPoints()
	regionSize := 0;
	for i := top; i <= bottom; i++ {
		for j :=  left; j <= right; j++ {
			total := 0
			for index := 0; index < pointsCount; index++ {
				total += manhattan(points[index],i,j)
			}
			if total < 10000 {
				regionSize++;
			}
		}
	}
	fmt.Printf("%d\n",regionSize)
}
