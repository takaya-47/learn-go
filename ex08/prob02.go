package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

// 構造体とそのメソッド、という形でも同じことができるが、今回はフィールドに相当するものが一つしか（実際の値だけ）扱わないので
// 独自型とそのメソッドという形で実装する。（フィールドが複数必要なら構造体が適切。）
type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

type MyFloat float64

func (mf MyFloat) String() string {
	return fmt.Sprintf("%v", float64(mf))
}

func Print[T Printable](v T) {
	fmt.Println(v.String())
}

func main() {
	var value1 MyInt = 42
	var value2 MyFloat = 3.14

	Print(value1)
	Print(value2)
}
