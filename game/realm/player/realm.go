package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/realm/dao"
	realmentity "fgame/fgame/game/realm/entity"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/realm/realm"
	realmtemplate "fgame/fgame/game/realm/template"

	"github.com/pkg/errors"
)

//天劫塔对象
type PlayerTianJieTaObject struct {
	player         player.Player
	Id             int64
	PlayerId       int64
	PlayerName     string
	Level          int32
	UsedTime       int64
	IsCheckReissue int32
	UpdateTime     int64
	CreateTime     int64
	DeleteTime     int64
}

func NewPlayerTianJieTaObject(pl player.Player) *PlayerTianJieTaObject {
	ptjto := &PlayerTianJieTaObject{
		player: pl,
	}
	return ptjto
}

func (ptjto *PlayerTianJieTaObject) GetPlayerId() int64 {
	return ptjto.PlayerId
}

func (ptjto *PlayerTianJieTaObject) GetDBId() int64 {
	return ptjto.Id
}

func (ptjto *PlayerTianJieTaObject) ToEntity() (e storage.Entity, err error) {
	e = &realmentity.PlayerTianJieTaEntity{
		Id:             ptjto.Id,
		PlayerId:       ptjto.PlayerId,
		PlayerName:     ptjto.PlayerName,
		Level:          ptjto.Level,
		UsedTime:       ptjto.UsedTime,
		IsCheckReissue: ptjto.IsCheckReissue,
		UpdateTime:     ptjto.UpdateTime,
		CreateTime:     ptjto.CreateTime,
		DeleteTime:     ptjto.DeleteTime,
	}
	return e, nil
}

func (ptjto *PlayerTianJieTaObject) FromEntity(e storage.Entity) error {
	ptjte, _ := e.(*realmentity.PlayerTianJieTaEntity)

	ptjto.Id = ptjte.Id
	ptjto.PlayerId = ptjte.PlayerId
	ptjto.PlayerName = ptjte.PlayerName
	ptjto.Level = ptjte.Level
	ptjto.UsedTime = ptjte.UsedTime
	ptjto.IsCheckReissue = ptjte.IsCheckReissue
	ptjto.UpdateTime = ptjte.UpdateTime
	ptjto.CreateTime = ptjte.CreateTime
	ptjto.DeleteTime = ptjte.DeleteTime
	return nil
}

func (ptjto *PlayerTianJieTaObject) SetModified() {
	e, err := ptjto.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "TianJieTa"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	ptjto.player.AddChangedObject(obj)
	return
}

//玩家境界管理器
type PlayerRealmDataManager struct {
	p player.Player
	//玩家天劫塔对象
	playerTianJieTaObject *PlayerTianJieTaObject
	inviteTime            int64
}

func (prdm *PlayerRealmDataManager) Player() player.Player {
	return prdm.p
}

//加载
func (prdm *PlayerRealmDataManager) Load() (err error) {
	//加载玩家天劫塔信息
	tianJieTaEntity, err := dao.GetRealmDao().GetTianJieTaEntity(prdm.p.GetId())
	if err != nil {
		return
	}
	if tianJieTaEntity == nil {
		prdm.initPlayerTianJieTaObject()
	} else {
		prdm.playerTianJieTaObject = NewPlayerTianJieTaObject(prdm.p)
		prdm.playerTianJieTaObject.FromEntity(tianJieTaEntity)
	}

	return nil
}

//第一次初始化
func (prdm *PlayerRealmDataManager) initPlayerTianJieTaObject() {
	ptjto := NewPlayerTianJieTaObject(prdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ptjto.Id = id
	//生成id
	ptjto.PlayerId = prdm.p.GetId()
	ptjto.PlayerName = prdm.p.GetName()
	ptjto.Level = int32(0)
	ptjto.UsedTime = int64(0)
	ptjto.IsCheckReissue = int32(1)
	ptjto.CreateTime = now
	prdm.playerTianJieTaObject = ptjto
	ptjto.SetModified()
}

//加载后
func (prdm *PlayerRealmDataManager) AfterLoad() (err error) {
	return nil
}

//天劫塔等级
func (prdm *PlayerRealmDataManager) GetTianJieTaLevel() int32 {
	return prdm.playerTianJieTaObject.Level
}

//心跳
func (prdm *PlayerRealmDataManager) Heartbeat() {
}

//是否邀请过于频繁
func (prdm *PlayerRealmDataManager) InviteFrequent() bool {
	if prdm.inviteTime == 0 {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	cdTime := realmtemplate.GetRealmTemplateService().GetInvitePairCdTime()
	if now-prdm.inviteTime < cdTime {
		return true
	}
	return false
}

func (prdm *PlayerRealmDataManager) InviteTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	prdm.inviteTime = now
	return now
}

//是否满级
func (prdm *PlayerRealmDataManager) IfFullLevel() bool {
	curLevel := prdm.playerTianJieTaObject.Level
	if curLevel == 0 {
		return false
	}
	to := realmtemplate.GetRealmTemplateService().GetTianJieTaTemplateByLevel(curLevel)
	if to.NextId == 0 {
		return true
	}
	return false
}

//提升天劫塔等级
func (prdm *PlayerRealmDataManager) UpgradeTianJieLevel(usedTime int64) bool {
	flag := prdm.IfFullLevel()
	if flag {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	prdm.playerTianJieTaObject.Level += 1
	prdm.playerTianJieTaObject.UsedTime = usedTime
	prdm.playerTianJieTaObject.UpdateTime = now
	prdm.playerTianJieTaObject.SetModified()
	return true
}

//玩家修改名字
func (prdm *PlayerRealmDataManager) PlayerChangeName(pl player.Player) {
	now := global.GetGame().GetTimeService().Now()
	prdm.playerTianJieTaObject.PlayerName = pl.GetName()
	prdm.playerTianJieTaObject.UpdateTime = now
	prdm.playerTianJieTaObject.SetModified()
}

//玩家检测补偿
func (prdm *PlayerRealmDataManager) IsCheckReissue() bool {
	return prdm.playerTianJieTaObject.IsCheckReissue != 0
}

//玩家检测补偿
func (prdm *PlayerRealmDataManager) SetCheckReissue() {
	now := global.GetGame().GetTimeService().Now()
	prdm.playerTianJieTaObject.IsCheckReissue = 1
	prdm.playerTianJieTaObject.UpdateTime = now
	prdm.playerTianJieTaObject.SetModified()
}

//gm使用
func (prdm *PlayerRealmDataManager) GmSetLevel(level int32) {
	prdm.playerTianJieTaObject.Level = level
	now := global.GetGame().GetTimeService().Now()
	prdm.playerTianJieTaObject.UpdateTime = now
	prdm.playerTianJieTaObject.SetModified()

	//发送事件
	gameevent.Emit(realmeventtypes.EventTypeRealmResult, prdm.p, true)

	//刷新天劫塔排名
	realm.GetRealmRankService().RefreshTianJieTaRank(prdm.p.GetId(), prdm.p.GetName(), level, 5000)
	return
}

//gm使用
func (prdm *PlayerRealmDataManager) GmSetCheckReissueOn() {
	now := global.GetGame().GetTimeService().Now()
	prdm.playerTianJieTaObject.IsCheckReissue = 0
	prdm.playerTianJieTaObject.UpdateTime = now
	prdm.playerTianJieTaObject.SetModified()
}

func CreatePlayerRealmDataManager(p player.Player) player.PlayerDataManager {
	prdm := &PlayerRealmDataManager{}
	prdm.p = p
	prdm.inviteTime = 0
	return prdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerRealmDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerRealmDataManager))
}
