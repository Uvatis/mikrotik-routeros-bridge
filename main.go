package main

import "fmt"

type HelloRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var name string
	name = "John"
	fmt.Println("H")
	fmt.Printf("Hello %s\n", name)

	age := 23

	fmt.Println(age)

	var t [5]int

	t[0] = 10
	fmt.Println(t)

	table := [2]int{1, 2}

	fmt.Println(table)

	a := map[string]string{
		"name": "John",
		"age":  "23",
	}

	fmt.Printf("%#v\n", a["name"])
}
