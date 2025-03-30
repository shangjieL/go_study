package main

import (
	//1._ 为匿名方式导包，只调用init方法 无法调用其他方法
	//不然编译不过
	_ "go_study/5-init/lib1"

	//2.别名导包 调方法时可直接使用别名
	myli2 "go_study/5-init/lib2"

	//3 .方式 将当前包的方法导入到本包下 全部方法可以直接使用 不需要lib2.api来调用
	// 不推荐使用，防止多包内方法名重复
	//. "go_study/5-init/lib2"
	
)
func main() {
	//方法大写为公用的，小写是私用的只能用在当前类
	//lib1.Lib1Test()
	myli2.Lib2Test()
}