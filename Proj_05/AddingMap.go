package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Phone"] = 1234568956
	employee["Pin"] = 78256
	fmt.Println(employee)
}
