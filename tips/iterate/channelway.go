package iterate

import (
	"context"
	"fmt"
	"log"
)

// To run:
// go run channel.go

// IntWithError combines an integer value and an error
type IntWithError struct {
	Int int
	Err error
}

/**
TODO: Channel way? What is Channel?
A Generator
*/
func generateEvenNumbers(max int) chan IntWithError {
	ch := make(chan IntWithError)
	go func() { //start channel
		defer close(ch)
		if max < 0 {
			ch <- IntWithError{
				Err: fmt.Errorf("'max' is %d and should be >= 0", max),
			}
			return
		}

		for i := 2; i <= max; i += 2 {
			ch <- IntWithError{
				Int: i,
			}
		}
	}()
	return ch
}

/**
TODO: Channel/Cancellable Channel
 */
func generateEvenNumbersInCancellableChannel(ctx context.Context, max int) chan IntWithError {
	ch := make(chan IntWithError)
	go func() {
		defer close(ch)
		if max < 0 {
			ch <- IntWithError{
				Err: fmt.Errorf("'max' is %d and should be >= 0", max),
			}
			return
		}

		for i := 2; i <= max; i += 2 {
			if ctx != nil {
				// if context was cancelled, we stop early
				select {
				case <-ctx.Done():
					return //exit point
				default: //do nothing
				}
			}
			ch <- IntWithError{
				Int: i,
			}
		}
	}()
	return ch
}

func printEvenNumbersCancellable(max int, stopAt int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := generateEvenNumbersInCancellableChannel(ctx, max)
	for val := range ch {
		if val.Err != nil {
			log.Fatalf("Error: %s\n", val.Err)
		}
		if val.Int > stopAt {
			cancel()
			// notice we keep going in order to drain the channel
			continue
		}
		// process the value
		fmt.Printf("%d\n", val.Int)
	}
}

func printEvenNumbersInChannel(max int) {
	for val := range generateEvenNumbers(max) {
		if val.Err != nil {
			log.Fatalf("Error: %s\n", val.Err)
		}
		fmt.Printf("%d\n", val.Int)
	}
}

func ChannelRun() {
	//fmt.Printf("Even numbers up to 8:\n")
	//printEvenNumbersInChannel(8)
	//fmt.Printf("Even numbers up to 9:\n")
	//printEvenNumbersInChannel(9)
	//fmt.Printf("Error: even numbers up to -1:\n")
	//printEvenNumbersInChannel(-1)
	//
	//fmt.Printf("Even numbers up to 20, cancel at 8:\n")
	printEvenNumbersCancellable(20, 8)
}
