package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit uses a pointer receiver so the account's balance is updated.
// Negative or zero deposits are ignored.
func (account *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit ignored: amount must be positive.")
		return
	}
	account.Balance += amount
}

// Withdraw uses a pointer receiver and returns an error if funds are insufficient.
func (account *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > account.Balance {
		return fmt.Errorf("insufficient funds: available %.2f, requested %.2f", account.Balance, amount)
	}
	account.Balance -= amount
	return nil
}

func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		value := readLine(reader, prompt)
		number, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 1: STRUCTS, METHODS & POINTERS")

	owner := readLine(reader, "Enter account holder name: ")
	openingBalance := readFloat(reader, "Enter opening balance: ")
	account := BankAccount{Owner: owner, Balance: openingBalance}

	deposit := readFloat(reader, "Enter deposit amount: ")
	account.Deposit(deposit)
	fmt.Printf("Balance after deposit: %.2f\n", account.Balance)

	withdrawal := readFloat(reader, "Enter withdrawal amount: ")
	if err := account.Withdraw(withdrawal); err != nil {
		fmt.Println("Withdrawal error:", err)
	} else {
		fmt.Printf("Balance after withdrawal: %.2f\n", account.Balance)
	}
	fmt.Printf("Final balance for %s: %.2f\n", account.Owner, account.Balance)
}
