package main

import ("fmt"
		"encoding/json"
)

type employee struct {
	Name string `json:name`
	Age  int 	`json:age`
	Address string `json:address`

}

func main  () {
	empJsonData := `{"Name":"Saikumar","Age":30,"Address":"SanAntonio,Tx"}`
	empBytes := []byte (empJsonData)
	var emp employee
	json.Unmarshal(empBytes,&emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}