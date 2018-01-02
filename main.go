package main

import (
	"httpClient/agency"
	"fmt"
	"time"
)


func main()  {

	for{
			t1 := time.Now()
			agency.Naive()
			elapsed1 := time.Since(t1)
			fmt.Println("Naive Approach Time: ", elapsed1)

			t2 := time.Now()
			agency.Optimized()
			elapsed2 := time.Since(t2)
			fmt.Println("Optimized Approach Time: ", elapsed2)

      break

		}

}
