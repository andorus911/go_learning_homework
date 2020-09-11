package main

import (
	"fmt"
	"go_learning_homework/ch04_structs/jsonio"
	"os"
)

type TenNumbers struct {
	Numbers [10]int
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a json file name")
		return
	}

	filename := arguments[1]

	var myNumbers TenNumbers
	if err := jsonio.LoadFromJson(filename, &myNumbers); err == nil {
		fmt.Println(myNumbers)
	} else {
		fmt.Println(err)
	}

	for i := range myNumbers.Numbers {
		myNumbers.Numbers[i]++
	}

	if err := jsonio.SaveToJson(os.Stdout, myNumbers); err != nil {
		fmt.Println(err)
	}
}
