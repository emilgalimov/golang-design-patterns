package mobilePhone

import "strings"

type Processor struct{}

func (Processor) ProcessRaw(rawMessage string) string {
	return strings.Replace(rawMessage, "_", " ", -1)
}

func NewProcessor() *Processor {
	return &Processor{}
}
