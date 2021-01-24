package iterate

import (
	"fmt"
	"log"
)

type CallBackFunc func(n int) error

func IterateEvenNumbers(max int, cb CallBackFunc) error {
	if max < 0 {
		return fmt.Errorf("'max' is %d, must be >= 0", max)
	}
	for i := 2; i <= max; i += 2 {
		err := cb(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func PrintEvenNumbers(max int) {
	err := IterateEvenNumbers(max, func(n int) error {
		fmt.Printf("%d\n", n)
		return nil
	})
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

func printIt(n int) error {
	fmt.Printf("%d\n", n)
	return nil
}

func PrintEvenNumbersWithDefinedCallBack(max int) {
	err := IterateEvenNumbers(max, printIt)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

func RunCallBackWay() {
	fmt.Printf("Even numbers up to 8:\n")
	PrintEvenNumbers(8)
	fmt.Printf("Even numbers up to 9:\n")
	PrintEvenNumbers(9)
	fmt.Printf("Error: even numbers up to -1:\n")
	PrintEvenNumbers(-1)
	PrintEvenNumbersWithDefinedCallBack(10)
}
