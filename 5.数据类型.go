package main

import (
	"fmt"
	"unsafe"
)
func main() {
	num3:=1
	fmt.Printf("num3: %T",num3) //不换行
	// 换行
	fmt.Println()
	// 占用的字节数 ，32位系统，占4字节，64 占8字节，和操作系统有关
	fmt.Println(unsafe.Sizeof(num3))



	// 字符串

	str1:="hello"
	fmt.Println(str1)
	str1="world"
	fmt.Println(str1)

	fmt.Println(`我是反引号///+山东福利`)
}