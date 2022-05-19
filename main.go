package main

import (
    "fmt"
    "tutorigo/funcs"
    "tutorigo/basics"
    "tutorigo/packageName"
)


var ZZ = funcs.Z

func init() {
    funcs.F("main.go --> 1st init")
}

func init() {
    funcs.F("main.go --> 2nd init")
    fmt.Scanln()
}


func main() {
    // // In order to call the function at same folder as "main.go", ...
    // // ... run the following CMD line:
    // // $ go run /Users/kcl/Documents/Go_Projects/tutorigo
    // Print()    // Defined in "others.go"

    // In order to call function defined in another package organized in ...
    // ... other folder, 1) import the package 2) call it here.
    packageName.Print()
    fmt.Printf("note version: %v\n", basics.VERSION)

    basics.Variable()
    basics.Constant()
    basics.IntData()
    basics.OtherData()
    basics.DataType()
    basics.Pointer()

    basics.IfElse()
    basics.SwitchCase1()
    basics.SwitchCase2()
    basics.ForLoop()
    basics.Label2Goto()

    fmt.Println(funcs.Add(1, 2))
    fmt.Println(funcs.SumDif(3, 2))
    fmt.Println(funcs.SumDif2(3, 2))
    funcs.Hidden()
    funcs.Defer()

    funcs.Array()
    funcs.Slice()
    funcs.Map()
    funcs.SelfDefineType()
    funcs.Struct()

    funcs.Method()
    funcs.Interface()
}