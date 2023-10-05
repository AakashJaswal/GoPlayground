package main

import "fmt"

func main() {
	// Can declare below.
	var var1 string
	var1 = "This is var1"

	var var2 = "This is var2"

	var3 := "This is var3"

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(testFunc("Hello, World!"))

	println("Calling function to append to and print slice elements with index")
	slice_test := []string{"One", "Two", "Three"}
	printSlice(slice_test)

	println("Original did not change")
	for idx, num := range slice_test {
		println(idx, num)
	}

}

// Functions
func testFunc(s string) string {
	return s
}
