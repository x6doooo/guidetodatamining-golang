package chapter4

import (
    . "base"
)

func ComputeNearestNeighbor(itemName string, itemVector []float64, items map[string][]float64) ItemList {
    // 按照距离排序，返回邻近物品列表
    distances := ItemList{}
    for otherItem, otherItemVector := range items {
        distance := Manhattan(itemVector, otherItemVector)
        distances = append(distances, Item{otherItem, distance})
    }
    distances.Sort(false)
    return distances
}
