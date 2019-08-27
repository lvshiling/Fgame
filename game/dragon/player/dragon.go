package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/dragon/dao"
	"fgame/fgame/game/dragon/dragon"
	dragonentity "fgame/fgame/game/dragon/entity"
	dragoneventtypes "fgame/fgame/game/dragon/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//神龙对象
type PlayerDragonObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	StageId     int32
	ItemInfoMap map[int32]int32
	Status      int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerDragonObject(pl player.Player) *PlayerDragonObject {
	pdo := &PlayerDragonObject{
		player: pl,
	}
	return pdo
}

func convertPlayerDragonObjectToEntity(pdo *PlayerDragonObject) (*dragonentity.PlayerDragonEntity, error) {
	itemInfoBytes, err := json.Marshal(pdo.ItemInfoMap)
	if err != nil {
		return nil, err
	}

	e := &dragonentity.PlayerDragonEntity{
		Id:         pdo.Id,
		PlayerId:   pdo.PlayerId,
		StageId:    pdo.StageId,
		ItemInfo:   string(itemInfoBytes),
		Status:     pdo.Status,
		UpdateTime: pdo.UpdateTime,
		CreateTime: pdo.CreateTime,
		DeleteTime: pdo.DeleteTime,
	}
	return e, err
}

func (pdo *PlayerDragonObject) GetPlayerId() int64 {
	return pdo.PlayerId
}

func (pdo *PlayerDragonObject) GetDBId() int64 {
	return pdo.Id
}

func (pdo *PlayerDragonObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerDragonObjectToEntity(pdo)
	return e, err
}

func (pdo *PlayerDragonObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*dragonentity.PlayerDragonEntity)
	itemInfoMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pse.ItemInfo), &itemInfoMap); err != nil {
		return err
	}
	pdo.Id = pse.Id
	pdo.PlayerId = pse.PlayerId
	pdo.StageId = pse.StageId
	pdo.ItemInfoMap = itemInfoMap
	pdo.Status = pse.Status
	pdo.UpdateTime = pse.UpdateTime
	pdo.CreateTime = pse.CreateTime
	pdo.DeleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerDragonObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Dragon"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}

//玩家神龙管理器
type PlayerDragonDataManager struct {
	p player.Player
	//玩家神龙对象
	dragonObject *PlayerDragonObject
}

func (pddm *PlayerDragonDataManager) Player() player.Player {
	return pddm.p
}

//加载
func (pddm *PlayerDragonDataManager) Load() (err error) {
	//加载玩家神龙
	dragonEntity, err := dao.GetDragonDao().GetDragonEntity(pddm.p.GetId())
	if err != nil {
		return
	}
	if dragonEntity == nil {
		pddm.initPlayerDragonObject()
	} else {
		pddm.dragonObject = NewPlayerDragonObject(pddm.p)
		pddm.dragonObject.FromEntity(dragonEntity)
	}

	return nil
}

//第一次初始化
func (pddm *PlayerDragonDataManager) initPlayerDragonObject() {
	pdo := NewPlayerDragonObject(pddm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.Id = id
	//生成id
	pdo.PlayerId = pddm.p.GetId()
	pdo.StageId = int32(1)
	pdo.ItemInfoMap = make(map[int32]int32)
	pdo.Status = int32(0)
	pdo.CreateTime = now
	pddm.dragonObject = pdo
	pdo.SetModified()
}

//加载后
func (pddm *PlayerDragonDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pddm *PlayerDragonDataManager) Heartbeat() {

}

//神龙信息
func (pddm *PlayerDragonDataManager) GetDragon() *PlayerDragonObject {
	return pddm.dragonObject
}

//参数校验
func (pddm *PlayerDragonDataManager) IsValid(curItemId int32) (flag bool) {
	if pddm.dragonObject.Status == 1 {
		return
	}
	_, exist := pddm.dragonObject.ItemInfoMap[curItemId]
	if exist {
		flag = true
		return
	}

	maxNum := pddm.getMaxNum(curItemId)
	if maxNum > 0 {
		flag = true
	}

	return
}

//神龙是否激活
func (pddm *PlayerDragonDataManager) IsActive() bool {
	return pddm.dragonObject.Status == 1
}

func (pddm *PlayerDragonDataManager) getMaxNum(curItemId int32) (maxNum int32) {
	maxNum = 0
	to := dragon.GetDragonService().GetDragonTemplate(pddm.dragonObject.StageId)
	for itemId, num := range to.GetNeedItemMap() {
		if curItemId == itemId {
			maxNum = num
		}
	}
	return
}

//能吃几个
func (pddm *PlayerDragonDataManager) CanEatNum(itemId int32) (num int32, flag bool) {
	flag = pddm.IsValid(itemId)
	if !flag {
		return
	}
	maxNum := pddm.getMaxNum(itemId)
	curNum, exist := pddm.dragonObject.ItemInfoMap[itemId]
	if !exist {
		num = maxNum
	} else {
		num = maxNum - curNum
	}
	if num > 0 {
		flag = true
	}

	return
}

//喂养
func (pddm *PlayerDragonDataManager) DragonFeed(itemId int32, num int32) (curStageEatFull bool) {
	if num <= 0 {
		return
	}
	flag := pddm.IsValid(itemId)
	if !flag {
		return
	}
	_, flag = pddm.CanEatNum(itemId)
	if !flag {
		return
	}
	curNum, exist := pddm.dragonObject.ItemInfoMap[itemId]
	if exist {
		num += curNum
	}
	pddm.dragonObject.ItemInfoMap[itemId] = num

	//是否升阶
	full, eatFull, mountId := pddm.IfFullAndEatFull()
	if !full && eatFull {
		pddm.dragonObject.StageId += 1
		pddm.dragonObject.ItemInfoMap = make(map[int32]int32)

	}
	if full && eatFull {
		pddm.dragonObject.Status = 1
		//发送事件
		gameevent.Emit(dragoneventtypes.EventTypeDragonActive, pddm.p, mountId)
	}
	if eatFull {
		gameevent.Emit(dragoneventtypes.EventTypeDragonAdvanced, pddm.p, pddm.dragonObject.StageId)
	}

	now := global.GetGame().GetTimeService().Now()
	pddm.dragonObject.UpdateTime = now
	pddm.dragonObject.SetModified()
	curStageEatFull = eatFull
	return
}

//是否达到最高阶和吃满
func (pddm *PlayerDragonDataManager) IfFullAndEatFull() (full, eatFull bool, mountId int32) {
	to := dragon.GetDragonService().GetDragonTemplate(pddm.dragonObject.StageId)
	if to == nil {
		return
	}
	if to.NextId == 0 {
		mountId = to.DragonMount
		full = true
	}
	for itemId, num := range to.GetNeedItemMap() {
		curNum, exist := pddm.dragonObject.ItemInfoMap[itemId]
		if !exist {
			return
		}
		if curNum == num {
			continue
		}
		return
	}
	eatFull = true
	return
}

//仅gm使用
func (pddm *PlayerDragonDataManager) GmSetDragonStage(stageId int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.dragonObject.StageId = stageId
	pddm.dragonObject.ItemInfoMap = make(map[int32]int32)
	pddm.dragonObject.UpdateTime = now
	pddm.dragonObject.SetModified()
	return
}

//仅gm使用
func (pddm *PlayerDragonDataManager) GmDragonActive() {
	now := global.GetGame().GetTimeService().Now()

	to := dragon.GetDragonService().GetDragonTemplate(pddm.dragonObject.StageId)
	for to.NextId == 0 {
		to = dragon.GetDragonService().GetDragonTemplate(pddm.dragonObject.StageId)
		pddm.dragonObject.StageId = int32(to.Id)
		if to.NextId == 0 {
			pddm.dragonObject.Status = 1
			for itemId, num := range to.GetNeedItemMap() {
				pddm.dragonObject.ItemInfoMap[itemId] = num
			}
		}
	}
	pddm.dragonObject.UpdateTime = now
	pddm.dragonObject.SetModified()
	return
}

func CreatePlayerDragonDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerDragonDataManager{}
	pddm.p = p
	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerDragonDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerDragonDataManager))
}
