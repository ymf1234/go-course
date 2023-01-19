package day02

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	var ch chan int        // 管道的声明
	ch = make(chan int, 8) // 管道的初始化，环形队列里可容纳8个int
	ch <- 1                //往管道里写入（send）数据
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6

	v := <-ch // 从管道里取走（recv）数据
	fmt.Println(v)
	v = <-ch // 从管道里取数据
	fmt.Println(v)
	//readOnly := make(<-chan int)  // 定义只读的channel
	//<-readOnly                    // 读
	//writeOnly := make(chan<- int) // 定义只写的channel
	//writeOnly <- 1                // 写

	close(ch) // 遍历前必须先关闭管道，禁止在写入元素

	// 遍历管道里剩下的元素
	for ele := range ch {
		fmt.Println(ele)
	}

}

// 只有向channel里写数据
func send(c chan<- int) {
	c <- 1
}

// 只能取channel中的数据
func receive(c <-chan int) int {
	return <-c
}

type Context struct {
}

// 返回一个只读channel
func (c *Context) Done() <-chan int {
	return nil
}
