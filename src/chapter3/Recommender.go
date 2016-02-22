package chapter3

import (
    . "base"
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Recommender struct {
    Data           map[string]Ratings
    Frequencies    map[string]Ratings
    Deviations     map[string]Ratings
    ProductId2Name map[string]string
    UserId2Name    map[string]string
    Username2Id    map[string]string
}

func (me *Recommender) ComputeDeviations() {
    // 无数据就跳出
    if me.Data == nil {
        return
    }
    if me.Frequencies == nil {
        me.Frequencies = map[string]Ratings{}
    }
    if me.Deviations == nil {
        me.Deviations = map[string]Ratings{}
    }
    // 获取每位用户的评分数据
    for _, ratings := range me.Data {
        // 对于该用户的每个评分项（歌手、分数）
        for item, rating := range ratings {
            if _, ok := me.Frequencies[item]; !ok {
                me.Frequencies[item] = Ratings{}
            }
            if _, ok := me.Deviations[item]; !ok {
                me.Deviations[item] = Ratings{}
            }
            tem_frequencies := me.Frequencies[item]
            tem_deviations := me.Deviations[item]
            // 再次遍历该用户的每个评分项
            for item2, rating2 := range ratings {
                if item != item2 {
                    // 将评分的差异保存到变量中
                    if _, ok := tem_frequencies[item2]; !ok {
                        tem_frequencies[item2] = float64(0)
                        tem_deviations[item2] = float64(0)
                    }
                    tem_frequencies[item2] += 1
                    tem_deviations[item2] += rating - rating2
                }
            }
        }
    }

    for item, ratings := range me.Deviations {
        for item2, deviation := range ratings {
            ratings[item2] = deviation / me.Frequencies[item][item2]
        }
    }
}

func (me *Recommender) SlopeOneRecommendations(userRatings Ratings) ItemList {
    recommendations := Ratings{}
    frequencies := Ratings{}
    // 遍历目标用户的评分项（歌手、分数）
    for userItem, userRating := range userRatings {
        // 对目标用户未评价的歌手进行计算
        for diffItem, diffRating := range me.Deviations {
            _, ok1 := userRatings[diffItem]
            _, ok2 := me.Deviations[diffItem][userItem]
            if !ok1 && ok2 {
                freq := me.Frequencies[diffItem][userItem]
                if _, ok := recommendations[diffItem]; !ok {
                    recommendations[diffItem] = float64(0)
                }
                if _, ok := frequencies[diffItem]; !ok {
                    frequencies[diffItem] = float64(0)
                }
                // 分子
                recommendations[diffItem] += (diffRating[userItem] + userRating) * freq
                // 分母
                frequencies[diffItem] += freq
            }
        }
    }

    list := ItemList{}
    for k, v := range recommendations {
        item := Item{
            Key: k,
            Val: v / frequencies[k],
        }
        list = append(list, item)
    }
    // 排序并返回
    list.Sort(true)
    return list
}

func (me *Recommender) LoadMovieLens(path string) {
    me.Data = map[string]Ratings{}

    // first load movie ratings
    i := 0

    // First load book ratings into self.data
    fileRatings, err := os.Open(path + "u.data")
    defer fileRatings.Close()
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(fileRatings)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        i++
        // separate line into fields
        line := scanner.Text()
        fields := strings.Split(line, "\t")
        user := fields[0]
        movie := fields[1]
        rating, _ := strconv.ParseFloat(fields[2], 64)
        if _, ok := me.Data[user]; !ok {
            me.Data[user] = Ratings{}
        }
        me.Data[user][movie] = rating
    }

    // Now load movie into self.productid2name
    // the file u.item contains movie id, title, release date among
    // other fields
    fileMovie, err := os.Open(path + "u.item")
    defer fileMovie.Close()
    if err != nil {
        panic(err)
    }
    scanner = bufio.NewScanner(fileMovie)
    scanner.Split(bufio.ScanLines)
    me.ProductId2Name = map[string]string{}
    for scanner.Scan() {
        i++
        line := scanner.Text()
        fields := strings.Split(line, "|")
        mid := fields[0]
        title := fields[1]
        me.ProductId2Name[mid] = title
    }

    // Now load user info into both self.userid2name
    // and self.username2id
    fileUser, err := os.Open(path + "u.user")
    defer fileUser.Close()
    if err != nil {
        panic(err)
    }
    scanner = bufio.NewScanner(fileUser)
    scanner.Split(bufio.ScanLines)
    me.UserId2Name = map[string]string{}
    for scanner.Scan() {
        i++
        line := scanner.Text()
        fields := strings.Split(line, "|")
        userId := fields[0]
        me.UserId2Name[userId] = line
    }
    fmt.Println(i)
}
