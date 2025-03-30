package main

import "fmt"

//父类
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat().....")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk().....")
}

//子类
type SupperMan struct {
	Human
	level int
}

//重定义父类的方法Eat()
func (this *SupperMan) Eat() {
	fmt.Println("SupperMan.Eat()...")
}

func (this *SupperMan) Fly() {
	fmt.Println("SupperMan.Fly()...")
}

func (this *SupperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s := SupperMan{Human{"li4", "female"}, 88}
	var s SupperMan
	s.name = "li4"
	s.sex = "female"
	s.level = 88

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法

	s.Print()

}
