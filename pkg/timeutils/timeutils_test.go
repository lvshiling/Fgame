package timeutils_test

import (
	. "fgame/fgame/pkg/timeutils"
	"fmt"
	"testing"
	"time"
)

func TestBeginOfNow(t *testing.T) {
	begin, _ := BeginOfNow()
	fmt.Println(begin)
}

func TestBeginOfYesterday(t *testing.T) {
	begin, _ := BeginOfYesterday()
	fmt.Println(begin)
}

func TestBeginOfWeek(t *testing.T) {
	begin, _ := BeginOfWeek()
	fmt.Println(begin)
}

func TestBeginOfLastWeekOfTime(t *testing.T) {
	begin, _ := BeginOfLastWeekOfTime(time.Now())
	fmt.Println(begin)
}

func TestBeginOfMonth(t *testing.T) {
	begin, _ := BeginOfMonth()
	fmt.Println(begin)
}

func TestIsSameWeek(t *testing.T) {

}

func TestParseYYYYMMDD(t *testing.T) {
	str := "20170102"
	val := int64(1483286400000)
	ms, err := ParseYYYYMMDD(str)
	if err != nil {
		t.Fatalf("timeutils:err %s", err.Error())
	}
	if ms != val {
		t.Fatalf("timeutils:the ms of %s should be %d,but get %d", str, val, ms)
	}
}

func TestGetHourMs(t *testing.T) {
	ms := int64(1483304461000)
	expectedHourMs := int64(1483304400000)
	hourMs := GetHourMs(ms)
	if hourMs != expectedHourMs {
		t.Fatalf("timeutils:the hour ms of %d should be %d,but get %d", ms, expectedHourMs, hourMs)
	}
}
