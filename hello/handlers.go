package main

import (
	"encoding/json"
	"net/http"
)

// Define Student struct with correct capitalized fields and JSON mapping
type Student struct {
	ID   int    `json:"id"`   // JSON key: "id"
	Name string `json:"name"` // JSON key: "name"
	Age  int    `json:"age"`  // JSON key: "age"
}

// Slice to act like an in-memory database
var student = []Student{
	{ID: 1, Name: "Priya", Age: 21},
	{ID: 2, Name: "Saurabh", Age: 24},
}

// Handler for GET and POST requests on /students
func studenthandler(w http.ResponseWriter, r *http.Request) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		// Send the list of students as JSON response
		json.NewEncoder(w).Encode(student)
		return
	}

	if r.Method == http.MethodPost {
		var newStudent Student

		// Parse the incoming JSON body into newStudent struct
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Assign next ID and add new student to the slice
		newStudent.ID = len(student) + 1
		student = append(student, newStudent)

		// Send 201 Created response and the new student
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)
		return
	}

	// Handle unsupported HTTP methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"  // print to console
// 	"net/http"   // for building http server and handeling request
// )

// type Student struct{
// 	Id int `json:"id"`  //field tag -- tells go t map ID to id in JSON
// 	Name string `json:"name"` // same here for Name
// 	Age int `json:"age"`   // and for AGe
// }

// var student = []Student{
// 	{ID:1, name: "Priya", Age:21},
// 	{ID:2, Name: "Saurabh", Age:24},
// }

// func studenthandler(w http.ResponseWriter, r *http.Request){
// 	//studenthandler will handle both get and post requests for /students woute
// 	//w is used to send response back to client
// 	//r cntains the HTTp Request info(Method, body, headers)
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == http.MethodGet{
// 		json.NewEncoder(w).Encode(student) // encode go slice to json and send as response
// 		return // exit fter sending response
// 	}

// 	if r.Method ==http.MethodPost{
// 		var newStudent Student
// 		err:= json.NewDecoder(r.Body).Decode(&newStudent)

// 		if err!=nil{
// 			http.Error(w,"Invalid Input", http.StatusBadRequest)
// 			return
// 		}
// 			newStudent.ID = len(Student)+1
// 			student = append(student, newStudent)

// 			w.WriteHeader(http.StatusCreated)
// 			json.newEncoder(w).Encode(newStudent)
// 			return

// 	}

// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

// }
