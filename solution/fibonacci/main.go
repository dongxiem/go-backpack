package main

// 计算斐波纳契数列（Fibonacci）的第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	println(fib(10))
}
