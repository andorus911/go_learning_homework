package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	text := `CHART A COURSE IN THE WILDS
Venture into the wilds with this kit for the world's greatest roleplaying game.
This kit equips the Dungeon Master with a screen and other tools that are perfect for running D&D adventures in the wilderness.
The Dungeon Master’s screen features a gorgeous painting of fantasy landscapes on the outside, and useful rules references cover the inside of the screen, with an emphasis on wilderness rules. The kit also includes the following:
5 dry-erase sheets, featuring hex maps, a food-and-water tracker, and rules references (wilderness chases, wilderness journeys, and the actions you can take in combat)
27 cards that make it easy to keep track of conditions, initiative, and environmental effects
1 box to hold the kit's cards`
	dict := WordCounter(text)
	fmt.Println(Top(10, dict))
}

func WordCounter(text string) map[string]int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	ltext := strings.ToLower(text)
	arr := strings.FieldsFunc(ltext, f)
	dict := make(map[string]int)
	for _, v := range arr {
		dict[v]++
	}
	return dict
}

type Pair struct {
	Count int
	Word  string
}

func Top(topNumber int, dict map[string]int) []Pair {
	pairList := make([]Pair, topNumber)
	for word, count := range dict {
		for i := 0; i < topNumber; i++ {
			if count > pairList[i].Count {
				pairList[i].Count = count
				pairList[i].Word = word
				sort.Slice(pairList, func(i, j int) bool {
					return pairList[i].Count < pairList[j].Count
				})
				break
			}
		}
	}
	return pairList
}
