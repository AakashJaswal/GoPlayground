package main

import "fmt"

func loop() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

}

func whileLoop() {
	idx := 1
	// Support break and continue
	for idx < 5 {
		fmt.Println(idx)
		idx++
	}
}
