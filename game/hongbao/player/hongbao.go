package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/hongbao/dao"
	hongbaoentity "fgame/fgame/game/hongbao/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	viplogic "fgame/fgame/game/vip/logic"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//玩家红包对象
type PlayerHongBaoObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	SnatchCount int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerHongBaoObject(pl player.Player) *PlayerHongBaoObject {
	pto := &PlayerHongBaoObject{
		player: pl,
	}
	return pto
}

func convertPlayerHongBaoObjectToEntity(pewo *PlayerHongBaoObject) (*hongbaoentity.PlayerHongBaoEntity, error) {
	e := &hongbaoentity.PlayerHongBaoEntity{
		Id:          pewo.Id,
		PlayerId:    pewo.PlayerId,
		SnatchCount: pewo.SnatchCount,
		UpdateTime:  pewo.UpdateTime,
		CreateTime:  pewo.CreateTime,
		DeleteTime:  pewo.DeleteTime,
	}
	return e, nil
}

func (pewo *PlayerHongBaoObject) GetPlayerId() int64 {
	return pewo.PlayerId
}

func (pewo *PlayerHongBaoObject) GetDBId() int64 {
	return pewo.Id
}

func (pewo *PlayerHongBaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerHongBaoObjectToEntity(pewo)
	return e, err
}

func (pewo *PlayerHongBaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*hongbaoentity.PlayerHongBaoEntity)

	pewo.Id = pse.Id
	pewo.PlayerId = pse.PlayerId
	pewo.SnatchCount = pse.SnatchCount
	pewo.UpdateTime = pse.UpdateTime
	pewo.CreateTime = pse.CreateTime
	pewo.DeleteTime = pse.DeleteTime
	return nil
}

func (pewo *PlayerHongBaoObject) SetModified() {
	e, err := pewo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerHongBao"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pewo.player.AddChangedObject(obj)
	return
}

//玩家红包管理器
type PlayerHongBaoDataManager struct {
	p player.Player
	//玩家膜拜
	playerHongBaoObject *PlayerHongBaoObject
}

func (pedm *PlayerHongBaoDataManager) Player() player.Player {
	return pedm.p
}

//加载
func (pedm *PlayerHongBaoDataManager) Load() (err error) {
	//加载玩家红包信息
	entity, err := dao.GetHongBaoDao().GetPlayerHongBaoEntity(pedm.p.GetId())
	if err != nil {
		return
	}
	if entity == nil {
		pedm.initPlayerHongBaoObject()
	} else {
		pedm.playerHongBaoObject = NewPlayerHongBaoObject(pedm.p)
		pedm.playerHongBaoObject.FromEntity(entity)
	}

	return nil
}

//第一次初始化
func (pedm *PlayerHongBaoDataManager) initPlayerHongBaoObject() {
	pewo := NewPlayerHongBaoObject(pedm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pewo.Id = id
	//生成id
	pewo.PlayerId = pedm.p.GetId()
	pewo.SnatchCount = 0
	pewo.CreateTime = now
	pedm.playerHongBaoObject = pewo
	pewo.SetModified()
}

//加载后
func (pedm *PlayerHongBaoDataManager) AfterLoad() (err error) {
	//刷新抢红包次数
	err = pedm.refreshSnatchCount()
	if err != nil {
		return err
	}
	return nil
}

//刷新抢红包次数
func (pedm *PlayerHongBaoDataManager) refreshSnatchCount() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := pedm.playerHongBaoObject.UpdateTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pedm.playerHongBaoObject.SnatchCount = 0
			pedm.playerHongBaoObject.UpdateTime = now
			pedm.playerHongBaoObject.SetModified()
		}
	}
	return nil
}

//心跳
func (pedm *PlayerHongBaoDataManager) Heartbeat() {
}

//获取抢红包次数
func (pedm *PlayerHongBaoDataManager) GetSnatchCount() int32 {
	pedm.refreshSnatchCount()
	return pedm.getSnatchCount()
}

func (pedm *PlayerHongBaoDataManager) getSnatchCount() int32 {
	return pedm.playerHongBaoObject.SnatchCount
}

//抢红包次数是否到达上限
func (pedm *PlayerHongBaoDataManager) IsSnatchCountReachLimit() bool {
	num := pedm.GetSnatchCount()
	// 可抢次数
	vipCount := viplogic.GetHongBaoSnatchCount(pedm.p)
	maxNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeHongBaoDayCount) + vipCount
	if num >= maxNum {
		return true
	}
	return false
}

//增加抢红包次数
func (pedm *PlayerHongBaoDataManager) AddSnatchCount() int32 {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerHongBaoObject.SnatchCount++
	pedm.playerHongBaoObject.UpdateTime = now
	pedm.playerHongBaoObject.SetModified()
	return pedm.playerHongBaoObject.SnatchCount
}

//仅gm使用
func (pedm *PlayerHongBaoDataManager) GMSetSnatchCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerHongBaoObject.SnatchCount = count
	pedm.playerHongBaoObject.UpdateTime = now
	pedm.playerHongBaoObject.SetModified()
	return
}

func CreatePlayerHongBaoDataManager(p player.Player) player.PlayerDataManager {
	pedm := &PlayerHongBaoDataManager{}
	pedm.p = p
	return pedm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerHongBaoDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerHongBaoDataManager))
}
