package main

type iShirt interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shirt struct {
	logo string
	size string
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size string) {
	s.size = size
}

func (s *shirt) getSize() string {
	return s.size
}
