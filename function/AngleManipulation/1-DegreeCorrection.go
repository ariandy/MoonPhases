package AngleManipulation

func DegreeCorrection(x float64) float64 {
	for x < 0 {
		x += 360
	}

	for x > 360 {
		x -= 360
	}

	return x
}