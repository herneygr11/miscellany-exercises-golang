package main

type Triangle struct {
	base   float64
	height float64
}

// Area ...
func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
