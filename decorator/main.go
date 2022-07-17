package main

import "fmt"

func main() {

	service := NewConcreteService()

	logRequestService := NewLogRequestDecorator(service, &logger{})
	logRequestService.processRequest("service with request logger")

	logRequestResponseService := NewLogResponseDecorator(logRequestService, &logger{})
	logRequestResponseService.processRequest("service with request and response logger")

}

type logger struct {
}

func (l *logger) print(message string) {
	fmt.Println(message)
}
