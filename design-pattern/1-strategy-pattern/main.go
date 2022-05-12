package main

import "fmt"

type Dog interface {
	Bark()
	Walk()
}

type DogImpl struct {
}

func (d *DogImpl) Bark() {
	fmt.Println("Woof")
}

func (d *DogImpl) Walk() {
	fmt.Println("Walking")
}

type Bulldog struct {
	Dog
}

type Flyable interface {
	Fly()
}

type FlyableImpl struct{}

func (f *FlyableImpl) Fly() {
	fmt.Println("Flying")
}

type RobotDog struct {
	Dog
	Flyable
}

func main() {
	dog := new(DogImpl)
	bulldog := new(Bulldog)

	bulldog.Dog = dog
	bulldog.Bark()
	bulldog.Walk()

	fly := new(FlyableImpl)
	robotDog := new(RobotDog)

	robotDog.Dog = dog
	robotDog.Flyable = fly
	robotDog.Bark()
	robotDog.Walk()
	robotDog.Fly()
}
