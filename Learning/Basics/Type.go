package main

type num_type []string

func (n num_type) appendAndPrint() {
	n = append(n, "type2")
	for idx, nm := range n {
		println(idx, nm)
	}
}

func (n num_type) print() {

	for idx, nm := range n {
		println(idx, nm)
	}
}
