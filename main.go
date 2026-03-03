package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
func main() {
	var name string
	var age int
	var job string

	//Create scanner

	scanner := bufio.NewScanner(os.Stdin)

	//Handle name
	fmt.Print("Enter Your name: ") //Print Instruction to guide what to input
	scanner.Scan() //Wait for the input
	name = scanner.Text() //Scans the input 

	//Handle age
	
	fmt.Print("Enter Your Age: ")
	scanner.Scan()
	age,_ = strconv.Atoi(strings.TrimSpace(scanner.Text()))

	//Handle Age
	fmt.Print("Enter Your profession: ")
	scanner.Scan()
	job = scanner.Text()

	fmt.Printf("\n-------Profile-------\n")
	fmt.Printf("%s, %d, %s \n", name, age, job)

}
