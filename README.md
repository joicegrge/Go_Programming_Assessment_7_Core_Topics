# Go Programming Assessment – 7 Independently Runnable Topics

Each topic is a separate Go package with its own `main()` function. You can run and test topics independently and enter your own values when prompted.

## Requirements
- Go 1.22 or later
- VS Code (optional)

## Run one topic at a time
Open a terminal in this project folder and run any one command:

```bash
go run ./topic1
go run ./topic2
go run ./topic3
go run ./topic4
go run ./topic5
go run ./topic6
go run ./topic7
```

Each program prompts you for input. Use normal numeric values when asked for numbers.

## Run the automated tests
```bash
go test ./...
```

Run only Topic 7 tests:
```bash
go test ./topic7 -v
```

Check for data races in Topic 6:
```bash
go run -race ./topic6
```

## Topic guide
1. **Structs, methods and pointers:** enter account holder, opening balance, deposit and withdrawal amounts.
2. **Interfaces and polymorphism:** enter a notification message and the email/phone destinations. Notifications are simulated in the terminal; no real email or SMS is sent.
3. **Error handling:** enter an age and see validation plus `errors.As` type checking.
4. **Defer and file handling:** enter a filename and message. The message is appended with a newline; the file is created if it does not exist.
5. **Goroutines:** enter comma-separated URL labels. Downloads are simulated with a short delay; no network requests are made.
6. **Mutex and race conditions:** enter worker and increment counts. The program reports the expected and actual total.
7. **Modules, packages and testing:** enter a price and discount percentage. A table-driven test suite covers valid and invalid cases.

## Short note: Modules and packages
A **package** is a collection of Go source files in the same directory that are compiled together and provide related functionality. Packages help organize code and make it reusable. The `package` declaration at the top of each `.go` file identifies its package.

A **Go module** is a collection of related Go packages managed as a unit. The `go.mod` file declares the module path and Go version. The module path is used when importing packages within the project or from other modules. Commands such as `go mod init`, `go run`, and `go test` help initialize, run, and test a module.
