package basics

import "fmt"
import "tutorigo/funcs"


func init() {
	fmt.Printf("value defined in calc.go: Z=%v\n", funcs.Z)
	funcs.F("initialization.go --> 1st init")
}

func init() {
	fmt.Printf("value defined in calc.go: Z=%v\n", funcs.Z)
	funcs.F("initialization.go --> 2nd init")
}