package main

import "fmt"

var (
	myRes =make(map[int]int,20)
)
func f1(n int,ch chan map[int]int ) {
	var res =1
	//ch :=make(chan map[int]int)
	for i :=1;i<=n;i++{
		res *=i
	}
myRes[n]=res
ch<-myRes
}

func main() {
	b:=make (chan map[int]int,20)
	for i:=1;i<=20;i++{
		go f1(i,b)
	}
	for v:=range b{
		fmt.Println(v)
	}
}