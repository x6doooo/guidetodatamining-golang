package chapter2

import (
    . "base"
    "math"
)

func Manhattan(rating1, rating2 Ratings) float64 {
    // 计算曼哈顿距离。rating1和rating2参数中存储的数据格式均为
    // {'The Strokes': 3.0, 'Slightly Stoopid': 2.5}
    var distance float64 = 0
    for key, val1 := range rating1 {
        if val2, ok := rating2[key]; ok {
            distance += math.Abs(val1 - val2)
        }
    }
    return distance
}
