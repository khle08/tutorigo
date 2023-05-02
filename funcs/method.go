////////////////////////////////////////////////////////////////////////

package funcs

import "fmt"
import "math"
import "sync"
import "time"
import _"strconv"

////////////////////////////////////////////////////////////////////////

type Car struct {
    gear int8
    door int8
    wheel int8
    name string
}

func (c Car) drive() {
    fmt.Println("\n6-1. method")
    c.gear = 1    // this is "copy" without sharing memory
    fmt.Println(c)
}

func (c *Car) back() {
    c.gear = 6
    fmt.Println(*c)
}


func Method() {
    car := Car {
        gear: 0,
        door: 4,
        // wheel: 3,
        name: "this is a three wheel car",
    }
    car.drive()
    fmt.Println("car gear:", car.gear)
    car.back()
    fmt.Println("car gear:", car.gear)
}

// =====================================================================

type geometry interface {
    area() float32
    edge() float32
}

type rect struct {
    width, height float32
}

func (r rect) area() float32 {
    return r.width * r.height
}

func (r rect) edge() float32 {
    return 2 * (r.width + r.height)
}

type circle struct {
    radius float32
}

func (c circle) area() float32 {
    return math.Pi * c.radius * c.radius
}

func (c circle) edge() float32 {
    return 2 * math.Pi * c.radius
}

func Measure(g geometry) {
    fmt.Printf("type: %T, obj: %v\n", g, g)
    fmt.Println("area:", g.area())
    fmt.Println("edge:", g.edge())
}

func Interface() {
    fmt.Println("\n6-2. interface")

    r := rect{width: 4, height: 5}
    Measure(r)

    c := circle{radius: 10}
    Measure(c)
}

// =====================================================================

func async(name string) {
    for i := 1; i <= 5; i++ {
        fmt.Println("async", name, ": ", i)
    }
}

func GoAsync() {
    fmt.Println("\n6-3. asyncronous method")

    // As there are "go" in front of the func, it can be run parallelly
    // without considering their sequence.
    go async("1st")
    go async("2nd")
    go async("3rd")

    // Anonymous function
    go func(msg string) {
        fmt.Println("this is an", msg)
    }("anonymous function")

    time.Sleep(time.Second)
}


var (
    c int
    lock sync.Mutex
)

func prime(n int) {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return
        }
    }
    fmt.Println("n:", n)
    
    // If there are multiple progess modifying the same variable, it will has some error.
    // To prevent this situation, lock the variable this way when operating it.
    lock.Lock()
    c++
    lock.Unlock()
}

func Goroutine() {
    fmt.Println("\n6-4. Goroutine")

    for i := 2; i < 100000; i++ {
        go prime(i)
    }

    // By using "go" at the beginning for a func, it will create a new thread.
    // However, this thread will be killed when the main thread is finished.
    // In order to make sure all the sub threads are well executed, put sth here
    // in the main thread, which is "Scanln".
    var key string
    fmt.Scanln(&key)
    fmt.Printf("\n Find %v numbers\n", c)
}


func Channel() {
    fmt.Println("\n6-5. Channel")

    var ch1 chan string = make(chan string, 1)  // buf = 1 here

    // When an async func is executed, the generated data will be put into "channel".
    go func() {
        for i := 1; i <= 3; i++ {
            ch1 <- fmt.Sprintf("%s <- %d", "ch1", i)
            fmt.Println("finish - ", i)
        }
    }()

    // If we do not setup buf for a channel, the object "ch1" will wait until someone
    // retrieves the saved data.
    res := ""
    res = <- ch1
    fmt.Println("res:", res)
    res = <- ch1
    fmt.Println("res:", res)
    res = <- ch1
    fmt.Println("res:", res)
}
