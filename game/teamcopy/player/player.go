package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
	"fgame/fgame/game/teamcopy/dao"
	teamcopyeventtypes "fgame/fgame/game/teamcopy/event/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家组队副本管理器
type PlayerTeamCopyDataManager struct {
	p player.Player
	//玩家组队副本对象
	teamCopyMap map[teamtypes.TeamPurposeType]*PlayerTeamCopyObject
}

func (ptm *PlayerTeamCopyDataManager) Player() player.Player {
	return ptm.p
}

//加载
func (ptm *PlayerTeamCopyDataManager) Load() (err error) {
	ptm.teamCopyMap = make(map[teamtypes.TeamPurposeType]*PlayerTeamCopyObject)
	//加载玩家组队副本信息
	teamCopyList, err := dao.GetTeamCopyDao().GetTeamCopyList(ptm.p.GetId())
	if err != nil {
		return
	}
	//组队副本信息
	for _, teamCopyEntity := range teamCopyList {
		pto := NewPlayerTeamCopyObject(ptm.p)
		pto.FromEntity(teamCopyEntity)
		purpose := teamtypes.TeamPurposeType(teamCopyEntity.PurPose)
		ptm.teamCopyMap[purpose] = pto
	}
	return nil
}

//加载后
func (ptm *PlayerTeamCopyDataManager) AfterLoad() (err error) {
	ptm.refreshAll()
	return nil
}

func (ptm *PlayerTeamCopyDataManager) newTeamCopyObj(purpose teamtypes.TeamPurposeType) (obj *PlayerTeamCopyObject) {
	obj, ok := ptm.teamCopyMap[purpose]
	if ok {
		return
	}
	obj = NewPlayerTeamCopyObject(ptm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.purpose = purpose
	obj.num = 0
	obj.rewTime = now
	obj.createTime = now
	ptm.teamCopyMap[purpose] = obj
	return
}

func (ptm *PlayerTeamCopyDataManager) refreshAll() (err error) {
	for purpose, _ := range ptm.teamCopyMap {
		ptm.refresh(purpose)
	}
	return
}

func (ptm *PlayerTeamCopyDataManager) refresh(purpose teamtypes.TeamPurposeType) (err error) {
	now := global.GetGame().GetTimeService().Now()
	teamCopyObj, ok := ptm.teamCopyMap[purpose]
	if !ok {
		return
	}
	lastTime := teamCopyObj.rewTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return err
	}
	if !flag {
		teamCopyObj.num = 0
		teamCopyObj.rewTime = now
		teamCopyObj.updateTime = now
		teamCopyObj.SetModified()
	}
	return
}

//心跳
func (ptm *PlayerTeamCopyDataManager) Heartbeat() {
}

func (ptm *PlayerTeamCopyDataManager) GetTeamCopyMap() map[teamtypes.TeamPurposeType]*PlayerTeamCopyObject {
	ptm.refreshAll()
	return ptm.teamCopyMap
}

//完成
func (ptm *PlayerTeamCopyDataManager) FinishPurpose(purpose teamtypes.TeamPurposeType, sucess bool) (obj *PlayerTeamCopyObject, isRew bool) {
	obj, ok := ptm.teamCopyMap[purpose]
	if !ok {
		obj = ptm.newTeamCopyObj(purpose)
	} else {
		ptm.refresh(purpose)
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		teamCopyTemplate := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(purpose)
		if teamCopyTemplate == nil {
			return
		}
		if obj.num >= teamCopyTemplate.RewardNumber {
			return
		}
		obj.num++
		obj.rewTime = now
		obj.SetModified()
		isRew = true
		gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyFinishSucess, ptm.p, purpose)
	}
	return
}

//仅gm 使用
func (ptm *PlayerTeamCopyDataManager) GmClearNum() {
	now := global.GetGame().GetTimeService().Now()
	for _, teamCopyObj := range ptm.teamCopyMap {
		teamCopyObj.num = 0
		teamCopyObj.rewTime = now
		teamCopyObj.updateTime = now
		teamCopyObj.SetModified()
	}
}

func CreatePlayerTeamCopyDataManager(p player.Player) player.PlayerDataManager {
	ptm := &PlayerTeamCopyDataManager{}
	ptm.p = p
	return ptm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTeamCopyDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTeamCopyDataManager))
}
