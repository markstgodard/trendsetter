package trend

// Trend calculates the slope and y-intercept of a trendline
// Reference: https://classroom.synonym.com/calculate-trendline-2709.html
func Calc(data []float64) (float64, float64) {
	var a, b, c, d, e, f, m, yx float64

	total := float64(len(data))

	if total < 1.0 {
		return 0, 0
	}

	if total == 1.0 {
		return float64(data[0]), 0
	}

	for i, v := range data {
		a += ((float64(i) + 1.0) * float64(v))
	}

	a = total * a

	var b1 float64
	for i, v := range data {
		b1 += (float64(i) + 1.0)
		// e will be sum of all values
		e += float64(v)
	}
	b = b1 * e

	var c1 float64
	for i, _ := range data {
		c1 += (float64(i) + 1.0) * (float64(i) + 1.0)
	}
	c = total * c1

	d = b1 * b1

	m = (a - b) / (c - d)

	f = m * b1

	yx = (e - f) / total

	return m, yx
}
