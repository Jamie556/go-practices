package main

import "fmt"

type author struct {
	name string
	branch string
	particles int
	salary int
}

func (a author) show() {
	fmt.Println("Author'name: ", a.name)
	fmt.Println("Branch name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

type data int

func(d1 data) multiply(d2 data) data {
	return d1 * d2
}

func(a *author) show2(abranch string) {
	(*a).branch = abranch
}


type author2 struct {
	name string
	branch string
}

func (a * author2) show_1(abranch string) {
	(*a).branch = abranch
}

func (a author2) show_2() {
	a.name = "Gourav"
	fmt.Println("Author'name(Before): ", a.name)
}


