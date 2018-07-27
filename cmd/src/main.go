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

	//in this second method we can see that the map is created and then added too. but whats important to note is that the key and value
	//are of difrent types , the key is of type string, and the value is type int. all keys must be of the same type, and all values must be
	//of the same type. this is a rule of maps.
	//2.
	letters := make(map[string]int)

	letters["a"] = 065
	letters["b"] = 066
	letters["c"] = 067

	fmt.Println(colours)
	fmt.Println(letters)

	//we can also print the valus of a key by useing the key as a refrence
	fmt.Println("The Ascii valus of A is : ", letters["a"])

	printMap(colours)
}

//itterates through the map
func printMap(m map[string]string) {

	//print a single key value pair from piont n in the range of m where m is the length of the map
	for key, value := range m {

		fmt.Printf("the key is : %v, the value is : %v \n", key, value)
	}
}
