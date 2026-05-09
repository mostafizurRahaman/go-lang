package main

import (
	"fmt"
	"time"
)

// For : only construct in go for looping
//


func main (){ 


	// While Loop (finite)
	// i:= 0

	// for i <= 3 {		
	// 	fmt.Println(i)
	// 	// i=i + 1
	// 	i++
	// 	// i+=1
	// }

	// Infinite while loop
	var startTime = time.Now().UnixMicro()
	j :=1
	for { 


		fmt.Println(j)
		j++

		if(j == 10000){ 
			break
		}

	}
	var endTime = time.Now().UnixMicro()

	fmt.Println("Start Time", startTime)
	fmt.Println("End Time", endTime)
	fmt.Println("Difference", endTime - startTime)


	// NOTE: Classic for loop : 
	fmt.Println("CLASSIC FOR LOOP")
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}


	// CASE : For loop with Continue: 
	fmt.Println("CONTINUE-------->")
	for i:=0; i < 5; i++{ 

		if(i == 2){ 
			continue
		}
		fmt.Println(i)
	}

	// CASE : For loop with braek statement: 
	fmt.Println("<-----------Break statement:-------->")
	for i := 0; i < 10; i++ { 

		if(i == 5){ 
			break
		}
		fmt.Println(i)
	}
	fmt.Println("<-----------Range:-------->")

	// CASE : RANGE  
	for i := range 5 { 
		fmt.Println(i)
	}

	// CASE : (Range) with an array: 
	nums := []int{1,10,5}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// CASE : RANGE FOR String: 
	str := "Mostafizur Rahaman"

	for i,v := range str {  
		fmt.Println("<-----------SEP ERATION:-------->")
		fmt.Println("I => " , i , " V =>", v)
		fmt.Println("I => " , i , " V =>", string(v))		
	}
}

