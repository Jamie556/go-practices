package main

import (
	"fmt"
	"reflect"
	"sort"
)

func testSlice() {

	arr := [7]string{ "this", "is", "the", "tutorial", "of", "go", "language" }
	fmt.Println("Array:", arr)

	mySlice := arr[1:6]
	fmt.Println("Slice:", mySlice)
	fmt.Printf("Length of slice is: %d", len(mySlice))
	fmt.Printf("\nCapacity of slice is: %d\n", cap(mySlice))

	mySlice2 := []string{"This", "is", "the", "tutorial", "of", "go", "languange"}

  for e := 0; e < len(mySlice2); e++ {
    fmt.Println(mySlice2[e])
  }


	mySlice3 := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}
  for index, ele := range mySlice3 {
    fmt.Printf("index: %d and element: %s\n", index, ele)
  }

	mySlice4 := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}

	for _, ele := range mySlice4 {
		fmt.Printf("Element: %s\n", ele)
	}

	var mySlice5 []string
  fmt.Printf("length: %d\n", len(mySlice5))
  fmt.Printf("capacity: %d\n", cap(mySlice5))


	arr2 := [6]int{55, 66, 77, 88, 99, 22}
  slc := arr2[0:4]

  fmt.Println("Origin Array", arr2);
  fmt.Println("Origin Slice", slc);

  slc[0] = 100
  slc[1] = 1000
  slc[2] = 1000


  fmt.Println("newArray", arr2);
  fmt.Println("newSlice", slc);

	s1 := []int{12, 34, 56}
  var s2 []int

  fmt.Println(s1 == nil)
  fmt.Println(s2 == nil)

	s2 = []int{12, 34, 56}
  fmt.Println(reflect.DeepEqual(s1, s2))

	s3 := [][]int{{12,13}, {14, 15}, {16, 17}}
  fmt.Println(s3)

  s4 := [][]string{[]string{"Geeks", "For"}, []string{"Geeks", "GFG"}, []string{"gfg", "geek"}}
  fmt.Println(s4)


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
