package adwinv

import (
	"math"
	"testing"
)

func TestSimilarity(t *testing.T) {
	v1 := []float64{1.0, 2.0, 3.0}
	v2 := []float64{10.0, 20.0, 30.0}

	if sim := similarity(v1, v2); sim != 1.0 {
		t.Errorf("Similarity should return 1.0, but %f", sim)
	}
}

func TestNormalize(t *testing.T) {
	v := []float64{1.0, 2.0, 3.0}
	actuals := normalize(v)
	expects := []float64{0.2672612419124244, 0.5345224838248488, 0.8017837257372732}
	for i, e := range expects {
		if !almostEqual(e, actuals[i]) {
			t.Errorf("normalize should return %f, but %f", e, actuals[i])
		}
	}
}

func TestOnlineMean(t *testing.T) {
	xs := [][]float64{
		[]float64{1.0, 2.0, 3.0},
		[]float64{2.0, 4.0, 6.0},
		[]float64{3.0, 6.0, 9.0},
	}

	expects := [][]float64{
		[]float64{1.0, 2.0, 3.0},
		[]float64{1.5, 3.0, 4.5},
		[]float64{2.0, 4.0, 6.0},
	}

	mean := make([]float64, 3)
	for i, _ := range xs {
		mean = onlineMean(xs[i], mean, i)
		for j, m := range mean {
			if m != expects[i][j] {
				t.Errorf("onlineMean should return %f, but %f", m, expects[i][j])
			}
		}
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
