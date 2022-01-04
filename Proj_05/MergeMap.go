package main

import "fmt"

func main() {
	var employee = map[string]int{"saikumar": 52, "bandiwala": 50}
	var student = map[string]int{"gun": 50, "test": 56}

	for k, v := range student {
		employee[k] = v
	}
	fmt.Println(employee)

}
