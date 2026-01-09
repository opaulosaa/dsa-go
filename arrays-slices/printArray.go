package main

import "fmt"

func printArray(list []int) {
	for _, item := range list {
		fmt.Printf("%d\t", item) //outra maneira de iterar o array/slice
	}

	//for i = 0; i < len(slice); i++ {
	// 		fmt.Printf("%d \t", slice[i])
	//}
}
