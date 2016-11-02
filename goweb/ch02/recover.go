package main

import "fmt"

func doPanic() string {
	// 异常捕获
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)

			// 在recover里面是不允许再return的
			// return "haha"
		}
	}()

	// 抛出异常
	panic("Just panicking for the sake of demo")
	// 下面的代码被不会输出
	fmt.Println("This will never be called")

	return "haha"
}

func main() {
	fmt.Println("Starting to panic")
	doPanic()

	// 由于异常被捕获，下面的代码可正常执行
	fmt.Println("Program regains control after panic recover")
}
