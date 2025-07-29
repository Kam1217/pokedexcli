package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string{
	/*
	The purpose of this function will be to split the user's input into "words" based on whitespace. 
	It should also lowercase the input and trim any leading or trailing whitespace.
	*/
	lowerStr := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(lowerStr)
	
	return words
}

func main(){
	fmt.Println("Hello, World!")
}
