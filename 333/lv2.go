package main

import "fmt"

func main() {
	ch:=make (chan int)
	go factorial(ch)
	for i :=range ch{
	fmt.Println(i)
	}
}

func factorial (n   chan int) {
	var res =1
	for i := 1; i <= 20; i++ {
		res *=i
		n<-res
		}
	close(n)
	}


