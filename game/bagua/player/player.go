package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	baguacommon "fgame/fgame/game/bagua/common"
	"fgame/fgame/game/bagua/dao"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	baguatemplate "fgame/fgame/game/bagua/template"
	gameevent "fgame/fgame/game/event"
)

//玩家八卦秘境管理器
type PlayerBaGuaDataManager struct {
	p player.Player
	//玩家八卦秘境对象
	playerBaGuaObject *PlayerBaGuaObject
}

func (pdm *PlayerBaGuaDataManager) Player() player.Player {
	return pdm.p
}

//加载
func (pdm *PlayerBaGuaDataManager) Load() (err error) {
	//加载玩家八卦秘境信息
	baGuaEntity, err := dao.GetBaGuaDao().GetBaGuaEntity(pdm.p.GetId())
	if err != nil {
		return
	}
	if baGuaEntity == nil {
		pdm.initPlayerTianJieTaObject()
	} else {
		pdm.playerBaGuaObject = NewPlayerBaGuaObject(pdm.p)
		pdm.playerBaGuaObject.FromEntity(baGuaEntity)
	}

	return nil
}

//第一次初始化
func (pdm *PlayerBaGuaDataManager) initPlayerTianJieTaObject() {
	po := NewPlayerBaGuaObject(pdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	po.id = id
	//生成id
	po.playerId = pdm.p.GetId()
	po.level = int32(0)
	po.isBuChang = int32(1)
	po.inviteTime = int64(0)
	po.createTime = now
	pdm.playerBaGuaObject = po
	po.SetModified()
}

//加载后
func (pdm *PlayerBaGuaDataManager) AfterLoad() (err error) {
	return nil
}

//八卦秘境等级
func (pdm *PlayerBaGuaDataManager) GetLevel() int32 {
	return pdm.playerBaGuaObject.level
}

//是否八卦秘境补偿
func (pdm *PlayerBaGuaDataManager) IsBuChang() bool {
	return pdm.playerBaGuaObject.isBuChang == 0
}

//设置以后不予发补偿奖励
func (pdm *PlayerBaGuaDataManager) SetIsBuChang() {
	now := global.GetGame().GetTimeService().Now()
	pdm.playerBaGuaObject.isBuChang = 1
	pdm.playerBaGuaObject.updateTime = now

	pdm.playerBaGuaObject.SetModified()
}

//心跳
func (pdm *PlayerBaGuaDataManager) Heartbeat() {
}

//是否邀请过于频繁
func (pdm *PlayerBaGuaDataManager) InviteFrequent() bool {
	if pdm.playerBaGuaObject.inviteTime == 0 {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	cdTime := baguatemplate.GetBaGuaTemplateService().GetInvitePairCdTime()
	if now-pdm.playerBaGuaObject.inviteTime < cdTime {
		return true
	}
	return false
}

func (pdm *PlayerBaGuaDataManager) InviteTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	pdm.playerBaGuaObject.inviteTime = now
	pdm.playerBaGuaObject.updateTime = now
	pdm.playerBaGuaObject.SetModified()
	return now
}

func (pdm *PlayerBaGuaDataManager) GetInviteLeftTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - pdm.playerBaGuaObject.inviteTime
	cdTime := baguatemplate.GetBaGuaTemplateService().GetInvitePairCdTime()
	return cdTime - elapse
}

//是否满级
func (pdm *PlayerBaGuaDataManager) IfFullLevel() bool {
	curLevel := pdm.playerBaGuaObject.level
	if curLevel == 0 {
		return false
	}
	to := baguatemplate.GetBaGuaTemplateService().GetBaGuaTemplateByLevel(curLevel)
	if to.NextId == 0 {
		return true
	}
	return false
}

//提升八卦秘境等级
func (pdm *PlayerBaGuaDataManager) UpgradeLevel() bool {
	flag := pdm.IfFullLevel()
	if flag {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pdm.playerBaGuaObject.level += 1
	pdm.playerBaGuaObject.updateTime = now
	pdm.playerBaGuaObject.SetModified()
	return true
}

func (pdm *PlayerBaGuaDataManager) ToBaGuaInfo() *baguacommon.BaGuaInfo {
	info := &baguacommon.BaGuaInfo{
		Level: pdm.playerBaGuaObject.level,
	}
	return info
}

// GM 补偿标志置零
func (pdm *PlayerBaGuaDataManager) GmSetBuChang() {
	pdm.playerBaGuaObject.isBuChang = 0
	now := global.GetGame().GetTimeService().Now()
	pdm.playerBaGuaObject.updateTime = now
	pdm.playerBaGuaObject.SetModified()
}

//gm使用
func (pdm *PlayerBaGuaDataManager) GmSetLevel(level int32) {
	pdm.playerBaGuaObject.level = level
	now := global.GetGame().GetTimeService().Now()
	pdm.playerBaGuaObject.updateTime = now
	pdm.playerBaGuaObject.SetModified()

	//发送事件
	gameevent.Emit(baguaeventtypes.EventTypeBaGuaResult, pdm.p, true)
	return
}

func CreatePlayerBaGuaDataManager(p player.Player) player.PlayerDataManager {
	pdm := &PlayerBaGuaDataManager{}
	pdm.p = p
	return pdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerBaGuaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerBaGuaDataManager))
}
