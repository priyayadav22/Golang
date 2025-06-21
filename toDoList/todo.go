package main

import (
	"errors"
	"fmt"
	"time"
	"os"
	"strconv"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Todo struct{
	Title string
	Completed bool
	CreatedAt time.Time  // always set when task is created
	CompletedAt *time.Time  // optional: maybe nil until completed
}


type Todos []Todo  // to store our todo 

func (todos *Todos) add(title string){

	todo :=Todo{
		Title:  title,  // setting up title
		Completed: false, // abhi to task create hua h isliye not completed yet
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo) // addd new todo directly to existig list of to-dos using a pointer(*todos).. this pointer allows us to access and modify the original list stored in memory

}



func (todos *Todos)  validateIndex(index int) error{  // it can return an error
	if(index <0 || index >= len(*todos)){ // trying to access something outsid ethe list
		err:=errors.New("Invalid Index")// creating new Error Messgae
		fmt.Println(err.Error())
		return err
	}
	return nil
}


func(todos *Todos) delte(index int) error{
	if err:= todos.validateIndex(index); err!=nil{
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	// before := (*todos)[:index] 
    //after := (*todos)[index+1:]
	return nil
}

func (todos *Todos) toggle(index int) error{
	if err:=todos.validateIndex(index); err!=nil{
		return err
	}

	//access the actual todo item
	todo := &(*todos)[index] // get pointer so that we can modify the item directly  
	todo.Completed = !todo.Completed


	if todo.Completed{
		//todo.CompletedAt = time.Now()// cant do it as it expects adress..
		now:=time.Now()
		todo.CompletedAt = &now
	}

	return nil
}

func (todos *Todos) print() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout) // print to terminal
	t.SetTitle("📝 To-Do List")
	t.AppendHeader(table.Row{"#", "Title", "Completed", "Created At", "Completed At"})

	for i, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format("02-Jan-2006 15:04:05")
			}
		}

		t.AppendRow(table.Row{
			strconv.Itoa(i),
			todo.Title,
			completed,
			todo.CreatedAt.Format("02-Jan-2006 15:04:05"),
			completedAt,
		})
	}

	t.Render() // print the final table
}

/*
struct is like a dabba holding ur data
Ex: todo is a tupperware box with title and completed isnide

* pointer
-- it gives the actual dabbba's address.. not the copy of it
--- you are modifying the actuall dabba not its copy
-- it is taking u to dabba directly(bu anywhere door)

& addressOf 
gives te locatio of daba.. ye lo dabba ka adress.. ab ja aur change kr

:= create and assign varible automatically
*/