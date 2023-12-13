package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

// fetchThirdPartyStuff retrieves third-party data.
func fetchThirdPartyStuff() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil
}

// fetchUserData retrieves user data based on the userID.
func fetchUserData(ctx context.Context, userID int) (int, error) {
	key := "key"
	val := ctx.Value(key)
	fmt.Println(val.(string))

	// Set a timeout of 200 ms.
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel() // Prevent context leakage.

	respChannel := make(chan Response)

	// Asynchronously fetch third-party data.
	go func() {
		val, err := fetchThirdPartyStuff()
		respChannel <- Response{
			value: val,
			err:   err,
		}
	}()

	// Wait for either context cancellation or data retrieval.
	select {
	case <-ctx.Done():
		return -1, fmt.Errorf("exceeded time limit for fetching third-party data")
	case response := <-respChannel:
		return response.value, response.err
	}
}

func main() {
	start := time.Now()

	// Create a starting context and add a value, so we can trace in the future for logging, debugging stuff
	ctx := context.WithValue(context.Background(), "key", "value")
	userID := 10

	// Fetch user data.
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	// Print the results.
	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}
