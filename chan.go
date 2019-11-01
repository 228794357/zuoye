package main

import "fmt"

func factorial(ch1 chan int) {
	for i := 1; i <= 20; i++{
		ch1 <-i
	}
	close(ch1)
}

func factorial2(ch1,ch2 chan int) {
	var res =1
	for {
		x, ok:=<-ch1
		if !ok{
			break
		}
		res *= x
		ch2 <-res
	}
	close(ch2)
}

func main() {
	var(
		a =make(chan int,20)
		b =make(chan int,20)
	)
	go factorial(a)
	go factorial2(a,b)
	for c :=range b{
		fmt.Println(c)
	}
}
