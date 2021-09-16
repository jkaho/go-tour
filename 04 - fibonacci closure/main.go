package main

import "fmt"

func fibonacci() func() int {
	prevNum := 0
	currNum := 1
	return func() int {
		prevNumRef := prevNum
		nextNum := prevNum + currNum
		prevNum = currNum
		currNum = nextNum
		return prevNumRef
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
