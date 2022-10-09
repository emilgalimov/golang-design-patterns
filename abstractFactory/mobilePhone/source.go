package mobilePhone

type Source struct {
}

func (s *Source) GetRaw() string {
	return "mobile_phone_message"
}

func NewSource() *Source {
	return &Source{}
}
