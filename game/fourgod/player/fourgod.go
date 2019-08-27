package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/fourgod/dao"
	fourgodentity "fgame/fgame/game/fourgod/entity"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"math"

	"github.com/pkg/errors"
)

//玩家四神遗迹对象
type PlayerFourGodObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	KeyNum     int32
	Exp        int64
	ItemMap    map[int32]int32
	EndTime    int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerFourGodObject(pl player.Player) *PlayerFourGodObject {
	pdo := &PlayerFourGodObject{
		player: pl,
	}
	return pdo
}

func (pdo *PlayerFourGodObject) GetPlayerId() int64 {
	return pdo.PlayerId
}

func (pdo *PlayerFourGodObject) GetDBId() int64 {
	return pdo.Id
}

func convertObjectToEntity(pdo *PlayerFourGodObject) (*fourgodentity.PlayerFourGodEntity, error) {

	itemsBytes, err := json.Marshal(pdo.ItemMap)
	if err != nil {
		return nil, err
	}

	e := &fourgodentity.PlayerFourGodEntity{
		Id:         pdo.Id,
		PlayerId:   pdo.PlayerId,
		KeyNum:     pdo.KeyNum,
		Exp:        pdo.Exp,
		ItemInfo:   string(itemsBytes),
		EndTime:    pdo.EndTime,
		UpdateTime: pdo.UpdateTime,
		CreateTime: pdo.CreateTime,
		DeleteTime: pdo.DeleteTime,
	}
	return e, nil
}

func (pdo *PlayerFourGodObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(pdo)
	return
}

func (pdo *PlayerFourGodObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fourgodentity.PlayerFourGodEntity)

	itemInfoMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(pse.ItemInfo), &itemInfoMap); err != nil {
		return err
	}

	pdo.Id = pse.Id
	pdo.PlayerId = pse.PlayerId
	pdo.KeyNum = pse.KeyNum
	pdo.Exp = pse.Exp
	pdo.EndTime = pse.EndTime
	pdo.ItemMap = itemInfoMap
	pdo.UpdateTime = pse.UpdateTime
	pdo.CreateTime = pse.CreateTime
	pdo.DeleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerFourGodObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FourGod"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}

//玩家四神遗迹管理器
type PlayerFourGodDataManager struct {
	p player.Player
	//玩家四神遗迹对象
	fourGodObject *PlayerFourGodObject
}

func (pddm *PlayerFourGodDataManager) Player() player.Player {
	return pddm.p
}

//加载
func (pddm *PlayerFourGodDataManager) Load() (err error) {
	//加载玩家四神遗迹
	fourGodEntity, err := dao.GetFourGodDao().GetFourGodEntity(pddm.p.GetId())
	if err != nil {
		return
	}
	if fourGodEntity == nil {
		pddm.initPlayerFourGodObject()
	} else {
		pddm.fourGodObject = NewPlayerFourGodObject(pddm.p)
		pddm.fourGodObject.FromEntity(fourGodEntity)
	}

	return nil
}

//第一次初始化
func (pddm *PlayerFourGodDataManager) initPlayerFourGodObject() {
	pdo := NewPlayerFourGodObject(pddm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.Id = id
	//生成id
	pdo.PlayerId = pddm.p.GetId()
	pdo.KeyNum = int32(0)
	pdo.Exp = int64(0)
	pdo.ItemMap = make(map[int32]int32)
	pdo.EndTime = int64(0)
	pdo.CreateTime = now
	pddm.fourGodObject = pdo
	pdo.SetModified()
}

//加载后
func (pddm *PlayerFourGodDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pddm *PlayerFourGodDataManager) Heartbeat() {

}

//获取钥匙
func (pddm *PlayerFourGodDataManager) GetKeyNum() int32 {
	return pddm.fourGodObject.KeyNum
}

//活动获得的经验和物品
func (pddm *PlayerFourGodDataManager) GetExpAndItemMap() (exp int64, itemMap map[int32]int32) {
	exp = pddm.fourGodObject.Exp
	itemMap = pddm.fourGodObject.ItemMap
	return
}

//活动结束时间
func (pddm *PlayerFourGodDataManager) EndTime(endTime int64) {
	if pddm.fourGodObject.EndTime == endTime {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pddm.fourGodObject.KeyNum = 0
	pddm.fourGodObject.Exp = 0
	pddm.fourGodObject.ItemMap = make(map[int32]int32)
	pddm.fourGodObject.EndTime = endTime
	pddm.fourGodObject.UpdateTime = now
	pddm.fourGodObject.SetModified()
	return
}

//玩家退出 清空钥匙
func (pddm *PlayerFourGodDataManager) ExitFourGod() {
	pddm.clearKeyNum()
	gameevent.Emit(fourgodeventtypes.EventTypeFourGodKeyChange, pddm.p, nil)
	return
}

//玩家开宝箱
func (pddm *PlayerFourGodDataManager) OpenBox(itemList []*droptemplate.DropItemData, useKeyNum int32) {
	if useKeyNum <= 0 {
		return
	}
	for _, itemData := range itemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()

		curNum, exist := pddm.fourGodObject.ItemMap[itemId]
		if exist {
			num += curNum
		}
		pddm.fourGodObject.ItemMap[itemId] = num
	}
	pddm.openBoxSubKey(useKeyNum)
	gameevent.Emit(fourgodeventtypes.EventTypeFourGodKeyChange, pddm.p, nil)
	return
}

//添加获得经验
func (pddm *PlayerFourGodDataManager) AddExp(exp int64) {
	if exp <= 0 {
		return
	}
	//zrc: 不保存,退出保存
	pddm.fourGodObject.Exp += exp

	return
}

//添加获得的物品
func (pddm *PlayerFourGodDataManager) AddItem(itemId int32, num int32) {
	if num <= 0 {
		return
	}
	//zrc: 不保存,退出保存
	curNum, exist := pddm.fourGodObject.ItemMap[itemId]
	if exist {
		num += curNum
	}
	pddm.fourGodObject.ItemMap[itemId] = num

	return
}

//清空钥匙
func (pddm *PlayerFourGodDataManager) clearKeyNum() {
	now := global.GetGame().GetTimeService().Now()
	pddm.fourGodObject.KeyNum = 0
	pddm.fourGodObject.UpdateTime = now
	pddm.fourGodObject.SetModified()
	return
}

func (pddm *PlayerFourGodDataManager) openBoxSubKey(useKeyNum int32) {
	now := global.GetGame().GetTimeService().Now()
	if pddm.fourGodObject.KeyNum < useKeyNum {
		useKeyNum = pddm.fourGodObject.KeyNum
	}
	pddm.fourGodObject.KeyNum -= useKeyNum
	pddm.fourGodObject.UpdateTime = now
	// pddm.fourGodObject.SetModified()
}

//增加钥匙
func (pddm *PlayerFourGodDataManager) AddKeyNum(num int32) {
	if num <= 0 {
		return
	}
	keyMax := template.GetFourGodTemplateService().GetFourGodConstTemplate().KeyMax
	addNum := keyMax - pddm.fourGodObject.KeyNum
	if addNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pddm.fourGodObject.KeyNum += num
	pddm.fourGodObject.UpdateTime = now
	// pddm.fourGodObject.SetModified()

	gameevent.Emit(fourgodeventtypes.EventTypeFourGodKeyChange, pddm.p, nil)
	return
}

//玩家死亡 钥匙减半
func (pddm *PlayerFourGodDataManager) KeyHalve() (dropNum int32) {
	beforeNum := pddm.fourGodObject.KeyNum
	if beforeNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	dropNum = int32(math.Floor(float64(beforeNum) / 2))
	pddm.fourGodObject.KeyNum = dropNum
	pddm.fourGodObject.UpdateTime = now
	// pddm.fourGodObject.SetModified()

	gameevent.Emit(fourgodeventtypes.EventTypeFourGodKeyChange, pddm.p, nil)

	return
}

func CreatePlayerFourGodDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerFourGodDataManager{}
	pddm.p = p
	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFourGodDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFourGodDataManager))
}
