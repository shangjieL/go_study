package main //程序的包名

import (
	"fmt"
	//"strconv"
)

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

func main() {
	//fmt.Println("hello Go!")
	//time.Sleep(1 * time.Second)

	//字符串 进制  位数
	//ui64, err := strconv.ParseUint(str, 10, 32)

	var a Supplier = &DigitSupplier{value: 1}
	fmt.Println(a.Get())

	b, ok := a.(*DigitSupplier)
	fmt.Println(b, ok)

}
