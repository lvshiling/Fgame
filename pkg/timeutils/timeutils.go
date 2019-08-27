package timeutils

import (
	"fmt"
	"time"
)

const (
	//秒
	SECOND = time.Second / time.Millisecond

	//分
	MINUTE = 60 * SECOND

	//时
	HOUR = 60 * MINUTE

	//24小时
	DAY = HOUR * 24
)

func MillisecondToTime(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*int64(time.Millisecond))
}

func TimeToMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func TimeToSecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Second)
}

func MillisecondForTime(t time.Time) int64 {
	return (int64(t.Hour())*int64(time.Hour) + int64(t.Minute())*int64(time.Minute) + int64(t.Second())*int64(time.Second)) / int64(time.Millisecond)
}

func BeginOfNow(now int64) (int64, error) {
	t := MillisecondToTime(now)
	return BeginOfDayOfTime(t)
}

func BeginOfDayOfTime(t time.Time) (int64, error) {
	timeStr := t.Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	timeNumber := t.UnixNano() / int64(time.Millisecond)
	return timeNumber, nil
}

func SecondToDuration(sec int64) time.Duration {
	return time.Duration(sec * int64(time.Second))
}

func BeginOfYesterday() (int64, error) {
	timeStr := time.Now().Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t = t.AddDate(0, 0, -1)
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

func BeginOfWeek() (int64, error) {
	return BeginOfWeekOfTime(time.Now())
}

func BeginOfWeekOfMillisecond(now int64) (int64, error) {
	t := MillisecondToTime(now)
	return BeginOfWeekOfTime(t)
}

func BeginOfWeekOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	days := 1 - int(t.Weekday())
	if days > 0 {
		days = -6
	}
	t = t.AddDate(0, 0, days)
	timeNumber := t.UnixNano() / int64(time.Millisecond)
	return timeNumber, nil
}

func BeginOfLastWeekOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	days := 1 - int(t.Weekday())
	if days > 0 {
		days = -6
	}
	days -= 7
	t = t.AddDate(0, 0, days)
	timeNumber := t.UnixNano() / int64(time.Millisecond)
	return timeNumber, nil
}

func BeginOfLastMonth() (int64, error) {
	return BeginOfLastMonthOfTime(time.Now())
}

func BeginOfLastMonthOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006-01")
	t, err := time.ParseInLocation("2006-01", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t = t.AddDate(0, -1, 0)
	timeNumber := t.UnixNano() / int64(time.Millisecond)
	return timeNumber, nil
}

func BeginOfMonth() (int64, error) {
	return BeginOfMonthOfTime(time.Now())
}

func BeginOfMonthOfMillisecond(now int64) (int64, error) {
	t := MillisecondToTime(now)
	return BeginOfMonthOfTime(t)
}

func BeginOfMonthOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006-01")
	t, err := time.ParseInLocation("2006-01", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

func BeginOfYear() (int64, error) {
	return BeginOfYearOfTime(time.Now())
}

func BeginOfLastYear() (int64, error) {
	return BeginOfLastYearOfTime(time.Now())
}

func BeginOfYearOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006")
	t, err := time.ParseInLocation("2006", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

func BeginOfLastYearOfTime(ti time.Time) (int64, error) {
	timeStr := ti.Format("2006")
	t, err := time.ParseInLocation("2006", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t = t.AddDate(-1, 0, 0)
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

func IsSameWeek(time1 int64, time2 int64) (flag bool, err error) {
	t1 := MillisecondToTime(time1)
	t2 := MillisecondToTime(time2)
	tt1, err := BeginOfWeekOfTime(t1)
	if err != nil {
		return
	}
	tt2, err := BeginOfWeekOfTime(t2)
	if err != nil {
		return
	}
	flag = tt1 == tt2
	return
}

func IsSameMonth(time1 int64, time2 int64) (flag bool, err error) {
	t1 := MillisecondToTime(time1)
	t2 := MillisecondToTime(time2)
	tt1, err := BeginOfMonthOfTime(t1)
	if err != nil {
		return
	}
	tt2, err := BeginOfMonthOfTime(t2)
	if err != nil {
		return
	}
	flag = tt1 == tt2
	return
}

func IsSameDay(time1 int64, time2 int64) (flag bool, err error) {
	t1 := MillisecondToTime(time1)
	t2 := MillisecondToTime(time2)
	tt1, err := BeginOfDayOfTime(t1)
	if err != nil {
		return
	}
	tt2, err := BeginOfDayOfTime(t2)
	if err != nil {
		return
	}
	flag = tt1 == tt2
	return
}

func fiveOfDayOfTime(t time.Time) (int64, error) {
	timeStr := t.Format("2006-01-02")
	nt, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t2 := nt.Add(time.Hour * 5)
	if t.Hour() < 5 {
		t2 = nt.Add(-time.Hour * 19)
	}
	timeNumber := t2.UnixNano() / int64(time.Millisecond)
	return timeNumber, nil
}

func IsSameFive(time1 int64, time2 int64) (flag bool, err error) {
	t1 := MillisecondToTime(time1)
	t2 := MillisecondToTime(time2)
	tt1, err := fiveOfDayOfTime(t1)
	if err != nil {
		return
	}

	tt2, err := fiveOfDayOfTime(t2)
	if err != nil {
		return
	}
	flag = tt1 == tt2
	return
}

func GetHourMs(now int64) (hourMs int64) {
	t := MillisecondToTime(now)
	return (t.Unix() - int64(t.Second()) - int64(t.Minute()*60)) * 1000
}

func GetIntervalTimeStampMs(now int64, intervalTime int64) (timeStamp int64, flag bool) {
	t := MillisecondToTime(now)

	dayMs := int64(24 * 60 * 60 * 1000)
	if dayMs%intervalTime != 0 {
		return
	}
	timeStamp = (t.Unix()*1000 - t.Unix()*1000%intervalTime)
	flag = true
	return
}

func ParseYYYYMMDD(value string) (int64, error) {
	quantum, err := time.ParseInLocation("20060102", value, time.Local)
	if err != nil {
		return 0, err
	}
	return TimeToMillisecond(quantum), nil
}

//格式化时分：0800（早上8点）
func ParseDayOfHHMM(hourStr string) (int64, error) {
	t := time.Now()
	timeStr := t.Format("20060102")
	dateStr := fmt.Sprint(timeStr, hourStr)

	timeNumber, err := ParseDayOfYYYYMMDDHHMM(dateStr)
	if err != nil {
		return 0, err
	}
	beginTime, err := BeginOfDayOfTime(t)
	if err != nil {
		return 0, err
	}
	return timeNumber - beginTime, nil
}

//格式化时分：0800（早上8点）
func ParseDayOfHHMMSS(hourStr string) (int64, error) {
	t := time.Now()
	timeStr := t.Format("20060102")
	dateStr := fmt.Sprint(timeStr, hourStr)

	timeNumber, err := ParseDayOfYYYYMMDDHHMMSS(dateStr)
	if err != nil {
		return 0, err
	}
	beginTime, err := BeginOfDayOfTime(t)
	if err != nil {
		return 0, err
	}
	return timeNumber - beginTime, nil
}

func ParseDayOfYYYYMMDDHHMM(timeStr string) (int64, error) {
	t, err := time.ParseInLocation("200601021504", timeStr, time.Local)
	if err != nil {
		return 0, err
	}

	timeNumber := TimeToMillisecond(t)
	return timeNumber, nil
}

func ParseDayOfYYYYMMDDHHMMSS(timeStr string) (int64, error) {
	t, err := time.ParseInLocation("20060102150405", timeStr, time.Local)
	if err != nil {
		return 0, err
	}

	timeNumber := TimeToMillisecond(t)
	return timeNumber, nil
}

//相差天数:time1>time2
func DiffDay(time1, time2 int64) (int32, error) {
	t1 := MillisecondToTime(time1)
	t2 := MillisecondToTime(time2)
	tt1, err := BeginOfDayOfTime(t1)
	if err != nil {
		return 0, err
	}
	tt2, err := BeginOfDayOfTime(t2)
	if err != nil {
		return 0, err
	}

	diffDay := (tt1 - tt2) / int64(DAY)

	return int32(diffDay), nil
}

//指定时间的前一天
func PreDayOfTime(now int64) (int64, error) {
	tn := MillisecondToTime(now)

	timeStr := tn.Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t = t.AddDate(0, 0, -1)
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

//指定时间后n天
func AfterNDayOfTime(now int64, afterDay int) (int64, error) {
	tn := MillisecondToTime(now)

	timeStr := tn.Format("2006-01-02")
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	t = t.AddDate(0, 0, afterDay)
	timeNumber := t.UnixNano() / int64(time.Millisecond)

	return timeNumber, nil
}

func GetYearMonthDay(now int64) (year int, month int, day int, hour int, min int) {
	t := time.Unix(now/int64(SECOND), 0)
	year = t.Year()
	month = int(t.Month())
	day = t.Day()
	hour = t.Hour()
	min = t.Minute()
	return
}

func MillisecondToSecondCeil(ms int64) (leftTime int64) {
	leftTime = ms / int64(SECOND)
	modValue := ms % int64(SECOND)
	if modValue != 0 {
		leftTime++
	}
	return
}

func MondayFivePointTime(ms int64) (timeStamp int64, err error) {
	t := MillisecondToTime(ms)
	mondayTime, err := BeginOfWeekOfTime(t)
	if err != nil {
		return
	}
	timeStamp = mondayTime + int64(5*HOUR)
	return
}
