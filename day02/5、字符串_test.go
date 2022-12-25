package day02

import (
	"fmt"
	"strings"
	"testing"
)

// 字符串常用操作
//
// |                方法                 |      介绍      |
// | :---------------------------------: | :------------: |
// |              len(str)               |     求长度     |
// |            strings.Split            |      分割      |
// |          strings.Contains           |  判断是否包含  |
// | strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
// | strings.Index(),strings.LastIndex() | 子串出现的位置 |
func TestString(t *testing.T) {
	s := " My "
	fmt.Println(s)
	s = "He say:\"I'm fine.\" \n\\Thank\tyou.\\"
	fmt.Println(s)

	s = `here is first line. 

  there is third line.
`
	fmt.Println(s)

	str := "born to win, born to die."
	fmt.Printf("sentence length %d\n", len(str))
	fmt.Printf("\"s\" length %d\n", len("s"))  //英文字母的长度为1
	fmt.Printf("\"中\"  length %d\n", len("中")) //一个汉字占3个长度

	arr := strings.Split(str, " ")
	fmt.Printf("arr[3]=%s\n", arr[3])
	fmt.Printf("contain die %t\n", strings.Contains(str, "die"))        //包含子串
	fmt.Printf("contain wine %t\n", strings.Contains(str, "wine"))      //包含子串
	fmt.Printf("first index of born %d\n", strings.Index(str, "born"))  //寻找子串第一次出现的位置
	fmt.Printf("last index of born %d\n", strings.LastIndex(s, "born")) //寻找子串最后一次出现的位置

	fmt.Printf("begin with born %t\n", strings.HasPrefix(s, "born")) //以xxx开头
	fmt.Printf("end with die. %t\n", strings.HasSuffix(s, "die."))   //以xxx结尾

	/*
		把多个字符串拼接成一个长的字符串有多种方式。
			1. 加号连接。
			2. func fmt.Sprintf(format string, a ...interface{}) string
			3. func strings.Join(elems []string, sep string) string
			4. 当有大量的string需要拼接时，用strings.Builder效率最高
	*/
	s1 := "Hello"
	s2 := "how"
	s3 := "are"
	s4 := "you"
	merged := s1 + " " + s2 + " " + s3 + " " + s4
	fmt.Println(merged)

	merged = fmt.Sprintf("%s %s %s %s", s1, s2, s3, s4)
	fmt.Println(merged)

	merged = strings.Join([]string{s1, s2, s3, s4}, " ")
	fmt.Println(merged)
	//当有大量的string需要拼接时，用strings.Builder效率最高
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	sb.WriteString(" ")
	sb.WriteString(s3)
	sb.WriteString(" ")
	sb.WriteString(s4)
	sb.WriteString(" ")
	merged = sb.String()
	fmt.Println(merged)

	/*
		&#8195;string中每个元素叫“字符”，字符有两种：

		1. byte：1个字节， 代表ASCII码的一个字符。
		2. rune：4个字节，代表一个UTF-8字符，一个汉字可用一个rune表示。

		string是常量，不能修改其中的字符。
		string可以转换为[]byte或[]rune类型。
		string底层是byte数组，string的长度就是该byte数组的长度， UTF-8编码下一个汉字占3个byte，即一个汉字占3个长度。
	*/

	str1 := "My name is 张"
	arr1 := []byte(str1)
	brr := []rune(str1)
	fmt.Printf("last byte %d\n", arr1[len(arr1)-1]) //string可以转换为[]byte或[]rune类型
	fmt.Printf("last byte %c\n", arr1[len(arr1)-1]) //byte或rune可以转为string
	fmt.Printf("last rune %d\n", brr[len(brr)-1])
	fmt.Printf("last rune %c\n", brr[len(brr)-1])
	L := len(s1)
	fmt.Printf("string len %d byte array len %d rune array len %d\n", L, len(arr), len(brr))
	for _, ele := range s1 {
		fmt.Printf("%c ", ele) //string中的每个元素是字符
	}
	fmt.Println()
	for i := 0; i < L; i++ {
		fmt.Printf("%c ", s1[i]) //[i]前面应该出现数组或切片，这里自动把string转成了[]byte（而不是[]rune）
	}
}
