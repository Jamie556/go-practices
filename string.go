package main

import (
	"fmt"
	"unicode/utf8"
)

func testString() {

	my_s_values := "Welecome to GeeksForGeeks"
	var my_value_5 string
	my_value_5 = "GeeksForGeeks"

	fmt.Println("string1: ", my_s_values)
	fmt.Println("string2: ", my_value_5)

	my_value_1 := "Welecome!\n GeeksForGeeks"
	my_value_2 := "Welecome to GeeksForGeeks"

	my_value_3 := `Hello GeeksForGeeks`
	my_value_4 := `Hello!\nGeeksForGeeks`

	fmt.Println("value1: ", my_value_1)
	fmt.Println("value2: ", my_value_2)
	fmt.Println("value3: ", my_value_3)
	fmt.Println("value4: ", my_value_4)

	mystr := "Welecome to GeeksForGeeks"
	fmt.Println("myStr: ", mystr)

	fmt.Println(mystr[1])

	// 使用range和For表达式进行循环迭代
	for index, s := range "GeeksForGeeks" {
		fmt.Printf("The Index Number of %c is %d\n", s, index)
	}

	// 获取字符串的bytes
	str  := "Welecome To GeeksForGeeks"
	for c := 0; c < len(str); c++ {
		fmt.Printf("\nCharacter = %c Bytes = %v\n", str[c], str[c])
	}

	// 通过slice创建字符串
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	myString := string(myslice1)
	fmt.Println("String 1: ", myString)

	myslice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	myString2 := string(myslice2)

	fmt.Println("String 2: ", myString2)

	// 统计字符串长度
	myStr := "Welecome to GeeksForGeeks 学习Go语言 ?????"
	// len函数返回字符串的字节数
	length1 := len(myStr)
	// utf8.RuneCountInString 返回字符总数
	length2 := utf8.RuneCountInString(myStr)
	fmt.Println("string", myStr)
	fmt.Println("length 1", length1)
	fmt.Println("length 2", length2)
	fmt.Println("length1 is equal length2:", length1 == length2)
}

