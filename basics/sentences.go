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


func Label2Goto() {
	fmt.Println("\n2-4. label and goto")

	outside:
	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {
			fmt.Print("+ ")
			if i == 6 && j == 3 {
				// call the label and stop the corresponding loop
				break outside
			}
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("label 1")
	goto lab  // skip the following lines and jump directly to the label
	fmt.Println("label 2")

	lab:  // make a name here
	fmt.Println("label 3")
	fmt.Println("label 4")
}

