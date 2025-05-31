package main

import "fmt"

func loop() {
	fmt.Println("*************Loop Started*****************")
	fmt.Println("Counting bits from 1 to 10")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("******************While loop Started***********")
	j := 1

	for j <= 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("****************** IF-ElSE ***************")

	age := 20

	if age >= 18 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are a minor")
	}
}
