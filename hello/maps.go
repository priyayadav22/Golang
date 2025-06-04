package main

import "fmt"

func runMaps(){
	fmt.Println("************maps*****************")

	marks := map[string]int{
		"Priya" : 97,
		"yadav" : 91,
		"Girl" : 90,
	}

	fmt.Println("map Marks: ", marks)

	fmt.Println("marks of Priya : ", marks["yadav"])

	marks["Ankit"] =99

	delete(marks, "Girl")
	fmt.Println("map Marks: ", marks)

	for name, score := range marks {
		fmt.Println(name, "scored", score)
	}

}