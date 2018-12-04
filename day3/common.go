package main

func contains(x int, y int, claim claim) bool{
	return (x >= claim.topx) && (y >= claim.topy) && (x < claim.topx + claim.width) && (y < claim.topy + claim.height);
}

func collision(x int, y int, claims []claim) bool{
	found := 0;
	for _, claim := range claims {
		if (contains(x, y, claim)){
			found++;
		}
		if(found == 2){
			return true;
		}
	}
	return false;
}

func overlap(claim1 claim, claim2 claim) bool {
	return !((claim1.topx > (claim2.topx + claim2.width - 1)) ||
		((claim1.topx + claim1.width - 1) < claim2.topx) ||
		(claim1.topy > (claim2.topy + claim2.height - 1)) ||
		((claim1.topy + claim1.height - 1) < claim2.topy));
}
