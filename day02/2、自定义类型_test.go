package day02

import "testing"

// 自定义类型
func TestCustomTypes(t *testing.T) {
	// 类型别名
	type byte = uint8
	type rune = int32
	type semaphore = uint8

	// 自定义类型
	type user struct {
		name string
		age  int
	} //用分号把多行代码隔开
	type signal uint8
	type ms map[string]string
	type add func(a, b int) int
}
