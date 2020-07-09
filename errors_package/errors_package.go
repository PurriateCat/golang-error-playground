package main

import "fmt"
import "errors"

func main() {
	//The errors.New function expect an error message and returns an error (from type error)
	err := errors.New("Huupsie...")
	fmt.Println(err)

	//Check the type of variable err
	//Type of err is *errors.errorString
	fmt.Printf("Type of err: %T\n", err)

	//Behind the scenes of errors.New() and *errors.errorString:
	//errors.New returns a pointer to the errors.errorString struct.
	//   func New(text string) error {
	//	     return &errorString{text}
	//   }

	//The errorString struct is implemented like this:
	//   type errorString struct {
	//       s string
	//   }
	//
	//   func (e *errorString) Error() string {
	//       return e.s
	//   }
}
