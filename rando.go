package main

import (
	"flag"
	"fmt"
	"os"
	"rando/characters"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Printf("The operation failed with the following error: %s\n", e)
		os.Exit(1)
	}
}

func main() {

	charCmd := flag.NewFlagSet("character", flag.ExitOnError)
	charSeed := charCmd.Int64("seed", 0, "an integer to regenerate a character from")

	if len(os.Args) < 2 {
		println("Expected a subcommand.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "character":
		charCmd.Parse(os.Args[2:])
		var seed int64
		if *charSeed == 0 {
			seed = time.Now().UnixNano()
		} else {
			seed = *charSeed
		}

		characters.GenerateCharacter(seed)
	}
}
