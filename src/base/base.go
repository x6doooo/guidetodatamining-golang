package base

import (
    "sort"
)

/**
作品->评分
eg. { 'artistname': 1.0, ... }
*/
type Ratings map[string]float64

/**
键值结构体, 键为string, 值为float64
算法中有若干个地方用到这种数据结构, 所以抽象出来
*/
type Item struct {
    Key string
    Val float64
}

/**
[]item的排序
*/
type ItemList []Item

// sort方法
func (list ItemList) Len() int {
    return len(list)
}
func (list ItemList) Swap(i, j int) {
    list[i], list[j] = list[j], list[i]
}

func (list ItemList) Less(i, j int) bool {
    return list[i].Val < list[j].Val
}

func (list ItemList) Sort(desc bool) {
    if desc {
        sort.Sort(sort.Reverse(list))
    } else {
        sort.Sort(list)
    }
}
