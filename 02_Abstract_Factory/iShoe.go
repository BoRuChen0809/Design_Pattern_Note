package main

type iShoe interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shoe struct {
	logo string
	size string
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size string) {
	s.size = size
}

func (s *shoe) getSize() string {
	return s.size
}
