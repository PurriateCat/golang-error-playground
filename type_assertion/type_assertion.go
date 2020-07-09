package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type HttpError struct {
	status  int
	method  string
	message string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("%v | %v | %v",
		httpError.method, httpError.status, httpError.message)
}

type InternalServerError struct {
	HttpError
}

func NewInternalServerError(method string) *InternalServerError {
	return &InternalServerError{HttpError{
		status:  500,
		method:  method,
		message: "Internal server error",
	}}
}

type MethodNotAllowedError struct {
	HttpError
}

func NewMethodNotAllowedError(method string) *MethodNotAllowedError {
	return &MethodNotAllowedError{HttpError{
		status:  405,
		method:  method,
		message: "Method not allowed",
	}}
}

func mockErrorRequest(method string) (string, error) {
	switch method {
	case "GET":
		return "", NewMethodNotAllowedError(method)
	case "POST":
		return "", NewInternalServerError(method)
	default:
		return "SUCCESS", nil
	}
}

func main() {
	result, err := mockErrorRequest("GET")
	//result, err := mockErrorRequest("POST")
	//result, err := mockErrorRequest("PATCH")

	switch err.(type) {
	case nil:
		fmt.Println(result)
		os.Exit(0)
	case *MethodNotAllowedError:
		log.Errorln("Wrong HTTP Method: " + err.Error())
		os.Exit(1)
	case *InternalServerError:
		log.Errorln("Error: " + err.Error())
		os.Exit(2)
	}
}
