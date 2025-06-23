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
	// "log"
	// "net/http"
	"fmt"
	"bufio"
	"os"
	
)

func main() {
	// http.HandleFunc("/students", studenthandler)
	// log.Println("🚀 Server started at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// var name string
	reader := bufio.NewReader(os.Stdin)
	//bufio package == buffered input/output from go's standard library
	//newReader is a function that creates a new "reader" object that can read from input
	//os.stdin: 	Stands for Standard Input (your keyboard input). It’s the source we're reading from
	fmt.Println("write your name sweetheart: ")
	// fmt.Scanln(&name) // it asks for input and wait until u press enter and take only first word from ur inputt
	name,_ := reader.ReadString('\n')
	// name, _ :=   Read and store the result in name. The _ is used to ignore the error
	fmt.Println("hello ", name)
}


/*
TaskFile.yml(often just called TaskFile) is not a Go built-in. it is part of popular tool in Go-Ecosystem called Task

taskfile is automation tool for Go... it defines custom tasks(like build, run, test, format, etc. so that you can run them with simple commands)


why use taskfile--- lets us automative repetative commands like running ur server, running tests, building a binary, linting Code   AND MOST IMPORTANT YOU CAN NORMALLY TYPE INTO TERMINAL AS--- 'RUN MY GO APP' AND IT WILL RUN YOUR APP etc


fORMAT OF  TASKFILE.YAML
version: '3'
tasks:
    run: 
        desc : run the Go Application
	    cmds:
	      - go run main.go
	test:
		desc: run all tests
		cmds:
			-go test ./.  


			

			
****************NOw instead of go run main.go.... just type : tasks run

*********** it is optional  .. not required by GO.. use it only if you want automation





*/