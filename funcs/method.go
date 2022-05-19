package funcs

import "fmt"
import "math"


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

// =====================================================================

type geometry interface {
	area() float32
	edge() float32
}

type rect struct {
	width, height float32
}

func (r rect) area() float32 {
	return r.width * r.height
}

func (r rect) edge() float32 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) edge() float32 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) {
	fmt.Printf("type: %T, obj: %v\n", g, g)
	fmt.Println("area:", g.area())
	fmt.Println("edge:", g.edge())
}

func Interface() {
	fmt.Println("\n6-2. interface")

	r := rect{width: 4, height: 5}
	Measure(r)

	c := circle{radius: 10}
	Measure(c)
}





