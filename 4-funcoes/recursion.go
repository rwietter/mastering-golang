package main

func fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fibo := fibonacci(uint64(50))
	println(fibo)
}
