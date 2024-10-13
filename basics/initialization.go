package basics

import "fmt"
import "tutorigo/funcs"


func init() {
	fmt.Printf("\n[initialization.go] value defined in calc.go: Z=%v\n", funcs.Z)
	funcs.F("initialization.go --> 1st init\n")
}

func init() {
	fmt.Printf("\n[initialization.go] value defined in calc.go: Z=%v\n", funcs.Z)
	funcs.F("initialization.go --> 2nd init\n")
}