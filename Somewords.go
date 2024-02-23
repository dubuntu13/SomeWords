package main

import (
	"fmt"
)

// Say Hello and return answer
func Hello(name string) string {
	// Now return answer
	message := fmt.Sprintf("Hi, %v, Welcome!", name)
	return message
}
