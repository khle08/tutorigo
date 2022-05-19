package funcs

import "fmt"


type Car struct {
	gear int8
	door int8
	wheel int8
	name string
}

func (c Car) drive() {
	fmt.Println("\n6-1. method")
	c.gear = 1    // this is "copy" without sharing memory
	fmt.Println(c)
}

func (c *Car) back() {
	c.gear = 6
	fmt.Println(*c)
}


func Method() {
	car := Car {
		gear: 0,
		door: 4,
		wheel: 3,
		name: "this is a three wheel car",
	}
	car.drive()
	fmt.Println("car gear:", car.gear)
	car.back()
	fmt.Println("car gear:", car.gear)
}