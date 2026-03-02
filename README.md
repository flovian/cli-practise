# CLI PRACTISE

This is personal learning project.This program should be able to capture the user bio helping with the big question tell us about yourself.
It should be able to take the user name,age,career,location,hobbies,Education/Achievements,fun fact

## Use fmt as an import

This is how the code will look:

```bash 
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
	fmt.Print("Enter Job descrription: ")
	fmt.Scan(&job)
    }
```
## Use Bufio,os plus fmt
