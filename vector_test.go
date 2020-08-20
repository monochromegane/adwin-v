package adwinv

import "testing"

func TestSimilarity(t *testing.T) {
	v1 := []float64{1.0, 2.0, 3.0}
	v2 := []float64{10.0, 20.0, 30.0}

	if sim := similarity(v1, v2); sim != 1.0 {
		t.Errorf("Similarity should return 1.0, but %f", sim)
	}
}
