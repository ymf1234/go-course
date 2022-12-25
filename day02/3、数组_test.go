package day02

import (
	"fmt"
	"testing"
)

// TestAnArrayOf 数组
// 数组是块连续的内存空间，在声明的时候必须指定长度，且长度不能改变。所以数组在声明的时候就可以把内存空间分配好，并赋上默认值，即完成了初始化。
func TestAnArrayOf(t *testing.T) {
	// 一维数组初始化
	var arr1 [5]int = [5]int{} // 数组必须指定长度和类型，且长度和类型指定后不可变
	fmt.Println("arr1: ", arr1)

	var arr2 = [5]int{}
	fmt.Println("arr2: ", arr2)

	var arr3 = [5]int{3, 2} // 给前2个元素赋值
	fmt.Println("arr3: ", arr3)

	var arr4 = [5]int{2: 15, 4: 30} // 指定index赋值
	fmt.Println("arr4: ", arr4)

	var arr5 = [...]int{3, 2, 6, 5, 4} //根据{}里元素的个数推断出数组的长度
	fmt.Println("arr5: ", arr5)

	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}} //数组的元素类型由匿名结构体给定
	fmt.Println("arr6: ", arr6)

	// 二维数组初始化
	//5行3列，只给前2行赋值，且前2行的所有列还没有赋满
	var arr7 = [5][3]int{{1}, {2, 3}}
	fmt.Println("arr7: ", arr7)
	//第1维可以用...推测，第2维不能用...
	var arr8 = [...][3]int{{1}, {2, 3}}
	fmt.Println("arr8: ", arr8)

	// 访问数组里的元素
	fmt.Println("首元素:", arr5[0])
	// 末元素
	fmt.Println("末元素:", arr5[len(arr5)-1])

	fmt.Println("遍历数组 --------")
	// 遍历数组里的元素
	fmt.Println("遍历数组里的元素 --------")
	for i, ele := range arr1 {
		fmt.Printf("index=%d, element=%d\n", i, ele)
	}

	// 者这样遍历数组
	fmt.Println("者这样遍历数组 --------")
	for i := 0; i < len(arr1); i++ { //len(arr)获取数组的长度
		fmt.Printf("index=%d, element=%d\n", i, arr1[i])
	}

	// 遍历二维数组
	fmt.Println("遍历二维数组 --------")
	for row, array := range arr8 { //先取出某一行
		for col, ele := range array { //再遍历这一行
			fmt.Printf("arr[%d][%d]=%d\n", row, col, ele)
		}
	}

	// 通过for range遍历数组时取得的是数组里每一个元素的拷贝
	fmt.Println("通过for range遍历数组时取得的是数组里每一个元素的拷贝 --------")
	arr := [...]int{1, 2, 3}
	for i, ele := range arr { //ele是arr中元素的拷贝
		arr[i] += 8 //修改arr里的元素，不影响ele
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
		ele += 1 //修改ele不影响arr
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}

// 参数必须是长度为5的int型数组（注意长度必须是5）
func update_array1(arr [5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 888

}

func update_array2(arr *[5]int) {
	fmt.Printf("array in function, address is %p\n", &((*arr)[0]))
	arr[0] = 888 //因为传的是数组指针，所以直接在原来的内存空间上进行修改
}
