package main

import (
	"fmt"

	"github.com/joseflores-io/design-patterns-with-go/creational/singleton"
)

func main() {
	counter := singleton.GetInstance()
	currentCount := counter.AddOne()

	fmt.Println(currentCount)

	counter2 := singleton.GetInstance()
	currentCount = counter2.AddOne()

	fmt.Println(currentCount)

}
