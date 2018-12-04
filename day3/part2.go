package main

import ("fmt");


func main(){
	claims := ReadClaims();

	for _, claim := range claims {
		overlapped := false
		for _, claim2 := range claims {
			if ((claim != claim2) && overlap(claim, claim2)) {
				overlapped = true;
			}
		}
		if (!overlapped){
			fmt.Printf("%d\n", claim.id)
			return;
		}
	}
}
