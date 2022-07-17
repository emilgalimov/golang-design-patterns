package main

import "fmt"

type concreteService struct {
}

func NewConcreteService() *concreteService {
	return &concreteService{}
}

func (s concreteService) processRequest(request string) string {
	return fmt.Sprintf("Response is <%s>", request)
}
