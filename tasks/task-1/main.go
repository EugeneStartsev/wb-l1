package main

import "fmt"

type Human struct {
	name   string
	weight int
	growth int
}

type Action struct {
	Human
}

func main() {
	action := *new(Action)

	action.growth = 75
	action.weight = 189
	action.name = "Jogn"

	fmt.Println(action)
}
