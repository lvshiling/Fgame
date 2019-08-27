package timeutils_test

import (
	"fmt"
	. "qipai/pkg/timeutils"
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
