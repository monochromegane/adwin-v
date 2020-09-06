# ADWIN-V [![Actions Status](https://github.com/monochromegane/adwin-v/workflows/Go/badge.svg)](https://github.com/monochromegane/adwin-v/actions)

ADWIN-V is an adaptive windowing algorithm for vector data.
It detects a change point using the vector magnitude and cosine similarity between mean vector and the vector.
The change detection algorithm is ADWIN or ADWIN2 which is from `Learning from time-changing data with adaptive windowing, Bifet, Albert, and Ricard Gavalda; Proceedings of the 2007 SIAM international conference on data mining. Society for Industrial and Applied Mathematics, 2007`.

See also:
- https://www.cs.upc.edu/~gavalda/DataStreamSeminar/files/Lecture6.pdf

## Usage

### ADWIN-V

ADWIN-V uses ADWIN for change detection of vector magnitude and cosine similarity.

```go
        dim := 2
	deltaM := 0.01
	deltaA := 0.01
	adwinv := NewAdwinV(dim, deltaM, deltaA)
	adwinv.Conservative(true) // if you wants
	adwinv.ScaleMagnitudes(0.1) // if you wants
	adwinv.ScaleAngles(1.0) // if you wants

	// Add stream data
	adwinv.Add([]float64{1.0, 2.0})

	// Show status
	adwinv.Size()
	adwinv.SizeMagnitudes()
	adwinv.SizeAngles()
	adwinv.Mean()
	adwinv.Detected()
```

### ADWIN2-V

ADWIN2-V uses ADWIN2 for change detection of vector magnitude and cosine similarity.
- Note: This version provides `SyncWindow` option to synchronize the bucket size of the exponential histogram.

```go
        dim := 2
	deltaM := 0.01
	deltaA := 0.01
	adwinv := NewAdwin2V(dim, deltaM, deltaA)
	adwinv.SyncWindow(true) // if you wants
	adwinv.Conservative(true) // if you wants
	adwinv.ScaleMagnitudes(0.1) // if you wants
	adwinv.ScaleAngles(1.0) // if you wants

	// Add stream data
	adwinv.Add([]float64{1.0, 2.0})

	// Show status
	adwinv.Size()
	adwinv.SizeMagnitudes()
	adwinv.SizeAngles()
	adwinv.Mean()
	adwinv.Detected()
```

## Installation

```sh
$ go get github.com/monochromegane/adwin-v
```

## License

[MIT](https://github.com/monochromegane/adwin-v/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
