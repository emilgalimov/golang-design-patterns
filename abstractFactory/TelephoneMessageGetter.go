package abstractFactory

import "fmt"

func writeMessageProcessHistory(factory PhoneFactory) string {
	source := factory.createSource()
	processor := factory.createProcessor()
	printer := factory.createPrinter()

	history := ""

	rawMessage := source.GetRaw()
	history += fmt.Sprintln(rawMessage)

	processedMessage := processor.ProcessRaw(rawMessage)
	history += fmt.Sprintln(processedMessage)

	printerMessage := printer.Print(processedMessage)
	history += fmt.Sprintln(printerMessage)
}
