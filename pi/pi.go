// Package pi contains Daily Coding Problem #14
//
// This problem was asked by Google.
//
// The area of a circle is defined as πr^2. Estimate π to
// 3 decimal places using a Monte Carlo method.
//
// Hint: The basic equation of a circle is x2 + y2 = r2.
//
// Note: There is problem with cancellation when I try to
// estimate the error, even by using batches of say 100.
// I do the horrible inverse-crime, and estimate with the
// exact value for pi (math.Pi), which makes the method
// completely useless, except for showing that you can
// estimate π by picking uniformly random points (x,y)
// and counting the fraction of points that are inside a
// circle of radius 1.
// Note2: We could try to use the central limit theorem,
// that loosly states that the variance is sigma^2/n.
// In other words, the precision scales as 1/sqrt{n}.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package pi

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Approx returns an approximation of Pi
func Approx() float64 {
	return approx(0.0005, 100000000)
}

func ApproxBatch(batchSize int) float64 {
	return approx(0.0005, batchSize)
}

func approx(prec float64, batchSize int) float64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	accPi := batch(batchSize, rnd)
	for cnt := 1.0; math.Abs(accPi-math.Pi) > prec; cnt += 1.0 {
		pi := batch(batchSize, rnd)
		accPi = (cnt*accPi + pi) / (cnt + 1.0)
		fmt.Println(accPi, pi, cnt)
	}
	return accPi
}

func batch(batchSize int, rnd *rand.Rand) float64 {
	insideCircle := 0
	var x, y float64
	for i := 0; i < batchSize; i++ {
		x, y = rnd.Float64(), rnd.Float64()
		if x*x+y*y <= 1.0 {
			insideCircle++
		}
	}
	return 4.0 * float64(insideCircle) / float64(batchSize)
}
