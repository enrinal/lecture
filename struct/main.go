package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
	Animal
	Barking bool
}

type Cat struct {
	Animal
	Meowing bool
}

func (d *Dog) Speak() {
	if d.Barking {
		println("Woof!")
	} else {
		println("Quack!")
	}
}

func (c *Cat) Speak() {
	if c.Meowing {
		println("Meow!")
	} else {
		println("Mew!")
	}
}

func (a *Animal) ChangeName(newName string) {
	a.Name = newName
}

func (a *Animal) Talk() {
	println("I'm an animal.")
}

func ChangeAge(a *Dog, newAge int) {
	a.Age = newAge
}

func main() {
	d := Dog{Animal{"Dog", 5}, true}
	c := Cat{Animal{"Cat", 2}, false}

	d.Speak()
	c.Speak()

	d.ChangeName("Doggo")
	fmt.Println(d)
	d.Talk()

	c.ChangeName("Kitty")
	fmt.Println(c)
	c.Talk()

	ChangeAge(&d, 10)
	fmt.Println(d)
}
