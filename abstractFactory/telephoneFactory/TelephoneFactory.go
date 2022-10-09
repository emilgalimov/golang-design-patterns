package telephoneFactory

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/Errors"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/landlinePhone"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/mobilePhone"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/phoneInterfaces"
)

type PhoneFactory interface {
	CreateSource() phoneInterfaces.Source
	CreateProcessor() phoneInterfaces.Processor
	CreatePrinter() phoneInterfaces.Printer
}

const (
	MobilePhoneType = iota
	LandlinePhoneType
)

func GetTelephoneFactory(phoneType int) (PhoneFactory, error) {
	switch phoneType {
	case MobilePhoneType:
		return mobilePhone.NewFactory(), nil
	case LandlinePhoneType:
		return landlinePhone.NewFactory(), nil
	default:
		return nil, Errors.NewPhoneTypeNotFoundError(phoneType)
	}
}
