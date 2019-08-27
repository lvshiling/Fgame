package timeutils

import "time"

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

func BeginOfNow() (int64, error) {
	return BeginOfDayOfTime(time.Now())
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
