package main

import "fmt"

type ChiCha interface {
	Pee(string)
	Eat(string)
}

type Dog interface {
	Pee(string)
	Eat(string)
}

type Haski struct {
	name   string
	weight int
}

func (h *Haski) Pee(amount string) {
	if h == nil {
		fmt.Println("Give your dog name! She is missery")
		return
	}
	fmt.Printf("%s is going to pee %s \n", h.name, amount)
}

func (h *Haski) Eat(food string) {
	if h == nil {
		fmt.Println("Give your dog some food! Sie hat hunger")
		return
	}
	fmt.Printf("%s is going to eat %s weight of %s is %d \n", h.name, food, h.name, h.weight)
}

// Here if we set Haski without pointer
// it would cause an error because
// methods on Haski type are defined only on
// *Haski
// error -> is interface
// here instead of only on type we can use all types which
// are fulfill interface methods
// func(a int) (int, error)
// here we expect to have type which returned from the function
// to have all methods in interface if it's not we would get error or panic

var a Dog

func main() {
	a = &Haski{name: "Chapi", weight: 18}
	var b *Haski
	// type assertion to check the type
	f, ok := a.(*Haski)
	fmt.Println("Type assertion: Value and Status", f, ok)

	a.Pee("much")
	a.Eat("much")

	a = b
	a.Pee("much")
	a.Eat("much")
}
