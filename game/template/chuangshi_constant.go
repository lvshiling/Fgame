package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ChuangShiConstantTemplate struct {
	*ChuangShiConstantTemplateVO
	//报名每周活动天
	signWeekdayArr []int32
	//选举每周活动天
	voteWeekdayArr []int32
	//报名开始时间
	signTimeHourArr []int64
	//选举开始时间
	voteTimeHourArr []int64

	//更换阵营消耗
	changedCampItemMap map[int32]int32
}

//产出次数
func (t *ChuangShiConstantTemplate) RewCout(lastTime, now int64) (rewCount int32, lastReceiveTime int64) {
	interval := now - lastTime
	count := interval / t.CityRewTime

	lastReceiveTime = lastTime + count*t.CityRewTime
	rewCount = int32(count)
	return
}

//神王报名阶段
func (t *ChuangShiConstantTemplate) IfShenWangSign(now int64) bool {
	//每周开启的日期
	day := timeutils.MillisecondToTime(now).Weekday()
	if day == time.Sunday {
		day = 7
	}
	dayInt := int32(day)

	if !utils.ContainInt32(t.signWeekdayArr, dayInt) {
		return false
	}

	//时间内
	beginDay, _ := timeutils.BeginOfNow(now)
	begin := beginDay + t.signTimeHourArr[0]
	end := beginDay + t.signTimeHourArr[1]
	if now < begin || now > end {
		return false
	}

	return true
}

//神王投票阶段
func (t *ChuangShiConstantTemplate) IfShenWangVote(now int64) bool {
	//每周开启的日期
	day := timeutils.MillisecondToTime(now).Weekday()
	if day == time.Sunday {
		day = 7
	}
	dayInt := int32(day)

	if !utils.ContainInt32(t.voteWeekdayArr, dayInt) {
		return false
	}

	//时间内
	beginDay, _ := timeutils.BeginOfNow(now)
	begin := beginDay + t.voteTimeHourArr[0]
	end := beginDay + t.voteTimeHourArr[1]
	if now < begin || now > end {
		return false
	}

	return true
}

func (t *ChuangShiConstantTemplate) GetChangedCampUseItemMap() map[int32]int32 {
	return t.changedCampItemMap
}

func (t *ChuangShiConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiConstantTemplate) FileName() string {
	return "tb_chuangshi_constant.json"
}

func (t *ChuangShiConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//每周几报名时间
	signWeekday, err := utils.SplitAsIntArray(t.BaomingWeekday)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.BaomingWeekday)
		return template.NewTemplateFieldError("BaomingWeekday", err)
	}
	t.signWeekdayArr = append(t.signWeekdayArr, signWeekday...)

	//每周几选举时间
	voteWeekday, err := utils.SplitAsIntArray(t.XuanjuWeekday)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.XuanjuWeekday)
		return template.NewTemplateFieldError("XuanjuWeekday", err)
	}
	t.voteWeekdayArr = append(t.voteWeekdayArr, voteWeekday...)

	//
	timeStr := strings.Split(t.BaomingTime, ",")
	if len(timeStr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.BaomingTime)
		return template.NewTemplateFieldError("BaomingTime", err)
	}
	for _, value := range timeStr {
		timeInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.BaomingTime)
			return template.NewTemplateFieldError("BaomingTime", err)
		}
		time := int64(0)
		if timeInt == 2400 {
			time = int64(common.DAY)
		} else {
			time, err = timeutils.ParseDayOfHHMM(value)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.BaomingTime)
				return template.NewTemplateFieldError("BaomingTime", err)
			}
		}

		t.signTimeHourArr = append(t.signTimeHourArr, time)
	}
	//
	voteTimeStr := strings.Split(t.XuanjuTime, ",")
	if len(voteTimeStr) != 2 {
		err = fmt.Errorf("[%s] invalid", t.XuanjuTime)
		return template.NewTemplateFieldError("XuanjuTime", err)
	}
	for _, value := range voteTimeStr {
		timeInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.XuanjuTime)
			return template.NewTemplateFieldError("XuanjuTime", err)
		}
		time := int64(0)
		if timeInt == 2400 {
			time = int64(common.DAY)
		} else {
			time, err = timeutils.ParseDayOfHHMM(value)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.XuanjuTime)
				return template.NewTemplateFieldError("XuanjuTime", err)
			}
		}

		t.voteTimeHourArr = append(t.voteTimeHourArr, time)
	}
	return
}

func (t *ChuangShiConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	err = validator.MinValidate(float64(t.CityRewTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CityRewTime)
		return template.NewTemplateFieldError("CityRewTime", err)
	}

	//
	to := template.GetTemplateService().Get(int(t.GenghunUseItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.GenghunUseItemId)
		return template.NewTemplateFieldError("GenghunUseItemId", err)
	}

	//数量
	err = validator.MinValidate(float64(t.GenghunUseItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GenghunUseItemCount)
		return template.NewTemplateFieldError("GenghunUseItemCount", err)
	}

	return
}

func (t *ChuangShiConstantTemplate) PatchAfterCheck() {
	t.changedCampItemMap = make(map[int32]int32)
	t.changedCampItemMap[t.GenghunUseItemId] = t.GenghunUseItemCount
}

func init() {
	template.Register((*ChuangShiConstantTemplate)(nil))
}
