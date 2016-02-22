package chapter2

//def minkowski(rating1, rating2, r):
//distance = 0
//for key in rating1:
//if key in rating2:
//distance += pow(abs(rating1[key] - rating2[key]), r)
//return pow(distance, 1.0 / r)

import (
	. "base"
	"math"
)

func Minkowski(rating1, rating2 Ratings, r float64) float64 {
	var distance float64 = 0
	for key, value1 := range rating1 {
		if value2, ok := rating2[key]; ok {
			distance += math.Pow(math.Abs(value1 - value2), r)
		}
	}
	return math.Pow(distance, 1 / r)
}