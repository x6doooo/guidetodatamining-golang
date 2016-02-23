package chapter4

import (
    . "base"
    c2 "chapter2"
    "testing"
    "fmt"
)

var music = map[string]Ratings {
    "Dr Dog/Fate": Ratings{
        "piano": 2.5, "vocals": 4, "beat": 3.5, "blues": 3, "guitar": 5, "backup vocals": 4, "rap": 1,
    },
    "Phoenix/Lisztomania": Ratings{
        "piano": 2, "vocals": 5, "beat": 5, "blues": 3, "guitar": 2, "backup vocals": 1, "rap": 1,
    },
    "Heartless Bastards/Out at Sea": Ratings{
        "piano": 1, "vocals": 5, "beat": 4, "blues": 2, "guitar": 4, "backup vocals":1, "rap": 1,
    },
    "Todd Snider/Don't Tempt Me": Ratings{
        "piano": 4, "vocals": 5, "beat": 4, "blues": 4, "guitar": 1, "backup vocals": 5, "rap": 1,
    },
    "The Black Keys/Magic Potion": Ratings{
        "piano": 1, "vocals": 4, "beat": 5, "blues": 3.5, "guitar": 5, "backup vocals": 1, "rap": 1,
    },
    "Glee Cast/Jessie's Girl": Ratings{
        "piano": 1, "vocals": 5, "beat": 3.5, "blues": 3, "guitar":4, "backup vocals": 5, "rap": 1,
    },
    "La Roux/Bulletproof": Ratings{
        "piano": 5, "vocals": 5, "beat": 4, "blues": 2, "guitar": 1, "backup vocals": 1, "rap": 1,
    },
    "Mike Posner": Ratings{
        "piano": 2.5, "vocals": 4, "beat": 4, "blues": 1, "guitar": 1, "backup vocals": 1, "rap": 1,
    },
    "Black Eyed Peas/Rock That Body": Ratings{
        "piano": 2, "vocals": 5, "beat": 5, "blues": 1, "guitar": 2, "backup vocals": 2, "rap": 4,
    },
    "Lady Gaga/Alejandro": Ratings{
        "piano": 1, "vocals": 5, "beat": 3, "blues": 2, "guitar": 1, "backup vocals": 2, "rap": 1,
    },
}

// 物品向量中的特征依次为：piano, vocals, beat, blues, guitar, backup vocals, rap
var items = map[string][]float64{
    "Dr Dog/Fate": []float64{2.5, 4, 3.5, 3, 5, 4, 1},
    "Phoenix/Lisztomania": []float64{2, 5, 5, 3, 2, 1, 1},
    "Heartless Bastards/Out": []float64{1, 5, 4, 2, 4, 1, 1},
    "Todd Snider/Don't Tempt Me": []float64{4, 5, 4, 4, 1, 5, 1},
    "The Black Keys/Magic Potion": []float64{1, 4, 5, 3.5, 5, 1, 1},
    "Glee Cast/Jessie's Girl": []float64{1, 5, 3.5, 3, 4, 5, 1},
    "La Roux/Bulletproof": []float64{5, 5, 4, 2, 1, 1, 1},
    "Mike Posner": []float64{2.5, 4, 4, 1, 1, 1, 1},
    "Black Eyed Peas/Rock That Body": []float64{2, 5, 5, 1, 2, 2, 4},
    "Lady Gaga/Alejandro": []float64{1, 5, 3, 2, 1, 2, 1},
}

var users = map[string]map[string]string{
    "Angelica": map[string]string{
        "Dr Dog/Fate": "L",
        "Phoenix/Lisztomania": "L",
        "Heartless Bastards/Out at Sea": "D",
        "Todd Snider/Don't Tempt Me": "D",
        "The Black Keys/Magic Potion": "D",
        "Glee Cast/Jessie's Girl": "L",
        "La Roux/Bulletproof": "D",
        "Mike Posner": "D",
        "Black Eyed Peas/Rock That Body": "D",
        "Lady Gaga/Alejandro": "L",
    },
    "Bill": map[string]string{
        "Dr Dog/Fate": "L",
        "Phoenix/Lisztomania": "L",
        "Heartless Bastards/Out at Sea": "L",
        "Todd Snider/Don't Tempt Me": "D",
        "The Black Keys/Magic Potion": "L",
        "Glee Cast/Jessie's Girl": "D",
        "La Roux/Bulletproof": "D",
        "Mike Posner": "D",
        "Black Eyed Peas/Rock That Body": "D",
        "Lady Gaga/Alejandro": "D",
    },}


func Test_Chapter2ComputeNearestNeighbor(t *testing.T) {
    list := c2.ComputeNearestNeighbor("The Black Keys/Magic Potion", music)
    fmt.Println(list)

    nearest := ComputeNearestNeighbor("Chris Cagle/I Breathe In. I Breathe Out", []float64{1, 5, 2.5, 1, 1, 5, 1}, items)
    fmt.Println(nearest)

    rating := Classify("Angelica", "Chris Cagle/I Breathe In. I Breathe Out", []float64{1, 5, 2.5, 1, 1, 5, 1}, items, users)
    fmt.Println(rating)
}