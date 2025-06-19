package main 

import "fmt"

type Ptudent struct {
	name string
	age int
	score int
}

func runStruct(){
	fmt.Println("****************Struct **********************")

	p:= Ptudent{name: "Priya", age:21, score:98}
// if u r passing printf then it expacts  (like %s, %d).
// fmt.Printf("Student name: %s\n", p.name)  // %s for string
 
	fmt.Print("Student name: ", p.name)
	fmt.Print("Student Age: ", p.age)
	fmt.Println("Student Score: ", p.score)
}