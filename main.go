package main

import (
	"time"
	"math/rand"
)

func init(){
	// Make sure that we can generate random numbers
	rand.Seed(time.Now().UnixNano())
}

func main() {

}
