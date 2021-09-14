package main

import "fmt"

type Animals struct {
	name string
}

func (a *Animals) move() {
	fmt.Println("animal name")
}

type Dogs struct {
	feet string
	*Animals
}

func (d *Dogs) say() {
	fmt.Println("dog is ...")
}

func main() {
	d := Dogs{
		feet: "aa",
		Animals: &Animals{
			name: "jiao",
		},
	}
	d.say()
	d.move()
}
