package day03

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// if
func TestIf(t *testing.T) {
	if 5 > 9 {
		fmt.Printf("5 > 9")
	}

	//- 如果逻辑表达式成立，就会执行{}里的内容。
	//- 逻辑表达式不需要加()。
	//- "{"必须紧跟在逻辑表达式后面，不能另起一行。
	if c, d, e := 5, 6, 1; c < d && (c > e || c > 3) { // 初始化多个局部变量。复杂的逻辑表达式
		fmt.Println("fit")
	}

	// - 逻辑表达中可以含有变量或常量。
	//- if句子中允许包含1个(仅1个)分号，在分号前初始化一些局部变量(即只在if块内可见)。
	// if-else的用法
	//

	color := "black"
	if color == "red" { //if只能有一个
		fmt.Println("stop")
	} else if color == "green" {
		fmt.Println("go")
	} else if color == "yellow" { //else if可以有0个、一个或者连续多个
		fmt.Println("stop")
	} else { //else有0个或1个
		fmt.Printf("invalid traffic signal: %s\n", strings.ToUpper(color))
	}
}

// switch
func TestSwitch(t *testing.T) {
	color := "black"
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	default:
		fmt.Printf("invalid traffic signal: %s\n", strings.ToUpper(color))
	}

	// switch type
	SwitchType()

	// fallthrough
	fallThrough(30)
}

// switch type
func SwitchType() {
	var num interface{} = 6.5
	switch num.(type) { //获取interface的具体类型。.(type)只能用在switch后面
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Println("neither")
	}

	switch value := num.(type) { //相当于在每个case内部申明了一个变量value
	case int: //value已被转换为int类型
		fmt.Printf("number is int %d\n", value)
	case float64: //value已被转换为float64类型
		fmt.Printf("number is float64 %f\n", value)
	case byte, string: //如果case后有多个类型，则value还是interface{}类型
		fmt.Printf("number is inerface %v\n", value)
	default:
		fmt.Println("neither")
	}

	//等价形式
	switch num.(type) {
	case int:
		value := num.(int)
		fmt.Printf("number is int %d\n", value)
	case float64:
		value := num.(float64)
		fmt.Printf("number is float64 %f\n", value)
	case byte:
		value := num.(byte)
		fmt.Printf("number is byte %d\n", value)
	default:
		fmt.Println("neither")
	}

}

// fallthrough，当命中某一个case时，强行进入下一个case。
func fallThrough(age int) {
	fmt.Printf("您的年龄是%d, 您可以：\n", age)
	switch {
	case age > 50:
		fmt.Println("出任国家首脑")
		fallthrough
	case age > 25:
		fmt.Println("生育子女")
		fallthrough
	case age > 22:
		fmt.Println("结婚")
		fallthrough
	case age > 38:
		fmt.Println("开车")
		fallthrough
	case age > 16:
		fmt.Println("参加工作")
	case age > 15:
		fmt.Println("上高中")
		fallthrough
	case age > 3:
		fmt.Println("上幼儿园")
	}
}

// for
func TestFor(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ { // 正序遍历切片
		fmt.Printf("%d: %d \n", i, arr[i])
	}

	nest_for()

	complex_break_continue()
}

func nest_for() {
	const SIZE = 4

	A := [SIZE][SIZE]float64{}
	//初始化二维数组
	//两层for循环嵌套
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			A[i][j] = rand.Float64() //[0,1)上的随机数
		}
	}

	B := [SIZE][SIZE]float64{}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			B[i][j] = rand.Float64() //[0,1)上的随机数
		}
	}

	rect := [SIZE][SIZE]float64{}
	//三层for循环嵌套
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			prod := 0.0
			for k := 0; k < SIZE; k++ {
				prod += A[i][k] * B[k][j]
			}
			rect[i][j] = prod
		}
	}

	i, j := 2, 1
	fmt.Println(A[i]) //二维数组第i行
	//打印二维数组的第j列
	//注意：B[:][j]这不是二维数组第j列，这是二维数组第j行！
	for _, row := range B {
		fmt.Printf("%g ", row[j])
	}
	fmt.Println()
	fmt.Println(rect[i][j])
}

// break与continue
// break和continue都是针对for循环的，不针对if或switch
// break和continue都是针对套在自己外面的最靠里的那层for循环，不针对更外层的for循环（除非使用Label）
func complex_break_continue() {
	const SIZE = 5
	arr := [SIZE][SIZE]int{}
	for i := 0; i < SIZE; i++ {
		fmt.Printf("开始检查第%d行\n", i)
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if arr[i][j]%2 == 0 {
					continue //针对第二层for循环
				}
				fmt.Printf("将要检查第%d列\n", j+1)
			}
			break //针对第一层for循环
		}
	}
}

// goto与Label
func TestGotoAndLabel(t *testing.T) {
	var i int = 4
	/*MY_LABEL:
	i += 3
	fmt.Println(i)
	goto MY_LABEL //返回定义MY_LABEL的那一行，把代码再执行一遍（会进入一个无限循环）*/
	fmt.Println(i % 2)
	if i%2 == 0 {
		goto L1 //Label指示的是某一行代码，并没有圈定一个代码块，所以goto L1也会执行L2后的代码
	} else {
		goto L2 //先使用Label
	}

L1:
	i += 3
L2: //后定义Label。Label定义后必须在代码的某个地方被使用
	i *= 3

	fmt.Println(i)
}
