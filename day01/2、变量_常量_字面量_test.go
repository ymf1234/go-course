package day01

import (
	"fmt"
	"testing"
)

// 变量类型
// |   类型   |                         go变量类型                         | fmt输出  |
// | :------: | :--------------------------------------------------------: | :------: |
// |   整型   | int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 |    %d    |
// |  浮点型  |                      float32 float64                       | %f %e %g |
// |  布尔型  |                            bool                            |    %t    |
// |   指针   |                          uintptr                           |    %p    |
// |   引用   |                     map slice channel                      |    %v    |
// |   字节   |                            byte                            |    %c    |
// | 任意字符 |                            rune                            |    %c    |
// |  字符串  |                           string                           |    %s    |
// |   错误   |                           error                            |    %v    |
func TestVariable(t *testing.T) {
	// 标题声明
	/*var name string
	var age int
	var isOk bool*/

	// 批量声明
	/*var (
		name string
		age  int
		isOk bool
	)*/

	//变量初始化
	var name string = "name"
	var age int = 3
	var isOk bool = true

	fmt.Println(name, age, isOk)
}

// 常量
func TestConst(t *testing.T) {
	const PI float32 = 3.14

	const (
		PI1 = 3.14
		E   = 2.71
	)

	const (
		a = 100
		b //100，跟上一行的值相同
		c //100，跟上一行的值相同
	)

	// iota
	const (
		e = iota //0
		f        //1
		_        //2
		g        //3
	)

	const (
		q = iota //0
		w = 30
		y = iota //2
		r        //3
	)

	const (
		_  = iota             // iota =0
		KB = 1 << (10 * iota) // iota =1
		MB = 1 << (10 * iota) // iota =2
		GB = 1 << (10 * iota) // iota =3
		TB = 1 << (10 * iota) // iota =4
	)

	const (
		aa, bb = iota + 1, iota + 2 //1,2  iota =0
		cc, dd                      //2,3  iota =1
		ee, ff                      //3,4  iota =2
	)
}

// 字面量
// 字面量--没有出现变量名，直接出现了值。基础类型的字面量相当于是常量。
// 字面量
func TestLiteral(t *testing.T) {
	fmt.Printf("%t\n", 04 == 4.00)      //用到了整型字面量和浮点型字面量
	fmt.Printf("%v\n", .4i)             //虚数字面量 0.4i
	fmt.Printf("%t\n", '\u4f17' == '众') //Unicode和rune字面量
	fmt.Printf("Hello\nWorld\n!\n")     //字符串字面量
}

// 变量作用域
//  对于全局变量，如果以大写字母开头，所有地方都可以访问，跨package访问时需要带上package名称；如果以小写字母开头，则本package内都可以访问。
//  函数内部的局部变量，仅本函数内可以访问。{}可以固定一个作用域。内部声明的变量可以跟外部声明的变量有冲突，以内部的为准--就近原则。

var (
	A = 3 //所有地方都可以访问
	b = 4 //本package内可以访问
)

func TestFoo(t *testing.T) {
	b := 5 //本函数内可以访问
	{
		b := 6 //本作用域内可以访问
		fmt.Println(b)
	}
	fmt.Println(b)
}
