package main

type Short interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo()  string
	GetSize()  int
}
type short struct {
	logo string
	size int
}

func (s *short)SetLogo(logo string){
	s.logo = logo
}
func (s *short)SetSize(size int){
	s.size = size
}

func (s *short)GetLogo()string{
	return s.logo
}

func (s *short)GetSize()int{
	return s.size
}




