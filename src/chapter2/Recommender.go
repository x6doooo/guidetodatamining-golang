package chapter2

import (
    . "base"
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

//推荐算法别名
type metricFunc func(Ratings, Ratings) float64

type Recommender struct {
    // 训练数据
    data           map[string]Ratings
    // K邻近算法中的值
    k              int
    // 使用何种距离计算方式
    metric         metricFunc
    // 推荐结果的数量
    n              int

    // 用户转换用户名和作品的id和名称
    username2id    map[string]string
    userid2name    map[string]string
    productid2name map[string]string
}

// 构造方法
func NewRecommender(k int, n int, data map[string]Ratings, metric metricFunc) *Recommender {
    if k == 0 {
        k = 1
    }
    if n == 0 {
        n = 5
    }
    if metric == nil {
        metric = Pearson
    }
    return &Recommender{
        data:   data,
        k:      k,
        metric: metric,
        n:      n,
    }
}

// 加载BX数据集，path是数据文件位置
func (me *Recommender) LoadBookDB(path string) {
    me.data = map[string]Ratings{}
    i := 0
    // 将书籍评分数据放入self.data
    file, err := os.Open(path + "BX-Book-Ratings.csv")
    defer file.Close()
    if err != nil {
        panic(err)
    }
    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = ';'
    reader.LazyQuotes = true
    for {
        line, err := reader.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        i++

        // 数据异常 跳出
        if len(line) != 3 {
            if len(line) == 2 {
                // 部分数据有转义符号 造成混淆 这里手动格式化
                line[1] = strings.Replace(line[1], "\\\"", "", -1)
                tem := strings.Split(line[1], ";")
                line[1] = tem[0]
                tem[1] = strings.Replace(tem[1], "\"", "", -1)
                line = append(line, tem[1])
            } else {
                continue
            }
        }
        user := line[0]
        book := line[1]
        rating, _ := strconv.ParseFloat(line[2], 64)

        if _, ok := me.data[user]; !ok {
            me.data[user] = Ratings{}
        }
        me.data[user][book] = rating
    }

    // 将书籍信息存入self.productid2name
    // 包括isbn号、书名、作者等
    fileBook, err := os.Open(path + "BX-Books.csv")
    defer fileBook.Close()
    if err != nil {
        panic(err)
    }
    reader = csv.NewReader(bufio.NewReader(fileBook))
    reader.Comma = ';'
    reader.LazyQuotes = true

    me.productid2name = map[string]string{}
    for {
        line, err := reader.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        i++
        isbn := line[0]
        title := line[1]
        author := line[2]
        me.productid2name[isbn] = title + " by " + author
    }

    // 将用户信息存入self.userid2name和self.username2id
    fileUser, err := os.Open(path + "BX-Users.csv")
    defer fileUser.Close()
    if err != nil {
        panic(err)
    }
    reader = csv.NewReader(bufio.NewReader(fileUser))
    reader.Comma = ';'
    reader.LazyQuotes = true

    me.userid2name = map[string]string{}
    me.username2id = map[string]string{}
    for {
        line, err := reader.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        i++
        userid := line[0]
        location := line[1]
        var value string
        if len(line) > 3 {
            value = location + "(age: " + line[2] + ")"
        } else {
            value = location
        }
        me.userid2name[userid] = value
        me.username2id[location] = userid
    }

    fmt.Println(i)

}

// 通过产品ID获取名称
func (me *Recommender) ConvertProductID2name(id string) string {
    if name, ok := me.productid2name[id]; ok {
        return name
    }
    return id
}

// 返回该用户评分最高的物品
func (me *Recommender) UserRatings(id string, n int) {
    fmt.Println("Ratings for " + me.userid2name[id])
    ratings := me.data[id]

    list := ItemList{}
    for itemId, rating := range ratings {
        tem := Item{me.ConvertProductID2name(itemId), rating}
        list = append(list, tem)
    }
    list.Sort(true)
    if len(list) < n {
        fmt.Println(list)
    } else {
        fmt.Println(list[:n])
    }
}

// 获取邻近用户
func (me *Recommender) ComputeNearestNeighbor(userId string) ItemList {
    list := ItemList{}
    for user, userRatings := range me.data {
        if user != userId {
            distance := me.metric(me.data[userId], userRatings)
            tem := Item{user, distance}
            list = append(list, tem)
        }
    }
    list.Sort(true)
    return list
}

func (me *Recommender) Recommend(userId string) ItemList {
    // 返回推荐列表
    recommendations := map[string]float64{}

    // 首先，获取邻近用户
    nearest := me.ComputeNearestNeighbor(userId)

    // 获取用户评价过的商品
    userRatings := me.data[userId]

    // 计算总距离
    var totalDistance float64 = 0
    for i := 0; i < me.k; i++ {
        totalDistance += nearest[i].Val
    }

    // 汇总K邻近用户的评分
    for i := 0; i < me.k; i++ {
        // 计算饼图的每个分片
        weight := nearest[i].Val / totalDistance
        // 获取用户名称
        id := nearest[i].Key
        // 获取用户评分
        neighborRatings := me.data[id]
        // 获得没有评价过的商品
        for artist, rating := range neighborRatings {
            if _, ok := userRatings[artist]; !ok {
                if _, ok := recommendations[artist]; !ok {
                    recommendations[artist] = rating * weight
                } else {
                    recommendations[artist] = recommendations[artist] + rating * weight
                }
            }
        }
    }

    // 开始推荐
    recommendationsList := ItemList{}
    for artist, rating := range recommendations {
        artistName := me.ConvertProductID2name(artist)
        tem := Item{artistName, rating}
        recommendationsList = append(recommendationsList, tem)
    }
    recommendationsList.Sort(true)

    if len(recommendationsList) < me.n {
        return recommendationsList
    }
    return recommendationsList[:me.n]

}
