package main

import "fmt"

func runSlice(){
	fmt.Println("\n ******** Slice*******")

	number := []int{1,2,3,4,5}
	fmt.Println("Original slice:", number)

	number = append(number, 5,6)
	fmt.Println("After append: " , number)

	part := number[1:4] // from index 1 to 3
	fmt.Println("Sliced part:", part)

    fmt.Println("***********For Each LOOP ************* syntax thoda alg")
	for _, value := range number{  // if we dont care about index er can use _ and if we need specific index,,.. then pass that
	fmt.Println("Value:", value)
}
}