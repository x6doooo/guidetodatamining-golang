package chapter2

import (
    . "base"
    "fmt"
    "testing"
)

var users = map[string]Ratings{
    "Angelica": Ratings{
        "Blues Traveler": 3.5, "Broken Bells": 2.0, "Norah Jones": 4.5, "Phoenix": 5.0, "Slightly Stoopid": 1.5, "The Strokes": 2.5, "Vampire Weekend": 2.0,
    },
    "Bill": Ratings{
        "Blues Traveler": 2.0, "Broken Bells": 3.5, "Deadmau5": 4.0, "Phoenix": 2.0, "Slightly Stoopid": 3.5, "Vampire Weekend": 3.0,
    },
    "Chan": Ratings{
        "Blues Traveler": 5.0, "Broken Bells": 1.0, "Deadmau5": 1.0, "Norah Jones": 3.0, "Phoenix": 5, "Slightly Stoopid": 1.0,
    },
    "Dan": Ratings{
        "Blues Traveler": 3.0, "Broken Bells": 4.0, "Deadmau5": 4.5, "Phoenix": 3.0, "Slightly Stoopid": 4.5, "The Strokes": 4.0, "Vampire Weekend": 2.0,
    },
    "Hailey": Ratings{
        "Broken Bells": 4.0, "Deadmau5": 1.0, "Norah Jones": 4.0, "The Strokes": 4.0, "Vampire Weekend": 1.0,
    },
    "Jordyn": Ratings{
        "Broken Bells": 4.5, "Deadmau5": 4.0, "Norah Jones": 5.0, "Phoenix": 5.0, "Slightly Stoopid": 4.5, "The Strokes": 4.0, "Vampire Weekend": 4.0,
    },
    "Sam": Ratings{
        "Blues Traveler": 5.0, "Broken Bells": 2.0, "Norah Jones": 3.0, "Phoenix": 5.0, "Slightly Stoopid": 4.0, "The Strokes": 5.0,
    },
    "Veronica": Ratings{
        "Blues Traveler": 3.0, "Norah Jones": 5.0, "Phoenix": 4.0, "Slightly Stoopid": 2.5, "The Strokes": 3.0,
    },
}

func Test_Manhattan(t *testing.T) {
    var distance float64
    distance = Manhattan(users["Hailey"], users["Veronica"])
    fmt.Println("manhattan distance -> Hailey and Veronica =", distance)
    distance = Manhattan(users["Hailey"], users["Jordyn"])
    fmt.Println("manhattan distance -> Hailey and Jordyn =", distance)
}

func Test_ComputeNearestNeighbor(t *testing.T) {
    distances := ComputeNearestNeighbor("Hailey", users)
    fmt.Println("manhattan distance -> Hailey and others =", distances)
}

func Test_Recommend(t *testing.T) {
    list := Recommend("Hailey", users)
    fmt.Println("recommendations -> Hailey =", list)
    list = Recommend("Chan", users)
    fmt.Println("recommendations -> Chan =", list)
    list = Recommend("Sam", users)
    fmt.Println("recommendations -> Sam =", list)
}

func Test_Pearson(t *testing.T) {
    d := Pearson(users["Angelica"], users["Bill"])
    fmt.Println("pearson distance -> Angelica and Bill =", d)
    d = Pearson(users["Angelica"], users["Hailey"])
    fmt.Println("pearson distance -> Angelica and Hailey =", d)
    d = Pearson(users["Angelica"], users["Jordyn"])
    fmt.Println("pearson distance -> Angelica and Jordyn =", d)
}

func Test_Minkowski(t *testing.T) {
    d := Minkowski(users["Angelica"], users["Bill"], 2)
    fmt.Println("minkowski distance -> Angelica and Bill =", d)
}

func Test_Cosine(t *testing.T) {
    d := Cosine(users["Angelica"], users["Veronica"])
    fmt.Println("cosine distance -> Angelica and Veronica", d)
}

func Test_Recommender(t *testing.T) {
    r := NewRecommender(0, 0, users, nil)
    d := r.Recommend("Jordyn")
    fmt.Println(d)
    d = r.Recommend("Hailey")
    fmt.Println(d)
    r.LoadBookDB("./BX-CSV-Dump/")
    d = r.Recommend("171118")
    fmt.Println(d)
    r.UserRatings("171118", 5)
    //fmt.Println(d)
}
