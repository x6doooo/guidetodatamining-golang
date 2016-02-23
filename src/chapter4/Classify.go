package chapter4


//def classify(user, itemName, itemVector):
//nearest = computeNearestNeighbor(itemName, itemVector, items)[0][1]
//rating = users[user][nearest]
//return rating


func Classify(
    user string,
    itemName string,
    itemVector []float64,
    items map[string][]float64,
    users map[string]map[string]string,
) string {
    nearest := ComputeNearestNeighbor(itemName, itemVector, items)[0].Key
    rating := users[user][nearest]
    return rating
}