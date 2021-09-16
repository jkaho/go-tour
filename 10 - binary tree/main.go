package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// ----- For reference -----
// type Tree struct {
// 	  Left *Tree
//	  Value int
//	  Right *Tree
// }
// -------------------------

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// close channel when function ends
	defer close(ch)

	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t != nil {
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
		}
		return
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// different number of nodes
		if ok1 != ok2 {
			return false
		}
		// different values
		if v1 != v2 {
			return false
		}
		// end of tree, no differences
		if !ok1 {
			return true
		}
	}
}

func main() {
	// test Walk()
	fmt.Println("----- test Walk() ----- ")
	fmt.Println("expected:\n [1 2 3 4 5 6 7 8 9 10]")
	ch := make(chan int)
	var a []int
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		a = append(a, <-ch)
	}
	fmt.Println("result:\n", a)

	// test Same()
	fmt.Println("\n----- test Same() ----- ")
	fmt.Println("expected: ")
	fmt.Println(" tree.New(1) == tree.New(1): true")
	fmt.Println(" tree.New(1) == tree.New(10): false")
	fmt.Println("result: ")
	fmt.Println(" tree.New(1) == tree.New(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println(" tree.New(1) == tree.New(10):", Same(tree.New(1), tree.New(10)))
}
