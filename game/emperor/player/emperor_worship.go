package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/emperor/dao"
	emperorentity "fgame/fgame/game/emperor/entity"
	emperorTemplate "fgame/fgame/game/emperor/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//玩家膜拜对象
type PlayerEmperorWorshipObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Num        int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerEmperorWorshipObject(pl player.Player) *PlayerEmperorWorshipObject {
	pto := &PlayerEmperorWorshipObject{
		player: pl,
	}
	return pto
}

func convertEmperorWorshipObjectToEntity(pewo *PlayerEmperorWorshipObject) (*emperorentity.PlayerEmperorWorshipEntity, error) {
	e := &emperorentity.PlayerEmperorWorshipEntity{
		Id:         pewo.Id,
		PlayerId:   pewo.PlayerId,
		Num:        pewo.Num,
		LastTime:   pewo.LastTime,
		UpdateTime: pewo.UpdateTime,
		CreateTime: pewo.CreateTime,
		DeleteTime: pewo.DeleteTime,
	}
	return e, nil
}

func (pewo *PlayerEmperorWorshipObject) GetPlayerId() int64 {
	return pewo.PlayerId
}

func (pewo *PlayerEmperorWorshipObject) GetDBId() int64 {
	return pewo.Id
}

func (pewo *PlayerEmperorWorshipObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertEmperorWorshipObjectToEntity(pewo)
	return e, err
}

func (pewo *PlayerEmperorWorshipObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*emperorentity.PlayerEmperorWorshipEntity)

	pewo.Id = pse.Id
	pewo.PlayerId = pse.PlayerId
	pewo.Num = pse.Num
	pewo.LastTime = pse.LastTime
	pewo.UpdateTime = pse.UpdateTime
	pewo.CreateTime = pse.CreateTime
	pewo.DeleteTime = pse.DeleteTime
	return nil
}

func (pewo *PlayerEmperorWorshipObject) SetModified() {
	e, err := pewo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "EmperorWorship"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pewo.player.AddChangedObject(obj)
	return
}

//玩家膜拜管理器
type PlayerEmperorDataManager struct {
	p player.Player
	//玩家膜拜
	playerEmperorWorshipObject *PlayerEmperorWorshipObject
}

func (pedm *PlayerEmperorDataManager) Player() player.Player {
	return pedm.p
}

//加载
func (pedm *PlayerEmperorDataManager) Load() (err error) {
	//加载玩家膜拜信息
	worshipEntity, err := dao.GetEmperorDao().GetEmperorWorshipEntity(pedm.p.GetId())
	if err != nil {
		return
	}
	if worshipEntity == nil {
		pedm.initPlayerEmperorWorshipObject()
	} else {
		pedm.playerEmperorWorshipObject = NewPlayerEmperorWorshipObject(pedm.p)
		pedm.playerEmperorWorshipObject.FromEntity(worshipEntity)
	}

	return nil
}

//第一次初始化
func (pedm *PlayerEmperorDataManager) initPlayerEmperorWorshipObject() {
	pewo := NewPlayerEmperorWorshipObject(pedm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pewo.Id = id
	//生成id
	pewo.PlayerId = pedm.p.GetId()
	pewo.Num = 0
	pewo.LastTime = 0
	pewo.CreateTime = now
	pedm.playerEmperorWorshipObject = pewo
	pewo.SetModified()
}

//加载后
func (pedm *PlayerEmperorDataManager) AfterLoad() (err error) {
	//刷新膜拜次数
	err = pedm.refreshWorshipNum()
	if err != nil {
		return err
	}
	return nil
}

//刷新膜拜次数
func (pedm *PlayerEmperorDataManager) refreshWorshipNum() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := pedm.playerEmperorWorshipObject.LastTime
	if lastTime != 0 {
		//flag, err := timeutils.IsSameDay(lastTime, now)
		flag, err := timeutils.IsSameFive(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pedm.playerEmperorWorshipObject.Num = 0
			pedm.playerEmperorWorshipObject.LastTime = 0
			pedm.playerEmperorWorshipObject.UpdateTime = now
			pedm.playerEmperorWorshipObject.SetModified()
		}
	}
	return nil
}

//心跳
func (pedm *PlayerEmperorDataManager) Heartbeat() {
}

//获取膜拜次数
func (pedm *PlayerEmperorDataManager) GetWorshipNum() int32 {
	pedm.refreshWorshipNum()
	return pedm.getWorshipNum()
}

func (pedm *PlayerEmperorDataManager) getWorshipNum() int32 {
	return pedm.playerEmperorWorshipObject.Num
}

//膜拜次数是否到达上限
func (pedm *PlayerEmperorDataManager) IfWorshipReachLimit() bool {
	num := pedm.GetWorshipNum()
	maxNum := emperorTemplate.GetEmperorTemplateService().GetEmperorWorshipNum()
	if num >= maxNum {
		return true
	}
	return false
}

//增加膜拜次数
func (pedm *PlayerEmperorDataManager) AddWorshipNum() (num int32) {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerEmperorWorshipObject.Num++
	pedm.playerEmperorWorshipObject.LastTime = now
	pedm.playerEmperorWorshipObject.UpdateTime = now
	pedm.playerEmperorWorshipObject.SetModified()
	return pedm.playerEmperorWorshipObject.Num
}

//仅gm使用
func (pedm *PlayerEmperorDataManager) GMClearWorshipNum() {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerEmperorWorshipObject.Num = 0
	pedm.playerEmperorWorshipObject.UpdateTime = now
	pedm.playerEmperorWorshipObject.SetModified()
	return
}

func CreatePlayerEmperorDataManager(p player.Player) player.PlayerDataManager {
	pedm := &PlayerEmperorDataManager{}
	pedm.p = p
	return pedm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerEmperorDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerEmperorDataManager))
}
