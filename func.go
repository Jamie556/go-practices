package main

import (
	"fmt"
	"strings"
	"sort"
	"time"
)

func testFunction() {
	// 传值
	fmt.Printf("\nArea of rectangle is: %d\n", area(2, 3))
	var p int = 10
	var q int = 20
	fmt.Printf("\np = %d and q = %d\n", p, q)
	var res int = swap(p, q)
	fmt.Printf("\nRes is: %d\n", res)
	fmt.Printf("\np = %d and q = %d\n", p, q)
	// 传指针
	fmt.Printf("\np = %d and q = %d\n", p, q)
	var res2 = swap2(&p, &q)
	fmt.Printf("\nRes is: %d\n", res2)
	fmt.Printf("\np = %d and q = %d\n", p, q)

	// 可变参数函数
	fmt.Println(joinStr())
	fmt.Println(joinStr("GEEK", "GFG"))
	fmt.Println(joinStr("Geeks", "for", "GeeKs"))
	fmt.Println(joinStr("G", "E", "E", "K", "s"))

	//传切片
	elements := []string{"geeks", "For", "geeks"}
	fmt.Println(joinStr(elements...))

	// 匿名函数
	func ()  {
		fmt.Println("Welecome! to GeeksForGeeks")
	}()

	// 为变量分配匿名函数
	value := func ()  {
		fmt.Println("assgin anymouns func to variables")
	}
	// 运行匿名函数变量
	value()

	// 在匿名函数中传递参数
	func (element string) {
		fmt.Println(element)
	}("pass params to anymouns funcs")

	// 将匿名函数作为参数传递给其他函数
	value2 := func (p, q string) string {
		return p + q + "c"
	}
	GFG(value2, "a", "b")

	// 从一个函数中返回一个匿名函数
	value3 := testAnymounsFuncs()
	var res3 string = value3("return", "funcs")
	fmt.Println("res3: ", res3)


	// main和init函数
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("sorted slice", s)

	// 字符的位置
	fmt.Println("Index value: ", strings.Index("GeeksForGeeks", "ks"))
	// unix时间
	fmt.Println("Time: ", time.Now().Unix())

	
}



func area(length, width int) int {
	Ar := length * width
	return Ar
}

func swap(a, b int) int {
	var o int
	o = a
	a = b
	b = o
	return 0
}


func swap2(a, b * int) int {
	var o int
	o = *a
	*a = *b
	*b = o
	return o
}


func joinStr(elements ...string) string {
	return strings.Join(elements, "-")
}


func GFG(i func(p, q string) string, x string, y string) {
	fmt.Println(i(x, y))
}

func testAnymounsFuncs() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "test"
	}
	return myf
}
