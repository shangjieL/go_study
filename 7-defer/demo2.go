/*
知识点二：defer和return谁前谁后
*/
package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDerfer() int {
	//return 先执行
	//defer 后执行
	defer deferFunc()

	return returnFunc()
}

func main() {
	returnAndDerfer()
}
