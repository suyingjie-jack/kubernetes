package main

import (
	"fmt"
)

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	//myArray := [5]int{1, 2, 3, 4, 5}
	myArray[2] = "smart"
	myArray[4] = "strong"
	for _, char := range myArray {
		fmt.Println(char)

	}

}
