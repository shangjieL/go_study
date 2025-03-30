package main

import "fmt"

func main() {
	//固定长度的数组
	var myArray1 [10]int

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10] int {1,2,3,4}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}
}
