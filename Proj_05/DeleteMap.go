package main 

import ("fmt")

func main  () {
	var employee = map [string] int{"saikumar":30,"bandiwala":58}
	delete(employee,"bandiwala")
	fmt.Println(employee)
}