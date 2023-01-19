package day03

import (
	"fmt"
	"testing"
	"time"
)

// 定义结构体
type User struct {
	id            int
	score         float32
	enrollment    time.Time
	name, address string // 多个字段类型相同时可以简写到一行里
}

// 结构体中含有匿名成员
type Student struct {
	Id      int
	string  // 匿名字段
	float32 // 直接使用数据类型作为字段名，所以匿名字段中不能出现重复的数据类型
}

// 为自定义类型添加方法
type UserMap map[int]User

// 可以给自定义类型添加任意方法
func (m UserMap) GetUser(id int) User {
	return m[id]
}

// 结构体创建、访问与修改
func TestStruct(t *testing.T) {
	// 声明和初始化结构体
	var u User // 声明结构头， 会用相应类型的默认值初始化struct里的每个字段
	u = User{} // 用相应类型的默认值初始化struct里的每个字段
	u = User{
		id:         1,
		score:      1.1,
		enrollment: time.Now(),
		name:       "john",
		address:    "123 Main St.",
	} // 赋值初始化

	user2 := User{2, 1.2, time.Now(), "john", "Golang"} // 赋值初始化，可以不写字段名。但需要跟结构体定义里的字段顺序一致
	fmt.Println(u, user2)

	// 访问与修改结构体
	u.enrollment = time.Now()
	fmt.Printf("id=%d enrollment=%v, name=%s \n", u.id, u.enrollment, u.name)

	// 成员方法
	u.hello("Man")
	u.think("Man")

	// 结构体的可见性：
	//- go语言关于可见的统一规则：大写字母开头跨package也可以访问；否则只能本package内部访问。
	//- 结构体名称以大写开头时，package外部可见，在此前提下，结构体中以大写开头在成员变量或成员方法在package外部也可见。

	// 匿名结构体
	var stu struct { // 声明stu是一个结构体，但这个结构体是匿名的
		Name string
		Addr string
	}
	stu.Name = "john"
	stu.Addr = "123 Main St."
	fmt.Println(stu)

	var stu1 = Student{Id: 1, string: "john", float32: 78.0}
	fmt.Printf("anonymous_member string member=%s float member=%f\n", stu1.string, stu1.float32) //直接使用数据类型访问匿名成员

}

// 函数里不需要访问user的成员，可以传匿名，甚至_也不传
func (_ User) think(man string) {
	fmt.Println("hi " + man + ", do you know my name?")
}

// 结构体指针
func TestStruct2(t *testing.T) {
	// 创建结构体指针
	var u User
	user := &u    // 通过取地址符&得到指针
	user = &User{ // 直接创建结构体指针
		id:      1,
		name:    "john",
		address: "123 Main St.",
	}
	fmt.Println(user)

	user = new(User) // 通过new()函数实体化一个结构体，并返回其指针
}

// 构造函数
// 构造函数。返回指针是为了避免值拷贝
func NewUser(id int, name string) *User {
	return &User{
		id:      id,
		name:    name,
		address: "123 Main St.",
		score:   59.0,
	}
}

// user传的是值，即传的是整个结构体的拷贝。在函数里修改结构体不会影响原来的结构体
func hello(u User, man string) {
	u.name = "杰克"
	fmt.Println("hi " + man + ", my name is " + u.name)
}

// 传的是user指针，在函数里修改user的成员会影响原来的结构体
func hello2(u *User, man string) {
	u.name = "杰克"
	fmt.Println("hi " + man + ", my name is " + u.name)
}

// //可以把user理解为hello函数的参数，即hello(u user, man string)
func (u User) hello(man string) {
	fmt.Println("hi " + man + ", my name is " + u.name)
}

// 可以理解为hello2(u *user, man string)
func (u *User) hello2(man string) {
	u.name = "杰克"
	fmt.Println("hi " + man + ", my name is " + u.name)
}

// 结构体嵌套
type user1 struct {
	name string
	sex  byte
}
type paper struct {
	name   string
	author user1
}

type vedio struct {
	length int
	name   string
	user1  //匿名字段,可用数据类型当字段名
}

func TestStruct3(t *testing.T) {
	p := new(paper)
	p.name = "标题"
	p.author.name = "姓名"
	p.author.sex = 0
	fmt.Println(p)

	v := new(vedio)
	v.length = 10
	v.name = "视频名称"
	v.user1.sex = 1       // 通过字段名逐级访问
	v.sex = 1             //对于匿名字段也可以跳过中间字段名，直接访问内部的字段名
	v.user1.name = "作者姓名" //由于内部、外部结构体都有name这个字段，名字冲突了，所以需要指定中间字段名

	fmt.Println(v)
}

// 深拷贝与浅拷贝
type User4 struct {
	Name string
}

type Vedio struct {
	Length int
	Author User4
}

// - 深拷贝，拷贝的是值
// - 浅拷贝，拷贝的是指针
// - 深拷贝开辟了新的内存空间，修改操作不影响原先的内存。
// - 浅拷贝指向的还是原来的内存空间，修改操作直接作用在原内存空间上。
// 传slice，对sclice的3个字段进行了拷贝，拷贝的是底层数组的指针，所以修改底层数组的元素会反应到原数组上。
func TestStruct4(t *testing.T) {

}
