package main

import ("fmt");

func main(){
	claims := ReadClaims();
	total := 0;
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if (collision(i,j, claims)) {
				total++;
			}
		}
	}
	fmt.Printf("%d\n", total)
}
