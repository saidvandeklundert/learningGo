package main

import "fmt"

func main() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)
	// allows one-liners with append:
	// s := append([]int(nil), 42)
	
	s = []string{} // recommended approach to creating slices with initial elements
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	/*
	   1: empty=true   nil=true
	   2: empty=true   nil=true
	   3: empty=true   nil=false
	   4: empty=true   nil=false

	   This can have a profound impact when serializing to JSON for instance.

	   
	   The slices that equal nil do not require any allocation.

	   You are able to use append to add to them. When returning a slice that is empty,
	   it is favored to return the slice that equals nil.

	   If we should produce an empty slice with a known lenght, then consider .4
	*/
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}
