package main

import "fmt"

func main() {
	var employee = map[string]int{"SaiKumar": 56}
	employee["SaiKumar"] = 40
	fmt.Println(employee["SaiKumar"])
}
