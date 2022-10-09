package phoneInterfaces

type Source interface {
	GetRaw() string
}

type Processor interface {
	ProcessRaw(rawMessage string) string
}

type Printer interface {
	Print(message string) string
}
