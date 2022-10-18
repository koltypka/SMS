package main

import (
	"fmt"

	"github.com/koltypka/SMS/system"
)

func main() {
	sys := system.NewSystem(1)
	sys.StartSimulation()

	fmt.Print("\n")
}
