package services

type Square struct {
	Side float64
}

// Area...
func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (S *Square) Perimeter() float64 {
	return S.Side * 4
}
