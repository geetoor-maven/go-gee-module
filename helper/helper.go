package helper

import (
	"errors"
	"fmt"
)

// function for take input
func InputHelper(message string) (interface{}, error){

	fmt.Print(message)
	var result string
	fmt.Scanln(&result);

	if len(result) == 0 {
		return nil, errors.New("Please input data....")
	}else {
		return result, nil	
	}
}

func LineBarrier(){
	fmt.Println("----------------------------")
}