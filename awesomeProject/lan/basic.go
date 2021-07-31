package lan

func test01() {
	var v1 int            // 整型
	var v2 string         // 字符串
	var v3 bool           // 布尔型
	var v4 [10]int        // 数组，数组元素类型为整型
	var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
		f float64
	}
	var v6 *int           // 指针，指向整型
	var v7 map[string]int   // map（字典），key为字符串类型，value为整型
	var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型
}