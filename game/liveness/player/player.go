package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/liveness/dao"
	livenesseventtypes "fgame/fgame/game/liveness/event/types"
	livenesstemplate "fgame/fgame/game/liveness/template"
)

//玩家活跃度管理器
type PlayerLivenessDataManager struct {
	p player.Player
	//活跃度对象
	livenessObject *PlayerLivenessObject
	//活跃度任务对象
	livenessQuestMap map[int32]*PlayerLivenessQuestObject
}

//心跳
func (psdm *PlayerLivenessDataManager) Heartbeat() {

}

func (psdm *PlayerLivenessDataManager) Player() player.Player {
	return psdm.p
}

//加载
func (psdm *PlayerLivenessDataManager) Load() (err error) {
	err = psdm.loadLiveness()
	if err != nil {
		return
	}
	err = psdm.loadLivenessQuest()
	if err != nil {
		return
	}
	return nil
}

func (psdm *PlayerLivenessDataManager) loadLiveness() (err error) {
	//加载活跃度
	livenessEntity, err := dao.GetLivenessDao().GetLivenessEntity(psdm.p.GetId())
	if err != nil {
		return
	}
	if livenessEntity == nil {
		psdm.initPlayerLivenessObject()
	} else {
		psdm.livenessObject = NewPlayerLivenessObject(psdm.p)
		psdm.livenessObject.FromEntity(livenessEntity)
	}
	return
}

func (psdm *PlayerLivenessDataManager) loadLivenessQuest() (err error) {
	//加载活跃度任务
	psdm.livenessQuestMap = make(map[int32]*PlayerLivenessQuestObject)
	livenessQuestList, err := dao.GetLivenessDao().GetLivenessQuestList(psdm.p.GetId())
	if err != nil {
		return
	}
	for _, livenessQuest := range livenessQuestList {
		obj := NewPlayerLivenessQuestObject(psdm.p)
		obj.FromEntity(livenessQuest)
		psdm.livenessQuestMap[livenessQuest.QuestId] = obj
	}
	return
}

//第一次初始化
func (psdm *PlayerLivenessDataManager) initPlayerLivenessObject() {
	psco := NewPlayerLivenessObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	psco.id = id
	//生成id
	psco.playerId = psdm.p.GetId()
	psco.liveness = 0
	psco.openBoxList = make([]int32, 0, 8)

	psco.lastTime = 0
	psco.createTime = now
	psdm.livenessObject = psco
	psco.SetModified()
}

func (psdm *PlayerLivenessDataManager) initPlayerLivenessQuestObject(questId int32) (psco *PlayerLivenessQuestObject) {
	_, exist := psdm.livenessQuestMap[questId]
	if exist {
		return
	}
	livenessTempalte := livenesstemplate.GetHuoYueTempalteService().GetHuoYueTemplate(questId)
	if livenessTempalte == nil {
		return
	}
	psco = NewPlayerLivenessQuestObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	psco.id = id
	//生成id
	psco.playerId = psdm.p.GetId()
	psco.questId = questId
	psco.num = 0

	psco.lastTime = 0
	psco.createTime = now
	psdm.livenessQuestMap[questId] = psco
	psco.SetModified()
	return
}

//加载后
func (psdm *PlayerLivenessDataManager) AfterLoad() (err error) {
	psdm.refreshLiveness()
	//刷新活跃度信息
	psdm.refreshLivenessMap()
	return nil
}

//刷新refresh活跃度
func (psdm *PlayerLivenessDataManager) refreshLivenessMap() error {
	//是否跨
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range psdm.livenessQuestMap {
		lastTime := obj.GetLastTime()
		if lastTime != 0 {
			flag, err := timeutils.IsSameFive(lastTime, now)
			if err != nil {
				return err
			}
			if !flag {
				obj.num = 0
				obj.lastTime = now
				obj.updateTime = now
				obj.SetModified()
			}
		}
	}
	return nil
}

func (psdm *PlayerLivenessDataManager) refreshLiveness() {
	now := global.GetGame().GetTimeService().Now()
	lastTime := psdm.livenessObject.GetLastTime()
	if lastTime != 0 {
		flag, _ := timeutils.IsSameFive(lastTime, now)
		if !flag {
			gameevent.Emit(livenesseventtypes.EventTypeLivenessCrossFive, psdm.p, psdm.livenessObject)
			psdm.livenessObject.liveness = 0
			psdm.livenessObject.openBoxList = make([]int32, 0, 8)
			psdm.livenessObject.lastTime = now
			psdm.livenessObject.updateTime = now
			psdm.livenessObject.SetModified()
		}
	}
}

func (psdm *PlayerLivenessDataManager) refreshLivenessByQuestId(questId int32) {
	obj, exist := psdm.livenessQuestMap[questId]
	if !exist {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	lastTime := obj.GetLastTime()
	if lastTime != 0 {
		flag, _ := timeutils.IsSameFive(lastTime, now)
		if !flag {
			obj.num = 0
			obj.lastTime = now
			obj.updateTime = now
			obj.SetModified()
		}
	}
}

//获取活跃度
func (psdm *PlayerLivenessDataManager) GetLiveness() *PlayerLivenessObject {
	psdm.refreshLiveness()
	return psdm.livenessObject
}

//获取活跃度完成次数
func (psdm *PlayerLivenessDataManager) GetLivenessQuestMap() map[int32]*PlayerLivenessQuestObject {
	//刷新活跃度信息
	psdm.refreshLivenessMap()
	return psdm.livenessQuestMap
}

//增加完成次数
func (psdm *PlayerLivenessDataManager) AddQuestNum(questId int32) (num int32) {
	psdm.refreshLiveness()
	psdm.refreshLivenessByQuestId(questId)
	now := global.GetGame().GetTimeService().Now()
	livenessTempalte := livenesstemplate.GetHuoYueTempalteService().GetHuoYueTemplate(questId)
	if livenessTempalte == nil {
		return
	}
	livenessLevelTempalte, flag := livenesstemplate.GetHuoYueTempalteService().GetHuoYueLevelTemplate(questId, psdm.p.GetLevel())
	if !flag {
		return
	}
	if livenessLevelTempalte == nil {
		return
	}
	questObj, exist := psdm.livenessQuestMap[questId]
	if !exist {
		questObj = psdm.initPlayerLivenessQuestObject(questId)
	}

	curNum := questObj.num
	if curNum >= livenessTempalte.RewardCountLimit {
		num = curNum
		return
	}
	//活跃度
 	psdm.livenessObject.liveness += int64(livenessLevelTempalte.HuoYue)
	psdm.livenessObject.lastTime = now
	psdm.livenessObject.updateTime = now
	psdm.livenessObject.SetModified()
	//完成次数
	questObj.num += 1
	questObj.updateTime = now
	questObj.lastTime = now
	questObj.SetModified()
	gameevent.Emit(livenesseventtypes.EventTypeLivenessChanged, psdm.p, questObj)
	return questObj.num
}

//活跃度能否领取星数奖励
func (psdm *PlayerLivenessDataManager) IfLivenessBoxRew(openBox int32) (sucess bool) {
	psdm.refreshLiveness()
	liveness := psdm.livenessObject.GetLiveness()
	for _, curOpenBox := range psdm.livenessObject.GetOpenBoxs() {
		if curOpenBox == openBox {
			return
		}
	}
	to := livenesstemplate.GetHuoYueTempalteService().GetHuoYueBoxTemplate(openBox)
	if to == nil {
		return
	}
	if liveness < int64(to.NeedStar) {
		return
	}
	sucess = true
	return
}

//宝箱开启
func (psdm *PlayerLivenessDataManager) LivenessOpenRew(openBox int32) (sucess bool) {
	sucess = psdm.IfLivenessBoxRew(openBox)
	if !sucess {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	psdm.livenessObject.openBoxList = append(psdm.livenessObject.openBoxList, openBox)
	psdm.livenessObject.updateTime = now
	psdm.livenessObject.lastTime = now
	psdm.livenessObject.SetModified()
	return
}

//仅Gm 命令使用
func (psdm *PlayerLivenessDataManager) GmSetLiveness(liveness int64) {
	now := global.GetGame().GetTimeService().Now()
	psdm.livenessObject.liveness = liveness
	psdm.livenessObject.updateTime = now
	psdm.livenessObject.lastTime = now
	psdm.livenessObject.SetModified()
}

func (psdm *PlayerLivenessDataManager) GmSetLivenessClearNum() {
	now := global.GetGame().GetTimeService().Now()
	for _, livenessQuestObj := range psdm.livenessQuestMap {
		livenessQuestObj.num = 0
		livenessQuestObj.updateTime = now
		livenessQuestObj.lastTime = now
		livenessQuestObj.SetModified()
	}
	psdm.livenessObject.openBoxList = make([]int32, 0, 8)
	psdm.livenessObject.updateTime = now
	psdm.livenessObject.lastTime = now
	psdm.livenessObject.SetModified()
}

func CreatePlayerLivenessDataManager(p player.Player) player.PlayerDataManager {
	psdm := &PlayerLivenessDataManager{}
	psdm.p = p
	return psdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerLivenessDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerLivenessDataManager))
}
