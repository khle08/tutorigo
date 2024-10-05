package packageName

import "fmt"


func Print() {
    fmt.Println("hello world")
}


func Print2[T any](value T) {
    fmt.Println(value)
}
