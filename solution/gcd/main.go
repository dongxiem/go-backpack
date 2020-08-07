package main

// 计算两个整数值的的最大公约数（GCD-greatest common divisor）
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	x, y := 2, 4
	// 得到返回的最大公约数
	res := gcd(x, y)
	// 进行打印
	println(res)
}
