/**

defer 关键字的主要特征：

延迟执行： 在含有 defer 的函数或方法返回时，不管以何种方式（正常返回、panic 或者其它控制流程），都会执行由 defer 关联的函数调用。

后进先出（LIFO）顺序执行： 如果在一个函数中有多个 defer 语句，它们会在函数退出时按照声明的逆序执行，即最后声明的 defer 语句对应的函数最先执行。

参数即时求值，执行延迟： defer 语句中的函数及其参数会在声明时立即求值，但函数的实际执行则会延迟至函数退出时。

**/

package main

import "fmt"
func defer1() string{
	fmt.Println("defer1")
	return "第一个执行的defer"
}
func defer2() string{
	fmt.Println("defer2")
	return "第二个执行的defer"
}

func main() {
	defer defer1()
	defer defer2()
	println("defer")
}
