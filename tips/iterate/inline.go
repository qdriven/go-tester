package iterate

import (
	"fmt"
	"log"
)

func InlinePrintNumbers(max int) {
	if max < 0 {
		log.Fatalf("'max' is %d, should be >= 0", max)
	}
	for i := 2; i <= max; i += 2 {
		fmt.Printf("%d\n", i)
	}
}

func InlineRun() {
	fmt.Printf("Even numbers up to 8:\n")
	InlinePrintNumbers(8)
	fmt.Printf("Even numbers up to 9:\n")
	InlinePrintNumbers(9)
	fmt.Printf("Error: even numbers up to -1:\n")
	InlinePrintNumbers(-1)
}