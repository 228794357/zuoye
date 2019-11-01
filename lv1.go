package main

import (
	"fmt"
)

type dove interface {
	gugugu()
}

type repeater interface {
	repeat(word string)
}

type lemon interface {
	suan()
}

type zhenxiangguai interface {
	zhenxiang()
}

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) gugugu() {
	fmt.Println(p.name,p.age,p.gender,"又鸽了")
}

func (p *person) repeat(word string) {
	fmt.Println(word )
}

func (p *person) suan() {
	fmt.Println(p.name,"酸了")
}

func (p *person) zhenxiang() {
	fmt.Println(p.name,"真香")
}

func main() {
	p :=  &person{
		name: "uzi",
		age: 26,
		gender: "male",
	}
var (
	a dove
	b repeater
	c lemon
	d zhenxiangguai
)
a = p
b = p
c = p
d = p
a.gugugu()
b.repeat("rng")
c.suan()
d.zhenxiang()
}