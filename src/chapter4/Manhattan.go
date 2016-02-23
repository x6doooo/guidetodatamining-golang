package chapter4

import (
    "math"
)

func Manhattan(vector1, vector2 []float64) float64 {
    var distance float64 = 0

    for i, val1 := range vector1 {
        distance += math.Abs(val1 - vector2[i])
    }

    return distance
}