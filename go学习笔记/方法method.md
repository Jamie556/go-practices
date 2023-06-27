
# 方法

## 基本概念

Go 语言支持方法。Go 方法类似于 Go 函数，但有一个区别，即该方法包含一个接收器参数。在接收器参数的帮助下，该方法可以访问接收器的属性。这里，接收者可以是结构类型或非结构类型。当您在代码中创建方法时，接收器和接收器类型必须存在于同一个包中。并且您不允许创建接收器类型已在另一个包中定义的方法，包括 int、string 等内置类型。如果您尝试这样做，编译器将给出错误。

---
```go
// 语法
func(reciver_name 类型) method_name(parameter_list)(return_type){
	// 代码
	}
```
---
示例一，使用结构体

---
```go
// Go program to illustrate the
// method with struct type receiver
package main

import "fmt"

// Author structure
type author struct {
	name	 string
	branch string
	particles int
	salary int
}

// Method with a receiver
// of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:	 "Sona",
		branch: "CSE",
		particles: 203,
		salary: 34000,
	}

	// Calling the method
	res.show()
}
```
---
示例二，使用非结构体

---
```go
// Go program to illustrate the method
// with non-struct type receiver
package main

import "fmt"

// Type definition
type data int

// Defining a method with
// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

/*
// if you try to run this code,
// then compiler will throw an error
func(d1 int)multiply(d2 int)int{
return d1 * d2
}
*/

// Main function
func main() {
	value1 := data(23)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}

```
---
示例三，使用带指针接收器的方法

---
```go
func (p *Type) method_name(...Type) 类型 {
	// 代码
}
```
---
---
```go
// Go program to illustrate pointer receiver
package main

import "fmt"

// Author structure
type author struct {
	name	 string
	branch string
	particles int
}

// Method with a receiver of author type
func (a *author) show(abranch string) {
	(*a).branch = abranch
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name: "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(Before): ", res.branch)

	// Creating a pointer
	p := &res

	// Calling the show method
	p.show("ECE")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(After): ", res.branch)
}

```
---

示例四，接受值和指针
在go的函数（func）中，接受的参数的值是唯一确定的，如果定义为指针的话只能接受指针，反之亦然
但是在go的方法（method）中，可以同时接受值和指针作为参数

---
```go
// Go program to illustrate how the
// method can accept pointer and value

package main

import "fmt"

// Author structure
type author struct {
	name string
	branch string
}

// Method with a pointer
// receiver of author type
func (a *author) show_1(abranch string) {
	(*a).branch = abranch
}

// Method with a value
// receiver of author type
func (a author) show_2() {

	a.name = "Gourav"
	fmt.Println("Author's name(Before) : ", a.name)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name: "Sona",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before): ", res.branch)

	// Calling the show_1 method
	// (pointer method) with value
	res.show_1("ECE")
	fmt.Println("Branch Name(After): ", res.branch)

	// Calling the show_2 method
	// (value method) with a pointer
	(&res).show_2()
	fmt.Println("Author's name(After): ", res.name)
}
```
---
