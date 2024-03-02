package cli

import (
	"errors"
	"fmt"

	"github.com/geetoor-maven/go-gee-module/v1/helper"
)

func LoginView(pwdInDB string, lengthPwd int)bool{
	isValidate := false;
	for {
		inputPwd, err := helper.InputHelper("Enter your password card number : ")
		if err == nil {
			if validateLogin(pwdInDB, inputPwd.(string), lengthPwd) {
				isValidate = true
				break
			}
		}
	}
	return isValidate
}

func FeaturesView(features []string)(string, error){

	if len(features) == 0 {
		return "", errors.New("Array Data is Empty in FeaturesView(...)");
	}

	for i, feature := range features{
		fmt.Println((i+1), feature)
	}

	input, err := helper.InputHelper("Choose a Number : ")
	if err == nil {
		return input.(string), nil;
	}
	return "", nil;
}

func ViewHistoryTransaction(historyTransaction []string){
	if len(historyTransaction) == 0 {
		fmt.Println("you haven't made a transaction yet")
	}else {
		for i := 0; i < len(historyTransaction); i++ {
			if historyTransaction[i] != "" {
				fmt.Println("- ", ":", historyTransaction[i]);
			}
		}
	}
}

func ViewBalance(balance interface{} ){
	fmt.Println("Your total balance : ", balance)
}


func validateLogin(pwd string, inputPwd string, length int)bool{
	return validatePWD(pwd, inputPwd, length)
}

func validatePWD(pwd string ,pwdValidate string, length int)bool{
	isValidate := false
	
	if  len(pwdValidate) < length {
		fmt.Println("The password must not be less than " , length , " numbers")
	}else if pwdValidate == pwd {
		isValidate = true
	}else {
		fmt.Println("Your password is wrong")
	}

	return isValidate
}

