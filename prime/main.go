package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneCh := make(chan struct{})
	defer close(doneCh)

	go readUserInput(doneCh)

	<-doneCh
	fmt.Println("Goodbye.")
}

func readUserInput(doneCh chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneCh <- struct{}{}
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number.", false
	}

	_, msg := isPrime(num)
	return msg, false
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we will tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition.", n)
	}

	if n < 0 {
		return false, "Negative numbers are not prime, by definition."
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d.", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number.", n)
}
