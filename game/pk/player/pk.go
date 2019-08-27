package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fgame/fgame/game/pk/dao"
	pkentity "fgame/fgame/game/pk/entity"
	pktypes "fgame/fgame/game/pk/types"

	"github.com/pkg/errors"
)

//pk对象
type PlayerPkObject struct {
	Id           int64
	player       player.Player
	PkValue      int32
	KillNum      int32
	LastKillTime int64
	OnlineTime   int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerPkObject(pl player.Player) *PlayerPkObject {
	pmo := &PlayerPkObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerPkObjectToEntity(o *PlayerPkObject) (*pkentity.PlayerPkEntity, error) {
	e := &pkentity.PlayerPkEntity{
		Id:           o.Id,
		PlayerId:     o.player.GetId(),
		PkValue:      o.PkValue,
		KillNum:      o.KillNum,
		LastKillTime: o.LastKillTime,
		OnlineTime:   o.OnlineTime,
		UpdateTime:   o.UpdateTime,
		CreateTime:   o.CreateTime,
		DeleteTime:   o.DeleteTime,
	}
	return e, nil
}

func (ppo *PlayerPkObject) GetPlayerId() int64 {
	return ppo.player.GetId()
}

func (ppo *PlayerPkObject) GetDBId() int64 {
	return ppo.Id
}

func (ppo *PlayerPkObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerPkObjectToEntity(ppo)
	return e, err
}

func (ppo *PlayerPkObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*pkentity.PlayerPkEntity)
	ppo.Id = pse.Id
	ppo.PkValue = pse.PkValue
	ppo.KillNum = pse.KillNum
	ppo.LastKillTime = pse.LastKillTime
	ppo.OnlineTime = pse.OnlineTime
	ppo.UpdateTime = pse.UpdateTime
	ppo.CreateTime = pse.CreateTime
	ppo.DeleteTime = pse.DeleteTime
	return nil
}

func (ppo *PlayerPkObject) SetModified() {
	e, err := ppo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Pk"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	ppo.player.AddChangedObject(obj)
	return
}

//玩家pk管理器
type PlayerPkDataManager struct {
	p player.Player
	//玩家pk对象
	playerPkObject *PlayerPkObject
	//玩家pk状态
	pkState pktypes.PkState
	//阵营id
	pkCamp pktypes.PkCamp
	//心跳处理器
	// heartbeatRunner heartbeat.HeartbeatTaskRunner
	//登陆时间
	loginTime int64
}

func (ppdm *PlayerPkDataManager) Player() player.Player {
	return ppdm.p
}

//加载
func (ppdm *PlayerPkDataManager) Load() (err error) {

	//加载玩家pk信息
	pkEntity, err := dao.GetPKDao().GetPKEntity(ppdm.p.GetId())
	if err != nil {
		return
	}
	if pkEntity == nil {
		ppdm.initPlayerPkObject()

	} else {
		ppdm.playerPkObject = NewPlayerPkObject(ppdm.p)
		ppdm.playerPkObject.FromEntity(pkEntity)
	}
	ppdm.pkState = pktypes.PkStatePeach
	ppdm.pkCamp = pktypes.PkCommonCampDefault
	return nil
}

//第一次初始化
func (ppdm *PlayerPkDataManager) initPlayerPkObject() {
	ppo := NewPlayerPkObject(ppdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ppo.Id = id
	ppo.CreateTime = now
	ppo.SetModified()

	ppdm.playerPkObject = ppo
	return
}

//加载后
func (ppdm *PlayerPkDataManager) AfterLoad() (err error) {
	// now := global.GetGame().GetTimeService().Now()
	// ppdm.loginTime = now
	// _, err = ppdm.refresh(false)
	// if err != nil {
	// 	return
	// }
	// if !flag {
	// 	err = ppdm.pkValueChanged()
	// 	if err != nil {
	// 		return
	// 	}
	// }
	//添加定时任务
	// ppdm.heartbeatRunner.AddTask(newPkTask(ppdm.p))
	return nil
}

const (
	clearTime = common.MINUTE * 3
)

//刷新
// func (ppdm *PlayerPkDataManager) refresh(nofity bool) (flag bool, err error) {
// 	if ppdm.GetPkRedState() == pktypes.PkRedStateInit {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	elapseTime := now - ppdm.loginTime
// 	elapseTime += ppdm.playerPkObject.OnlineTime
// 	clearNum := int32(elapseTime / int64(clearTime))
// 	if clearNum <= 0 {
// 		ppdm.playerPkObject.OnlineTime = elapseTime
// 		ppdm.loginTime = now
// 		ppdm.playerPkObject.UpdateTime = now
// 		err = ppdm.playerPkObject.SetModified()
// 		if err != nil {
// 			return
// 		}
// 		return
// 	}
// 	if clearNum > ppdm.playerPkObject.PkValue {
// 		ppdm.playerPkObject.PkValue = 0
// 		ppdm.playerPkObject.OnlineTime = 0
// 		ppdm.playerPkObject.UpdateTime = now
// 		err = ppdm.playerPkObject.SetModified()
// 		if err != nil {
// 			return
// 		}

// 	} else {
// 		ppdm.playerPkObject.PkValue -= clearNum
// 		ppdm.playerPkObject.OnlineTime = elapseTime - (int64(clearTime) * int64(clearNum))
// 		ppdm.loginTime = now
// 		ppdm.playerPkObject.UpdateTime = now
// 		err = ppdm.playerPkObject.SetModified()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	flag = true
// 	if nofity {
// 		err = ppdm.pkValueChanged()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return
// }

func (ppdm *PlayerPkDataManager) GetPkState() pktypes.PkState {
	return ppdm.pkState
}

func (ppdm *PlayerPkDataManager) GetPkCamp() pktypes.PkCamp {
	return ppdm.pkCamp
}

func (ppdm *PlayerPkDataManager) GetPkRedState() pktypes.PkRedState {
	return pktypes.PkRedStateFromValue(ppdm.playerPkObject.PkValue)
}

func (ppdm *PlayerPkDataManager) GetPkValue() int32 {
	return ppdm.playerPkObject.PkValue
}

func (ppdm *PlayerPkDataManager) GetOnlineTime() int64 {
	return ppdm.playerPkObject.OnlineTime
}

func (ppdm *PlayerPkDataManager) GetKillNum() int32 {
	return ppdm.playerPkObject.KillNum
}

func (ppdm *PlayerPkDataManager) GetLastKillTime() int64 {
	return ppdm.playerPkObject.LastKillTime
}

func (ppdm *PlayerPkDataManager) GetLoginTime() int64 {
	return ppdm.loginTime
}

func (ppdm *PlayerPkDataManager) Save() {
	ppdm.playerPkObject.PkValue = ppdm.p.GetPkValue()
	ppdm.playerPkObject.KillNum = ppdm.p.GetKillNum()
	ppdm.playerPkObject.LastKillTime = ppdm.p.GetLastKillTime()
	ppdm.playerPkObject.OnlineTime = ppdm.p.GetPkOnlineTime()
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPkObject.UpdateTime = now
	ppdm.playerPkObject.SetModified()

	
}

//切换pk模式
// func (ppdm *PlayerPkDataManager) SwitchPkState(pkState pktypes.PkState, camp pktypes.PkCamp) (flag bool) {
// 	if !pkState.Valid() {
// 		panic(fmt.Errorf("pk:state [%d] invalid", pkState))
// 	}
// 	if ppdm.pkState == pkState && ppdm.pkCamp == camp {
// 		return
// 	}
// 	ppdm.pkState = pkState
// 	ppdm.pkCamp = camp
// 	//发送事件pk状态改变
// 	gameevent.Emit(pkeventtypes.EventTypePkStateSwitch, ppdm.p, nil)
// 	flag = true
// 	return
// }

// const (
// 	maxPKValue = 80
// )

//击杀
// func (ppdm *PlayerPkDataManager) Kill(white bool) (err error) {
// 	if white {
// 		firstKill := ppdm.playerPkObject.PkValue == 0
// 		if ppdm.playerPkObject.PkValue < maxPKValue {
// 			ppdm.playerPkObject.PkValue += 1
// 			now := global.GetGame().GetTimeService().Now()
// 			ppdm.playerPkObject.KillNum += 1
// 			ppdm.playerPkObject.LastKillTime = now
// 			ppdm.playerPkObject.UpdateTime = now
// 			if firstKill {
// 				ppdm.playerPkObject.OnlineTime = 0
// 				ppdm.loginTime = now
// 			}
// 			err = ppdm.playerPkObject.SetModified()
// 			if err != nil {
// 				return
// 			}
// 			err = ppdm.pkValueChanged()
// 			if err != nil {
// 				return
// 			}
// 		}
// 	}
// 	//TODO 添加杀人数
// 	//TODO 发送事件
// 	return
// }

// func (ppdm *PlayerPkDataManager) pkValueChanged() (err error) {
// 	//发送事件pk值改变
// 	gameevent.Emit(pkeventtypes.EventTypePkValueChanged, ppdm.p, nil)

// 	return
// }

//心跳
func (ppdm *PlayerPkDataManager) Heartbeat() {
	// ppdm.heartbeatRunner.Heartbeat()
}

func CreatePlayerPkDataManager(p player.Player) player.PlayerDataManager {
	ppdm := &PlayerPkDataManager{}
	ppdm.p = p
	// ppdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return ppdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerPkDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerPkDataManager))
}
