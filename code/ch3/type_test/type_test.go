package type_test

import "testing"
//go不支持隐式类型转换
//别名
type MyInt int64
//基本数据结构类型
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

//指针 （go不支持指针运算）
func TestPoint(t *testing.T) {
	a := 1
	//指针
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

//字符串
func TestString(t *testing.T) {
	//默认初始值为空字符串
	var s string
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))

}
