package fib

func Fibonacci() func() int{
	a, b := 0, 1
	return func() int {
		a, b = a + b, a
		return a
	}
}
