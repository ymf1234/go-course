package day04

import (
	"errors"
	"fmt"
	"testing"
)

// 函数的基本形式
func argf(a int, b int) {
	a = a + b
}

// - 形参是函数内部的局部变量，实参的值会拷贝给形参。
// - 函数定义时的第一个的大括号不能另起一行。
// - 形参可以有0个或多个。
// - 参数类型相同时可以只写一次，比如argf(a,b int)。
// - 在函数内部修改形参的值，实参的值不受影响。
// - 如果想通过函数修改实参，就需要指针类型。
func argf1(a, b *int) {
	*a = *a + *b
	*b = 888
}

// slice、map、channel都是引用类型，它们作为函数参数时其实跟普通struct没什么区别，都是对struct内部的各个字段做一次拷贝传到函数内部。
func sliceArg1(arr []int) { // slice作为参数，实际上是把slice的arrayPointer、len、cap拷贝了一份进来
	arr[0] = 1           // 修改底层数据里的首元素
	arr = append(arr, 1) // arr的len和cap发生了变化，不会影响实参
}

// 关于函数返回值
// - 可以返回0个或多个参数。
// - 可以在func行直接声明要返回的变量。
// - return后面的语句不会执行。
// - 无返回参数时return可以不写。
func returnFunc(a, b int) (c int) { // 返回变量c已经声明好了
	a = a + b
	c = a  // 直接使用c
	return // 由于函数要求有返回值，即使给c赋过值了，也需要显示写return
}

// 不定长参数实际上是slice类型。
func variableEngthArg(a int, other ...int) int {
	sum := a
	for _, ele := range other {
		sum += ele
	}
	fmt.Printf("len %d cap %d \n", len(other), cap(other))
	return sum
}

// 递归函数
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n // 凡是递归，一定要有终止条件，否则会进入无限循环
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 函数的基本形式
func TestFunction(t *testing.T) {
	var x, y int = 3, 6
	argf(x, y) // 函数调用。 x,y是实参

	argf1(&x, &y)
	fmt.Printf("x=%d  y=%d \n", x, y)

	arr := []int{8}
	sliceArg1(arr)
	fmt.Println(arr[0])   //1
	fmt.Println(len(arr)) //1

	fmt.Printf("returnFunc %d\n", returnFunc(x, y))

	variableEngthArg(1)
	variableEngthArg(1, 2, 3, 4)

	// append函数接收的就是不定长参数。
	fmt.Println("append函数接收的就是不定长参数。")

	arr = append(arr, 1, 2, 3)
	fmt.Printf("arr1 %v\n", arr)
	arr = append(arr, 7)
	fmt.Printf("arr2 %v\n", arr)
	arr = append(arr)
	fmt.Printf("arr3 %v\n", arr)

	slice := append([]byte("hello"), "world"...) // ...自动把"world"转出byte切片，等价与[]byte("world")...
	fmt.Printf("slice %v \n", slice)
	fmt.Printf("string(slice) %v \n", string(slice))
	slice2 := append([]rune("hello"), []rune("world")...) // 需要显示把"world" 转成rune切片
	fmt.Printf("slice2 %v \n", slice2)
	fmt.Printf("string(slice2) %v \n", string(slice2))

	fmt.Println("递归函数")
	fmt.Printf("Fibonacci(): %d\n", Fibonacci(10))

}

// 匿名函数
func functionArg1(f func(a, b int) int, b int) int { //f参数是一种函数类型
	a := 2 * b
	return f(a, b)
}

type foo func(a, b int) int //foo是一种函数类型

func functionArg2(f foo, b int) int { //参数类型看上去简洁多了
	a := 2 * b
	return f(a, b)
}

type User struct {
	Name  string
	bye   foo                      //bye的类型是foo，而foo代表一种函数类型
	hello func(name string) string //使用匿名函数来声明struct字段的类型
}

func TestAnonymousFunction(t *testing.T) {
	ch := make(chan func(string) string, 10)
	ch <- func(name string) string { //使用匿名函数
		return "hello " + name
	}
	getCh := <-ch
	fmt.Println(getCh("getCH"))
}

// 闭包
// 闭包（Closure）是引用了自由变量的函数，自由变量将和函数一同存在，即使已经离开了创造它的环境。闭包复制的是原对象的指针。
func sub() func() {
	i := 10
	fmt.Printf("%p \n", &i)
	b := func() {
		fmt.Printf("i addr %p \n", &i) // 闭环复制的是原对象的指针
		i--                            // b函数内部引用了变量i
		fmt.Println(i)
	}
	fmt.Printf("b 函数外的i: %d \n", i)
	return b //返回了b函数，变量i和b函数将一起存在，即使已经离开函数sub()
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("base addr %p \n", &base)
		base += i
		return base
	}
}
func TestClosure(t *testing.T) {
	b := sub()
	b()
	b()
	fmt.Println()

	fmt.Println("tmp1")
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2)) // 11, 12

	fmt.Println("tmp2")
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2)) //101, 103
}

// 延迟调用defer
// - defer用于注册一个延迟调用（在函数返回之前调用）。
// - defer典型的应用场景是释放资源，比如关闭文件句柄，释放数据库连接等。
// - 如果同一个函数里有多个defer，则后注册的先执行。
// - defer后可以跟一个func，func内部如果发生panic，会把panic暂时搁置，当把其他defer执行完之后再来执行这个。
// - defer后不是跟func，而直接跟一条执行语句，则相关变量在注册defer时被拷贝或计算。

func basic() {
	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")
	// 如果一个函数里有多个defer, 则后注册的先执行
	defer fmt.Println(2)
	fmt.Println("C")

}

func deferExeTime() (i int) {
	i = 9
	defer func() { // defer 后可以跟一个func
		fmt.Printf("first i = %d \n", i) //打印5，而非9。充分理解“defer在函数返回前执行”的含义，不是在“return语句前执行defer”
	}()

	defer func(i int) {
		fmt.Printf("second i=%d\n", i) //打印9
	}(i)

	defer fmt.Printf("third i=%d\n", i) //defer后不是跟func，而直接跟一条执行语句，则相关变量在注册defer时被拷贝或计算

	return 5
}
func TestDefer(t *testing.T) {
	basic()

	deferExeTime()
}

// 异常处理
// go语言没有try catch，它提倡返回error。

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("divide by zero")
	}
	return a / b, nil
}

func TestError(t *testing.T) {
	if res, err := divide(3, 0); err != nil { //函数调用方判断error是否为nil
		fmt.Println(err.Error())
		fmt.Println(res)
	}
}
