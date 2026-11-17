package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Notifier interface {
	SendNotification(message string) error
}

type EmailNotifier struct {
	EmailAddress string
}

func (notifier EmailNotifier) SendNotification(message string) error {
	if strings.TrimSpace(notifier.EmailAddress) == "" {
		return fmt.Errorf("email address cannot be empty")
	}
	fmt.Printf("Simulated email to %s: %s\n", notifier.EmailAddress, message)
	return nil
}

type SMSNotifier struct {
	PhoneNumber string
}

func (notifier SMSNotifier) SendNotification(message string) error {
	if strings.TrimSpace(notifier.PhoneNumber) == "" {
		return fmt.Errorf("phone number cannot be empty")
	}
	fmt.Printf("Simulated SMS to %s: %s\n", notifier.PhoneNumber, message)
	return nil
}

var _ Notifier = EmailNotifier{}
var _ Notifier = SMSNotifier{}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 2: INTERFACES & POLYMORPHISM")
	message := ask(reader, "Enter notification message: ")
	email := EmailNotifier{EmailAddress: ask(reader, "Enter recipient email: ")}
	phone := SMSNotifier{PhoneNumber: ask(reader, "Enter recipient phone number: ")}

	// Both concrete types are used through the same Notifier interface.
	var notifiers []Notifier = []Notifier{email, phone}
	for _, notifier := range notifiers {
		if err := notifier.SendNotification(message); err != nil {
			fmt.Println("Notification error:", err)
		}
	}
	fmt.Println("Note: notifications are simulated; no real email or SMS is sent.")
}
