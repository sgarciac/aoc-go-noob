package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"unicode"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	var runes []rune;
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			if c != '\n' {
				if(len(runes) == 0){
					runes = append(runes, c);
				} else {
					last := runes[len(runes) - 1]
					sameLetter := unicode.ToLower(c) == unicode.ToLower(last)
					if (sameLetter &&
						((unicode.IsLower(c) && unicode.IsUpper(last)) ||
						 ((unicode.IsLower(last) && (unicode.IsUpper(c)))))) {
						 runes = runes[:len(runes) - 1]
					 } else {
						 runes = append(runes, c);
					 }
				}
			}
		}
	}
	fmt.Printf("%d",len(runes));
}

