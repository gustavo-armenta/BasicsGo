package main

import "fmt"
import "errors"

import "time"

func main() {
	// learnFunctions()
	// learnMaps()
	// learnInterfaces()
	// learnDeferPanicRecover()
	// learnGoroutines()
	learnChannels()
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
	defer doRecover()
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
	go receiver(ch)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("exit")
}

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("sending:", i)
		ch <- i
	}
}

func receiver(ch chan int) {
	i := 0
	for i < 4 {
		i := <-ch
		fmt.Println("received:", i)
	}
}
