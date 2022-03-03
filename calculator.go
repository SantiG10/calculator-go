package main

import (
	"fmt"
	"strconv"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func main() {
	strAdd := strconv.Itoa(Add(5, 5))

	fmt.Println(strAdd)
	fmt.Println(strconv.Itoa(Subtract(10, 5)))

}
