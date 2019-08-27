package player

import (
	"encoding/json"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sort"

	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	wingcommon "fgame/fgame/game/wing/common"
	"fgame/fgame/game/wing/dao"
	wingentity "fgame/fgame/game/wing/entity"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	wingtypes "fgame/fgame/game/wing/types"
	"fgame/fgame/game/wing/wing"

	"github.com/pkg/errors"
)

//战翼对象
type PlayerWingObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	AdvanceId   int
	WingId      int32
	UnrealLevel int32
	UnrealNum   int32
	UnrealPro   int32
	UnrealList  []int
	TimesNum    int32
	Bless       int32
	BlessTime   int64
	FeatherId   int32
	FeatherNum  int32
	FeatherPro  int32
	Hidden      int32
	Power       int64
	FPower      int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerWingObject(pl player.Player) *PlayerWingObject {
	pwo := &PlayerWingObject{
		player: pl,
	}
	return pwo
}

func convertNewPlayerWingObjectToEntity(pwo *PlayerWingObject) (*wingentity.PlayerWingEntity, error) {
	unrealInfoBytes, err := json.Marshal(pwo.UnrealList)
	if err != nil {
		return nil, err
	}
	e := &wingentity.PlayerWingEntity{
		Id:          pwo.Id,
		PlayerId:    pwo.PlayerId,
		AdvancedId:  pwo.AdvanceId,
		WingId:      pwo.WingId,
		UnrealLevel: pwo.UnrealLevel,
		UnrealNum:   pwo.UnrealNum,
		UnrealPro:   pwo.UnrealPro,
		UnrealInfo:  string(unrealInfoBytes),
		TimesNum:    pwo.TimesNum,
		Bless:       pwo.Bless,
		BlessTime:   pwo.BlessTime,
		FeatherId:   pwo.FeatherId,
		FeatherNum:  pwo.FeatherNum,
		FeatherPro:  pwo.FeatherPro,
		Hidden:      pwo.Hidden,
		Power:       pwo.Power,
		FPower:      pwo.FPower,
		UpdateTime:  pwo.UpdateTime,
		CreateTime:  pwo.CreateTime,
		DeleteTime:  pwo.DeleteTime,
	}
	return e, nil
}

func (pwo *PlayerWingObject) GetPlayerId() int64 {
	return pwo.PlayerId
}

func (pwo *PlayerWingObject) GetDBId() int64 {
	return pwo.Id
}

func (pwo *PlayerWingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerWingObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerWingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*wingentity.PlayerWingEntity)

	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}
	pwo.Id = pse.Id
	pwo.PlayerId = pse.PlayerId
	pwo.AdvanceId = pse.AdvancedId
	pwo.WingId = pse.WingId
	pwo.UnrealLevel = pse.UnrealLevel
	pwo.UnrealNum = pse.UnrealNum
	pwo.UnrealPro = pse.UnrealPro
	pwo.UnrealList = unrealList
	pwo.TimesNum = pse.TimesNum
	pwo.Bless = pse.Bless
	pwo.BlessTime = pse.BlessTime
	pwo.FeatherId = pse.FeatherId
	pwo.FeatherNum = pse.FeatherNum
	pwo.FeatherPro = pse.FeatherPro
	pwo.Hidden = pse.Hidden
	pwo.Power = pse.Power
	pwo.FPower = pse.FPower
	pwo.UpdateTime = pse.UpdateTime
	pwo.CreateTime = pse.CreateTime
	pwo.DeleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerWingObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("wing: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

//战翼非进阶对象
type PlayerWingOtherObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Typ        wingtypes.WingType
	WingId     int32
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerWingOtherObject(pl player.Player) *PlayerWingOtherObject {
	pwo := &PlayerWingOtherObject{
		player: pl,
	}
	return pwo
}

func convertWingOtherObjectToEntity(pwo *PlayerWingOtherObject) (*wingentity.PlayerWingOtherEntity, error) {

	e := &wingentity.PlayerWingOtherEntity{
		Id:         pwo.Id,
		PlayerId:   pwo.PlayerId,
		Typ:        int32(pwo.Typ),
		WingId:     pwo.WingId,
		Level:      pwo.Level,
		UpNum:      pwo.UpNum,
		UpPro:      pwo.UpPro,
		UpdateTime: pwo.UpdateTime,
		CreateTime: pwo.CreateTime,
		DeleteTime: pwo.DeleteTime,
	}
	return e, nil
}

func (pwo *PlayerWingOtherObject) GetPlayerId() int64 {
	return pwo.PlayerId
}

func (pwo *PlayerWingOtherObject) GetDBId() int64 {
	return pwo.Id
}

func (pwo *PlayerWingOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertWingOtherObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerWingOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*wingentity.PlayerWingOtherEntity)

	pwo.Id = pse.Id
	pwo.PlayerId = pse.PlayerId
	pwo.Typ = wingtypes.WingType(pse.Typ)
	pwo.WingId = pse.WingId
	pwo.Level = pse.Level
	pwo.UpNum = pse.UpNum
	pwo.UpPro = pse.UpPro
	pwo.UpdateTime = pse.UpdateTime
	pwo.CreateTime = pse.CreateTime
	pwo.DeleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerWingOtherObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("wing: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

//战翼试用卡获得试用对象
type PlayerWingTrialObject struct {
	player       player.Player
	Id           int64
	PlayerId     int64
	TrialOrderId int32
	ActiveTime   int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerWingTrialObject(pl player.Player) *PlayerWingTrialObject {
	pwco := &PlayerWingTrialObject{
		player: pl,
	}
	return pwco
}

func (pwco *PlayerWingTrialObject) GetPlayerId() int64 {
	return pwco.PlayerId
}

func (pwco *PlayerWingTrialObject) GetDBId() int64 {
	return pwco.Id
}

func (pwco *PlayerWingTrialObject) ToEntity() (e storage.Entity, err error) {
	e = &wingentity.PlayerWingTrialEntity{
		Id:           pwco.Id,
		PlayerId:     pwco.PlayerId,
		TrialOrderId: pwco.TrialOrderId,
		ActiveTime:   pwco.ActiveTime,
		UpdateTime:   pwco.UpdateTime,
		CreateTime:   pwco.CreateTime,
		DeleteTime:   pwco.DeleteTime,
	}
	return e, nil
}

func (pwco *PlayerWingTrialObject) FromEntity(e storage.Entity) error {
	psce, _ := e.(*wingentity.PlayerWingTrialEntity)

	pwco.Id = psce.Id
	pwco.PlayerId = psce.PlayerId
	pwco.TrialOrderId = psce.TrialOrderId
	pwco.ActiveTime = psce.ActiveTime
	pwco.UpdateTime = psce.UpdateTime
	pwco.CreateTime = psce.CreateTime
	pwco.DeleteTime = psce.DeleteTime
	return nil
}

func (pwco *PlayerWingTrialObject) SetModified() {
	e, err := pwco.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Wing"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwco.player.AddChangedObject(obj)
	return
}

//玩家战翼管理器
type PlayerWingDataManager struct {
	p player.Player
	//玩家战翼对象
	PlayerWingObject *PlayerWingObject
	//玩家非进阶战翼对象
	playerOtherMap map[wingtypes.WingType]map[int32]*PlayerWingOtherObject
	//战翼试用卡获得阶数
	playerWingTrialObject *PlayerWingTrialObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (pwdm *PlayerWingDataManager) Player() player.Player {
	return pwdm.p
}

//加载
func (pwdm *PlayerWingDataManager) Load() (err error) {
	pwdm.playerOtherMap = make(map[wingtypes.WingType]map[int32]*PlayerWingOtherObject)

	//加载玩家战翼试用阶数信息
	wingTrialEntity, err := dao.GetWingDao().GetWingTrialEntity(pwdm.p.GetId())
	if err != nil {
		return
	}
	if wingTrialEntity == nil {
		pwdm.initPlayerWingTrialObject()
	} else {
		pwdm.playerWingTrialObject = NewPlayerWingTrialObject(pwdm.p)
		pwdm.playerWingTrialObject.FromEntity(wingTrialEntity)
		if pwdm.WingTrialIsOverdued() {
			pwdm.restWingTrial()
		}
	}

	//加载玩家战翼信息
	wingEntity, err := dao.GetWingDao().GetWingEntity(pwdm.p.GetId())
	if err != nil {
		return
	}
	if wingEntity == nil {
		pwdm.initPlayerWingObject()
	} else {
		pwdm.PlayerWingObject = NewPlayerWingObject(pwdm.p)
		pwdm.PlayerWingObject.FromEntity(wingEntity)
	}

	//加载玩家非进阶战翼信息
	wingOtherList, err := dao.GetWingDao().GetWingOtherList(pwdm.p.GetId())
	if err != nil {
		return
	}

	//非进阶战翼信息
	for _, wingOther := range wingOtherList {
		pwo := NewPlayerWingOtherObject(pwdm.p)
		pwo.FromEntity(wingOther)

		typ := wingtypes.WingType(wingOther.Typ)

		playerOtherMap, exist := pwdm.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerWingOtherObject)
			pwdm.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pwo.WingId] = pwo
	}

	return nil
}

func (pwdm *PlayerWingDataManager) initPlayerWingTrialObject() {
	pwto := NewPlayerWingTrialObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwto.Id = id
	pwto.PlayerId = pwdm.p.GetId()
	pwto.TrialOrderId = 0
	pwto.ActiveTime = 0
	pwto.CreateTime = now
	pwto.SetModified()
	pwdm.playerWingTrialObject = pwto
}

//第一次初始化
func (pwdm *PlayerWingDataManager) initPlayerWingObject() {
	pwo := NewPlayerWingObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pwdm.p.GetRole(), pwdm.p.GetSex())
	advanceId := playerCreateTemplate.Wing
	featherId := playerCreateTemplate.Feather
	//生成id
	pwo.PlayerId = pwdm.p.GetId()
	pwo.AdvanceId = int(advanceId)
	pwo.WingId = int32(0)
	pwo.UnrealLevel = int32(0)
	pwo.UnrealNum = int32(0)
	pwo.UnrealPro = int32(0)
	pwo.UnrealList = make([]int, 0, 8)
	pwo.TimesNum = int32(0)
	pwo.Bless = int32(0)
	pwo.BlessTime = int64(0)
	pwo.FeatherId = featherId
	pwo.FeatherNum = 0
	pwo.FeatherPro = 0
	pwo.Hidden = 0
	pwo.Power = int64(0)
	pwo.FPower = int64(0)
	pwo.CreateTime = now
	pwdm.PlayerWingObject = pwo
	pwo.SetModified()
}

func (pwdm *PlayerWingDataManager) WingTrialIsOverdued() (overdueFlag bool) {
	if pwdm.playerWingTrialObject.TrialOrderId == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	activeTime := pwdm.playerWingTrialObject.ActiveTime
	trialTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeWingTrialTime)
	existTime := now - activeTime
	if existTime >= int64(trialTime) {
		overdueFlag = true
	}
	return
}

//重置战翼试用
func (pwdm *PlayerWingDataManager) restWingTrial() {
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerWingTrialObject.TrialOrderId = 0
	pwdm.playerWingTrialObject.ActiveTime = 0
	pwdm.playerWingTrialObject.UpdateTime = now
	pwdm.playerWingTrialObject.SetModified()
}

//增加非进阶战翼
func (pwdm *PlayerWingDataManager) newPlayerWingOtherObject(typ wingtypes.WingType, wingId int32) (err error) {

	playerOtherMap, exist := pwdm.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerWingOtherObject)
		pwdm.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[wingId]
	if exist {
		return
	}

	pwo := NewPlayerWingOtherObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.Id = id
	//生成id
	pwo.PlayerId = pwdm.p.GetId()
	pwo.Typ = typ
	pwo.WingId = wingId
	pwo.Level = 0
	pwo.UpNum = 0
	pwo.UpPro = 0
	pwo.CreateTime = now
	playerOtherMap[wingId] = pwo
	pwo.SetModified()
	return
}

func (pwdm *PlayerWingDataManager) refreshBless() (err error) {
	now := global.GetGame().GetTimeService().Now()
	number := int32(pwdm.PlayerWingObject.AdvanceId)
	nextNumber := number + 1
	wingTemplate := wing.GetWingService().GetWingNumber(nextNumber)
	if wingTemplate == nil {
		return
	}
	if !wingTemplate.GetIsClear() {
		return
	}
	lastTime := pwdm.PlayerWingObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pwdm.PlayerWingObject.Bless = 0
			pwdm.PlayerWingObject.BlessTime = 0
			pwdm.PlayerWingObject.TimesNum = 0
			pwdm.PlayerWingObject.SetModified()
		}
	}
	return
}

//加载后
func (pwdm *PlayerWingDataManager) AfterLoad() (err error) {
	err = pwdm.refreshBless()
	pwdm.heartbeatRunner.AddTask(CreateWingTask(pwdm.p))
	return nil
}

//战翼信息对象
func (pwdm *PlayerWingDataManager) GetWingInfo() *PlayerWingObject {
	pwdm.refreshBless()
	return pwdm.PlayerWingObject
}

func (pwdm *PlayerWingDataManager) GetWingAdvancedId() int32 {
	return int32(pwdm.PlayerWingObject.AdvanceId)
}

//战翼试用对象
func (pwdm *PlayerWingDataManager) GetWingTrialInfo() *PlayerWingTrialObject {
	return pwdm.playerWingTrialObject
}

func (m *PlayerWingDataManager) GetWingId() int32 {
	if m.PlayerWingObject.Hidden != 0 {
		return 0
	}
	if m.playerWingTrialObject.TrialOrderId != 0 {
		return m.playerWingTrialObject.TrialOrderId
	}
	if m.PlayerWingObject.WingId != 0 {
		return m.PlayerWingObject.WingId
	}
	wingTemplate := wing.GetWingService().GetWingNumber(int32(m.PlayerWingObject.AdvanceId))
	if wingTemplate == nil {
		return 0
	}
	return int32(wingTemplate.TemplateId())
}

//获取玩家非进阶战翼对象
func (pwdm *PlayerWingDataManager) GetWingOtherMap() map[wingtypes.WingType]map[int32]*PlayerWingOtherObject {
	return pwdm.playerOtherMap
}

// func (pwdm *PlayerWingDataManager) EatUnrealDan(pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}

// 	if sucess {
// 		pwdm.PlayerWingObject.UnrealLevel += 1
// 		pwdm.PlayerWingObject.UnrealNum = 0
// 		pwdm.PlayerWingObject.UnrealPro = pro
// 	} else {
// 		pwdm.PlayerWingObject.UnrealNum += 1
// 		pwdm.PlayerWingObject.UnrealPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	pwdm.PlayerWingObject.UpdateTime = now
// 	pwdm.PlayerWingObject.SetModified()
// 	return
// }

func (pwdm *PlayerWingDataManager) EatUnrealDan(level int32) {
	if pwdm.PlayerWingObject.UnrealLevel == level || level <= 0 {
		return
	}
	huanHuaTemplate := wing.GetWingService().GetWingHuanHuaTemplate(level)
	if huanHuaTemplate == nil {
		return
	}
	pwdm.PlayerWingObject.UnrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	return
}

//心跳
func (pwdm *PlayerWingDataManager) Heartbeat() {
	pwdm.heartbeatRunner.Heartbeat()
}

func (pwdm *PlayerWingDataManager) IsWingHidden() bool {
	return pwdm.PlayerWingObject.Hidden == 1
}

//进阶
func (pwdm *PlayerWingDataManager) WingAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := pwdm.PlayerWingObject.AdvanceId + 1
		wingTemplate := wing.GetWingService().GetWingNumber(int32(nextAdvancedId))
		if wingTemplate == nil {
			return
		}
		pwdm.PlayerWingObject.AdvanceId += 1
		pwdm.PlayerWingObject.TimesNum = 0
		pwdm.PlayerWingObject.Bless = 0
		pwdm.PlayerWingObject.BlessTime = 0
		pwdm.PlayerWingObject.WingId = int32(0)
		gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
		gameevent.Emit(wingeventtypes.EventTypeWingAdvanced, pwdm.p, pwdm.PlayerWingObject.AdvanceId)
		if pwdm.playerWingTrialObject.TrialOrderId != 0 {
			pwdm.RemoveWingTrial(false)
		}
	} else {
		pwdm.PlayerWingObject.TimesNum += addTimes
		if pwdm.PlayerWingObject.Bless == 0 {
			pwdm.PlayerWingObject.BlessTime = now
		}
		pwdm.PlayerWingObject.Bless += pro
	}
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	return
}

//直升券进阶
func (pwdm *PlayerWingDataManager) WingAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := pwdm.PlayerWingObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		wingTemplate := wing.GetWingService().GetWingNumber(int32(nextAdvancedId))
		if wingTemplate == nil {
			return
		}
		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.AdvanceId += canAddNum
	pwdm.PlayerWingObject.TimesNum = 0
	pwdm.PlayerWingObject.Bless = 0
	pwdm.PlayerWingObject.BlessTime = 0
	pwdm.PlayerWingObject.WingId = int32(0)
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()

	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	gameevent.Emit(wingeventtypes.EventTypeWingAdvanced, pwdm.p, pwdm.PlayerWingObject.AdvanceId)
	if pwdm.playerWingTrialObject.TrialOrderId != 0 {
		pwdm.RemoveWingTrial(false)
	}
	return
}

// 战翼激活卡
func (pwdm *PlayerWingDataManager) WingActivate() (err error) {
	if pwdm.PlayerWingObject.AdvanceId > 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.AdvanceId = 1
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingAdvanced, pwdm.p, pwdm.PlayerWingObject.AdvanceId)
	if pwdm.playerWingTrialObject.TrialOrderId != 0 {
		pwdm.RemoveWingTrial(false)
	}

	return
}

//展示隐藏战翼
func (pwdm *PlayerWingDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		pwdm.PlayerWingObject.Hidden = 1
	} else {
		pwdm.PlayerWingObject.Hidden = 0
	}
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	if !hiddenFlag {
		gameevent.Emit(wingeventtypes.EventTypeWingUse, pwdm.p, nil)
	}
	return
}

//护体仙羽培养
func (pwdm *PlayerWingDataManager) FeatherFeed(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	if sucess {
		nextAdvancedId := pwdm.PlayerWingObject.FeatherId + 1
		featherTemplate := wing.GetWingService().GetFeather(nextAdvancedId)
		if featherTemplate == nil {
			return
		}
		pwdm.PlayerWingObject.FeatherId += 1
		pwdm.PlayerWingObject.FeatherNum = 0
		pwdm.PlayerWingObject.FeatherPro = pro

		gameevent.Emit(wingeventtypes.EventTypeFeatherAdvanced, pwdm.p, pwdm.PlayerWingObject.FeatherId)
	} else {
		pwdm.PlayerWingObject.FeatherNum += addTimes
		pwdm.PlayerWingObject.FeatherPro += pro
	}

	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	return
}

//护体仙羽培养
func (pwdm *PlayerWingDataManager) FeatherFeedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := pwdm.PlayerWingObject.FeatherId + 1
	for addAdvancedNum > 0 {
		featherTemplate := wing.GetWingService().GetFeather(nextAdvancedId)
		if featherTemplate == nil {
			return
		}

		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.FeatherId += int32(canAddNum)
	pwdm.PlayerWingObject.FeatherNum = 0
	pwdm.PlayerWingObject.FeatherPro = 0
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeFeatherAdvanced, pwdm.p, pwdm.PlayerWingObject.FeatherId)
	return
}

//设置幻化信息
func (pwdm *PlayerWingDataManager) AddUnrealInfo(wingId int) {
	if wingId <= 0 {
		return
	}
	pwdm.PlayerWingObject.UnrealList = append(pwdm.PlayerWingObject.UnrealList, wingId)
	sort.Ints(pwdm.PlayerWingObject.UnrealList)
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	pwdm.WingOtherGet(int32(wingId))

	gameevent.Emit(wingeventtypes.EventTypeWingUnrealActivate, pwdm.p, wingId)
	return
}

//是否能幻化
func (pwdm *PlayerWingDataManager) IsCanUnreal(wingId int) bool {
	wingTemplate := wing.GetWingService().GetWing(wingId)
	if wingTemplate == nil {
		return false
	}
	curAdvancedId := pwdm.PlayerWingObject.AdvanceId
	//食用幻化丹等级
	curUnrealLevel := pwdm.PlayerWingObject.UnrealLevel
	//幻化条件
	for condType, cond := range wingTemplate.GetMagicParamXUMap() {
		scondType := wingtypes.WingUCondType(condType)
		switch scondType {
		//战翼阶别
		case wingtypes.WingUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹等级条件
		case wingtypes.WingUCondTypeU:
			if int32(curUnrealLevel) < cond {
				return false
			}
		default:
			break
		}
	}
	return true
}

//是否已幻化
func (pwdm *PlayerWingDataManager) IsUnrealed(wingId int) bool {
	uList := pwdm.PlayerWingObject.UnrealList
	for _, v := range uList {
		if v == wingId {
			//容错处理
			pwdm.WingOtherGet(int32(wingId))
			return true
		}
	}
	return false
}

//幻化
func (pwdm *PlayerWingDataManager) Unreal(wingId int) (flag bool) {
	if !pwdm.IsUnrealed(wingId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.WingId = int32(wingId)
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	flag = true
	return
}

//卸下
func (pwdm *PlayerWingDataManager) Unload() {
	pwdm.PlayerWingObject.WingId = 0
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	return
}

//战翼战斗力
func (pwdm *PlayerWingDataManager) WingPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := pwdm.PlayerWingObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.Power = power
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()

	gameevent.Emit(wingeventtypes.EventTypeWingPowerChanged, pwdm.p, power)
	return
}

//护体仙羽战斗力
func (pwdm *PlayerWingDataManager) WingFeatherPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := pwdm.PlayerWingObject.FPower
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.FPower = power
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()

	gameevent.Emit(wingeventtypes.EventTypeFeatherPowerChanged, pwdm.p, power)
	return
}

//非进阶战翼激活
func (pwdm *PlayerWingDataManager) WingOtherGet(wingId int32) {
	wingTemplate := wing.GetWingService().GetWing(int(wingId))
	if wingTemplate == nil {
		return
	}
	typ := wingTemplate.GetTyp()
	if typ == wingtypes.WingTypeAdvanced {
		return
	}
	pwdm.newPlayerWingOtherObject(typ, wingId)
	return
}

//战翼试用
func (pwdm *PlayerWingDataManager) WingTrialOrder() {
	now := global.GetGame().GetTimeService().Now()
	trialOrderId := wing.GetWingService().GetWingTrialOrderId()
	pwdm.playerWingTrialObject.TrialOrderId = trialOrderId
	pwdm.playerWingTrialObject.ActiveTime = now
	pwdm.playerWingTrialObject.UpdateTime = now
	pwdm.playerWingTrialObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)

}

//战翼试用过期
func (pwdm *PlayerWingDataManager) RemoveWingTrial(bResult bool) {
	trialWingId := pwdm.playerWingTrialObject.TrialOrderId
	pwdm.restWingTrial()
	eventData := wingeventtypes.CreateWingTrialOverdueEventData(trialWingId, bResult)
	gameevent.Emit(wingeventtypes.EventTypeWingTrialOverdue, pwdm.p, eventData)
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
}

//是否已拥有该战翼皮肤
func (pwdm *PlayerWingDataManager) IfWingSkinExist(wingId int32) (*PlayerWingOtherObject, bool) {
	wingTemplate := wing.GetWingService().GetWing(int(wingId))
	if wingTemplate == nil {
		return nil, false
	}
	typ := wingTemplate.GetTyp()
	playerOtherMap, exist := pwdm.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	wingOtherObj, exist := playerOtherMap[wingId]
	if !exist {
		return nil, false
	}
	return wingOtherObj, true
}

//是否能升星
func (pwdm *PlayerWingDataManager) IfCanUpStar(wingId int32) (*PlayerWingOtherObject, bool) {
	wingTemplate := wing.GetWingService().GetWing(int(wingId))
	if wingTemplate == nil {
		return nil, false
	}

	wingOtherObj, flag := pwdm.IfWingSkinExist(wingId)
	if !flag {
		return nil, false
	}

	if wingTemplate.WingUpstarBeginId == 0 {
		return nil, false
	}

	level := wingOtherObj.Level
	if level <= 0 {
		return wingOtherObj, true
	}
	nextTo := wingTemplate.GetWingUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return wingOtherObj, true
	}
	return nil, false
}

//战翼皮肤升星
func (pwdm *PlayerWingDataManager) Upstar(wingId int32, pro int32, sucess bool) bool {
	obj, flag := pwdm.IfCanUpStar(wingId)
	if !flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		to := wing.GetWingService().GetWing(int(wingId))
		if to == nil {
			return false
		}
		wingUpstarTemplate := to.GetWingUpstarByLevel(obj.Level + 1)
		if wingUpstarTemplate == nil {
			return false
		}
		obj.Level += 1
		obj.UpNum = 0
		obj.UpPro = pro
	} else {
		obj.UpNum += 1
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return true
}

func (pwdm *PlayerWingDataManager) ToWingInfo() *wingcommon.WingInfo {
	info := &wingcommon.WingInfo{
		AdvanceId:   int32(pwdm.PlayerWingObject.AdvanceId),
		WingId:      pwdm.PlayerWingObject.WingId,
		UnrealLevel: pwdm.PlayerWingObject.UnrealLevel,
		UnrealPro:   pwdm.PlayerWingObject.UnrealPro,
	}
	for _, typM := range pwdm.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &wingcommon.WingSkinInfo{
				WingId: otherObj.WingId,
				Level:  otherObj.Level,
				UpPro:  otherObj.UpPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func (pwdm *PlayerWingDataManager) ToFeatherInfo() *wingcommon.FeatherInfo {
	info := &wingcommon.FeatherInfo{
		FeatherId: pwdm.PlayerWingObject.FeatherId,
		Progress:  pwdm.PlayerWingObject.FeatherPro,
	}
	return info
}

func (pwdm *PlayerWingDataManager) IfFullAdvanced() (flag bool) {
	if pwdm.PlayerWingObject.AdvanceId == 0 {
		return
	}
	wingTemplate := wing.GetWingService().GetWingNumber(int32(pwdm.PlayerWingObject.AdvanceId))
	if wingTemplate == nil {
		return
	}
	if wingTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 战翼进阶
func (pwdm *PlayerWingDataManager) GmSetWingAdvanced(advancedId int) {
	pwdm.PlayerWingObject.AdvanceId = advancedId
	pwdm.PlayerWingObject.TimesNum = int32(0)
	pwdm.PlayerWingObject.Bless = int32(0)
	pwdm.PlayerWingObject.BlessTime = int64(0)
	pwdm.PlayerWingObject.WingId = int32(0)
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	gameevent.Emit(wingeventtypes.EventTypeWingAdvanced, pwdm.p, pwdm.PlayerWingObject.AdvanceId)
	return
}

//仅gm使用 战翼幻化
func (pwdm *PlayerWingDataManager) GmSetWingUnreal(wingId int) {
	now := global.GetGame().GetTimeService().Now()
	if !pwdm.IsUnrealed(wingId) {
		pwdm.PlayerWingObject.UnrealList = append(pwdm.PlayerWingObject.UnrealList, wingId)
		sort.Ints(pwdm.PlayerWingObject.UnrealList)
		pwdm.PlayerWingObject.UpdateTime = now
		pwdm.PlayerWingObject.SetModified()
	}

	pwdm.PlayerWingObject.WingId = int32(wingId)
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
	gameevent.Emit(wingeventtypes.EventTypeWingUnrealActivate, pwdm.p, wingId)
}

//仅gm使用 战翼食幻化丹
func (pwdm *PlayerWingDataManager) GmSetWingUnrealDanLevel(level int32) {
	pwdm.PlayerWingObject.UnrealLevel = level
	pwdm.PlayerWingObject.UnrealNum = 0
	pwdm.PlayerWingObject.UnrealPro = 0

	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
}

//仅gm使用 护体仙羽进阶
func (pwdm *PlayerWingDataManager) GmSetFeatherId(featherId int32) {
	pwdm.PlayerWingObject.FeatherId = featherId
	pwdm.PlayerWingObject.FeatherNum = int32(0)
	pwdm.PlayerWingObject.FeatherPro = int32(0)
	now := global.GetGame().GetTimeService().Now()
	pwdm.PlayerWingObject.UpdateTime = now
	pwdm.PlayerWingObject.SetModified()
	gameevent.Emit(wingeventtypes.EventTypeWingChanged, pwdm.p, nil)
}

func CreatePlayerWingDataManager(p player.Player) player.PlayerDataManager {
	pwdm := &PlayerWingDataManager{}
	pwdm.p = p
	pwdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return pwdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerWingDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerWingDataManager))
}
