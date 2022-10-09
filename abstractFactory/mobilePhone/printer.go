package mobilePhone

type Printer struct {
}

func (*Printer) Print(message string) string {
	return "mobile phone says: " + message
}

func NewPrinter() *Printer {
	return &Printer{}
}
