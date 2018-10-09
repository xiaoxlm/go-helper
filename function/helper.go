package function

// mac1
// win1
// win2
// mac2
// mac3
// mac4
// win3
// win4
import (
	"time"
	"strconv"
	"fmt"
	"math/rand"
	glog "log"
	"io/ioutil"
	"bytes"
	"runtime"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
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

// 异常捕获
func PanicRecover(skip int) {
	if err := recover(); err != nil {
		stack := Stack(skip)
		glog.Printf("[Recovery] panic recovered:\n%s\n%s", err, stack)
		/*
		* 利用log包里设置logger, 将panic输出到指定的目录文件
		var logger *glog.Logger
		out, _ := log.NewDailyWriteLog("", "system_err")
		logger = glog.New(out, "\n\n", glog.LstdFlags)
		if logger != nil {
			stack := Stack(3)
			logger.Printf("[Recovery] panic recovered:\n%s\n%s", err, stack)
		}
		*/
		return
	}
}

func Stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
