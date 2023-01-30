package day04

import (
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
