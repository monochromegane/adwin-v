package adwinv

import "math"

func dotProduct(a, b []float64) float64 {
	p := 0.0
	for i, _ := range a {
		p += a[i] * b[i]
	}
	return p
}

func l2Norm(x []float64) float64 {
	p := dotProduct(x, x)
	return math.Sqrt(p)
}

func similarity(a, b []float64) float64 {
	p := dotProduct(a, b)
	magA := l2Norm(a)
	magB := l2Norm(b)

	return p / (magA * magB)
}

func normalize(x []float64) []float64 {
	normalized := make([]float64, len(x))
	l2norm := l2Norm(x)
	for i, _ := range x {
		normalized[i] = x[i] / l2norm
	}
	return normalized
}
