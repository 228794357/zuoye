package main

import "fmt"

func abc(n int, b chan int) {
	var res =1
	for i:=1; i<=20;i++{
		res *=i
		b<-res
	}
}

func main() {
	b :=make(chan int,20)
	go abc(cap(b),b)
	for v:=range b{
		fmt.Println(v)
	}
}