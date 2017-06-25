package main

import "fmt"
import "errors"
import "sync"
import "time"

func main() {
	learnArrays()
	learnBasicTypes()
	learnConditionals()
	learnLoops()
	learnFunctions()
	learnMaps()
	learnInterfaces()
	learnDeferPanicRecover()
	learnGoroutines()
	learnChannels()
	learnMutex()
}

func learnArrays() {
	array := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(array)
	slice := array[:]
	slice = append(slice, 17)
	fmt.Println(slice, len(slice), cap(slice))
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board)
}

func learnBasicTypes() {
	var b bool = true
	var by byte = 1
	var i int = 1
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var f32 float32 = 1.0
	var f64 float64 = 1.0
	var c64 complex64 = 1.0 + 1.0i
	var c128 complex128 = 1.0 + 1.0i
	var r rune
	fmt.Println(b, by, i, i8, i16, i32, i64, f32, f64, c64, c128, r)
}

func learnConditionals() {
	x := 5
	if x < 1 {
		fmt.Println("if")
	} else if x > 3 && x < 5 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	switch x {
	case 1:
		fmt.Println("case 1")
	case 5:
		fmt.Println("case 5")
	default:
		fmt.Println("case default")
	}
}

func learnLoops() {
	a := 0
	for i := 0; i < 10; i++ {
		a += i
	}
	for a > 10 {
		a--
	}
	array := [6]int{2, 3, 5, 7, 11, 13}
	slice := array[:]
	for i, v := range slice {
		fmt.Println("for:", i, v)
	}
	for _, v := range slice {
		fmt.Println("for:", v)
	}
}

func learnMaps() {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m)
	fmt.Println(m["one"])
	m["three"] = 3
	fmt.Println(m)
	delete(m, "three")
	fmt.Println(m)
	value, ok := m["three"]
	fmt.Println(value, ok)
	fmt.Println(m["three"])
}

func learnFunctions() {
	c := closure()
	r := 0
	for i := 0; i < 10; i++ {
		r = c(i)
	}
	fmt.Println(r)

	v, err := returnError(true)
	fmt.Println(v, err)
	v, err = returnError(false)
	fmt.Println(v, err)
}

func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func returnError(b bool) (int, error) {
	if b {
		return 0, errors.New("this function returns an error")
	} else {
		return 1, nil
	}
}

func learnInterfaces() {
	d := dog{Message: "buf!"}
	var a animal
	a = &d
	fmt.Println(a.Speak())
}

type animal interface {
	Speak() string
}

type dog struct {
	Message string
}

func (a dog) Speak() string {
	return a.Message
}

func learnDeferPanicRecover() {
	handlePanic()
}

func returnPanic() {
	fmt.Println("before panic")
	panic("panic!")
}

func doRecover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in handlePanic:", r)
	}
}

func handlePanic() {
	fmt.Println("before defer")
	defer func() { fmt.Println("defer1") }()
	defer doRecover()
	defer func() { fmt.Println("defer3") }()
	fmt.Println("after defer")
	returnPanic()
	fmt.Println("this line is not executed due to panic")
}

func learnGoroutines() {
	go say("async")
	say("sync")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func learnChannels() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
	fmt.Println("channel closed and receiver completed")
}

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("sending:", i)
		ch <- i
	}
	close(ch)
}

func receiver(ch chan int) {
	for i := range ch {
		fmt.Println("received:", i)
	}
}

func learnMutex() {
	c := safeCounter{}
	go inc(&c, "a")
	go inc(&c, "b")
	go inc(&c, "c")
	for c.Value() < 30 {
		time.Sleep(100 * time.Millisecond)
	}
}

type safeCounter struct {
	v int
	m sync.Mutex
}

func (c *safeCounter) Increment() int {
	defer c.m.Unlock()
	c.m.Lock()
	c.v++
	return c.v
}

func (c *safeCounter) Value() int {
	defer c.m.Unlock()
	c.m.Lock()
	return c.v
}

func inc(c *safeCounter, s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(c.Increment(), s)
		time.Sleep(100 * time.Microsecond)
	}
}
