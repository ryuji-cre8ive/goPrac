package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main () {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	
	switch s + 1 {
	case 6:
		fmt.Println("大吉")
	case 4,5:
		fmt.Println("中吉")
	case 2,3:
		fmt.Println("小吉")
	default:
		fmt.Println("凶")
	}
}