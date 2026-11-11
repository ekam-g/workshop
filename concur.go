package main

import (
	"fmt"
	"time"
)

func sendSpam(user string) {
	fmt.Printf("Starting email to %s...\n", user)
	time.Sleep(1 * time.Second)
	fmt.Printf("EXPLODED email to %s!\n", user)
}

func main() {
	start := time.Now()
	users := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	// TODO to make fast
	for _, user := range users {
		sendSpam(user)
	}

	fmt.Printf("All emails sent in %v\n", time.Since(start))
}
