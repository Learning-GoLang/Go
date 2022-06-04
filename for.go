package main

import "fmt"

func main() {

	space_1 := " "

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(space_1)
	for j := 5; j <= 11; j++ {
		fmt.Println(j)
	}
	fmt.Println(space_1)
	for n := 3; n <= 20; n++ {
		if n%3 == 0 {
			fmt.Println(n)
		}
		continue
	}
}
