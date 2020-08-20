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

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
