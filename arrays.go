package main

import "fmt"

func main() {

	var a [9]int
	fmt.Println("empty:", a)

	a[4] = 100
	a[6] = 20
	a[8] = 19
	fmt.Println("set array:", a)
	fmt.Println("get array:", a[4], a[6], a[8])

	fmt.Println("length:", len(a))

	var matris [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			matris[i][j] = i + j
		}
	}
	fmt.Println(matris)
}
