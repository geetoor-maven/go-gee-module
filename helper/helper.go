package helper

import (
	"errors"
	"fmt"
	"time"
)

func GeeHappyHelper(name string){
	happy, err := geeHappyHelper(name)
	if err == nil {
		fmt.Println(happy)
	}else {
		fmt.Println(err.Error())
	}
}

func geeHappyHelper(name string)(string, error){

	timeNow := time.Now();
	hour, _ , _ := timeNow.Clock();

	if name != "" {
		
		if hour > 2 && hour < 10 {
			return "Hello.. Good Morning..." + name, nil;
		}else if hour > 10 && hour < 18 {
			return "Hello.. Good Afternoon..." + name, nil;
		}else{
			return "Hello.. Good Night..." + name, nil;
		}	
	
	}else {
		return "", errors.New("Please set name in params")	
	}
}

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