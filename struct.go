package main

type Address struct {
	Name string
	city string
	Pincode int
}

type Car struct {
	Name, Model, Color string
	WeightInKg float32
}

type Empployee struct {
	firstName, lastName string
	age, salary int
}

type Author3 struct {
	name string
	branch string
	year int
}

type HR struct {
	details Author3
}

type Student struct {
	name string
	branch string
	year int
}

type Teacher struct {
	name string
	subject string
	exp int
	details Student
}

type student struct {
	int
	string
	float64
}
