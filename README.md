<p align="center">
    <h1 align="center">GO GEE MODULE BANK CLI APPS</h1>
</p>

GO GEE MODULE BANK CLI APPS is code framework written in Go(https://go.dev/). features in this code framework to help streamline command line application development (BANK CLI Apps)

### Getting GO GEE MODULE CLI APPS
run the following Go command to install package:

```sh
$ go get github.com/geetoor-maven/go-gee-module/v1
```

### Use Go gee module bank cli

First you need to import go gee module package for using module bank cli, one simplest example likes the follow `main.go`:

### displays the sentences good morning, afternoon, evening to the user

```go
package main

import (
	"github.com/geetoor-maven/go-gee-module/v1/helper"
)

func main()  {
	helper.GeeHappyHelper("agus")
}
```

### Input Helper and validation

```go
package main

import (
	"fmt"

	"github.com/geetoor-maven/go-gee-module/v1/helper"
)

func main(){
	input, err := helper.InputHelper("message to the user : ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(input)
}
```

### Login View and set data Feature if login Success

```go
package main

import (
	"github.com/geetoor-maven/go-gee-module/v1/cli"
)

func main()  {
   isLogin := cli.LoginView("agus", 4) // params 1 : password from db, params 2 : set length password 
	if isLogin {
		features := []string{}
		result, err := cli.FeaturesView(features[:])
		if err == nil {
			// validate result 
			if result == "1" {
				fmt.Println("Feature 1")
			}else {
				fmt.Println("Feature 2")
			}			
		}else {
			fmt.Println(err.Error())
		}

	 }
}
```