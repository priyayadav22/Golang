package main

import "fmt"

type Developer struct {
	name string
	language string 
}

func (d Developer) introduce(){
	fmt.Println("Hi i am", d.name, "and i love", d.language)
}

func runStructMethod(){
	fmt.Println(" *********** struct with methods*************")

	dev := Developer{"priya", "Hindi"}
	dev.introduce()
}