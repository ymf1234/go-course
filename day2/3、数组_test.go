package day2

import (
	"fmt"
	"testing"
)

// TestAnArrayOf 数组
// 数组是块连续的内存空间，在声明的时候必须指定长度，且长度不能改变。所以数组在声明的时候就可以把内存空间分配好，并赋上默认值，即完成了初始化。
func TestAnArrayOf(t *testing.T) {
	// 一维数组初始化
	var arr1 [5]int = [5]int{} // 数组必须指定长度和类型，且长度和类型指定后不可变
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}            // 给前2个元素赋值
	var arr4 = [5]int{2: 15, 4: 30}    // 指定index赋值
	var arr5 = [...]int{3, 2, 6, 5, 4} //根据{}里元素的个数推断出数组的长度
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}} //数组的元素类型由匿名结构体给定

	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)
	fmt.Println("arr4: ", arr4)

	// 二维数组初始化
	//5行3列，只给前2行赋值，且前2行的所有列还没有赋满
	var arr7 = [5][3]int{{1}, {2, 3}}
	//第1维可以用...推测，第2维不能用...
	var arr8 = [...][3]int{{1}, {2, 3}}

}
