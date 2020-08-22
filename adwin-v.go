package adwinv

import (
	"math"

	"github.com/monochromegane/adwin"
)

type AdwinV struct {
	version         int
	dim             int
	magnitudes      adwin.AdaptiveWindow
	angles          adwin.AdaptiveWindow
	scaleMagnitudes float64
	scaleAngles     float64

	total int
	means []float64

	sync bool
}

func NewAdwinV(dim int, deltaM, deltaA float64) *AdwinV {
	return newAdwinV(dim, deltaM, deltaA, 1)
}

func NewAdwin2V(dim int, deltaM, deltaA float64) *AdwinV {
	return newAdwinV(dim, deltaM, deltaA, 2)
}

func newAdwinV(dim int, deltaM, deltaA float64, version int) *AdwinV {
	var adwinM adwin.AdaptiveWindow
	var adwinA adwin.AdaptiveWindow
	switch version {
	case 1:
		adwinM = adwin.NewAdwin(deltaM)
		adwinA = adwin.NewAdwin(deltaA)
	case 2:
		adwinM = adwin.NewAdwin2(deltaM)
		adwinA = adwin.NewAdwin2(deltaA)
	}
	return &AdwinV{
		version:    version,
		dim:        dim,
		magnitudes: adwinM,
		angles:     adwinA,

		scaleMagnitudes: 1.0,
		scaleAngles:     1.0,
		means:           make([]float64, dim),
	}
}

func (a *AdwinV) Add(x []float64) {
	magnitude := l2Norm(x)
	a.magnitudes.Add(magnitude * a.scaleMagnitudes)

	angle := 1.0 - similarity(a.means, x)
	if math.IsNaN(angle) {
		angle = 0.0
	}
	a.angles.Add(angle * a.scaleAngles)

	a.updateMean(x)

	if a.sync {
		a.syncWindow()
	}
}

func (a *AdwinV) Detected() bool {
	return a.magnitudes.Detected() || a.angles.Detected()
}

func (a *AdwinV) Size() int {
	magnitudesSize := a.magnitudes.Size()
	anglesSize := a.angles.Size()
	if magnitudesSize < anglesSize {
		return magnitudesSize
	}
	return anglesSize
}

func (a *AdwinV) ScaleMagnitudes(scale float64) {
	a.scaleMagnitudes = scale
}

func (a *AdwinV) ScaleAngles(scale float64) {
	a.scaleAngles = scale
}

func (a *AdwinV) SizeMagnitudes() int {
	return a.magnitudes.Size()
}

func (a *AdwinV) SizeAngles() int {
	return a.angles.Size()
}

func (a *AdwinV) Mean() []float64 {
	return a.means
}

func (a *AdwinV) Conservative(t bool) {
	a.magnitudes.Conservative(t)
	a.angles.Conservative(t)
}

func (a *AdwinV) SyncWindow(t bool) {
	a.sync = t
}

func (a *AdwinV) syncWindow() {
	if a.version == 2 {
		if a.magnitudes.Detected() && !a.angles.Detected() {
			a.angles.(*adwin.Adwin2).Drop()
		}
		if !a.magnitudes.Detected() && a.angles.Detected() {
			a.magnitudes.(*adwin.Adwin2).Drop()
		}
	}
}

func (a *AdwinV) updateMean(x []float64) {
	if a.angles.Detected() {
		a.total = 1
		a.means = x
		return
	}
	a.means = onlineMean(x, a.means, a.total)
	a.total += 1
}
