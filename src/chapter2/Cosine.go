package chapter2

import (
	"math"
	. "base"
)

// 余弦相似计算
func Cosine(rating1, rating2 Ratings) float64 {

	// 使用goroutine计算
	x_pow2_sum_ch := make(chan float64)
	y_pow2_sum_ch := make(chan float64)
	xy_sum_ch := make(chan float64)

	go func() {
		var xy_sum float64 = 0
		var x_pow2_sum float64 = 0
		for key, value1 := range rating1 {
			if value2, ok := rating2[key]; ok {
				xy_sum += value1 * value2
			}
			x_pow2_sum += math.Pow(value1, 2)
		}
		x_pow2_sum_ch <- x_pow2_sum
		xy_sum_ch <- xy_sum
	}()

	go func() {
		var y_pow2_sum float64 = 0
		for _, value2 := range rating2 {
			y_pow2_sum += math.Pow(value2, 2)
		}
		y_pow2_sum_ch <- y_pow2_sum
	}()

	var x_pow2_sum float64
	var y_pow2_sum float64
	var xy_sum float64
	for i := 0; i < 3; i++ {
		select {
		case x_pow2_sum = <- x_pow2_sum_ch:
		case y_pow2_sum = <- y_pow2_sum_ch:
		case xy_sum = <- xy_sum_ch:
		}
	}


	x_mod := math.Sqrt(x_pow2_sum)
	y_mod := math.Sqrt(y_pow2_sum)

	return xy_sum / (x_mod * y_mod)

}
