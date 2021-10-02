package recover

import "fmt"

/**
recover 只能在 defer 修饰的函数,用于取得 panic 调用中传递过来的错误值
 */

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		// catch panic
		if e := recover(); e != nil {
			// handle error
			fmt.Printf("Panicing %v\n", e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test completed")
}
