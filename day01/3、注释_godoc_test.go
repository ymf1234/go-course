package day01

//- 单行注释，以//打头。
//- 多行注释有2种形式：
//1. 连续多行以//打头，注意多行注释之间不能出现空行。
//2. 在段前使用/\*，段尾使用*/。
//- 注释行前加缩进即可写go代码。
//- 注释中给定的关键词。NOTE: 引人注意，TODO: 将来需要优化，Deprecated: 变量或函数强烈建议不要再使用。

// Add1 2个整数相加
// 返回和。
//
// NOTE: 注释可以有多行，但中间不能出现空行（仅有//不算空行）。
func Add1(a, b int) int {
	return a + b
}

/*
Sub 函数使用示例：

	  for i:=0;i<3;i++{
		  Sub(i+1, i)
	  }

看到了吗？只需要行前缩进，注释里就可以写go代码，是不是很简单。
*/
func Sub(a, b int) int {
	return a - b
}

// TODO: Prod 该函数不能并发调用，需要优化
func Prod(a, b int) int {
	return a * b
}

// Deprecated: Div 不要再调用了
func Div(a, b int) int {
	return a / b
}
