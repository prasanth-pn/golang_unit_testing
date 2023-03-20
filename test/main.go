package main

import (
	"fmt"
)

func main() {

	// scanner := bufio.NewScanner(os.Stdin)
	// s := scanner
	// fmt.Println("the text is :", s.Text())
	// s.Scan()
	// fmt.Println("scanned data is :", s.Text())
	// if strings.EqualFold(s.Text(), "q") {
	// 	fmt.Println("recieved")
	// }
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "please enter a whole number"}}
	for i, v := range tests {
		fmt.Println(v.name)
		fmt.Println(i, v)
	}
}
