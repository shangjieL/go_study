/*
	知识点一：defer的执行顺序
*/
package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	//先进后出 压栈
	defer func1()
	defer func2()
	defer func3()
}