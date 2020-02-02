package test

import (
	"fmt"
	"testing"
)

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {

	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func TestInterface(t *testing.T) {

	//实例化file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 将接口赋值f，也就是*file类型
	writer = f

	//writer := new(file)


	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")
	//sort.Interface()

	//sort.IntSlice{1,2,3}
}

func TestEmptyInterface(t *testing.T){
	//var b int = 1
	var a  interface{} = 1

	value,ok := a.(int)

	if(ok){
		print(value)
	}
}