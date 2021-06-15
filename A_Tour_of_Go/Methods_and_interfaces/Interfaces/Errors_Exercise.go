package main

import (
	"fmt"
)

// 自定义的 float64 类型
type ErrNegativeSqrt float64

// Note: A call to `fmt.Sprint(e)` inside the `Error` method
// will send the program into an infinite loop.
// You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?
//
// 上述的死循环，个人理解是: main函数中的`fmt.Println`方法调用了 ErrNegativeSqrt 的 Error() 方法，
// 而 Error() 方法中的`fmt.Sprintf`方法又调用了 ErrNegativeSqrt 的 Error() 方法,
// 如此下去产生了无法跳出的递归调用。因为`fmt`包在打印时会自动寻找 `error` 接口。
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := 1.0
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z) // Newton's method.
			// fmt.Printf("Sqrt(%d)---%g\n", i, z)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
