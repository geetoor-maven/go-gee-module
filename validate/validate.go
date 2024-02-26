package validate

import "fmt"

func ValidatePWD(pwd string ,pwdValidate string, length int)bool{
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