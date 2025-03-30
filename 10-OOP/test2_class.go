package main

import "fmt"

//如果类名首字母大写， 表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能类的内部访问
	Name  string
	Ad    int
	level int
}

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Ad)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) setName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝 ）
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, level: 1}

	hero.Show()

	hero.setName("li4")

	hero.Show()
}