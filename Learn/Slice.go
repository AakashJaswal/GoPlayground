package main

func slice() {

	slice_test := []string{"One", "Two", "Three"}

	for idx, num := range slice_test {
		println(idx, num)
	}

}
