pointer is a variable that store the memory adress of another variable


var a int = 10
var p *int = &a   // p is a pointer to int, and stores the address of a

fmt.Println("Value of a:", a)
fmt.Println("Address of a:", &a)
fmt.Println("Value stored in pointer p:", p)
fmt.Println("Value at address p points to:", *p)


*****************************8
a = 10
p = &a    ---> 🧭 pointer stores address of a
*p        ---> dereferencing pointer (go to address & get the value)
*********************************

we use pointer  to save memory as we dont need to copy big data agai and again
share data safely
modify variable inside function

Unlike C/C++, Go doesn’t allow adding/subtracting pointers:
p + 1  // ❌ invalid in Go
