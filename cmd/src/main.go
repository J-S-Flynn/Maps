package main

import (
	"fmt"
)

//the piont of this file is to deal with key value pairs
//or maps. which are not disimiler to objects in jscript
func main() {

	//how we declare maps
	//we declare the map wit a key and a value
	//1. key of type string , value of type string
	colours := map[string]string{
		"red":   "ff0000",
		"green": "008000",
		"blue":  "0000ff",
	}

	fmt.Println(colours)
}
