package main

/*
	四种变量的声明方式
*/

import "fmt"

//声明全局变量 方法一、方法二、方法三是可以的
var gA int = 100
var gB = 200

//用方法四来声明全局变量
//:= 只能够 用在函数体内声明
//gC := 300

func main() {
	//方法一：声明一个变量 默认值是0
	var a int
	fmt.Println("a =", a)

	//打印变量的类型
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量并赋值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 1000
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：(常用的方法) 省去var关键字，直接自动匹配
	//e 冒得 100
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f =", f)
	fmt.Printf("type of f = %T\n", f)

	//float 8字节浮点类型
	g := 3.14
	fmt.Println("g =", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA = ", gA, "gB = ", gB)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk, "ll = ", ll)

	//多行的多变量声明
	var (
		vv = 100
		jj = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
