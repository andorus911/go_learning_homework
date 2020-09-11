package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TenNumbers struct {
	Numbers [10]int
}

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	return nil
}

func saveToJson(output *os.File, key interface{}) error {
	encodeJSON := json.NewEncoder(output)
	err := encodeJSON.Encode(key)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a json file name")
		return
	}

	filename := arguments[1]

	var myNumbers TenNumbers
	if err := loadFromJson(filename, &myNumbers); err == nil {
		fmt.Println(myNumbers)
	} else {
		fmt.Println(err)
	}

	for i := range myNumbers.Numbers {
		myNumbers.Numbers[i]++
	}

	if err := saveToJson(os.Stdout, myNumbers); err != nil {
		fmt.Println(err)
	}
}
