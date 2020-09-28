package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Arguments:", os.Args)
	fmt.Println("Env:", os.Environ())
	os.Exit(69)
}
