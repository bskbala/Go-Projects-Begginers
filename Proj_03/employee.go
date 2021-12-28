package main

import ("fmt")

type employee struct {
	firstName string
	lastName  string
	city      string
	state     string
	street    string
	pincode   int
	age       int
}

func main() {
	emp1 := employee {
			"saikumar",
			"balam",
			"hyderabad",
			"AndhraPradesh",
			"srinivas nagar", 
			500060, 
			26,
	}

	fmt.Println("employee ",emp1)
}

