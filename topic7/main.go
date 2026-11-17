package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateDiscount(price float64, percentage float64) (float64, error) {
	if price < 0 {
		return 0, fmt.Errorf("price cannot be negative")
	}
	if percentage < 0 || percentage > 100 {
		return 0, fmt.Errorf("discount percentage must be between 0 and 100")
	}
	return price - (price * percentage / 100), nil
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil && strings.TrimSpace(line) == "" {
		return 0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(line), 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 7: MODULES, PACKAGES & DISCOUNT")
	price, err := readFloat(reader, "Enter original price: ")
	if err != nil {
		fmt.Println("Invalid price:", err)
		return
	}
	percentage, err := readFloat(reader, "Enter discount percentage (0-100): ")
	if err != nil {
		fmt.Println("Invalid percentage:", err)
		return
	}

	discountedPrice, err := CalculateDiscount(price, percentage)
	if err != nil {
		fmt.Println("Discount error:", err)
		return
	}
	fmt.Printf("Original price: %.2f\nDiscount: %.2f%%\nFinal price: %.2f\n",
		price, percentage, discountedPrice)
	fmt.Println("Run `go test ./topic7 -v` to execute the table-driven tests.")
}
