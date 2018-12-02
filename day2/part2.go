package main

import ("fmt");

// assuming same size and single byte runes
func closeEnough(source string, target string) int {
	differences := 0;
	diffpos := -1;
	for pos, _ := range source {
		if (target[pos] != source[pos]) {
			differences++;
			diffpos = pos;
		}
		if(differences>1){
			return -1;
		}
	}
	return diffpos;
}

func main(){
	lines := ReadLines();
	for pos, line := range lines {
		for _, line2 := range lines[pos +1 :] {
			diffpos := closeEnough(line, line2);
			if diffpos != -1 {
				fmt.Printf("%s%s\n", line[:diffpos], line[diffpos+1:]);
			} 
		}
	}
}
