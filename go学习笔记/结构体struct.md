
## 结构体struct

### 基本概念

A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. Any real-world entity which has some set of properties/fields can be represented as a struct. This concept is generally compared with the classes in object-oriented programming. It can be termed as a lightweight class that does not support inheritance but supports composition.

在Go语言中，结构体的定义是将其他的类型组合成为一种新的类型的概念。所有现实世界中的对象实体都具有一些基本的属性，这些属性可以用结构体来表示。这个概念类似于面向对象编程语言中的class。struct可以被称为一个不支持继承但是可以进行属性组合的class

---
```go
// 基本语法
type Address struct {
	name string 
	street string
	city string
	state string
	Pincode int
}
```
---

### 关键字

type
struct

### 初始化（创建实体）

#### 创建一个零值的结构体

```go
var a Address
```

#### 使用结构体字面量创建一个结构体

##### 注意

- 使用字面量创建一个结构体，需要按照结构体定义的属性顺序传入值，也可以用结构体定义的属性的子集
- 结构体也可以使用键值对的形式进行创建

```go
var a = Address("Akshay", "Prenagar", "Dehradun", "Uttarakhand", 252636)
var a = Address{Name:”Akshay”, street:”PremNagar”, state:”Uttarakhand”, Pincode:252636} //city:””
```




### 嵌套结构体（Nested Structure in GoLang）

嵌套结构体是指一个结构体可以作为另一个结构体的字段，这种嵌套形似的结构体叫做嵌套结构体。

```go
// 嵌套结构体的语法语法
type struct_name_1 struct{
  // Fields
} 
type struct_name_2 struct{
  variable_name  struct_name_1
}
```

### 匿名结构体(anymouse structure)

```go
// 语法
variable_name := struct{
// fields
}{// Field_values}
```

### 匿名字段（anymouse field）

```go
type struct_name struct{
    int
    bool
    float64
}

type student struct{
 name int
 price int
 string
}
```
