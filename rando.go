package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seedPtr := flag.Int64("seed", 0, "an integer to seed the process with")

	flag.Parse()

	var seed rand.Source

	if *seedPtr == 0 {
		seed = rand.NewSource(time.Now().UnixNano())
	} else {
		seed = rand.NewSource(*seedPtr)
	}

	r := rand.New(seed)

	fmt.Println(r.Int())
}
