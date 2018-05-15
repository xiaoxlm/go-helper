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

// 每个月最后一天时间戳
func MonthLastDay () int64 {
	yearMonth := time.Now().Format("2006-01")
	timestr := yearMonth + "-01"

	t, _ := time.Parse("2006-01-02", timestr)
	date := t.AddDate(0, 1, -1)
	return date.Add(15 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second).Unix()
}

// 每天开始的时间戳
func DayBegin () int64 {
	yearMonthDay := time.Now().Format("2006-01-02")
	timestr := yearMonthDay + " 00:00:00"

	t, _ := time.Parse("2006-01-02 15:04:05", timestr)
	return t.Add(-time.Hour * 8).Unix()
}

// 每天结束的时间戳
func DayEnd () int64 {
	yearMonthDay := time.Now().Format("2006-01-02")
	timestr := yearMonthDay + " 00:00:00"

	t, _ := time.Parse("2006-01-02 15:04:05", timestr)
	return t.Add(15 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second).Unix()
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
