package main

import (
	"fmt"
)
func main() {
	var name string
	var age int
	var job string

	fmt.Print("Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Your age: ")
	fmt.Scan(&age)
	fmt.Print("Enter Career : ")
	fmt.Scan(&job)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(job)
	

}
