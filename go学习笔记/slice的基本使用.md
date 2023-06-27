
# slice 的基本使用和基本概念


1. 概念

1) In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure.
   在 go 语言中，slice 是一种比数组更强大、灵活、方便，并且更加的轻量的数据结构。
   2）Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice
   slice 是一种不固定长度的序列，用于存储同一种数据类型的数据
   3）It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array
   就像数组一样拥有下标、值和长度，但是 slice 的大小不像是数组一样是固定的
   4）Internally, slice and an array are connected with each other, a slice is a reference to an underlying array
   在底层，slice 和数组是相互关联的，一个 slice 指向一个潜在的数组
   5）It is allowed to store duplicate elements in the slice. The first index position in a slice is always 0 and the last one will be (length of slice – 1).
   slice 允许存储重复的元数据，slice 的第一个下标的位置是 0，最后一个下标的位置是 length - 1
2) A slice is declared just like an array, but it doesn’t contain the size of the slice. So it can grow or shrink according to the requirement.
   slice 的定义跟数组大致一样，但是在 slice 的定义中不包含序列的大小（容量），因此它可以根据需要增加或者缩减它的大小（容量）


2. 定义语法

---
```go
// T是slice中元素的类型
[]T
[]T{}
[]T{value1, value2, value3, ....valuen}
```
---
3. 组成要素

1) pointer（指针）、Length（长度）、Capacity（容量）

---
```go
package main

import "fmt"

func main() {
  arr := [7]string{"this", "is", "the", "tutorial", "of", "go", "Go", "language"}
  fmt.Println("Array: ", arr)

  mySlice := arr[1:6]
  fmt.Println("Slice: ", mySlice)

  fmt.Printf("Length: %d", len(mySlice))
  fmt.Printf("\nCapacity: %d", cap(mySlice))
}
```
---
2）slice 结构示意图
![](./img/slice%E7%BB%93%E6%9E%84%E7%A4%BA%E6%84%8F%E5%9B%BE.jpeg)


4. 创建和初始化

1) 当通过字面量创建 slice 时，go 会先创建一个数组，然后返回一个指向 slice 的指针

---
```go
// 通过字面量创建，不需要指定slice的大小
var my_slice_1 := []string{"Geeks", "For", "Geeks"}
```
---

---
```go
package mian

import "fmt"

func main() {
  // 使用var关键字
  var my_slice_1 := []string{"Geeks", "For", "Geeks"}
  fmt.Println("my slice 1: ", my_slice_1)

  my_slice_2 := []int{1,2,3,4,5,6,7,8,9}
  fmt.Println("my slice 2: ", my_slice_2)
}
```
---

2. 通过指定的数组创建

- slice 是指向数组的，因此 slice 可以通过数组创建
- 默认最小下标是 0，最大下标是给定数组元素的数量大小

---
```go
array_name[low:high]
```
---
---
```go
package main

import "fmt"

func main() {
  arr := [4]string{"Geeks", "For", "Geeks", "GFG"}

  var slice_1 := arr[1:2]
  slice_2 := arr[0:]
  slice_3 := arr[:2]
  slice_4 := arr[:]

  fmt.Println("Array: ", arr)
  fmt.Println("slice1: ", slice_1)
  fmt.Println("slice2: ", slice_2)
  fmt.Println("slice3: ", slice_3)
  fmt.Println("slice4: ", slice_4)
}
```
---

3. 通过指定 slice 创建

- 最小下标默认是 0，最大小标是给定 slice 的最大元素的数量

---
```go
slice_name[low:high]
```
---
---
```go
package main

import "fmt"

func main() {
  origin_slice := []int{90, 60, 40, 50, 34, 49, 30}

  var my_slice_1 = origin_slice[1:5]
  my_slice_2 = origin_slice[0:]
  my_slice_3 = origin_slice[:6]
  my_slice_4 = origin_slice[:]
  my_slice_5 = origin_slice[2:4]

  fmt.Println("origin_slice", origin_slice)
  fmt.Println("my_slice_1", my_slice_1)
  fmt.Println("my_slice_2", my_slice_2)
  fmt.Println("my_slice_3", my_slice_3)
  fmt.Println("my_slice_4", my_slice_4)
  fmt.Println("my_slice_5", my_slice_5)
}
```
---

4. 使用 make 函数

---
```go
func make([]T, len, cap) []T
```
---

---
```go
package main

import "fmt"

func main() {

  // 创建一个大小为7的数组，对这个数组进行切片知道下标为4
  // 返回这个slice的指针
  var my_slice_1 := make([]int, 4, 7)
  fmt.Printf("Slice 1 = %v, \nlength: %d, \ncapacity: %d\n", my_slice_1, len(my_slice_1), cap(my_slice_1))

  // 创建一个大小为7的数组
  // 返回他的slice的指针
  var my_slice_2 := make([]int, 7)
  fmt.Printf("Slice2 = %v, \nlength: %d, \ncapacity: %d\n", my_slice_2, len(my_slice_2, cap(my_slice_2)))
}
```
---

5）如何遍历一个 slice

- 使用 for 循环

---
```go
package main

import "fmt"

func main() {

  mySlice := []string{"This", "is", "the", "tutorial", "of", "go", "languange"}

  for e := 0; e < len(mySlice); e++ {
    fmt.Println(mySlice[e])
  }
}
```
---
- 在 for 循环中是 range 方法

---
```go
package main

import "fmt"

func main() {

  myslice := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}
  for index, ele := range mySlice {
    fmt.Printf("index: %d and element: %s\n", index, ele)
  }

}
```
---

- 在 for 循环中使用空白标识符(\_)

---
```go
package main

import "fmt"

func main() {
  myslice := []string{"This", "is", "the",
        "tutorial", "of", "Go", "language"}

  for _, ele := range myslice {
    fmt.Printf("Element: %s\n", ele)
  }
}
```
---

6. 注意
   在 go 语言中，你可以创建一个值为 nil 的 slice，这个 slice 不包含任何元素；
   并且这个 slice 的容量和长度都为 0，值为 nil 的 slice 没有指向数组的指针。

---
```go
package main

import "fmt"

func main() {

  var mySlice []string
  fmt.Println("length: %d", len(mySlice))
  fmt.Println("capacity: %d", cap(mySlice))
}
```
---

7. 修改 slice

- 正如我们所知，slice 是指针类型，当修改 slice 的元素值时，它指向的数组的值也会相应的被修改

---
```go
package main

import "fmt"
func main() {
  arr := [6]int{55, 66, 77, 88, 99, 22}
  slc := arr[0:4]

  fmt.Println("Origin Array", arr);
  fmt.Println("Origin Slice", slc);

  slc[0] = 100
  slc[1] = 1000
  slc[2] = 1000


  fmt.Println("newArray", arr);
  fmt.Println("newSlice", slc);
}
```
---

8. slice 之间的比较, 使用“==”操作符和 DeepEqual 函数

---
```go
package main

import (
  "fmt"
  "refelect"
)

func main() {
  s1 := []int{12, 34, 56}
  var s2 []int

  fmt.Println(s1 == nil)
  fmt.Println(s2 == nil)

  s2 = []int{12, 34, 56}
  fmt.Println(reflect.DeepEqual(s1, s2))
}
```
---
- 反射（reflect）是指程序在运行时可以访问、检测、修改它本身状态或者行为的一种能力。
  用比喻来说，反射就是程序能够在运行的时候“观察”并修改自己的行为
  https://segmentfault.com/a/1190000021401057


9. 多维 slice

- 多维 slice 与多维数组类似，但是初始化时不包含大小

---
```go
package main

import "fmt"

func main() {
  s1 := [][]int{{12,13}, {14, 15}, {16, 17}}
  fmt.Println(s1)

  s2 := [][]string{[]string{"Geeks", "For"}, []string{"Geeks", "GFG"}, []string{"gfg", "geek"}}
  fmt.Println(s2)
}
```
---

10. slice 排序

- go 语言的标准库中提供了 sort 包，提供了多种不同类型的方法用于排序

---
```go
package main

import (
  "fmt"
  "sort"
)
func main() {
  slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
  slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

  fmt.Println("Before sorting:")
  fmt.Println("Slice 1: ", slc1)
  fmt.Println("Slice 2: ", slc2)

  sort.Strings(slc1)
  sort.Ints(slc2)

  fmt.Println("\nAfter sorting:")
  fmt.Println("Slice 1: ", slc1)
  fmt.Println("Slice 2: ", slc2)
}
```
---
