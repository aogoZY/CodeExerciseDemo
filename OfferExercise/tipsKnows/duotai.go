package main

import "fmt"

func main() {

}

type AnimalIf struct {
	Age() int
	Run()
	Eat()
}

type Animal struct {
	Age int
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

//cat
func (d *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (d *Cat) Run() {
	fmt.Println("cat is running")
}

func (d *Cat) Eat() {
	fmt.Println("cat is sleeping")
}

//cat
func (d *Dog) Sleep() {
	fmt.Println("cat is sleeping")
}

func (d *Dog) Run() {
	fmt.Println("cat is running")
}

func (d *Dog) Eat() {
	fmt.Println("cat is sleeping")
}
