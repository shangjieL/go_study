package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine退出子就不执行了
func main() {
	//创建一个go程 去执行newTask() 流程
	go newTask()

	fmt.Println("main goroutine exit")

	i := 0
	for {
		i++
		fmt.Printf("main goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
