package chapter2


import (
	. "base"
)


func ComputeNearestNeighbor(username string, users map[string]Ratings) ItemList {
	// 计算所有用户至username用户的距离，倒序排列并返回结果列表
	var currentUserRatings = users[username];
	allDistances := ItemList{}
	for user, userRatings := range users {
		if user != username {
			distance := Manhattan(userRatings, currentUserRatings)
			tem := Item{user, distance}
			allDistances = append(allDistances, tem)
		}
	}
	allDistances.Sort(false)
	return allDistances
}