package main

import (
	"os"

	"pault.ag/go/technicolor"
)

func main() {
	output := technicolor.NewWriter(os.Stdout)
	output.Bold().Red().Printf("U")
	output.White().Printf("S")
	output.Cyan().Printf("A\n")
	output.ResetColor().Write([]byte{})
}
