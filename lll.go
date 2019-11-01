package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
a:=make (chan int)
	a<-i
	go Factorial1(i,a)
	for V:=range a{
		fmt.Println(V)
	}}
}

func Factorial1(n int, ch chan int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
}