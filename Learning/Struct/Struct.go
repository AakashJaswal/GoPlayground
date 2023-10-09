package main

import "fmt"

type human struct {
	Name   string
	Age    int
	Gender string
}

type student struct {
	basic human
	GPA   float32
}

func main() {

	jim := student{basic: human{Name: "Jim", Age: 34, Gender: "M"}, GPA: 4.00}

	fmt.Println(jim)

}
