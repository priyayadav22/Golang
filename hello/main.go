// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello priya, U started Golang")
// 	fmt.Println("****************Variable Started*********************")

// 	var name string = "Priya Yadav"
// 	var age int = 21
// 	const isFemale bool = true

// 	var score int
// 	fmt.Println("Score : 😁", score)

// 	language := "Golang h bhai"

// 	fmt.Println("Name: ", name)
// 	fmt.Println("Age: ", age)
// 	fmt.Println("is Female : ", isFemale)
// 	fmt.Println("Language : ", language)

// 	fmt.Println("****************Variable Ended*********************")

// 	loop()

// 	runFunction()

// 	runSlice()

// 	runArray()

// 	runMaps()

// 	runStruct()

// 	runStructMethod()

// }


package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/students", studenthandler)
	log.Println("🚀 Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
