package basics

import "fmt"
import "reflect"


const (
	VERSION = 1.0    // the first letter should be upper case
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

func IntData() {
	var (
		i1 = 5
		i2 int8 = 127
		i3 uint16
		i4 = 0b0101    //  2: binary
		i5 = 0O77      //  8: octal
		i6 = 0xAF      // 16: hexadecimal
	)
	var f1 = float32(i1)

	fmt.Println("\n1-3. Integer")
    fmt.Printf("i1=%v, i2=%v, i3=%v, i4=%v, i5=%v, i6=%v, f1=%v\n",
               i1, i2, i3, i4, i5, i6, f1)
    fmt.Printf("i1=%T, i2=%T, i3=%T, i4=%T, i5=%T, i6=%T, f1=%T\n",
               i1, i2, i3, i4, i5, i6, f1)
}

func OtherData() {
	var (
		i8 int8 = 8
		i16 int16 = 16
		i32 int32 = 32
		i64 int64 = 64

		u8 uint8 = 8
		u16 uint16 = 16
		u32 uint32 = 32
		u64 uint64 = 64

		f32 float32 = 32.32
		f64 float64 = 64.64

		c64 complex64 = complex(3, 2)
		c128 complex128 = complex(4, 3)

		t, f bool = true, false
		s string = "string"
		r rune = 'r'
	)

	fmt.Println("\n1-4. Other Data")
    fmt.Printf("i8=%v, i16=%v, i32=%v, i64=%v, u8=%v, u16=%v, u32=%v, u64=%v\n",
               i8, i16, i32, i64, u8, u16, u32, u64)
    fmt.Printf("f32=%v, f64=%v, c64=%v, c128=%v\n", f32, f64, c64, c128)
    fmt.Printf("t=%v, f=%v, s=%v, r=%v\n", t, f, s, r)
    fmt.Printf("length of string: %v\n", len(s))

    str := `var (
    	abc byte = 78
    	c2       = '0'
    	c3 rune  = 23454
    )`
    fmt.Println(str)
}

func DataType() {
	// There are 3 ways to find out the data type of a variable
	c64 := complex(3, 2)

	fmt.Println("\n1-5. Data Type")
	fmt.Printf("%T\n", c64)
	fmt.Println(reflect.TypeOf(c64))

	var c64_t reflect.Kind = reflect.ValueOf(c64).Kind()
	fmt.Println(c64_t)
}

func Pointer() {
	fmt.Println("\n1-6. Pointer")

	var n1 = 100
	fmt.Printf("Original value n=%v\n", n1)
	increase(n1)
	fmt.Printf("After 'increase' n=%v\n", n1)
	increase_ptr(&n1)  // & get the address of the current value
	fmt.Printf("After 'increase_ptr' n=%v\n", n1)

	var ptr = new(int)  // Create an address with zero value.
	fmt.Printf("ptr: %v, address: %v\n", *ptr, ptr)
}

func increase(n int) {
	n++    // n = n + 1
}

func increase_ptr(n *int) {
	*n += 100  // * get the value with respect to the address
}

