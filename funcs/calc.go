////////////////////////////////////////////////////////////////////////

package funcs

import "fmt"
import "errors"
import _ "reflect"    // not use but also no error

////////////////////////////////////////////////////////////////////////

var Z = 0
var F = func(s string) int {
	fmt.Printf("func is called by: %s\n", s)
	Z++
	fmt.Printf("current value Z: %v\n", Z)
	return Z
}


func Add(a int, b int) int {  // (a, b int) also ok when they are same type
	fmt.Println("\n3-1. add function")
	return a + b
}


func SumDif(a, b int) (int, int) {  // the return type can also be named
	fmt.Println("\n3-2. multiple returned values")

	fmt.Printf("Add: %v, Type: %T\n", Add, Add)  // func is also one type of datatype
	return a + b, a - b
}


func SumDif2(a, b int) (sum, dif int) {  // Exactly the same as above "SumDif"
	fmt.Println("\n3-3. name the returned values")

	sum = a + b
	dif = a - b
	return  // Since names already defined above, go know what to return
}


func Hidden() {
	fmt.Println("\n3-4. hidden function")

	var hid = func(a, b int) (sum, dif int) {
		sum = a + b
		dif = a - b
		return
	}
	r1, r2 := hid(5, 2)
	fmt.Printf("r1: %v, r2: %v\n", r1, r2)
}


func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		// 闭包的写法
		fmt.Printf("hidden func is run %v times\n", i)
		i++
		fmt.Printf("current value i: %v, n: %v\n", i, n)
		return i
	}
}


func Defer() {
	fmt.Println("\n3-5. defer function")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()  // execute the hidden func directly when defining it

	du := deferUtil()
	defer du(1)     // 3rd execute
	defer du(100)   // 2nd execute
	defer du(2)     // 1st execute

	n := 0
	fmt.Println(3 / n)  // this will cause "error" and the error will be
	// deposited by the hidden function with "recover()" func.
}


func errorDemo(num int) (int, error) {
	// Default error type provided by golang.
	if num < 0 {
		return -1, errors.New("Negative number")
	}

	return 2 * num, nil
}

type ErrStruct struct {
	// Self-defined error type
	arg int
	msg string
}

func (e *ErrStruct) Error() string {
	// As long as "Error" method defined in a struct, it works the
	// same way as default error handler.
	return fmt.Sprintf("%d -> %s", e.arg, e.msg)
	// This is useful when recording the error log info in the backend.
}

func errorDemo2(num int) (int, error) {
	if num < 0 {
		return -1, &ErrStruct{num, "can not be negative value"}
	}

	return 2 * num, nil
}

func ErrorHandler() {
	fmt.Println("\n3-6. Error handling")

	// default
	res, err := errorDemo(10)
	fmt.Println(res, err)
	res, err = errorDemo(-10)
	fmt.Println(res, err)

	// self-defined
	res, err = errorDemo2(10)
	fmt.Println(res, err)
	res, err = errorDemo2(-10)
	fmt.Println(res, err)
}


