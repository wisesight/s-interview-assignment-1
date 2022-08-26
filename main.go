package main

import (
	"fmt"
	"math/rand"
	"time"
)

func checkAccountExist(accountID string) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}

func main() {
	fmt.Println("Hi K.Prem")
}
