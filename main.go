package main

import "fmt"

func main () {
	x:=[5]float64{
	98,
	93,
	77,
	82,
	83,
}
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

