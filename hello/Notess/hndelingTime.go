it is very handeling time is practical and powerful topic and not as scary as it looks

import "time"

//get current time
now:= time.now()
fmt.println(now)   // 2025-06-22 18:52:13.123456 +0530 IST

u can also do. 
now.year()
now.month()
now.Minute()  //print minute


Go time formatting is already unique and we dont need to pass dd-mm-yyyy etc.  but iu can do 

now.Format("02-Jan-2006 15:04:09")  // it will print current date and time with this format:--- 22-Jun-2025 18:52:19

//add/subttract time

tomorrow:= now.Add(24*time.Hour)
yesterdat:= now.Add(-24 * time.Hour)
fmt.println("tomorrow: ", tomorrow)
fmt.println("yesterday: ", yesterday)
**** you can also use time. Minute, time.Second, time. Millisecond etc


// sleep/Delay
fmt.println("Sleeping for 2 second....")
time.sleep(2*time.Second)
fmt.Println("Awake!...")

// compare 2 lines

ti1:= time.Now()
t2:= t1.Add(1*time.Hour)

fmt.Println("t1 Before t2?", t1.Before(t2))  // true
fmt.Println("t1 After t2?", t1.After(t2))    // false
fmt.Println("Equal?", t1.Equal(t2))          // false