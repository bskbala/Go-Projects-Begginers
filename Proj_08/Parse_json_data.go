package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Employees struct {
	Employees []Employee `json:"employee"`
}

type Employee struct {
	Name   string `json:"name"`
	Gender string `json:gender`
	Age    int    `json:Age`
}

func main() {
	jsonFile, err := os.Open("employee.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SuccessFully Opened Json file")
	defer jsonFile.Close()

	byteEmpValue, _ := ioutil.ReadAll(jsonFile)

	var employees Employees
	json.Unmarshal(byteEmpValue, &employees)

	for i := 0; i < len(employees.Employees); i++ {
		fmt.Println("Employee Name:" + employees.Employees[i].Name)
		fmt.Println("Employee Gender:" + employees.Employees[i].Gender)
		fmt.Println("Employee Age:" + strconv.Itoa(employees.Employees[i].Age))

	}
}
