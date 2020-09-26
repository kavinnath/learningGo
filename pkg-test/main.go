package main

import (
	"fmt"

	"github.com/kavinnath/learningGo/librarypkg"
)

func main() {
	fmt.Println("This is from main pkg")
	hello()
	bye()
	fmt.Println(tell)
	fmt.Println(librarypkg.Test)
}
