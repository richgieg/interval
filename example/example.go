package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/richgieg/interval"
)

func main() {
	c := 1
	i := interval.NewInterval(time.Millisecond*450, func() {
		fmt.Println("tick", c)
		c++
	})

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	i.Start()
	fmt.Println("Interval started")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)

	i.Stop()
	fmt.Println("Interval stopped")
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	i.Start()
	fmt.Println("Interval started")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)

	i.Stop()
	fmt.Println("Interval stopped")
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
