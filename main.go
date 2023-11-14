package main

//factored import statement
import (
	"fmt"
	"math"
)

var isTrue bool = true //class global var

func main() {
	var i int = 27 //local var
	j := 42 //implicit short var declaration
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(add(2, 3))
	fmt.Println(split(16))
	fmt.Println(math.Pi) // Exported Name (Public)
	fmt.Printf("%d is %t, %d",i,isTrue, j)
}

// some function, shortened function parameters
func add(x, y int) int {
	return x + y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum
	return
}
