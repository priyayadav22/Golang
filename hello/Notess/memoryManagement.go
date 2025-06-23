memory mangement in golang is simple on the surface because go does a lot for you but powerful once you understand the internals..

** imagine you are in a room called Memory.. So you need chairs (variables) to sit your data on 
-- go gives you chairs using var, new or make.
-- when you are done. go sends a cleaning robot called Garbage collector to throw away unused chairs.
you dont need to manually free memory unlike in c or c++

Techinal View:  Go manages memory through--
1. Stack Memory -- for short-lived values (like function variables)
2. heap Memory -- for long lived values (like pointers, slice, maps)
3. Garbage Collector --- automatic cleanerr
4. escape analysis- decide wheather a variable goes to stack or heap


// how you allocate Memory:  
basic variable :- stack
		var a int=5

heap allocation via new 
		p:= new(int)
		*p = 100

🚿 Garbage Collection (GC)
var data = make([]int, 1_000_000) // huge memory
data = nil                        // now unreachable
// GC will eventually clean this up
you dont need to free() manually --- go tracks what is still in use


💡 Bonus Tips
Don’t worry about memory unless you are building high-performance code.
If using big slices or maps, set them to nil when done (to help GC).
Use sync.Pool for object reuse if you are doing high-speed systems.

sync.Pool is a memory reuse mechanism in Go. It’s like a recycle bin for temporary objects.

--- instead of creating and destroying short-lived objects again and again which is costly.. go will say lets reuse them using sync pool