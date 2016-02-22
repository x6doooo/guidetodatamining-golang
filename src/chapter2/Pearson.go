package chapter2

import (
    . "base"
    "math"
)

func Pearson(rating1, rating2 Ratings) float64 {
    var sum_xy float64 = 0
    var sum_x float64 = 0
    var sum_y float64 = 0
    var sum_x2 float64 = 0
    var sum_y2 float64 = 0
    var n float64 = 0
    for key, x := range rating1 {
        if y, ok := rating2[key]; ok {
            n += 1
            sum_xy += x * y
            sum_x += x
            sum_y += y
            sum_x2 += math.Pow(x, 2)
            sum_y2 += math.Pow(y, 2)
        }
    }
    if n == 0 {
        return 0
    }
    denominator := math.Sqrt(sum_x2 - math.Pow(sum_x, 2) / n) * math.Sqrt(sum_y2 - math.Pow(sum_y, 2) / n)
    if denominator == 0 {
        return 0
    } else {
        return (sum_xy - (sum_x * sum_y) / n) / denominator
    }
}
