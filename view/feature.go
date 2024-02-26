package view

import "fmt"

func VariadicFeatureDashboard(features ...string){
	for i, feature := range features{
		fmt.Println((i+1), feature)
	}
}

