package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	livenesstemplate "fgame/fgame/game/liveness/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/dao"
	questentity "fgame/fgame/game/quest/entity"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//活跃度跨5点对象
type PlayerLivenessCrossFiveObject struct {
	player       player.Player
	id           int64
	crossDayTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerLivenessCrossFiveObject(pl player.Player) *PlayerLivenessCrossFiveObject {
	pmo := &PlayerLivenessCrossFiveObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerLivenessCrossFiveObjectToEntity(pqo *PlayerLivenessCrossFiveObject) (e *questentity.PlayerLivenessCrossFiveEntity, err error) {
	e = &questentity.PlayerLivenessCrossFiveEntity{
		Id:            pqo.id,
		PlayerId:      pqo.player.GetId(),
		CrossFiveTime: pqo.crossDayTime,
		UpdateTime:    pqo.updateTime,
		CreateTime:    pqo.createTime,
		DeleteTime:    pqo.deleteTime,
	}
	return
}

func (pqo *PlayerLivenessCrossFiveObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerLivenessCrossFiveObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerLivenessCrossFiveObject) GetCrossFiveTime() int64 {
	return pqo.crossDayTime
}

func (pqo *PlayerLivenessCrossFiveObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerLivenessCrossFiveObjectToEntity(pqo)
	return
}

func (pqo *PlayerLivenessCrossFiveObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerLivenessCrossFiveEntity)

	pqo.id = pqe.Id
	pqo.updateTime = pqe.UpdateTime
	pqo.crossDayTime = pqe.CrossFiveTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerLivenessCrossFiveObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "liveness_cross_five"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}

func (pqo *PlayerLivenessCrossFiveObject) IsCrossFive() (isCrossFive bool, err error) {
	if !pqo.player.IsFuncOpen(funcopentypes.FuncOpenTypeLiveness) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	lastTime := pqo.crossDayTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return false, err
	}
	if !flag {
		pqo.crossDayTime = now
		pqo.updateTime = now
		pqo.SetModified()
		isCrossFive = true
	}
	return
}

//数据库加载
func (pqdm *PlayerQuestDataManager) loadLiveness() (err error) {
	//加载活跃度信息
	livenessEntity, err := dao.GetQuestDao().GetLivenessCrossFive(pqdm.p.GetId())
	if err != nil {
		return
	}
	if livenessEntity == nil {
		pqdm.initPlayerLivenessObject()
	} else {
		pqdm.playerLivenessObject = NewPlayerLivenessCrossFiveObject(pqdm.p)
		pqdm.playerLivenessObject.FromEntity(livenessEntity)
	}
	return
}

//第一次初始化
func (pqdm *PlayerQuestDataManager) initPlayerLivenessObject() {
	ptmo := NewPlayerLivenessCrossFiveObject(pqdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ptmo.id = id
	//生成id
	ptmo.crossDayTime = now
	ptmo.createTime = now
	pqdm.playerLivenessObject = ptmo
	ptmo.SetModified()
}

//活跃度
func (pqdm *PlayerQuestDataManager) afterLoadLivenessQuest() (err error) {
	livenessAllMap := questtemplate.GetQuestTemplateService().GetQuestHuoYueMap()
	flag, err := pqdm.playerLivenessObject.IsCrossFive()
	if err != nil {
		return err
	}
	if !flag {
		return
	}
	for questId, _ := range livenessAllMap {
		quest := pqdm.GetQuestById(questId)
		if quest == nil {
			//判断功能是否开启
			livenessTemplate := livenesstemplate.GetHuoYueTempalteService().GetHuoYueTemplate(questId)
			if livenessTemplate == nil {
				continue
			}
			if !pqdm.p.IsFuncOpen(livenessTemplate.GetFuncOpenTyp()) {
				continue
			}
			pqdm.AddQuest(questId)
			continue
		}
		pqdm.resetLivenessInit(quest)
	}
	return
}

//活跃度任务过天跨5点
func (pqdm *PlayerQuestDataManager) resetLivenessCrossFive() (err error) {
	flag, err := pqdm.playerLivenessObject.IsCrossFive()
	if err != nil {
		return
	}
	if !flag {
		return
	}
	for _, questStateMap := range pqdm.stateQuestMap {
		for questId, quest := range questStateMap {
			questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
			if questTemplate == nil {
				continue
			}
			switch questTemplate.GetQuestType() {
			case questtypes.QuestTypeLiveness:
				break
			default:
				continue
			}
			pqdm.resetLivenessInit(quest)
		}
	}

	gameevent.Emit(questeventtypes.EventTypeQuestLivenessCrossFive, pqdm.p, nil)
	return
}

//重置活跃度状态
func (pqdm *PlayerQuestDataManager) resetLivenessInit(quest *PlayerQuestObject) {
	questId := quest.QuestId
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	if questTemplate.GetQuestType() != questtypes.QuestTypeLiveness {
		return
	}
	pqdm.restQuestInit(quest)
}

//活跃度任务重置
func (pqdm *PlayerQuestDataManager) CommitLivenessResetInit(questId int32) (quest *PlayerQuestObject) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	if questTemplate.GetQuestType() != questtypes.QuestTypeLiveness {
		return
	}

	quest = pqdm.GetQuestByIdAndState(questtypes.QuestStateCommit, questId)
	if quest == nil {
		return
	}

	pqdm.resetLivenessInit(quest)
	return
}

//功能开启活跃度任务
func (pqdm *PlayerQuestDataManager) CheckLivenessQuest() (questList []*PlayerQuestObject) {
	livenessAllMap := questtemplate.GetQuestTemplateService().GetQuestHuoYueMap()
	for questId, _ := range livenessAllMap {
		quest := pqdm.GetQuestById(questId)
		if quest == nil {
			//判断功能是否开启
			livenessTemplate := livenesstemplate.GetHuoYueTempalteService().GetHuoYueTemplate(questId)
			if livenessTemplate == nil {
				continue
			}
			if !pqdm.p.IsFuncOpen(livenessTemplate.GetFuncOpenTyp()) {
				continue
			}
			quest, flag := pqdm.AddQuest(questId)
			if !flag {
				continue
			}
			questList = append(questList, quest)
		}
	}
	return
}
