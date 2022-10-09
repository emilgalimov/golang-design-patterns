package landlinePhone

type Source struct {
}

func (s *Source) GetRaw() string {
	return "landline-phone-message"
}

func NewSource() *Source {
	return &Source{}
}
