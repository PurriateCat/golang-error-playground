package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type HttpSimpleError struct {
	status int
	method string
}

func (httpError *HttpSimpleError) Error() string {
	return fmt.Sprintf("Something went wrong with %v request. Status: %v",
		httpError.method,
		httpError.status)
}

func CreateSomething() (string, error) {
	return "", &HttpSimpleError{404, "GET"}
}

func main() {
	result, err := CreateSomething()

	if err != nil {
		log.Errorln(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
