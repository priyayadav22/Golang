package main

func main() {
	todos:= Todos{} // initialize an empty todos list
	todos.add("Love yourself")
	todos.add("Buy Bread")
	todos.add("Study")
	todos.delte(1)
	todos.toggle(0)
	todos.print()
}