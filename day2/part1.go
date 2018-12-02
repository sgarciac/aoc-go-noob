package main

import ("fmt");

func hasRepetitions(source string, target int) bool {
	counts := make(map[rune]int);
	for _, char := range source {
		counts[char]++;
	}
	for _, char := range source {
		if(counts[char] == target){
			return true;
		}
	}
	return false;
}

func main(){
	lines := ReadLines();
	twos := 0;
	threes := 0;
	for _, line := range lines {
		if (hasRepetitions(line, 2)) {
			twos++;
		}
		if(hasRepetitions(line, 3)){
			threes++;
		}
	}

	fmt.Printf("%d\n",twos * threes);
}
