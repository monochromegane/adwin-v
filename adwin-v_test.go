package adwinv

import (
	"testing"
)

func TestAdwin2VSyncWindowByAngles(t *testing.T) {
	// If SyncWindow flag is set,
	// Adwin2 sync window size of magnitudes to window size of angles when it detects angle change.
	adwinv := NewAdwin2V(1, 0.1, 0.1)
	adwinv.SyncWindow(true)

	for i := 0; i < 100; i++ {
		adwinv.Add([]float64{1.0})
	}
	for i := 0; i < 100; i++ {
		adwinv.Add([]float64{-1.0})
		if adwinv.Detected() {
			if !adwinv.angles.Detected() || adwinv.magnitudes.Detected() {
				t.Errorf("Adwin2 should detect angle change only")
			}
			if adwinv.SizeAngles() != adwinv.SizeMagnitudes() {
				t.Errorf("Adwin2 should sync two window size")
			}
			break
		}
	}
}

func TestAdwin2VSyncWindowByMagnitudes(t *testing.T) {
	// If SyncWindow flag is set,
	// Adwin2 sync window size of angles to window size of magnitudes when it detects angle change.
	adwinv := NewAdwin2V(1, 0.1, 0.1)
	adwinv.SyncWindow(true)

	for i := 0; i < 100; i++ {
		adwinv.Add([]float64{1.0})
	}
	for i := 0; i < 100; i++ {
		adwinv.Add([]float64{2.0})
		if adwinv.Detected() {
			if adwinv.angles.Detected() || !adwinv.magnitudes.Detected() {
				t.Errorf("Adwin2 should detect magnitudes change only")
			}
			if adwinv.SizeAngles() != adwinv.SizeMagnitudes() {
				t.Errorf("Adwin2 should sync two window size")
			}
			break
		}
	}
}
