package main

import (
	"fmt"
	"rsc.io/quote"
)

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2
}

func show() {
	fmt.Println("Hello, GeeksforGeeks")
}

func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	testSlice()
	testString()
	testFunction()

	mul1, _ := mul_div(105, 7)
	fmt.Println("105 * 7 = ", mul1)

	mul(23, 45)
	defer mul(23, 45)
	show()

	fmt.Println("=== start")
	defer fmt.Println("END")
	defer add(34, 56)
	defer add(10, 10)

	res := author{
		name: "sona",
		branch: "CSE",
		particles: 203,
		salary: 34000,
	}
	res.show()

	value3 := data(34)
	value4 := data(20)

	res1 := value3.multiply(value4)
	fmt.Println("FInal res is: ", res1)

	res2 := author {
		name: "Sona",
		branch: "CSE",
	}

	fmt.Println("name: ", res2.name)
	fmt.Println("branch(before): ", res2.branch)

	p := &res2

	p.show2("ECE")
	fmt.Println("name: ", res2.name)
	fmt.Println("branch(before): ", res2.branch)

	res3 := author2 {
		name: "Sona",
		branch: "CSE",
	}

	fmt.Println("1", res3.branch)
	res3.show_1("ECE")
	fmt.Println("2", res3.branch)
	(&res3).show_2()
	fmt.Println("3", res3.branch)

	// 结构体
	var a Address
	fmt.Println(a)

	a1 := Address{"a1", "b2", 123}
	fmt.Println("a1: ", a1)

	a2 := Address{Name: "a2", city: "广州", Pincode: 222}
	fmt.Println("a2: ", a2)

	a3 := Address{Name: "a3", city: "长沙"}
	fmt.Println("a3: ", a3)

	c := Car{Name: "Ferrari", Model: "GTC4", Color: "Red", WeightInKg: 1920}

	fmt.Println("car Name: ", c.Name)
	fmt.Println("car Color: ", c.Color)

	c.Color = "Black"
	fmt.Println("Car: ", c)

	emp8 := &Empployee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name1: ", (*emp8).firstName)
	fmt.Println("First Name2: ", emp8.firstName)
	// 结构体取值法
	fmt.Println("Age1: ", (*emp8).age)
	// 指针取值法
	fmt.Println("Age2: ", emp8.age)

	// 嵌套结构体
	result := HR{
		 details: Author3{"Sona", "ECE", 2022},
	}
	fmt.Println("details: ", result)

	result1 := Teacher{
		name: "Suman",
		subject: "Java",
		exp: 5,
		details: Student{"Bongo", "CSE", 2},
	}
	fmt.Println("t'name: ", result1.name)
	fmt.Println("t'subject: ", result1.subject)
	fmt.Println("t'exp: ", result1.exp)
	fmt.Println("t'exp: ", result1.exp)

	fmt.Println("t'details: ", result1.details)
	fmt.Println("s'name: ", result1.details.name)
	fmt.Println("s'branch: ", result1.details.branch)
	fmt.Println("s'year: ", result1.details.year)

	// 匿名结构体
	Element := struct {
		name string
		branch string
		language string
		Particles int
	} {
		name: "Pikachu",
		branch: "ECE",
		language: "C++",
		Particles: 498,
	}

	fmt.Println(Element)


	value5 := student{123, "bud", 8900.23}
	fmt.Println("num:", value5.int)
	fmt.Println("name:", value5.string)
	fmt.Println("price:", value5.float64)

	arr:= [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}


	arr2 := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},{"C++", "Go", "HTML"}}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr2[x][y])
		}
 	}

	var arr3 [2][2]int
	arr3[0][0] = 100
	arr3[0][1] = 200
	arr3[1][0] = 300
	arr3[1][1] = 400

	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr3[p][q])
		}
	}

	var value6 []int
	fmt.Println("value6: ", value6)


	arr6 := [3]int{9,7,6}
	arr7 := [...]int{1,2,3,4,5,6,7,8,9,0}
	arr8 := [3]int{1,2,3}

	fmt.Println("arr6 len: ", len(arr6))
	fmt.Println("arr7 len: ", len(arr7))
	fmt.Println("arr8 len: ", len(arr8))


	myarray:= [...]string{"GFG", "gfg", "geeks","GeeksforGeeks", "GEEK"}
	fmt.Println("Elements of the array: ", myarray)
	fmt.Println("Length of the array is:", len(myarray))

	myarra := [...]int{29, 79, 49, 39, 20, 49, 48, 49}
	for x := 0; x < len(myarra); x++ {
		fmt.Printf("%d\n", myarra[x])
	}
	tstArr := [3]string{"1", "2", "3"}
	tstSlice := tstArr[:0]
	fmt.Println("tstSlice", cap(tstSlice))
}


