package main

import "fmt"

//声明一种新的数据类型myInt, 是int的一个别名
type myInt int

//定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book) {
	//传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	//指针传递
	book.auth = "777"
}

func main() {
	/*
	var a myInt = 10
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)
	*/

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}