package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Printf("The operation failed with the following error: %s\n", e)
		os.Exit(1)
	}
}

func main() {
	seed := flag.Int64("seed", 0, "an integer to seed the process with")
	flag.Parse()

	var s int64
	if *seed == 0 {
		s = time.Now().UnixNano()
	} else {
		s = *seed
	}

	r := rand.New(rand.NewSource(s))

	fmt.Println(s)
	fmt.Println(r.Int())
}
