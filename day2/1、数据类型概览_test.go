package day2

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"unsafe"
)

// 基础数据类型
// |     类型      | 长度(字节) | 默认值 |                   说明                    |
// | :-----------: | :--------: | :----: | :---------------------------------------: |
// |     bool      |     1      | false  |                                           |
// |     byte      |     1      |   0    |          uint8，取值范围[0,255]           |
// |     rune      |     4      |   0    |         Unicode Code Point, int32         |
// |   int, uint   |    4或8    |   0    |        32 或 64 位，取决于操作系统        |
// |  int8, uint8  |     1      |   0    |            -128 ~ 127, 0 ~ 255            |
// | int16, uint16 |     2      |   0    |         -32768 ~ 32767, 0 ~ 65535         |
// | int32, uint32 |     4      |   0    | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名 |
// | int64, uint64 |     8      |   0    |                                           |
// |    float32    |     4      |  0.0   |                                           |
// |    float64    |     8      |  0.0   |                                           |
// |   complex64   |     8      |        |                                           |
// |  complex128   |     16     |        |                                           |
// |    uintptr    |    4或8    |        |    以存储指针的 uint32 或 uint64 整数     |

// 复合数据类型
// |   类型    |             默认值             |     说明     |
// | :-------: | :----------------------------: | :----------: |
// |   array   |   取每个元素对应类型的默认值   |    值类型    |
// |  struct   | 取每个成员变量对应类型的默认值 |    值类型    |
// |  string   |               ""               | UTF-8 字符串 |
// |   slice   |              nil               |   引用类型   |
// |    map    |              nil               |   引用类型   |
// |  channel  |              nil               |   引用类型   |
// | interface |              nil               |     接口     |
// | function  |              nil               |     函数     |
func TestBasicDataTypes(tt *testing.T) {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节，取决于操作系统是32位还是64位
	var a int = 5
	var b int8 = 5
	var c int16 = 5
	var d int32 = 5
	var e int64 = 5
	var f uint = 5
	var g uint8 = 5
	var h uint16 = 5
	var i uint32 = 5
	var j uint64 = 5
	fmt.Printf("a=%d, b=%d, c=%d, d=%d, e=%d, f=%d, g=%d, h=%d, i=%d, j=%d\n", a, b, c, d, e, f, g, h, i, j)

	var k float32 = 5
	var l float64 = 5
	fmt.Printf("k=%f, l=%.2f\n", k, l) //%.2f保留2位小数
	var m complex128 = complex(4, 7)
	var n complex64 = complex(4, 7)
	fmt.Printf("type of m is %T, type of n is %T\n", m, n) //%T输出变量类型
	fmt.Printf("m=%v, n=%v\n", m, n)                       //按值的本来值输出
	fmt.Printf("m=%+v, n=%+v\n", m, n)                     //在 %v 基础上，对结构体字段名和值进行展开
	fmt.Printf("m=%#v, n=%#v\n", m, n)                     //输出 Go 语言语法格式的值
	fmt.Printf("m的实部%f, m的虚部%f\n", real(m), imag(m))
	fmt.Printf("m的实部%e, m的虚部%g\n", real(m), imag(m)) //%e科学计数法，%g根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	o := true                                        //等价于var o bool = true
	fmt.Printf("o=%t\n", o)                          //%t布尔变量
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int = &a
	fmt.Printf("p=%x pointer=%p ptr=%p\n", p, pointer, ptr) //%p输出地址，%x十六进制
	var q byte = 100                                        //byte是uint，取值范围[0,255]
	fmt.Printf("q=%d, binary of q is %b\n", q, q)           //%b输出二进制
	var r rune = '☻'                                        //rune实际上是int32，即可以表示2147483647种字符，包括所有汉字和各种特殊符号
	fmt.Printf("r=%d, r=%U\n", r, r)                        //%U Unicode 字符
	var s string = "I'm"
	fmt.Printf("s=%s\n", s)
	var t error = errors.New("my error")
	fmt.Printf("error is %v\n", t)
	fmt.Printf("error is %+v\n", t) //在 %v 基础上，对结构体字段名和值进行展开
	fmt.Printf("error is %#v\n", t) //输出 Go 语言语法格式的值

	// 数值型变量的默认值是0，
	// 字符串的默认值是空字符串
	// 布尔型变量的默认值是false
	// 引用类型、函数、指针、接口的默认值是nil
	// 数组的默认值取每个元素对应类型的默认值
	// 结构体的默认值取每个成员变量对应类型的默认值

	var aa int
	var bb byte
	var ff float32
	var t1 bool
	var ss string
	var rr rune
	var arr [3]int
	var slc []int

	fmt.Printf("default value of int %d\n", aa)
	fmt.Printf("default value of byte %d\n", bb)
	fmt.Printf("default value of float %.2f\n", ff)
	fmt.Printf("default value of bool %t\n", t1)
	fmt.Printf("default value of string [%s]\n", ss)
	fmt.Printf("default value of rune %d, [%c]\n", rr, rr)
	fmt.Printf("default int array is %v\n", arr) //取每个元素对应类型的默认值
	fmt.Printf("default slice is nil %t\n", slc == nil)

}
