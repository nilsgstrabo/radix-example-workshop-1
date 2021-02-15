package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("computing")
	time.Sleep(time.Duration(5 * time.Second))
	fmt.Println("done")
}
