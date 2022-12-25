package day02

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 切片的初始化
	var s1 []int // 切片声明 len=cap=0
	fmt.Println("s1:", s1)
	s1 = []int{} // 初始化 len=cap=0
	fmt.Println("s2:", s1)
	s1 = make([]int, 3) // 初始化 len=cap=3
	fmt.Println("s3:", s1)
	s1 = make([]int, 3, 5) // 初始化 len=3, cap=5
	fmt.Println("s4:", s1)
	s1 = []int{1, 2, 3, 4, 5} // 初始化 len=5,cap=5
	fmt.Println("s5:", s1)

	s2d := [][]int{
		{1}, {2, 3}, //二维数组各行的列数相等，但二维切片各行的len可以不等
	}
	fmt.Println("s2d:", s2d)

	s2 := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s2[i] = i + 1
	} //s=[1,2,3]
	fmt.Printf("s[0] address %p, s=%v\n", &s2[0], s1)

	/*
		capacity还够用，直接把追加的元素放到预留的内存空间上
	*/
	s2 = append(s2, 4, 5) //可以一次append多个元素
	fmt.Printf("s[0] address %p, s=%v\n", &s2[0], s2)

	/*
		capacity不够用了，得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作
	*/
	s2 = append(s2, 6)
	fmt.Printf("s[0] address %p, s=%v\n", &s2[0], s2)

	fmt.Println("探究capacity扩容规律")
	expansion()

	sub_slice()
}

/**
  切片相对于数组最大的特点就是可以追加元素，
  可以自动扩容。追加的元素放到预留的内存空间里，
  同时len加1。如果预留空间已用完，
  则会重新申请一块更大的内存空间，
  capacity大约变成之前的2倍(cap<1024)或1.25倍(cap>1024)。
  把原内存空间的数据拷贝过来，在新内存空间上执行append操作。
*/
// 探究capacity扩容规律
func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			//每次扩容都是扩到原先的2倍
			fmt.Printf("capacity从%d变成%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}

func sub_slice() {
	fmt.Println("sub_slice ------")
	/*
		截取一部分，创造子切片，此时子切片与母切片(或母数组)共享底层内存空间，母切片的capacity子切片可能直接用
	*/
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	} //s=[1,2,3]
	fmt.Printf("s[1] address %p\n", &s[1])
	sub_slice := s[1:3] //从切片创造子切片，len=cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	/*
		母切片的capacity还允许子切片执行append操作
	*/
	sub_slice = append(sub_slice, 6, 7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("s=%v, sub_slice=%v, s[1] address %p, sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])
	/*
		母切片的capacity用完了，子切片再执行append就得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作。此时的append操作跟母切片没有任何关系
	*/
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("s=%v, sub_slice=%v, s[1] address %p, sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr[1] address %p\n", &arr[1])
	sub_slice = arr[1:3] //从数组创造子切片，len=cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	/*
		母数组的capacity还允许子切片执行append操作
	*/
	sub_slice = append(sub_slice, 6, 7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("arr=%v, sub_slice=%v, arr[1] address %p, sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])
	/*
		母数组的capacity用完了，子切片再执行append就得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作。此时的append操作跟母数组没有任何关系
	*/
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("arr=%v, sub_slice=%v, arr[1] address %p, sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])
}
