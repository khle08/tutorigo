package funcs

import "fmt"


func Array() {
	fmt.Println("\n4-1. array")

	var arr1 [3]int = [3]int{1, 200, 4000}
	// arr1[3] = 30  // array has fixed length that can not be extended this way.
	
	// if we do not know how many elements in the array, use "..."
	var arr2 = [...]int{3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println("\nfor loop iteration")
	for i:=0; i<len(arr2); i++ {
		fmt.Printf("arr2[%v] : %v\n", i, arr2[i])
		if arr2[i] == 6 {
			// the element can be replaced by other values
			arr2[i] = 1000
		}
	}
	fmt.Println()

	for i, v := range arr2 {
		// this way of "v" is a new var with different address
		fmt.Printf("arr2[%v] : %v\n", i, v)
	}
	fmt.Println(arr2)
	fmt.Println()

	var arr3 = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	for i, vi := range arr3{
		for j, vj := range vi {
			fmt.Printf("arr[%v][%v] : %v\t", i, j, vj)
		}
		fmt.Println()
	}
}


func Slice() {
	fmt.Println("\n4-2. Slice")

	arr := [5]int{3, 4, 5, 6, 7}
	// slice is the reference of the array. they share the same memory
	var s1 []int = arr[:3]
	fmt.Printf("before changes arr: %v\n", arr)
	fmt.Printf("before changes s1: %v\n", s1)

	s1[1] = 100
	fmt.Printf("after changes arr: %v\n", arr)
	fmt.Printf("after changes s1: %v\n", s1)

	var s2 []int
	fmt.Println("s2:", s2, "is nil ? ", s2 == nil)

	s2 = make([]int, 3, 5)    // make([]type, len, cap)
	fmt.Printf("len(s2)=%v, cap(s2)=%v\n", len(s2), cap(s2))

	s3 := []int{10, 20, 30}   // more elements can be added to slice
	copy(s3, s1)              // mind that array can not be copied to slice
	fmt.Println("s3: ", s3)   // the copied elements will only replace to the same len as s3

	s1 = append(s1, 100, 101, 102)
	fmt.Printf("before changes arr: %v\n", arr)
	fmt.Printf("after changes s1: %v\n", s1)
	// when new elements are added, s1 would be in diff address
	
	str := "hello world"
	fmt.Printf("str: %s\nbyte: %v\n", str, []byte(str))
	for i, v := range str {
		fmt.Printf("str[%v]=%c\n", i, v)
	}
}


func Map() {
	fmt.Println("\n4-3. Map")

	var m1 map[string]string
	fmt.Println("m1:", m1, "is nil ? ", m1 == nil)
	// if you know how many key at the beginning already, write it to gain better performance
	m1 = make(map[string]string, 2)  // make(type, size)

	m1["monday"] = "sunny"
	m1["tuesday"] = "rainny"
	m1["wednesday"] = "cloudy"
	fmt.Println("m1:", m1)

	val, ok := m1["friday"]
	if ok {
		fmt.Println("val =", val)
	} else {
		fmt.Println("key not exist")
	}
	delete(m1, "monday")
	fmt.Println("monday deleted - m1:", m1)

	for key, val := range m1 {
		fmt.Printf("m1[\"%s\"]=%s\n", key, val)
	}
}


func SelfDefineType() {
	fmt.Println("\n5-1. Define customized data type")

	type MesType uint16
	var u1 uint16 = 1000
	var msg1 MesType = MesType(u1)
	var msg2 MesType = 200
	fmt.Printf("msg1: %v, type: %T\nmsg2: %v, type: %T\n", msg1, msg1, msg2, msg2)
}


type User struct {
	name string
	id uint32
}
type Account struct {
	User    // inherit from User struct
	password string
}
type Contact struct {
	*User   // inherit pointer from User struct
	note string
}

func Struct() {
	fmt.Println("\n5-2. struct")

	var u1 User = User{
		name: "jack",
		id: 2022,
	}
	fmt.Println("u1:", u1)

	// struct with pointer
	var u2 *User = &User {
		name: "Bob",
		id: 2021,
	}
	fmt.Println("u2:", *u2)

	var a1 = Account {
		User: u1,         // this is "copy" which does not share the same memory
		password: "abc",
	}
	u1.name = "Allen"
	fmt.Println("copy from u1:", u1)
	fmt.Println("a1:", a1)

	var a2 = Account {
		User: *u2,        // this also "copy" which does not share the same memory
		password: "abc",
	}
	u2.name = "Pop"
	fmt.Println("shared with u2:", *u2)
	fmt.Println("a2:", a2)

	var c1 = Contact {
		User: u2,
		note: "share",
	}
	u2.name = "Modified"
	fmt.Println("shared with u2:", *u2)
	fmt.Println("c1:", c1, "c1.User:", *(c1.User))

}



