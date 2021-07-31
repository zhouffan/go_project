package main

import "fmt"

func GetName() (p1 string,p2 string){
	p1 = "hello"
	p2 = "world"
	return p1, p2
	//return "hello", "world"
}
func main() {
	_, str := GetName()
	fmt.Println(str)

	Param()
}

func Param() {
	var v1 int = 10           // 整型
	v1_2 := 20
	var v2 string         // 字符串
	var v3 bool           // 布尔型
	//var v4 [10]int        // 数组，数组元素类型为整型
	//var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
	//	f float64
	//}
	//var v6 *int           // 指针，指向整型
	//var v7 map[string]int   // map（字典），key为字符串类型，value为整型
	//var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型

	v1, v1_2 = v1_2, v1
	fmt.Println(v1, v1_2, v2, v3)

	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (          // 通过一个 const 关键字定义多个常量，和 var 类似
		size int64 = 1024
		eof = -1  // 无类型整型常量
	)
	const u, v float32 = 0, 3  // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量

	fmt.Println(size, eof, u, a)

	//出现一次 iota，值自增 1。下一个 const出现时重置为0
	const (    // iota 被重置为 0
		c0 = iota   // c0 = 0
		c1 = iota   // c1 = 1
		c2 = iota   // c2 = 2
	)
	const (
		uu = iota * 2  // u = 0
		vv = iota * 2  // v = 2
		ww = iota * 2  // w = 4
	)
	const x = iota  // x = 0
	const y = iota  // y = 0


	//枚举     以大写字母开头的常量在包外可见
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)


	fmt.Println(Monday)
	foo(Policy_MAX)
}

type PolicyType int32
const (
	Policy_MIN      PolicyType = 0
	Policy_MAX      PolicyType = 1
	Policy_MID      PolicyType = 2
	Policy_AVG      PolicyType = 3
)
func (p PolicyType) String() string {
	switch (p) {
	case Policy_MIN: return "MIN"
	case Policy_MAX: return "MAX"
	case Policy_MID: return "MID"
	case Policy_AVG: return "AVG"
	default:         return "UNKNOWN"
	}
}
func foo(p PolicyType) {
	fmt.Printf("enum value: %v\n", p)
}