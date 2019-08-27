package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

type ArenapvpTemplate struct {
	*ArenapvpTemplateVO
	arenapvpType    arenapvptypes.ArenapvpType
	winMailItemMap  map[int32]int32
	loseMailItemMap map[int32]int32
	beginTime       int64
	endTime         int64
	mapTemp         *MapTemplate
	nextTemp        *ArenapvpTemplate
	pos1            coretypes.Position
	pos2            coretypes.Position
}

func (t *ArenapvpTemplate) GetPos1() coretypes.Position {
	return t.pos1
}

func (t *ArenapvpTemplate) GetPos2() coretypes.Position {
	return t.pos2
}

func (t *ArenapvpTemplate) IfCanGuess(now int64) bool {
	begin := t.GetBeginTime(now)
	if now >= begin {
		return false
	}

	return true
}

func (t *ArenapvpTemplate) GetArenapvpType() arenapvptypes.ArenapvpType {
	return t.arenapvpType
}

func (t *ArenapvpTemplate) GetNextTemp() *ArenapvpTemplate {
	return t.nextTemp
}

func (t *ArenapvpTemplate) GetMapTemp() *MapTemplate {
	return t.mapTemp
}

func (t *ArenapvpTemplate) GetPVPNum() int32 {
	return t.PVPPlayerCount / t.PVPCount
}

func (t *ArenapvpTemplate) GetWinnerCount() int32 {
	if t.nextTemp == nil {
		return 0
	}

	return t.nextTemp.PVPPlayerCount / t.PVPCount
}

func (t *ArenapvpTemplate) GetWinMailItemMap() map[int32]int32 {
	return t.winMailItemMap
}

func (t *ArenapvpTemplate) GetLoseMailItemMap() map[int32]int32 {
	return t.loseMailItemMap
}

func (t *ArenapvpTemplate) GetEndTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.endTime
}

func (t *ArenapvpTemplate) GetBeginTime(now int64) int64 {
	beginDay, _ := timeutils.BeginOfNow(now)
	return beginDay + t.beginTime
}

func (t *ArenapvpTemplate) GetRemainReliveTimes(useTimes int32) int32 {
	if useTimes < 0 {
		return t.RebornCountMax
	}

	remain := t.RebornCountMax - useTimes
	if remain < 0 {
		return 0
	}
	return remain
}

func (t *ArenapvpTemplate) IsOnArenaTime(now int64) bool {
	if now < t.GetBeginTime(now) {
		return false
	}

	if now > t.GetEndTime(now) {
		return false
	}

	return true
}

func (t *ArenapvpTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	itemArr, err := coreutils.SplitAsIntArray(t.JingcaiLoseGetItemId)
	if err != nil {
		return template.NewTemplateFieldError("JingcaiLoseGetItemId", err)
	}
	numArr, err := coreutils.SplitAsIntArray(t.JingcaiLoseGetItemCount)
	if err != nil {
		return template.NewTemplateFieldError("JingcaiLoseGetItemCount", err)
	}
	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("[%s] or [%s] invalid ", t.JingcaiLoseGetItemId, t.JingcaiLoseGetItemCount)
		return template.NewTemplateFieldError("JingcaiLoseGetItemId or JingcaiLoseGetItemCount", err)
	}
	t.loseMailItemMap = make(map[int32]int32)
	for i := 0; i < len(itemArr); i++ {
		t.loseMailItemMap[itemArr[i]] = numArr[i]
	}

	//
	item2Arr, err := coreutils.SplitAsIntArray(t.JingcaiWinGetItemId)
	if err != nil {
		return template.NewTemplateFieldError("JingcaiWinGetItemId", err)
	}
	num2Arr, err := coreutils.SplitAsIntArray(t.JingcaiWinGetItemCount)
	if err != nil {
		return template.NewTemplateFieldError("JingcaiWinGetItemCount", err)
	}
	if len(item2Arr) != len(num2Arr) {
		err = fmt.Errorf("[%s] or [%s] invalid ", t.JingcaiWinGetItemId, t.JingcaiWinGetItemCount)
		return template.NewTemplateFieldError("JingcaiWinGetItemId or JingcaiWinGetItemCount", err)
	}

	t.winMailItemMap = make(map[int32]int32)
	for i := 0; i < len(item2Arr); i++ {
		t.winMailItemMap[item2Arr[i]] = num2Arr[i]
	}

	//
	mapTo := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mapTo == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.mapTemp = mapTo.(*MapTemplate)

	//
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*ArenapvpTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*ArenapvpTemplate)
		if t.nextTemp.Type-t.Type != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

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
	return
}
func (t *ArenapvpTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.arenapvpType = arenapvptypes.ArenapvpType(t.Type)
	if !t.arenapvpType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	if err = validator.MinValidate(float64(t.JingchaiUseBindgold), 0, true); err != nil {
		return template.NewTemplateFieldError("JingchaiUseBindgold", err)
	}

	if err = validator.MinValidate(float64(t.RebornCountMax), 1, true); err != nil {
		return template.NewTemplateFieldError("RebornCountMax", err)
	}

	if err = validator.MinValidate(float64(t.PVPPlayerCount), 1, true); err != nil {
		return template.NewTemplateFieldError("PVPPlayerCount", err)
	}

	if err = validator.MinValidate(float64(t.PVPCount), 1, true); err != nil {
		return template.NewTemplateFieldError("PVPCount", err)
	}

	if err = validator.MinValidate(float64(t.ZhanDouTime), 0, true); err != nil {
		return template.NewTemplateFieldError("ZhanDouTime", err)
	}

	if t.PVPPlayerCount%t.PVPCount != 0 {
		err = fmt.Errorf("[%d][%d] invalid", t.PVPPlayerCount, t.PVPCount)
		return template.NewTemplateFieldError("PVPPlayerCount or PVPCount", err)
	}

	//活动开始时间
	beginTime, err := timeutils.ParseDayOfHHMM(t.BeginTime)
	if err != nil {
		return template.NewTemplateFieldError("BeginTime", fmt.Errorf("[%s] invalid", t.BeginTime))
	}
	t.beginTime = beginTime

	//活动结束时间
	endInt, err := strconv.ParseInt(t.EndTime, 10, 64)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.EndTime)
		return template.NewTemplateFieldError("EndTime", err)
	}
	if endInt == 2400 {
		t.endTime = int64(common.DAY)
	} else {
		endTime, err := timeutils.ParseDayOfHHMM(t.EndTime)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.EndTime)
			return template.NewTemplateFieldError("EndTime", err)
		}
		t.endTime = endTime
	}

	for itemId, itemNum := range t.winMailItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.JingcaiWinGetItemId)
			return template.NewTemplateFieldError("JingcaiWinGetItemId", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.JingcaiWinGetItemCount)
			return template.NewTemplateFieldError("JingcaiWinGetItemCount", err)
		}
	}

	for itemId, itemNum := range t.loseMailItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.JingcaiLoseGetItemId)
			return template.NewTemplateFieldError("JingcaiLoseGetItemId", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.JingcaiLoseGetItemCount)
			return template.NewTemplateFieldError("JingcaiLoseGetItemCount", err)
		}
	}

	//验证
	err = validator.MinValidate(float64(t.JingcaiWinBindgold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.JingcaiWinBindgold)
		err = template.NewTemplateFieldError("JingcaiWinBindgold", err)
		return
	}
	//
	if !t.mapTemp.GetMap().IsMask(t.pos1.X, t.pos1.Z) {
		err = fmt.Errorf("BirthPos1 pos invalid")
		return template.NewTemplateFieldError("BirthPos1", err)
	}
	t.pos1.Y = t.mapTemp.GetMap().GetHeight(t.pos1.X, t.pos1.Z)
	//
	if !t.mapTemp.GetMap().IsMask(t.pos2.X, t.pos2.Z) {
		err = fmt.Errorf("BirthPos2 born pos invalid")
		return template.NewTemplateFieldError("BirthPos2", err)
	}
	t.pos2.Y = t.mapTemp.GetMap().GetHeight(t.pos2.X, t.pos2.Z)

	return
}

func (t *ArenapvpTemplate) PatchAfterCheck() {
	if t.JingcaiWinBindgold > 0 {
		t.winMailItemMap[constanttypes.BindGoldItem] = t.JingcaiWinBindgold
	}
}

func (t *ArenapvpTemplate) TemplateId() int {
	return t.Id
}

func (at *ArenapvpTemplate) FileName() string {
	return "tb_biwudahui_time.json"
}

func init() {
	template.Register((*ArenapvpTemplate)(nil))
}
