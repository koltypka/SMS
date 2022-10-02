package main

import (
	"fmt"

	"github.com/koltypka/SMS/rand"
	"github.com/koltypka/SMS/system"
)

func main() {
	fmt.Print(rand.MakeRavnr(7))
	fmt.Print("\n")

	fmt.Print(system.MakeVisitors())
}
