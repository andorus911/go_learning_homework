package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]

	for i, v := range str {

		if v <= '9' && v >= '1' {
			if i == 0 {
				fmt.Println("Incorrect string.")
				os.Exit(1)
			}

			for j := 0; j < (int(v) - '0' - 1); j++ {

				fmt.Printf("%s", string(str[i-1]))
			}
		} else {
			fmt.Printf("%s", string(v))
		}
	}
	fmt.Println()
}
