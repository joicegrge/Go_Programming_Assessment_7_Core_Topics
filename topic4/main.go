package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AppendMessage(filename, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, message); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 4: DEFER & FILE HANDLING")
	filename := ask(reader, "Enter filename (example: notes.txt): ")
	if filename == "" {
		fmt.Println("Filename cannot be empty.")
		return
	}
	message := ask(reader, "Enter message to append: ")

	if err := AppendMessage(filename, message); err != nil {
		fmt.Println("File error:", err)
		return
	}
	fmt.Printf("Message appended successfully to %q.\n", filename)
	fmt.Println("Run the program again with the same filename to append another line.")
}
