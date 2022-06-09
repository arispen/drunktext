package main

import (
	"fmt"

	"github.com/arispen/drunktext/drunktext"
)

func main() {

	drunk := drunktext.MakeDrunk("yesterday I went to the park with my son and my dog", 70)

	fmt.Print(drunk)
}
