package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Args should be in 'quotes' for Ubuntu and without for Windows
	str := UnpackString(os.Args[1])

	fmt.Println(str)
}

func UnpackString(str string) string {

	if unicode.IsDigit(rune(str[0])) {

		log.Println("Incorrect string.")
		return ""
	}

	var unpackedString strings.Builder
	var lastRune rune

	inEscape := false

	for _, v := range str {

		if inEscape {

			fmt.Fprintf(&unpackedString, "%v", string(v))
			lastRune = v
			inEscape = false
			continue
		}

		if v == '\\' {

			inEscape = true
			continue
		}

		if unicode.IsDigit(v) {

			for j := 0; j < (int(v) - '0' - 1); j++ {

				fmt.Fprintf(&unpackedString, "%v", string(lastRune))
			}
		} else {

			fmt.Fprintf(&unpackedString, "%v", string(v))
			lastRune = v
		}
	}
	return unpackedString.String()
}
