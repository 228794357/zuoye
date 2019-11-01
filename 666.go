package main

import (
	"fmt"
	"time"
)

func Factorial(n int, ch  chan int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch<-res
}

func main() {
	b := make(chan int,20)
	for i := 1; i <= 20; i++ {
		go Factorial(i, b)
	}
	for  v := range b {
		fmt.Println(v)
	}
	time.Sleep(2 * time.Second )
}
