package main

import "go-plugin-benchmark/fib"

func RandInt() int {
	return fib.FibonacciSum(10000)
	//return fib.FibonacciSum(10000)
}
