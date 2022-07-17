package main

type appService interface {
	processRequest(string) string
}
