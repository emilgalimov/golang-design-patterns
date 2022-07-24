package abstractFactory

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/Errors"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/landlinePhone"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/mobilePhone"
)

type Source interface {
	GetRaw() string
}

type Processor interface {
	ProcessRaw(rawMessage string) string
}

type Printer interface {
	Print(message string) string
}

type PhoneFactory interface {
	createSource() Source
	createProcessor() Processor
	createPrinter() Printer
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
