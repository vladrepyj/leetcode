func fib(n int) int {
	if n == 0 {
		return 0
	}
	f1, f2 := 0, 1
	for i := 2; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1 + f2
}