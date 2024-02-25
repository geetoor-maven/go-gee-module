package helper

import "fmt"

// function for take input
func InputHelper(message string) string{
	fmt.Print(message)
	var result string;
	fmt.Scanln(&result);
	return result 
}

func Line(){
	fmt.Println("----------------------------")
}