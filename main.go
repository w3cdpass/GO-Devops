package main

import "fmt"

func main(){

	// String in GO 
	doubleQuoteString := "Hello Go!";

	// Rune ? 
	singleQouteString := 'G';

	// output 

	fmt.Println(doubleQuoteString, singleQouteString) 
	// output := Hello Go! 71

	fmt.Println(doubleQuoteString, string(singleQouteString)) 
	// Hello Go! G
};