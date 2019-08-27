package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
	"strings"
)

type ChuangShiCityConstantTemplate struct {
	*ChuangShiCityConstantTemplateVO
	mapTemp   *MapTemplate
	beginTime int64 //城池进入时间
	endTime   int64 //城池进入时间
	pos1      coretypes.Position
	pos2      coretypes.Position
}

func (t *ChuangShiCityConstantTemplate) GetEndTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.endTime
}

func (t *ChuangShiCityConstantTemplate) GetBeginTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.beginTime
}

func (t *ChuangShiCityConstantTemplate) GetPos1() coretypes.Position {
	return t.pos1
}

func (t *ChuangShiCityConstantTemplate) GetPos2() coretypes.Position {
	return t.pos2
}

func (t *ChuangShiCityConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiCityConstantTemplate) FileName() string {
	return "tb_chuangshi_city_constant.json"
}

func (t *ChuangShiCityConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	mapTo := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mapTo == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.mapTemp = mapTo.(*MapTemplate)

	//
	pos1Arr, err := coreutils.SplitAsFloatArray(t.BirthPos1)
	if err != nil {
		return
	}
	if len(pos1Arr) != 3 {
		err = fmt.Errorf("[%s] invalid", t.BirthPos1)
		return template.NewTemplateFieldError("BirthPos1", err)
	}
	pos1 := coretypes.Position{
		X: pos1Arr[0],
		Y: pos1Arr[1],
		Z: pos1Arr[2],
	}
	t.pos1 = pos1

	//
	pos2Arr, err := coreutils.SplitAsFloatArray(t.BirthPos2)
	if err != nil {
		return
	}
	if len(pos2Arr) != 3 {
		err = fmt.Errorf("[%s] invalid", t.BirthPos2)
		return template.NewTemplateFieldError("BirthPos2", err)
	}
	pos2 := coretypes.Position{
		X: pos2Arr[0],
		Y: pos2Arr[1],
		Z: pos2Arr[2],
	}
	t.pos2 = pos2

	//
	if len(t.JinRuTime) > 0 {
		timeStr := strings.Split(t.JinRuTime, ",")
		if len(timeStr) != 2 {
			err = fmt.Errorf("[%s] invalid", t.JinRuTime)
			return template.NewTemplateFieldError("JinRuTime", err)
		}
		for index, value := range timeStr {
			timeInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				err = fmt.Errorf("[%s] invalid", t.JinRuTime)
				return template.NewTemplateFieldError("JinRuTime", err)
			}
			time := int64(0)
			if timeInt == 2400 {
				time = int64(common.DAY)
			} else {
				time, err = timeutils.ParseDayOfHHMM(value)
				if err != nil {
					err = fmt.Errorf("[%s] invalid", t.JinRuTime)
					return template.NewTemplateFieldError("JinRuTime", err)
				}
			}

			if index == 0 {
				t.beginTime = time
			} else {
				t.endTime = time
			}
		}
	}
	return
}

func (t *ChuangShiCityConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 出生位置1
	if !t.mapTemp.GetMap().IsMask(t.pos1.X, t.pos1.Z) {
		err = fmt.Errorf("BirthPos1 pos invalid")
		return template.NewTemplateFieldError("BirthPos1", err)
	}
	t.pos1.Y = t.mapTemp.GetMap().GetHeight(t.pos1.X, t.pos1.Z)

	// 出生位置2
	if !t.mapTemp.GetMap().IsMask(t.pos2.X, t.pos2.Z) {
		err = fmt.Errorf("BirthPos2 born pos invalid")
		return template.NewTemplateFieldError("BirthPos2", err)
	}
	t.pos2.Y = t.mapTemp.GetMap().GetHeight(t.pos2.X, t.pos2.Z)

	//
	if len(t.JinRuTime) > 0 {
		if t.endTime <= t.beginTime {
			err = fmt.Errorf("[%s] invalid", t.JinRuTime)
			return template.NewTemplateFieldError("JinRuTime", err)
		}
	}

	return
}

func (t *ChuangShiCityConstantTemplate) PatchAfterCheck() {

}

func init() {
	template.Register((*ChuangShiCityConstantTemplate)(nil))
}
