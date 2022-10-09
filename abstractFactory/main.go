package main

import (
	"fmt"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/telephoneFactory"
)

func main() {
	factoryLandline, _ := telephoneFactory.GetTelephoneFactory(telephoneFactory.LandlinePhoneType)
	landlineProcessingHistory := writeMessageProcessHistory(factoryLandline)
	fmt.Println(landlineProcessingHistory)

	factoryMobile, _ := telephoneFactory.GetTelephoneFactory(telephoneFactory.MobilePhoneType)
	mobileProcessingHistory := writeMessageProcessHistory(factoryMobile)
	fmt.Println(mobileProcessingHistory)

	_, err := telephoneFactory.GetTelephoneFactory(3)

	fmt.Sprintln(err)
}

func writeMessageProcessHistory(factory telephoneFactory.PhoneFactory) string {
	source := factory.CreateSource()
	processor := factory.CreateProcessor()
	printer := factory.CreatePrinter()

	history := ""

	rawMessage := source.GetRaw()
	history += fmt.Sprintln(rawMessage)

	processedMessage := processor.ProcessRaw(rawMessage)
	history += fmt.Sprintln(processedMessage)

	printerMessage := printer.Print(processedMessage)
	history += fmt.Sprintln(printerMessage)

	return history
}
