package test

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"
)

//使用常量组的定义常量
const (
	a = "A"
	b        //如果不提供初始值，则表示将使用上行的表达式
	c = iota //常量的计数器
	d
)

//初始化结构体（type <Name> struct{} 定义结构体）
type Person struct {
	Name   string
	Gender string
	Age    uint8
	Hobby  Hobby          //支持匿名结构，但是匿名结构不能直接初始化该变量
	Other  map[int]string //黑盒模型：隐藏处理过程细节去构造更大的盒子 SICP
}
type Hobby struct {
	Name string
}

//初始化接口
type Animal interface {
	Grow()
	UpdateName(string) string
}

//为结构体定义方法(func 和 函数名之间添加绑定对象的声明)
func (person *Person) Grow() {
	person.Age++
}
func (person *Person) UpdateName(name string) string {
	person.Name = name
	return person.Name
}

func main() {
	var name string = "huangjiangang"
	fmt.Println((string(name[1])))

	//变量赋值
	fmt.Println("############### 变量 ###############")
	var test1 int //变量的声明格式：var <变量名称> <变量类型>
	test1 = 3     //变量的赋值格式：<变量名称> = <表达式>
	fmt.Printf("变量声明后赋值 %d \n", test1)

	var test2 int = 4 //声明的同时赋值：var <变量名称> [变量类型] = <表达式>
	fmt.Printf("变量直接负值 %d \n", test2)

	test3 := 5 //系统类型推断赋值：<变量名称> = <表达式>
	fmt.Printf("系统类型推断 %d \n", test3)

	//Go 语言中并没有引用传递，因为在 Go 语言中连引用变量也没有。
	//https://www.jianshu.com/p/d28db9b2acd5
	test4 := &test3
	fmt.Printf("int 的指针类型 %d \n", *test4)

	fmt.Println("############### string 类型 ###############")
	var string1 string //变量的声明格式：var <变量名称> <变量类型>
	string1 = "Go 语言核心 36 讲"     //变量的赋值格式：<变量名称> = <表达式>
	fmt.Printf("变量 %s 长度 %d  %d\n", string1, len(string1),
		utf8.RuneCountInString(string1))

	fmt.Println("############### 标准库 strings 代码包 ###############")
	var strBuild strings.Builder
	strBuild.WriteByte('a')
	strBuild.WriteRune('a')
	strBuild.WriteString("你好")
	fmt.Println(strBuild.String())

	//https://blog.csdn.net/u014270740/article/details/89436719 strReader
	//strReader := strings.NewReader("abcdefg")

	//数组初始化 (知道初始值的数组 | 不知道初始值的初始化)
	fmt.Println("############### 数组 ###############")
	var ary1 [3]int
	ary1 = [3]int{1, 2, 3}
	fmt.Print("不知道数组初始值：")
	fmt.Println(ary1)

	ary3 := [3]int{1, 2, 3}
	fmt.Print("知道数组初始值：")
	fmt.Println(ary3)

	var ary4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //省略数组大小值
	fmt.Println(ary4)

	//数组切片(引用类型的变长数组)
	fmt.Println("############### 数组切片 ###############")
	var slice0 []int
	fmt.Println(slice0)

	var slice1 = make([]int, 3, 10) //一般使用 make 初始化切片
	fmt.Println(slice1)

	var slice2 = ary4[1:2:4] //限制容量上界
	fmt.Println(slice2)
	fmt.Printf("length: %d capcaity: %d \n", len(slice2), cap(slice2))

	fmt.Println("使用 append 函数向切片追加容量，并超出切片上限")
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7) //当 append 追加的数据超出限制，切片会指向新分配的数组
	fmt.Println(slice2)

	//哈希表字典类型（比线性搜素快，比使用索引访问的数据类型慢 100 倍）
	fmt.Println("############### map 哈希表 ###############")
	var map1 = make(map[int]string, 3) //超出容量自动扩容
	map1[1] = "你好"
	delete(map1, 1)
	fmt.Println(map1)

	map2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	fmt.Println(map2)
	fmt.Printf("map len: %d \n", len(map2))

	mapValue1, ok := map2["hello"] //为了判断 map 的 key 是否存在，所以增加了一个返回结果
	fmt.Printf("key 对应的 value = %d 判断值是否存在：%t \n", mapValue1, ok)

	//通道 Channel 类型，在不同Goroutine之间传递类型化的数据（并发安全）
	fmt.Println("############### Channel 通道类型 ###############")
	var chan1 = make(chan string, 5) //只能通过 make 来初始化（缓冲通道 | 非缓冲通道）
	chan1 <- "Hello"
	chanValue, ok := <-chan1
	fmt.Printf("通道返回的结果 %s,通道是否有数据 %t \n", chanValue, ok)

	fmt.Println("############### func 函数类型 ###############")
	fmt.Printf("3.1 is Rain %t \n", isRain("3.1"))
	isLoveMusic("末日飞船", "象牙州")

	fmt.Println("############### struct 结构体类型 ###############")
	pawn := Person{"pawn", "M", 23, Hobby{"ping"}, map[int]string{1: "Hello"}}
	pawn.Grow()
	pawn.UpdateName("ben")
	fmt.Println(pawn)

	fmt.Println("############### interface 接口类型 ###############")
	interPawn := interface{}(&pawn) //在接口类型上使用断言判断对象是否属于某一个接口
	/*
		指针类型(*Person)拥有以它以及以它的基底类型(Person)为接收者类型的所有方法，
		而它的基底类型(Person)却只拥有以它本身为接收者类型的方法。
	*/
	obj, ok := interPawn.(Animal)
	fmt.Println(obj, ok)
	fmt.Println(reflect.TypeOf(obj))

	fmt.Println("############### if 流程控制语句 ###############")
	if number := 4; number < 100 { //初始化字句 + 条件表达式
		number += 100
		fmt.Printf("if 流程控制语句 %d \n", number)
	}

	fmt.Println("############### switch 流程控制语句 ###############")
	//switch 的表达式与 case 的表达式进行比较，相同的进行执行
	switch number := 4; number {
	case 4:
		fmt.Println("switch 判断小于 100")
		fallthrough //继续向下执行
	default:
		fmt.Println("结束 switch 判断")
	}

	fmt.Println("############### for 流程控制语句 ###############")
	loop := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	//数组、字典、字符串会迭代出两个值
	//通道只会接受一个值
	fmt.Println("map 数据结构")
	for i, v := range loop {
		fmt.Printf("%d => %s\n", i, v)
	}

	fmt.Println("############## container list 数据结构 ##############")
	l := list.New()
	l.PushBack(1)
	l.PushBack("123123")
	l.PushBack(map[int]string{1:"1231313", 3:"12312"})
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}

	fmt.Println("############## test goroutine ##############")
	//study GPM 的基础概念
	loopInGoRoutine()
	time.Sleep(1*time.Second)

	fmt.Println("\n############## sync 并发编程包 ##############")
	var mu sync.Mutex
	mu.Lock()
	var atomInt int32 = 1
	atomic.AddInt32(&atomInt, 123)
	fmt.Println(atomInt)
	mu.Unlock()

	//hand := &test.ListNode{Val:1}
	fmt.Println("############## 数组 ##############")
	fmt.Println(test.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}

//函数声明：func + 括号包裹的参数声明列 + 可以使用括号包裹的结果声明列表
//不支持 嵌套、重载和默认参数（支持：不定长度变参、多返回值、）
func isRain(date string) bool {
	if date == "3.2" {
		return true
	}
	return false
}

//使用不定长变参
func isLoveMusic(music ...string) {
	fmt.Print("arg music is a slice : ")
	fmt.Println(music) //slice 切片的数据结构
}

func loopInGoRoutine(){
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			//函数式编程的特点
			fn := func() {
				fmt.Print(i)
			}
			trigger(i, fn)
		}(i)
	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
