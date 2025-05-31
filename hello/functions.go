package main

import "fmt"

func runFunction(){
	fmt.Println("********* functions************")
	greet("Saurabh")
	sum:= add(5,11)
	fmt.Println("Sum: ", sum)
}

func add(a, b int) int {
	return a+b
}

func greet(name string){
	fmt.Println("Hello,  My dear" , name)
}