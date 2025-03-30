package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1， number len = 4, [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2， number len = 5, [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的追加元素 cap会在原来基础上*2
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("========")
	numbers2 := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}