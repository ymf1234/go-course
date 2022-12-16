package day01

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// | 运算符 | 描述 |
// | :----: | :--: |
// |   +    | 相加 |
// |   -    | 相减 |
// |   *    | 相乘 |
// |   /    | 相除 |
// |   %    | 求余 |
// arithmetic 算术运算
func TestArithmetic(t *testing.T) {
	type arithmeticTest struct {
		numericalValue1 float32
		numericalValue2 float32
		arithmetic      func(a, b float32) float32
		symbol          string
		result          float32
	}
	var testData = []arithmeticTest{
		{8, 3, Add, "+", 11},
		{8, 3, Deductions, "-", 5},
		{8, 3, Punishment, "*", 24},
		{8, 3, Division, "/", 2.666667},
	}

	for _, tt := range testData {
		result := tt.arithmetic(tt.numericalValue1, tt.numericalValue2)
		fmt.Printf("%f %s %f, except:%f, actual:%f \n", tt.numericalValue1, tt.symbol, tt.numericalValue2, tt.result, result)
	}

}

// Add 加
func Add(a, b float32) float32 {
	return a + b
}

// Deductions 减
func Deductions(a, b float32) float32 {
	return a - b
}

// Punishment 乘
func Punishment(a, b float32) float32 {
	return a * b
}

// Division 除
func Division(a, b float32) float32 {
	return a / b
}

// relational 关系运算符
// | 运算符 | 描述                                                         |
// | :----: | :----------------------------------------------------------- |
// |   ==   | 检查两个值是否相等，如果相等返回 True 否则返回 False         |
// |   !=   | 检查两个值是否不相等，如果不相等返回 True 否则返回 False     |
// |   >    | 检查左边值是否大于右边值，如果是返回 True 否则返回 False     |
// |   >=   | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False |
// |   <    | 检查左边值是否小于右边值，如果是返回 True 否则返回 False     |
// |   <=   | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False |
func TestRelational(t *testing.T) {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8

	fmt.Printf("a==b  %t\n", a == b)
	fmt.Printf("a!=b  %t\n", a != b)
	fmt.Printf("a>b  %t\n", a > b)
	fmt.Printf("a>=b  %t\n", a >= b)
	fmt.Printf("a<c  %t\n", a < b)
	fmt.Printf("a<=c  %t\n", a <= c)
}

// 逻辑运算符
// | 运算符 | 描述                                                         |
// | :----: | :----------------------------------------------------------- |
// |   &    | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False |
// |  \|\|  | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False |
// |   !    | 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True   |
func TestLogistic(t *testing.T) {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a>b && b>c %t\n", a > b && b > c)
	fmt.Printf("a>b || b>c %t\n", a > b || b > c)
	fmt.Printf("!(a>b)  %t\n", !(a > b))
	fmt.Printf("!(b>c)  %t\n", !(b > c))
}

// 位运算符
// | 运算符 | 描述                                                         |
// |   &    | 参与运算的两数各对应的二进位相与（两位均为1才为1）           |
// |   \|   | 参与运算的两数各对应的二进位相或（两位有一个为1就为1）       |
// |   ^    | 参与运算的两数各对应的二进位相异或，当两对应的二进位相同时为0，不同时为1。作为一元运算符时表示按位取反，，符号位也跟着变 |
// |   <<   | 左移n位就是乘以2的n次方。a<<b是把a的各二进位全部左移b位，高位丢弃，低位补0。通过左移，符号位可能会变 |
// |   >>   | 右移n位就是除以2的n次方。a>>b是把a的各二进位全部右移b位，正数高位补0，负数高位补1 |
func TestBitOp(t *testing.T) {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节，取决于操作系统是32位还是64位
	var a int32 = 260
	fmt.Printf("260     %s\n", BinaryFormat(a))
	fmt.Printf("-260    %s\n", BinaryFormat(-a)) //负数用补码表示。在对应正数二进制表示的基础上，按拉取反，再末位加1
	fmt.Printf("260&4   %s\n", BinaryFormat(a&4))
	fmt.Printf("260|3   %s\n", BinaryFormat(a|3))
	fmt.Printf("260^7   %s\n", BinaryFormat(a^7))     //^作为二元运算符时表示异或
	fmt.Printf("^-260   %s\n", BinaryFormat(^-a))     //^作为一元运算符时表示按位取反，符号位也跟着变
	fmt.Printf("-260>>10 %s\n", BinaryFormat(-a>>10)) //正数高位补0，负数高位补1
	fmt.Printf("-260<<3 %s\n", BinaryFormat(-a<<3))   //负数左移，可能变成正数
	//go语言没有循环（无符号）左/右移符号   >>>  <<<
}

// 输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2, 31)) //最高位上是1，其他位全是0
	for i := 0; i < 32; i++ {
		if a&c != 0 { //判断n的当前位上是否为1
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 //"1"往右移一位
	}
	return sb.String()
}

// 赋值运算符
// | 运算符 | 描述                                           |
// |   =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
// |   +=   | 相加后再赋值                                   |
// |   -=   | 相减后再赋值                                   |
// |   *=   | 相乘后再赋值                                   |
// |   /=   | 相除后再赋值                                   |
// |   %=   | 求余后再赋值                                   |
// |  <<=   | 左移后赋值                                     |
// |  >>=   | 右移后赋值                                     |
// |   &=   | 按位与后赋值                                   |
// |  \|=   | 按位或后赋值                                   |
// |   ^=   | 按位异或后赋值                                 |
// assignment 赋值运算
func TestAssignment(t *testing.T) {
	var a, b int = 8, 3
	a += b
	fmt.Printf("a+=b %d\n", a)
	a, b = 8, 3
	a -= b
	fmt.Printf("a-=b %d\n", a)
	a, b = 8, 3
	a *= b
	fmt.Printf("a*=b %d\n", a)
	a, b = 8, 3
	a /= b
	fmt.Printf("a/=b %d\n", a)
	a, b = 8, 3
	a %= b
	fmt.Printf("a%%=b %d\n", a) //%在fmt里有特殊含意，所以需要前面再加个%转义一下
	a, b = 8, 3
	a <<= b
	fmt.Printf("a<<=b %d\n", a)
	a, b = 8, 3
	a >>= b
	fmt.Printf("a>>=b %d\n", a)
	a, b = 8, 3
	a &= b
	fmt.Printf("a&=b %d\n", a)
	a, b = 8, 3
	a |= b
	fmt.Printf("a|=b %d\n", a)
	a, b = 8, 3
	a ^= b
	fmt.Printf("a^=b %d\n", a)
}
