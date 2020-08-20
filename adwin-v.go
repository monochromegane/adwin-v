package adwinv

import "github.com/monochromegane/adwin"

type AdwinV struct {
	dim        int
	magnitudes adwin.AdaptiveWindow
	angles     adwin.AdaptiveWindow

	total int
	means []float64
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
		dim:        dim,
		magnitudes: adwinM,
		angles:     adwinA,

		means: make([]float64, dim),
	}
}

func (a *AdwinV) Add(x []float64) {
	magnitude := l2Norm(x)
	a.magnitudes.Add(magnitude)

	angle := 1.0 - similarity(a.means, x)
	a.angles.Add(angle)

	a.updateMean(x)
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

func (a *AdwinV) Conservative(t bool) {
	a.magnitudes.Conservative(t)
	a.angles.Conservative(t)
}

func (a *AdwinV) Sum() float64 {
	return 0.0
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
