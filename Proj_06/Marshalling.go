//Marshal Employee Struct

package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name    string
	Age     int
	Address string
}

func main() {
	emp := employee{Name: "SaiKumar", Age: 25, Address: "San Antonio,TX"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
}
