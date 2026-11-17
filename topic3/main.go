package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field  string
	Reason string
}

func (validationError *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", validationError.Field, validationError.Reason)
}

func ValidateAge(age int) error {
	if age < 0 || age > 120 {
		return &ValidationError{Field: "age", Reason: "must be between 0 and 120"}
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 3: ERROR HANDLING")
	fmt.Print("Enter age: ")
	input, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Input error: please enter a whole number.")
		return
	}

	err = ValidateAge(age)
	if err == nil {
		fmt.Println("Age is valid.")
		return
	}

	fmt.Println("Validation failed:", err)
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("errors.As confirmed a ValidationError.")
		fmt.Println("Field:", validationErr.Field)
		fmt.Println("Reason:", validationErr.Reason)
	}
}
