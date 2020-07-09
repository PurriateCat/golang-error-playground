package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type HttpError struct {
	status int
	method string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with %v request. Status: %v",
		httpError.method,
		httpError.status)
}

func CreateSomething() (string, error) {
	return "", &HttpError{404, "GET"}
}

func main() {
	result, err := CreateSomething()

	if err != nil {
		log.Errorln(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
