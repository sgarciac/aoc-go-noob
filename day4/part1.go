package main

import ("fmt");

func main(){
	log := ReadLog();
	for _, logentry := range log {
		fmt.Printf("%s\n", logentry.date);
	}
}
