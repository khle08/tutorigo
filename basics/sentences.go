package basics

import "fmt"
// import "reflect"


func IfElse() {
	fmt.Println("\n2-1. if ... else ...")

	var age uint8
	fmt.Println("Please type in your age: ")
	fmt.Scanln(&age)

	if age < 13 {
		fmt.Println("Elementary school")
	} else if age >= 13 && age < 20 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}
}


func SwitchCase1() {
	fmt.Println("\n2-2-1. switch ... case ...")

	var age uint8
	fmt.Println("Please type in your age: ")
	fmt.Scanln(&age)

	// Same as "IfElse" function
	switch {
	case age < 13:
		fmt.Println("Elementary school")
		// Even though this case is triggered, jump to the next condition below
		fallthrough
	case age >= 13 && age < 20:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}


func SwitchCase2() {
	fmt.Println("\n2-2-2. switch ... case ...")

	var weekday uint8
	fmt.Println("What is this week: ")
	fmt.Scanln(&weekday)

	// Same as "IfElse" function
	switch weekday{
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Wed, Thu, Fri, Sat, Sun")
	}
}


func ForLoop() {
	fmt.Println("\n2-3. for loop")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 10 {
			fmt.Printf("i=%v\n", i)
			break
		}
	}

	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()    // make a new line

	for k := 1; k < 11; k++ {
		fmt.Print(k, "\t")
	}
	fmt.Println()    // make a new line
}


