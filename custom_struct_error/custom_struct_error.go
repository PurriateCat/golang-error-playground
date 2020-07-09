package main

import (
	"fmt"
)

//This is my custom error as a struct
type MyCustomError struct{}

//This just implements the method Error() of interface error in MyCustomError
func (myCustomError *MyCustomError) Error() string {
	return "Something gone horribly wrong, yo"
}

func main() {
	myCstmErr := &MyCustomError{}
	fmt.Println(myCstmErr)
}
