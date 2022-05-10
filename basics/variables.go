package basics

import "fmt"


const (
	VERSION = 1.0
)


func Variable() {
	// there are 4 ways to define a variable.
	var v1 int
	v1 = 1

	var v2 int = 2
	var v3 = 3
	v4 := 4

	var (
		v5 = 5
		v6 int  // default value is 0
	)

	fmt.Println("\n1-1. Variable")
    fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v\n",
               v1, v2, v3, v4, v5, v6)
}


func Constant() {
	const (
		c1 = 8
		c2 = iota  // Fill the row index for the const name
		c3         // Add sequential value ...
		c4
		c5 = 2
		c6
	)

	fmt.Println("\n1-2. Constant")
    fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v\n",
               c1, c2, c3, c4, c5, c6)
}

