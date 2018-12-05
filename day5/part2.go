package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
	"unicode"
	"strings"
)



func reduction(target string) int {
	r := strings.NewReader(target)
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
	return len(runes);
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan();
	text := scanner.Text()
	letters := "abcdefghijklmnopqrstuvwxyz"
	minimumResult := reduction(text)
	for _, char := range letters {
		newTarget := strings.Replace(text, string(char), "", -1)
		newTarget = strings.Replace(newTarget, string(unicode.ToUpper(char)), "", -1)
		newResult := reduction(newTarget);
		if(newResult < minimumResult){
			minimumResult = newResult
		}
	}
	fmt.Printf("%d\n",minimumResult)
}

