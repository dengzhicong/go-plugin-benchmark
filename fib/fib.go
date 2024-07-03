package fib

func tailRecursiveFibonacci(n, a, b int) int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}

	return tailRecursiveFibonacci(n-1, b, (a+b)%1000000007)
}

func FibonacciRecursive(n int) int {
	return tailRecursiveFibonacci(n, 0, 1)
}

func FibonacciSum(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var a, b int = 0, 1
	for i := 2; i <= n; i++ {
		b = a + b
		a = b - a
		b = b % 1000000007
	}
	return b
}
