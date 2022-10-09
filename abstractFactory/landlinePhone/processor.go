package landlinePhone

import "strings"

type Processor struct{}

func (Processor) ProcessRaw(rawMessage string) string {
	return strings.Replace(rawMessage, "-", " ", -1)
}

func NewProcessor() *Processor {
	return &Processor{}
}
