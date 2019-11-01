package main

import "fmt"

func sushu (ch1 chan int) {
	for i:=2;i<=10000;i++{
		ch1 <- i
	}
	close(ch1)
}

func sushu2 (n int,ch1 , ch2 chan int) {
	for{
		i:=<-ch1
		if i%n!=0{
			ch2<-i
		}
	}
}

func main() {
	b:=make(chan int,10000)
	c :=make(chan int, 10000)
	go sushu(b)
	for i:=2;i<=10000;i++{
		n:=<-b
		fmt.Println(n)
		go sushu2(n,b,c)
		b=c
	}
}
