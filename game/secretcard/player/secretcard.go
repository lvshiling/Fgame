package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/secretcard/dao"
	secretcardentity "fgame/fgame/game/secretcard/entity"
	secretcardeventtypes "fgame/fgame/game/secretcard/event/types"
	"fgame/fgame/game/secretcard/secretcard"
	secretcardtypes "fgame/fgame/game/secretcard/types"
)

//天机牌对象
type PlayerSecretCardObject struct {
	player       player.Player
	Id           int64
	PlayerId     int64
	TotalNum     int64
	Num          int32
	TotalStar    int32
	OpenBoxList  []int32
	CardId       int32
	Star         int32
	CardMap      map[int32]int32
	UsedCardList []int32
	LastTime     int64
	UpdateTime   int64
	CreateTime   int64
	DeleteTime   int64
}

func NewPlayerSecretCardObject(pl player.Player) *PlayerSecretCardObject {
	pmo := &PlayerSecretCardObject{
		player: pl,
	}
	return pmo
}

func convertObjectToEntity(psco *PlayerSecretCardObject) (*secretcardentity.PlayerSecretCardEntity, error) {
	usedCardBytes, err := json.Marshal(psco.UsedCardList)
	if err != nil {
		return nil, err
	}
	cardsBytes, err := json.Marshal(psco.CardMap)
	if err != nil {
		return nil, err
	}

	openBoxBytes, err := json.Marshal(psco.OpenBoxList)
	if err != nil {
		return nil, err
	}

	e := &secretcardentity.PlayerSecretCardEntity{
		Id:         psco.Id,
		PlayerId:   psco.PlayerId,
		TotalNum:   psco.TotalNum,
		Num:        psco.Num,
		TotalStar:  psco.TotalStar,
		OpenBoxs:   string(openBoxBytes),
		CardId:     psco.CardId,
		Star:       psco.Star,
		Cards:      string(cardsBytes),
		UsedCards:  string(usedCardBytes),
		LastTime:   psco.LastTime,
		UpdateTime: psco.UpdateTime,
		CreateTime: psco.CreateTime,
		DeleteTime: psco.DeleteTime,
	}
	return e, nil
}

func (psco *PlayerSecretCardObject) GetPlayerId() int64 {
	return psco.PlayerId
}

func (psco *PlayerSecretCardObject) GetDBId() int64 {
	return psco.Id
}

func (psco *PlayerSecretCardObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertObjectToEntity(psco)
	return e, err
}

func (psco *PlayerSecretCardObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*secretcardentity.PlayerSecretCardEntity)
	usedCardList := make([]int32, 0, 8)
	cardMap := make(map[int32]int32)
	openBoxList := make([]int32, 0, 8)
	if err := json.Unmarshal([]byte(pse.UsedCards), &usedCardList); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(pse.Cards), &cardMap); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(pse.OpenBoxs), &openBoxList); err != nil {
		return err
	}

	psco.Id = pse.Id
	psco.PlayerId = pse.PlayerId
	psco.TotalNum = pse.TotalNum
	psco.Num = pse.Num
	psco.TotalStar = pse.TotalStar
	psco.OpenBoxList = openBoxList
	psco.CardId = pse.CardId
	psco.Star = pse.Star
	psco.CardMap = cardMap
	psco.UsedCardList = usedCardList
	psco.LastTime = pse.LastTime
	psco.UpdateTime = pse.UpdateTime
	psco.CreateTime = pse.CreateTime
	psco.DeleteTime = pse.DeleteTime
	return nil
}

func (psco *PlayerSecretCardObject) SetModified() {
	e, err := psco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("SecretCard: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psco.player.AddChangedObject(obj)
	return
}

//玩家天机牌管理器
type PlayerSecretCardDataManager struct {
	p player.Player
	//天机牌对象
	secretCardObject *PlayerSecretCardObject
}

//心跳
func (psdm *PlayerSecretCardDataManager) Heartbeat() {

}

func (psdm *PlayerSecretCardDataManager) Player() player.Player {
	return psdm.p
}

//加载
func (psdm *PlayerSecretCardDataManager) Load() (err error) {
	//加载天机牌
	secretCardEntity, err := dao.GetSecretCardDao().GetSecretCardEntity(psdm.p.GetId())
	if err != nil {
		return
	}
	if secretCardEntity == nil {
		psdm.initPlayerSecertCardObject()
	} else {
		psdm.secretCardObject = NewPlayerSecretCardObject(psdm.p)
		psdm.secretCardObject.FromEntity(secretCardEntity)
	}
	return nil
}

//第一次初始化
func (psdm *PlayerSecretCardDataManager) initPlayerSecertCardObject() {
	psco := NewPlayerSecretCardObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	psco.Id = id
	//生成id
	psco.PlayerId = psdm.p.GetId()
	psco.TotalNum = int64(0)
	psco.Num = int32(0)
	psco.TotalStar = int32(0)
	psco.OpenBoxList = make([]int32, 0, 8)
	psco.CardId = int32(0)
	psco.Star = int32(0)
	psco.CardMap = make(map[int32]int32)
	psco.UsedCardList = make([]int32, 0, 8)
	psco.LastTime = int64(0)
	psco.CreateTime = now
	psdm.secretCardObject = psco
	psco.SetModified()
}

//加载后
func (psdm *PlayerSecretCardDataManager) AfterLoad() (err error) {
	//刷新天机牌信息
	psdm.refreshSecretCard()

	return nil
}

//刷新refresh天机牌
func (psdm *PlayerSecretCardDataManager) refreshSecretCard() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := psdm.secretCardObject.LastTime
	if lastTime != 0 {
		//flag, err := timeutils.IsSameDay(lastTime, now)
		flag, err := timeutils.IsSameFive(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			//cardId 不清空,否则过天任务领取不了
			psdm.secretCardObject.Num = 0
			psdm.secretCardObject.TotalStar = 0
			psdm.secretCardObject.OpenBoxList = make([]int32, 0, 8)
			psdm.secretCardObject.Star = 0
			psdm.secretCardObject.CardMap = make(map[int32]int32)
			psdm.secretCardObject.UsedCardList = make([]int32, 0, 8)
			psdm.secretCardObject.LastTime = 0
			psdm.secretCardObject.UpdateTime = now
			psdm.secretCardObject.SetModified()
		}
	}
	return nil
}

//获取天机牌对象
func (psdm *PlayerSecretCardDataManager) GetSecretCard() *PlayerSecretCardObject {
	//刷新天机牌
	psdm.refreshSecretCard()
	return psdm.secretCardObject
}

//校验参数
func (psdm *PlayerSecretCardDataManager) IsValidSecretCard(cardId int32) (sucess bool) {
	for id, _ := range psdm.secretCardObject.CardMap {
		if id == cardId {
			sucess = true
			break
		}
	}
	return
}

//接取天机牌任务
func (psdm *PlayerSecretCardDataManager) PickUpSecretCard(cardId int32) (questId int32, sucess bool) {
	sucess = psdm.IsValidSecretCard(cardId)
	if !sucess {
		return
	}
	questId, _, sucess = secretcard.GetSecretCardService().GetQuestIdByCardId(cardId)
	if !sucess {
		return
	}
	star := psdm.secretCardObject.CardMap[cardId]
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.CardId = cardId
	psdm.secretCardObject.Star = star
	psdm.secretCardObject.CardMap = make(map[int32]int32)
	typ, sucess := secretcard.GetSecretCardService().GetQuestPoolType(cardId)
	if !sucess {
		return
	}
	if typ != secretcardtypes.SecretCardPoolTypePoll {
		psdm.secretCardObject.UsedCardList = append(psdm.secretCardObject.UsedCardList, cardId)
	}
	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.secretCardObject.SetModified()

	sucess = true
	return
}

//天机牌能否领取星数奖励
func (psdm *PlayerSecretCardDataManager) IfSecretStarRew(openBox int32) (sucess bool) {
	psdm.refreshSecretCard()
	totalStar := psdm.secretCardObject.TotalStar
	for _, curOpenBox := range psdm.secretCardObject.OpenBoxList {
		if curOpenBox == openBox {
			return
		}
	}
	to := secretcard.GetSecretCardService().GetStarTemplate(openBox)
	if to == nil {
		return
	}
	if totalStar < to.NeedStar {
		return
	}
	sucess = true
	return
}

//运势箱开启
func (psdm *PlayerSecretCardDataManager) SecretStarRew(openBox int32) (sucess bool) {
	sucess = psdm.IfSecretStarRew(openBox)
	if !sucess {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.OpenBoxList = append(psdm.secretCardObject.OpenBoxList, openBox)
	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.secretCardObject.SetModified()
	return
}

//能否窥探天机
func (psdm *PlayerSecretCardDataManager) IfCanSecretSpy() (sucess bool) {
	psdm.refreshSecretCard()
	maxNum := secretcard.GetSecretCardService().GetConstSecretCardNum()
	if psdm.secretCardObject.Num >= maxNum {
		return
	}
	sucess = true
	return
}

//放弃任务
func (psdm *PlayerSecretCardDataManager) DiscardQuest() {
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.CardId = 0
	psdm.secretCardObject.Star = 0
	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.secretCardObject.SetModified()
	return
}

func (psdm *PlayerSecretCardDataManager) ImmediateFinish() {
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.CardId = 0
	psdm.secretCardObject.TotalStar += psdm.secretCardObject.Star
	psdm.secretCardObject.Star = 0
	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.secretCardObject.SetModified()
	return
}

//获取剩余次数
func (psdm *PlayerSecretCardDataManager) GetLeftCardNum() (leftNum int32) {
	psdm.refreshSecretCard()
	num := psdm.secretCardObject.Num
	maxNum := secretcard.GetSecretCardService().GetConstSecretCardNum()
	leftNum = maxNum - num
	return
}

func (psdm *PlayerSecretCardDataManager) SecretCardSyp(cardMap map[int32]int32) (num int32) {
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.TotalNum += 1
	psdm.secretCardObject.Num += 1
	num = psdm.secretCardObject.Num
	for cardId, star := range cardMap {
		psdm.secretCardObject.CardMap[cardId] = star
	}

	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.secretCardObject.SetModified()
	return
}

//一键完成
func (psdm *PlayerSecretCardDataManager) SecretCardFinishAll(leftNum int32, addStar int32, leftBoxIdList []int32) {
	for _, boxId := range leftBoxIdList {
		psdm.secretCardObject.OpenBoxList = append(psdm.secretCardObject.OpenBoxList, boxId)
	}
	if len(psdm.secretCardObject.CardMap) != 0 {
		psdm.secretCardObject.CardMap = make(map[int32]int32)
	}
	psdm.secretCardObject.TotalNum += int64(leftNum)
	psdm.secretCardObject.Num += leftNum
	psdm.SecretCardFinish(addStar)
	gameevent.Emit(secretcardeventtypes.EventTypeSecretCardFinishAll, psdm.p, nil)
	return
}

//天机任务完成
func (psdm *PlayerSecretCardDataManager) SecretCardFinish(addStar int32) {
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.Star = 0
	psdm.secretCardObject.CardId = 0
	psdm.secretCardObject.TotalStar += addStar
	psdm.secretCardObject.UpdateTime = now
	psdm.secretCardObject.LastTime = now
	psdm.refreshSecretCard()
	psdm.secretCardObject.SetModified()
	return
}

//仅gm使用
func (psdm *PlayerSecretCardDataManager) GMClearNum() {
	now := global.GetGame().GetTimeService().Now()
	psdm.secretCardObject.Num = 0
	psdm.secretCardObject.UpdateTime = now
	psdm.refreshSecretCard()
	psdm.secretCardObject.SetModified()
	return
}

func CreatePlayerSecretCardDataManager(p player.Player) player.PlayerDataManager {
	psdm := &PlayerSecretCardDataManager{}
	psdm.p = p
	return psdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSecretCardDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSecretCardDataManager))
}
