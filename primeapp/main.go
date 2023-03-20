package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	//print a welcome  message
	intro()

	//create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)
	//start a gorutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)
	//block  utill the donChan gets a value
	<-doneChan
	//close the channel
	close(doneChan)
	//say goodbye
	fmt.Println("good bye brother")

}
func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	//read user input
	scanner.Scan()
	//check to see if the user wants to quit

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	//try to convert what the user typed into an int
	numTocheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "please enter a whole number", false
	}
	_, msg := isPrime(numTocheck)
	return msg, false
}

func intro() {
	p := fmt.Println
	p("it is a prime")
	p("-----------")
	p("Enter a whole number, and we'll tell you if is prime number or not")
	prompt()
}
func prompt() {
	fmt.Print("->")
}
func isPrime(n int) (bool, string) {
	//0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime by definition", n)
	}
	//negative numbers are not prime
	if n < 0 {
		return false, "negative number are not prime by definition"
	}
	//use the modules operator repeadtedly to see if we have prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)

		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)

}
