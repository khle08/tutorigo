package main

import (
    "fmt"
    "tutorigo/basics"
    "tutorigo/packageName"
)


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
}