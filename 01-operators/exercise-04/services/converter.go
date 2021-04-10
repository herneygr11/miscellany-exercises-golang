package services

// Converter...
func Converter(amount float64) float64 {
	totalDolares := amount * 1.19
	return float64(int64(totalDolares/0.05+0.5)) * 0.05
}
