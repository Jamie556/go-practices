
## 数组

### 数组的创建

```go
Var array_name[length]Type
array_name:= [length]Type{item1, item2, item3,...itemN}
```

### 多维数组

```go
Array_name[Length1][Length2]..[LengthN]Type
```

### 数组复制

```go
// 通过值复制
// creating a copy of an array by value
arr := arr1

// 通过指针复制
// Creating a copy of an array by reference
arr := &arr1
```

### 数组作为函数的参数

```go
// For sized array
func function_name(variable_name [size]type){
// Code
}
```


