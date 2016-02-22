package chapter3

import (
    . "base"
    "math"
)

func ComputeSimilarity(band1, band2 string, userRatings map[string]Ratings) float64 {
    // 计算每个用户给出分数的平均分
    averages := Ratings{}
    for u, rs := range userRatings {
        var count float64 = 0
        for _, r := range rs {
            count += r
        }
        l := len(rs)
        averages[u] = count / float64(l)
    }
    // 相似度
    var (
        num float64 = 0
        dem1 float64 = 0
        dem2 float64 = 0
    )
    for u, rs := range userRatings {
        band1_rating, band1_ok := rs[band1]
        band2_rating, band2_ok := rs[band2]
        if band1_ok && band2_ok {
            avg := averages[u]
            num += (band1_rating - avg) * (band2_rating - avg)
            dem1 += math.Pow(band1_rating - avg, 2)
            dem2 += math.Pow(band2_rating - avg, 2)
        }
    }
    return num / (math.Sqrt(dem1) * math.Sqrt(dem2))
}
