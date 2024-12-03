package main

import "fmt"

/*
All the things slices
*/
func main() {

	s := make([]int, 3, 6) // create a three-lenght, six-capacity slice. 6 capacity is allocated
	// 0 0 0 x x x
	s[1] = 1
	// 0 1 0 x x x
	s = append(s, 2)
	// 0 1 0 2 x x
	s = append(s, 3, 4, 5) // slice automatically grows (doubles) the backing array
	// 0 1 0 2 3 4 5
	fmt.Println(s)
	s2 := s[1:3] // the first index is included, whereas the second is excluded
	// both slices reference the same backing array!!
	fmt.Println(s2)
	// 1 0
	s2 = append(s2, 51)
	// operations on s2 impact the original s as well:
	fmt.Println(s)
	fmt.Println(s2)
	// s: [0 1 0 51 3 4 5]
	// s2: [1 0 51]
	// the backing array is internal and not available
	// directly to the Go developer. The only exception is when a slice is created
	// from slicing an existing array.

	/*
		If you expand s2 and force it to grow, it will double the backing array
		BUT
		only for s2. Additionally, the index is changed.

		The original slice will continue to use the original backing array.
	*/
	s1 := make([]int, 3, 4)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s2 = s1[0:1]
	s2 = append(s2, 222)
	s2 = append(s2, 333)
	s2 = append(s2, 444)
	s2 = append(s2, 555)
	s2 = append(s2, 666)
	s2 = append(s2, 777)
	s2 = append(s2, 888)
	s1 = append(s1, 999)
	fmt.Println(s1)
	fmt.Println(s2)
	// both s1 and s2 now have a different array. Notice how '999' is not in s2:
	// s1: [1 222 333 999]
	// s2: [1 222 333 444 555 666 777 888]

}
