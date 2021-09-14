package main

import "fmt"

func Calculator(x int) (result int) {
	result = x + 2
	return result
}

func main()  {
	fmt.Println("Go Testing tutorial")
	result := Calculator(2);
	fmt.Println(result)
}
