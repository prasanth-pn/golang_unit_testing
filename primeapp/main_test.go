package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"not prime", 0, false, "0 is not a prime by definition"},
		{"not prime", 1, false, "1 is not a prime by definition"},
		{"not prime", -10, false, "negative number are not prime by definition"},
	}
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expexpected true but got false", e.name)

		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s :expected %s but got %s", e.name, e.msg, msg)
		}
	}

}

// func Test_prompt(t *testing.T){
// //save a copy of os.Stdout
// 	oldOut:=os.Stdout

// 	//create a  read and write pipe
// 	r,w,_:=os.Pipe()
// 	//set os.Stdout to out write pipe
// 	os.Stdout=w
// 	prompt()
// _,=w.Close()
// //reset os.Stdout to what it was before
// os.Stdout=oldOut

// //read the output of out prompt () func from our read pipe
// out,_:=io.RadAll(r)
// //perform our test

// if string (out)!="->"{
// 	t.Errorf()
// }

// }
