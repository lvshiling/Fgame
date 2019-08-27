package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/chat/dao"
	chatentity "fgame/fgame/game/chat/entity"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//玩家聊天对象
type PlayerChatObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	ChatCount  int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerChatObject(pl player.Player) *PlayerChatObject {
	pto := &PlayerChatObject{
		player: pl,
	}
	return pto
}

func convertPlayerChatObjectToEntity(pewo *PlayerChatObject) (*chatentity.PlayerChatEntity, error) {
	e := &chatentity.PlayerChatEntity{
		Id:         pewo.Id,
		PlayerId:   pewo.PlayerId,
		ChatCount:  pewo.ChatCount,
		UpdateTime: pewo.UpdateTime,
		CreateTime: pewo.CreateTime,
		DeleteTime: pewo.DeleteTime,
	}
	return e, nil
}

func (pewo *PlayerChatObject) GetPlayerId() int64 {
	return pewo.PlayerId
}

func (pewo *PlayerChatObject) GetDBId() int64 {
	return pewo.Id
}

func (pewo *PlayerChatObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerChatObjectToEntity(pewo)
	return e, err
}

func (pewo *PlayerChatObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chatentity.PlayerChatEntity)

	pewo.Id = pse.Id
	pewo.PlayerId = pse.PlayerId
	pewo.ChatCount = pse.ChatCount
	pewo.UpdateTime = pse.UpdateTime
	pewo.CreateTime = pse.CreateTime
	pewo.DeleteTime = pse.DeleteTime
	return nil
}

func (pewo *PlayerChatObject) SetModified() {
	e, err := pewo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PlayerChat"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pewo.player.AddChangedObject(obj)
	return
}

//玩家聊天管理器
type PlayerChatDataManager struct {
	p player.Player
	//玩家膜拜
	playerChatObject *PlayerChatObject
}

func (pedm *PlayerChatDataManager) Player() player.Player {
	return pedm.p
}

//加载
func (pedm *PlayerChatDataManager) Load() (err error) {
	//加载玩家红包信息
	entity, err := dao.GetChatDao().GetPlayerChatEntity(pedm.p.GetId())
	if err != nil {
		return
	}
	if entity == nil {
		pedm.initPlayerChatObject()
	} else {
		pedm.playerChatObject = NewPlayerChatObject(pedm.p)
		pedm.playerChatObject.FromEntity(entity)
	}

	return nil
}

//第一次初始化
func (pedm *PlayerChatDataManager) initPlayerChatObject() {
	pewo := NewPlayerChatObject(pedm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pewo.Id = id
	//生成id
	pewo.PlayerId = pedm.p.GetId()
	pewo.ChatCount = 0
	pewo.CreateTime = now
	pedm.playerChatObject = pewo
	pewo.SetModified()
}

//加载后
func (pedm *PlayerChatDataManager) AfterLoad() (err error) {
	//刷新抢红包次数
	err = pedm.refreshChatCount()
	if err != nil {
		return err
	}
	return nil
}

//刷新抢红包次数
func (pedm *PlayerChatDataManager) refreshChatCount() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := pedm.playerChatObject.UpdateTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pedm.playerChatObject.ChatCount = 0
			pedm.playerChatObject.UpdateTime = now
			pedm.playerChatObject.SetModified()
		}
	}
	return nil
}

//心跳
func (pedm *PlayerChatDataManager) Heartbeat() {
}

//获取抢红包次数
func (pedm *PlayerChatDataManager) GetChatCount() int32 {
	pedm.refreshChatCount()
	return pedm.getChatCount()
}

func (pedm *PlayerChatDataManager) getChatCount() int32 {
	return pedm.playerChatObject.ChatCount
}

//抢红包次数是否到达上限
func (pedm *PlayerChatDataManager) IsChatCountReachLimit() bool {
	num := pedm.GetChatCount()
	maxNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChatAwardDayCount)
	if num >= maxNum {
		return true
	}
	return false
}

//增加抢红包次数
func (pedm *PlayerChatDataManager) AddChatCount() int32 {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerChatObject.ChatCount++
	pedm.playerChatObject.UpdateTime = now
	pedm.playerChatObject.SetModified()
	return pedm.playerChatObject.ChatCount
}

//仅gm使用
func (pedm *PlayerChatDataManager) GMSetChatCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pedm.playerChatObject.ChatCount = count
	pedm.playerChatObject.UpdateTime = now
	pedm.playerChatObject.SetModified()
	return
}

func CreatePlayerChatDataManager(p player.Player) player.PlayerDataManager {
	pedm := &PlayerChatDataManager{}
	pedm.p = p
	return pedm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerChatDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerChatDataManager))
}
