package main

import (
	"fmt"
)

type Hound interface {
	Hunt()
}

type Poodle interface {
	Bark() string
}

type Decorative interface {
	Sleep(hours int)
}

type Bloodhound struct {
	name string
}

func (b Bloodhound) Hunt() {
	fmt.Printf("%s is hunting...\n", b.name)
}

type TeacupPoodle struct {
	name string
}

func (p TeacupPoodle) Bark() string {
	return fmt.Sprintf("%s is barking...\n", p.name)
}

type ShihTzu struct {
	name string
}

func (p ShihTzu) Sleep(horts int) {
	fmt.Printf("%s is sleeping...\n", p.name)
}

func zoo(dog interface{}) {
	switch dog.(type) {
	case Hound:
		dog.(Hound).Hunt()
	case Poodle:
		fmt.Print(dog.(Poodle).Bark())
	case Decorative:
		dog.(Decorative).Sleep(1)
	default:
		fmt.Println("Unknown dog.")
	}
}

func main() {
	zoo(Bloodhound{"misha"})
	zoo(TeacupPoodle{"vasya"})
	zoo(ShihTzu{"kakulya"})
}
