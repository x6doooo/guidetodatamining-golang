package chapter3

import (
    . "base"
    "fmt"
    "testing"
)

var users2 = map[string]Ratings{
    "Amy":   {"Taylor Swift": 4, "PSY": 3, "Whitney Houston": 4},
    "Ben":   {"Taylor Swift": 5, "PSY": 2},
    "Clara": {"PSY": 3.5, "Whitney Houston": 4},
    "Daisy": {"Taylor Swift": 5, "Whitney Houston": 3},
}

var users3 = map[string]Ratings{
    "David": {"Imagine Dragons": 3, "Daft Punk": 5, "Lorde": 4, "Fall Out Boy": 1},
    "Matt":  {"Imagine Dragons": 3, "Daft Punk": 4, "Lorde": 4, "Fall Out Boy": 1},
    "Ben":   {"Kacey Musgraves": 4, "Imagine Dragons": 3, "Lorde": 3, "Fall Out Boy": 1},
    "Chris": {"Kacey Musgraves": 4, "Imagine Dragons": 4, "Daft Punk": 4, "Lorde": 3, "Fall Out Boy": 1},
    "Tori":  {"Kacey Musgraves": 5, "Imagine Dragons": 4, "Daft Punk": 5, "Fall Out Boy": 3},
}

func Test_ComputeSimilarity(t *testing.T) {
    s := ComputeSimilarity("Kacey Musgraves", "Lorde", users3)
    fmt.Println("Similarity, Kacey Musgraves & Lorde =", s)
    s = ComputeSimilarity("Imagine Dragons", "Lorde", users3)
    fmt.Println("Similarity, Imagine Dragons & Lorde =", s)
    s = ComputeSimilarity("Daft Punk", "Lorde", users3)
    fmt.Println("Similarity, Daft Punk & Lorde =", s)
}

func Test_Recommender(t *testing.T) {
    r := Recommender{}
    r.Data = users2
    r.ComputeDeviations()
    fmt.Println("ComputeDeviations", r.Deviations)
    list := r.SlopeOneRecommendations(users2["Ben"])
    fmt.Println("SlopeOneRecommendations", list)

    r2 := Recommender{}
    r2.LoadMovieLens("./ml-100k/")
    fmt.Println("---计算量较大,需要几分钟---")
    r2.ComputeDeviations()
    list = r2.SlopeOneRecommendations(r2.Data["25"])
    fmt.Println(list)
}
