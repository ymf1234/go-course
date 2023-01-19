package day02

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestTransformation(t *testing.T) {
	// 强制类型转换的基本方法就是把目标类型放在变量前面，把变量括起来。
	var i int = 9
	var by byte = byte(i) //int转为byte
	fmt.Println("int转为byte：", by, reflect.ValueOf(by).Kind(), reflect.TypeOf(by))
	i = int(by) //byte转为int
	fmt.Println("byte转为int：", i, reflect.ValueOf(i).Kind(), reflect.TypeOf(i))

	//- 低精度向高精度转换没问题，高精度向低精度转换会丢失位数。
	//- 无符号向有符号转换，最高位是符号位。
	//- byte和int可以互相转换。
	//- float和int可以互相转换，小数位会丢失。
	//- bool和int不能相互转换。
	//- 不同长度的int或float之间可以相互转换。

	//高精度向低精度转换，数字很小时这种转换没问题
	fmt.Println("高精度向低精度转换，数字很小时这种转换没问题")
	var ua uint64 = 1
	i8 := int8(ua)
	fmt.Printf("i8=%d  - 数据类型%T\n", i8, i8)
	fmt.Printf("ua=%d  - 数据类型%T\n", ua, ua)

	// 最高位的1变成了符号位
	fmt.Println("最高位的1变成了符号位")
	ua = uint64(math.MaxUint64)
	i64 := int64(ua)
	fmt.Printf("i64=%d  - 数据类型%T\n", i64, i64)
	fmt.Printf("ua=%d  - 数据类型%T\n", ua, ua)

	// 位数丢失
	fmt.Println()
	fmt.Println("位数丢失")
	ui32 := uint32(ua)
	fmt.Printf("ui32=%d  - 数据类型%T\n", ui32, ui32)
	fmt.Printf("ua=%d  - 数据类型%T\n", ua, ua)

	// 单个字符可以转为int
	fmt.Println()
	fmt.Println("单个字符可以转为int")
	var j int = int('a')
	fmt.Printf("j=%d - 数据类型%T \n", j, j)

	//bool和int不能相互转换

	//byte和int可以互相转换
	var by1 byte = byte(i)
	i = int(by1)
	fmt.Printf("i=%d\n", i)

	//float和int可以互相转换，小数位会丢失
	var ft float32 = float32(i)
	i = int(ft)
	fmt.Printf("i=%d\n", i)

	// string和其他数据类型互转。
	var err error
	var i1 int = 8
	var i64i int64 = int64(i1)
	//int转string
	var s string = strconv.Itoa(i1) //内部调用FormatInt
	s = strconv.FormatInt(i64i, 10)
	//string转int
	i, err = strconv.Atoi(s)
	//string转int64
	i64i, err = strconv.ParseInt(s, 10, 64)

	//float转string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) //保留2位小数
	fmt.Println(s)
	//string转float
	f, err = strconv.ParseFloat(s, 64)

	//string<-->[]byte
	var arr []byte = []byte(s)
	s = string(arr)

	//string<-->[]rune
	var brr []rune = []rune(s)
	s = string(brr)

	fmt.Printf("err %v\n", err)

}
