package fmtvsstrconv

import (
	"fmt"
	"strconv"
)

func Slow() string {
	return fmt.Sprintf("%d", 42) // Reflection :(
}

func Fast() string {
	return strconv.Itoa(42)
}
