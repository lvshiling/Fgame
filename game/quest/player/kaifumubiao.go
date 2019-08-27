package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
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

//开服目标对象
type PlayerKaiFuMuBiaoObject struct {
	player     player.Player
	id         int64
	kaiFuDay   int32
	finishNum  int32
	isReward   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerKaiFuMuBiaoObject(pl player.Player) *PlayerKaiFuMuBiaoObject {
	pmo := &PlayerKaiFuMuBiaoObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerKaiFuMuBiaoObjectToEntity(pqo *PlayerKaiFuMuBiaoObject) (e *questentity.PlayerKaiFuMuBiaoEntity, err error) {

	e = &questentity.PlayerKaiFuMuBiaoEntity{
		Id:         pqo.id,
		KaiFuDay:   pqo.kaiFuDay,
		PlayerId:   pqo.player.GetId(),
		FinishNum:  pqo.finishNum,
		IsReward:   pqo.isReward,
		UpdateTime: pqo.updateTime,
		CreateTime: pqo.createTime,
		DeleteTime: pqo.deleteTime,
	}
	return
}

func (pqo *PlayerKaiFuMuBiaoObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerKaiFuMuBiaoObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerKaiFuMuBiaoObject) GetKaiFuDay() int32 {
	return pqo.kaiFuDay
}

func (pqo *PlayerKaiFuMuBiaoObject) GetFinishNum() int32 {
	return pqo.finishNum
}

func (pqo *PlayerKaiFuMuBiaoObject) IsGroupRewardGet() bool {
	return pqo.isReward == 1
}

func (pqo *PlayerKaiFuMuBiaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerKaiFuMuBiaoObjectToEntity(pqo)
	return
}

func (pqo *PlayerKaiFuMuBiaoObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerKaiFuMuBiaoEntity)

	pqo.id = pqe.Id
	pqo.isReward = pqe.IsReward
	pqo.finishNum = pqe.FinishNum
	pqo.updateTime = pqe.UpdateTime
	pqo.kaiFuDay = pqe.KaiFuDay
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerKaiFuMuBiaoObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "kaifumubiao"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}

//数据库加载
func (pqdm *PlayerQuestDataManager) loadKaiFuMuBiao() (err error) {
	pqdm.playerMuBiaoMap = make(map[int32]*PlayerKaiFuMuBiaoObject)
	//加载开服目标信息
	kaiFuMuBiaoEntityList, err := dao.GetQuestDao().GetKaiFuMuBiaoList(pqdm.p.GetId())
	if err != nil {
		return
	}
	for _, kaiFuMuBiaoEntity := range kaiFuMuBiaoEntityList {
		obj := NewPlayerKaiFuMuBiaoObject(pqdm.p)
		err = obj.FromEntity(kaiFuMuBiaoEntity)
		if err != nil {
			return
		}
		pqdm.playerMuBiaoMap[obj.GetKaiFuDay()] = obj
	}
	return
}

func (pqdm *PlayerQuestDataManager) initKaiFuMuBiao(kaiFuDay int32) {
	obj, ok := pqdm.playerMuBiaoMap[kaiFuDay]
	if ok {
		return
	}

	kaiFuMuBiaoTemplate := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoTemplate(kaiFuDay)
	if kaiFuMuBiaoTemplate == nil {
		return
	}
	obj = NewPlayerKaiFuMuBiaoObject(pqdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	//生成id
	obj.kaiFuDay = kaiFuDay
	obj.finishNum = 0
	obj.isReward = 0
	obj.createTime = now
	obj.SetModified()
	pqdm.playerMuBiaoMap[kaiFuDay] = obj
}

//开服目标任务
func (pqdm *PlayerQuestDataManager) afterLoadKaiFuMuBiaoQuest() (err error) {
	if len(pqdm.playerMuBiaoMap) >= questtypes.KaiFuMuBiaoDayMax {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	diff, _ := timeutils.DiffDay(now, openTime)
	// //开服区间外&开服区间玩家登录过
	// if diff+1 > questtypes.KaiFuMuBiaoDayMax && len(pqdm.playerMuBiaoMap) == 0 {
	// 	return
	// }
	maxOpenDay := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiFuMuBiaoMaxOpenDay)
	if maxOpenDay != 0 && diff+1 > maxOpenDay && len(pqdm.playerMuBiaoMap) == 0 {
		return
	}
	pqdm.refreshKaiFuMuBiaoQuest(diff)
	return
}

func (pqdm *PlayerQuestDataManager) refreshKaiFuMuBiaoQuest(diff int32) {
	kaiFuMuBiaoTemplateMap := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoMap()
	for kaiFuDay, _ := range kaiFuMuBiaoTemplateMap {
		_, ok := pqdm.playerMuBiaoMap[kaiFuDay]
		if ok {
			continue
		}
		if diff+1 < kaiFuDay {
			continue
		}
		pqdm.initKaiFuMuBiao(kaiFuDay)
		pqdm.initKaiFuMuBiaoQuest(kaiFuDay)
	}
}

func (pqdm *PlayerQuestDataManager) initKaiFuMuBiaoQuest(kaiFuDay int32) {
	_, ok := pqdm.playerMuBiaoMap[kaiFuDay]
	if !ok {
		return
	}
	kaiFuMuBiaoTemplate := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoTemplate(kaiFuDay)
	if kaiFuMuBiaoTemplate == nil {
		return
	}
	questMap := kaiFuMuBiaoTemplate.GetQuestMap()
	for questId, _ := range questMap {
		pqdm.addKaiFuMuBiaoQuest(questId)
	}
}

//添加开服目标任务
func (pqdm *PlayerQuestDataManager) addKaiFuMuBiaoQuest(questId int32) {
	questObj := pqdm.GetQuestById(questId)
	//策划改数据 建议清表 (正式数据不应该出现)
	if questObj != nil {
		return
	}
	questObj = createQuest(pqdm.p, questId)
	questObj.SetModified()
	pqdm.addQuest(questObj)
	return
}

func (pqdm *PlayerQuestDataManager) GetKaiFuMuBiaoMap() map[int32]*PlayerKaiFuMuBiaoObject {
	return pqdm.playerMuBiaoMap
}

func (pqdm *PlayerQuestDataManager) GetKaiFuMuBiao(kaiFuDay int32) *PlayerKaiFuMuBiaoObject {
	if kaiFuDay < 1 || kaiFuDay > questtypes.KaiFuMuBiaoDayMax {
		return nil
	}

	obj, ok := pqdm.playerMuBiaoMap[kaiFuDay]
	if !ok {
		return nil
	}
	return obj
}

func (pqdm *PlayerQuestDataManager) IfCanReceive(kaiFuDay int32) (flag bool) {
	kaiFuMuBiaoObj := pqdm.GetKaiFuMuBiao(kaiFuDay)
	if kaiFuMuBiaoObj == nil {
		return
	}
	if kaiFuMuBiaoObj.IsGroupRewardGet() {
		return
	}
	kaiFuMuBiaoTemplate := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoTemplate(kaiFuDay)
	if kaiFuMuBiaoTemplate == nil {
		return
	}
	if kaiFuMuBiaoObj.GetFinishNum() < kaiFuMuBiaoTemplate.FinishQuestCount {
		return
	}
	flag = true
	return
}

func (pqdm *PlayerQuestDataManager) GroupReward(kaiFuDay int32) (flag bool) {
	if !pqdm.IfCanReceive(kaiFuDay) {
		return
	}
	kaiFuMuBiaoObj := pqdm.GetKaiFuMuBiao(kaiFuDay)
	if kaiFuMuBiaoObj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	kaiFuMuBiaoObj.isReward = 1
	kaiFuMuBiaoObj.updateTime = now
	kaiFuMuBiaoObj.SetModified()
	flag = true
	return
}

func (pqdm *PlayerQuestDataManager) AddFinishNum(kaiFuDayList []int32) {
	now := global.GetGame().GetTimeService().Now()
	var kaiFuTimeList []int32
	for _, kaiFuDay := range kaiFuDayList {
		obj, ok := pqdm.playerMuBiaoMap[kaiFuDay]
		if !ok {
			continue
		}
		obj.finishNum += 1
		obj.updateTime = now
		obj.SetModified()
		kaiFuTimeList = append(kaiFuTimeList, kaiFuDay)
	}

	if len(kaiFuTimeList) != 0 {
		gameevent.Emit(questeventtypes.EventTypeQuestKaiFuMuBiaoFinishChanged, pqdm.p, kaiFuTimeList)
	}
}
