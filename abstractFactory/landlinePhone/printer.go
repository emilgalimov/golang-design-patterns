package landlinePhone

type Printer struct{}

func (Printer) Print(message string) string {
	return "landline phone says:" + message
}

func NewPrinter() *Printer {
	return &Printer{}
}
