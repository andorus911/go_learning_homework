package main

import (
	"fmt"
	"go_learning_homework/ch04_structs/json-xml-io/jsonio"
	"go_learning_homework/ch04_structs/json-xml-io/xmlio"
	"log"
	"os"
	"strings"
)

type tenNumbers struct {
	Numbers []int
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a json file name")
		return
	}

	filename := arguments[1]
	var myNumbers tenNumbers
	expansion := strings.Split(filename, ".")

	switch expansion[len(expansion)-1] {
	case "json":
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
	case "xml":
		if err := xmlio.LoadFromXml(filename, &myNumbers); err == nil {
			fmt.Println(myNumbers)
		} else {
			fmt.Println(err)
		}

		for i := range myNumbers.Numbers {
			myNumbers.Numbers[i]++
		}

		if err := xmlio.SaveToXml(os.Stdout, myNumbers); err != nil {
			fmt.Println(err)
		}
	default:
		log.Println("Unsupported file extension.")
	}
}
