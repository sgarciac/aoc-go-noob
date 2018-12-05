package main

import (
	"io"
	"bufio"
	"os"
	"fmt"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	for {
		if c, sz, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			fmt.Printf("%q [%d]\n", string(c), sz)
		}
	}
}

