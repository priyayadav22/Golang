package main

import "fmt"

func runArray() {
	fmt.Println("************Array************")

	var nums [3]int
	nums[0] = 100
	// nums[1] =101
	nums[2] = 102
	// nums[3] = 1001   //out of bound

	fmt.Println("nums: ", nums)

	colors := [3]string{"Red", "Green", "Blue"}
	// colors1 := [3]string{"Red", "Green","Eellow", "Blue"}        //out of bound

	fmt.Println("colors: ", colors)

	for i := 0; i < len(nums); i++ {
		fmt.Println("Index", i, "Value: ", nums[i])
	}

}
