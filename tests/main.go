package tests

func facto(n int) (x int) {
	x = 1
	for i := 1; i <= n; i++ {
		x *= i
	}
	return
}

func main() {
	println("factorial of 8 is", facto(8))
}
