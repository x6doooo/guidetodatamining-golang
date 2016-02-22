package chapter2

import (
    . "base"
)

func Recommend(username string, users map[string]Ratings) ItemList {
    // 返回推荐结果列表

    // 找到距离最近的用户
    nearest := ComputeNearestNeighbor(username, users)[0].Key

    recommendations := ItemList{}

    // 找出这位用户评价过、但自己未曾评价的乐队
    neighborRatings := users[nearest]
    userRatings := users[username]

    for artist, rating := range neighborRatings {
        if _, ok := userRatings[artist]; !ok {
            tem := Item{artist, rating}
            recommendations = append(recommendations, tem)
        }
    }

    recommendations.Sort(true)
    return recommendations
}
