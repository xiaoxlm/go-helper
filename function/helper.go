package function

import (
	"time"
	"strconv"
	"fmt"
	"math/rand"
)

// 24小时前的时间戳
func Before24HourTimestamp() int64 {
	return time.Now().AddDate(0, 0, -1).Unix()
}

// 每个月的第一天
func MonthFirstDay () int64 {
	yearMonth := time.Now().Format("2006-01")
	timestr := yearMonth + "-01"

	t, _ := time.Parse("2006-01-02", timestr)
	return t.Add(-time.Hour * 8).Unix()
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
	//n10 := math.Pow10(n)
	//return math.Trunc((f+0.5/n10)*n10) / n10
}

// 随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rand.Seed(time.Now().Unix())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
