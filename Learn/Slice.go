package main

func printSlice(s []string) {

	s = append(s, "Four")

	for idx, num := range s {
		println(idx, num)
	}

}
