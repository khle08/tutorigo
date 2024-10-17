////////////////////////////////////////////////////////////////////////

// Date: 2022.5.30
// Author: Kuo Chun Lin
// ref : https://www.youtube.com/playlist?list=PLBjZhzRvV2ChPTPNDx_apHdKa9Ha7LVpN

////////////////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "tutorigo/funcs"
    "tutorigo/tools"
    // "tutorigo/basics"
    "tutorigo/template"
    "tutorigo/packageName"
)

////////////////////////////////////////////////////////////////////////


var ZZ = funcs.Z

func init() {
    funcs.F("main.go --> 1st init")
}

func init() {
    funcs.F("main.go --> 2nd init")
    fmt.Scanln()
}


////////////////////////////////////////////////////////////////////////


func main() {
    // // In order to call the function at same folder as "main.go", ...
    // // ... run the following CMD line:
    // // $ go run /Users/kcl/Documents/Go_Projects/tutorigo
    // Print()    // Defined in "others.go"

    // In order to call function defined in another package organized in ...
    // ... other folder, 1) import the package 2) call it here.
    packageName.Print()
    // packageName.Print2(123)
    // packageName.Print2("this is a string")
    // fmt.Printf("note version: %v\n", basics.VERSION)

    // fmt.Printf("var L: %v\n", template.L)
    intStack := template.Stack[int]{}
    intStack.Push(1)
    // intStack.Push(2)
    // fmt.Println(intStack.Items)
    // fmt.Println(intStack.GetItems())

    // basics.Variable()
    // basics.Constant()
    // basics.IntData()
    // basics.OtherData()
    // basics.DataType()
    // basics.Pointer()

    // basics.IfElse()
    // basics.SwitchCase1()
    // basics.SwitchCase2()
    // basics.ForLoop()
    // basics.Label2Goto()

    // fmt.Println(funcs.Add(1, 2))
    // fmt.Println(funcs.SumDif(3, 2))
    // fmt.Println(funcs.SumDif2(3, 2))
    // funcs.Hidden()
    // funcs.Defer()
    // funcs.ErrorHandler()  // needs some knowledge of "Method"

    // funcs.Array()
    // funcs.Slice()
    // funcs.Map()
    // funcs.SelfDefineType()
    // funcs.Struct()

    // funcs.Method()
    // funcs.Interface()
    // funcs.GoAsync()
    // // funcs.Goroutine()
    // funcs.Channel()
    
    // tools.ReadFile()
    // tools.ParamTest()
    // tools.SaveViper()
    // tools.OpenCV()
    tools.TestIP()
}