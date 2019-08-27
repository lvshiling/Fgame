package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/dan/dan"
	"fgame/fgame/game/dan/dao"
	danentity "fgame/fgame/game/dan/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"math"

	dantypes "fgame/fgame/game/dan/types"

	"github.com/pkg/errors"
)

//食药对象
type PlayerDanObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	LevelId    int32
	DanInfoMap map[int]int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerDanObject(pl player.Player) *PlayerDanObject {
	pdo := &PlayerDanObject{
		player: pl,
	}
	return pdo
}

func convertNewPlayerDanObjectToEntity(pdo *PlayerDanObject) (*danentity.PlayerDanEntity, error) {
	danInfoBytes, err := json.Marshal(pdo.DanInfoMap)
	if err != nil {
		return nil, err
	}

	e := &danentity.PlayerDanEntity{
		Id:         pdo.Id,
		PlayerId:   pdo.PlayerId,
		LevelId:    pdo.LevelId,
		DanInfo:    string(danInfoBytes),
		UpdateTime: pdo.UpdateTime,
		CreateTime: pdo.CreateTime,
		DeleteTime: pdo.DeleteTime,
	}
	return e, err
}

func (pdo *PlayerDanObject) GetPlayerId() int64 {
	return pdo.PlayerId
}

func (pdo *PlayerDanObject) GetDBId() int64 {
	return pdo.Id
}

func (pdo *PlayerDanObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerDanObjectToEntity(pdo)
	return e, err
}

func (pdo *PlayerDanObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*danentity.PlayerDanEntity)
	danInfoMap := make(map[int]int32)
	if err := json.Unmarshal([]byte(pse.DanInfo), &danInfoMap); err != nil {
		return err
	}
	pdo.Id = pse.Id
	pdo.PlayerId = pse.PlayerId
	pdo.LevelId = pse.LevelId
	pdo.DanInfoMap = danInfoMap
	pdo.UpdateTime = pse.UpdateTime
	pdo.CreateTime = pse.CreateTime
	pdo.DeleteTime = pse.DeleteTime
	return nil
}

func (pdo *PlayerDanObject) SetModified() {
	e, err := pdo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Dan"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pdo.player.AddChangedObject(obj)
	return
}

//炼丹对象
type PlayerAlchemyObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	KindId     int
	Num        int
	StartTime  int64
	State      dantypes.AlchemyState
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerAlchemyObject(pl player.Player) *PlayerAlchemyObject {
	pao := &PlayerAlchemyObject{
		player: pl,
	}
	return pao
}

func (pao *PlayerAlchemyObject) GetPlayerId() int64 {
	return pao.PlayerId
}

func (pao *PlayerAlchemyObject) GetDBId() int64 {
	return pao.Id
}

func (pao *PlayerAlchemyObject) ToEntity() (e storage.Entity, err error) {
	e = &danentity.PlayerAlchemyEntity{
		Id:         pao.Id,
		PlayerId:   pao.PlayerId,
		KindId:     pao.KindId,
		Num:        pao.Num,
		StartTime:  pao.StartTime,
		State:      int(pao.State),
		UpdateTime: pao.UpdateTime,
		CreateTime: pao.CreateTime,
		DeleteTime: pao.DeleteTime,
	}
	return e, nil
}

func (pao *PlayerAlchemyObject) FromEntity(e storage.Entity) error {
	pae, _ := e.(*danentity.PlayerAlchemyEntity)
	pao.Id = pae.Id
	pao.PlayerId = pae.PlayerId
	pao.KindId = pae.KindId
	pao.Num = pae.Num
	pao.StartTime = pae.StartTime
	pao.State = dantypes.AlchemyState(pae.State)
	pao.UpdateTime = pae.UpdateTime
	pao.CreateTime = pae.CreateTime
	pao.DeleteTime = pae.DeleteTime
	return nil
}

func (pao *PlayerAlchemyObject) SetModified() {
	e, err := pao.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Alchemy"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pao.player.AddChangedObject(obj)
	return
}

//玩家食丹管理器
type PlayerDanDataManager struct {
	p player.Player
	//玩家食丹对象
	playerDanObject *PlayerDanObject
	//玩家炼丹
	alchemyListMap map[int]*PlayerAlchemyObject
	//玩家领取过丹药,未再炼丹
	alchemyUsedListMap map[int]*PlayerAlchemyObject
}

func (pddm *PlayerDanDataManager) Player() player.Player {
	return pddm.p
}

//加载
func (pddm *PlayerDanDataManager) Load() (err error) {
	//加载玩家食丹
	danEntity, err := dao.GetDanDao().GetDanEntity(pddm.p.GetId())
	if err != nil {
		return
	}
	if danEntity == nil {
		pddm.initPlayerDanObject()
	} else {
		pddm.playerDanObject = NewPlayerDanObject(pddm.p)
		pddm.playerDanObject.FromEntity(danEntity)
	}

	//加载玩家炼丹
	achemyItems, err := dao.GetDanDao().GetAlchemyList(pddm.p.GetId())
	if err != nil {
		return
	}

	//炼丹信息
	for _, item := range achemyItems {
		pao := NewPlayerAlchemyObject(pddm.p)
		pao.FromEntity(item)
		if pao.State != dantypes.AlchemyStateReceive {
			//添加炼丹信息
			pddm.alchemyListMap[item.KindId] = pao
			continue
		}
		pddm.alchemyUsedListMap[item.KindId] = pao
	}
	return nil
}

//第一次初始化
func (pddm *PlayerDanDataManager) initPlayerDanObject() {
	pdo := NewPlayerDanObject(pddm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.Id = id
	//生成id
	pdo.PlayerId = pddm.p.GetId()
	pdo.LevelId = int32(1)
	pdo.DanInfoMap = make(map[int]int32)
	pdo.CreateTime = now
	pddm.playerDanObject = pdo
	pdo.SetModified()
}

//加载后
func (pddm *PlayerDanDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (ppdm *PlayerDanDataManager) Heartbeat() {

}

//获取食丹信息
func (pddm *PlayerDanDataManager) GetDanInfo() *PlayerDanObject {
	return pddm.playerDanObject
}

//能吃几个
func (pddm *PlayerDanDataManager) WhatEat() (dans map[int32]int32, flag bool) {
	flag = pddm.CheckFullLevel()
	if flag {
		return nil, false
	}
	id := pddm.playerDanObject.LevelId + 1
	to := dan.GetDanService().GetEatDan(int(id))
	dans = make(map[int32]int32)
	for danId, dannum := range to.GetAllEatDan() {
		currentNum := int32(0)
		if v, ok := pddm.playerDanObject.DanInfoMap[danId]; ok {
			currentNum = v
		}
		num := dannum - currentNum
		if num <= 0 {
			continue
		}
		dans[int32(danId)] = num
		flag = true
	}
	return dans, flag
}

//全部食用
func (pddm *PlayerDanDataManager) EatDan(dans map[int32]int32) {
	for id, num := range dans {
		currentNum := int32(0)
		if v, ok := pddm.playerDanObject.DanInfoMap[int(id)]; ok {
			currentNum = v
		}
		pddm.playerDanObject.DanInfoMap[int(id)] = currentNum + int32(num)
	}
	now := global.GetGame().GetTimeService().Now()
	pddm.playerDanObject.UpdateTime = now
	pddm.playerDanObject.SetModified()
	return
}

//升级
func (pddm *PlayerDanDataManager) Upgrade() bool {
	flag := pddm.CheckFullLevel()
	if flag {
		return false
	}
	flag = pddm.IfEatEnough()
	if !flag {
		return false
	}
	pddm.playerDanObject.LevelId += int32(1)
	now := global.GetGame().GetTimeService().Now()
	pddm.playerDanObject.UpdateTime = now
	pddm.playerDanObject.SetModified()
	pddm.clearData()
	return true
}

//校验已食用的丹药数
func (pddm *PlayerDanDataManager) IfEatEnough() (flag bool) {
	id := pddm.playerDanObject.LevelId + 1
	eatDanTemplate := dan.GetDanService().GetEatDan(int(id))
	for danId, num := range eatDanTemplate.GetAllEatDan() {
		currentNum := int32(0)
		if v, ok := pddm.playerDanObject.DanInfoMap[danId]; ok {
			currentNum = v
		}
		if num != currentNum {
			flag = false
		}
	}
	flag = true
	return
}

//校验食丹是否满级
func (pddm *PlayerDanDataManager) CheckFullLevel() bool {
	flag := false
	id := pddm.playerDanObject.LevelId
	danTemplate := dan.GetDanService().GetEatDan(int(id))
	if danTemplate.NextId == 0 {
		flag = true
	}
	return flag
}

//清空数据
func (pddm *PlayerDanDataManager) clearData() {
	for id, _ := range pddm.playerDanObject.DanInfoMap {
		delete(pddm.playerDanObject.DanInfoMap, id)
	}
}

func (pddm *PlayerDanDataManager) refreshAchemy() {
	now := global.GetGame().GetTimeService().Now()
	for id, alchemy := range pddm.alchemyListMap {
		if alchemy.State != dantypes.AlchemyStateStart {
			continue
		}
		alchemyTemplate := dan.GetDanService().GetAlchemy(id)
		costTime := alchemyTemplate.GetAchemyTime()

		totalTime := int64(alchemy.Num * int(costTime))
		diffTime := now - alchemy.StartTime - totalTime
		if diffTime < 0 {
			continue
		}
		alchemy.UpdateTime = now
		alchemy.State = dantypes.AlchemyStateEnd
		alchemy.SetModified()
	}
}

//获取炼丹信息
func (pddm *PlayerDanDataManager) GetAlchemyInfo() (achemyInfo map[int]*PlayerAlchemyObject) {
	pddm.refreshAchemy()
	return pddm.alchemyListMap
}

//获取丹药信息根据
func (pddm *PlayerDanDataManager) GetAlchemy(kindId int) *PlayerAlchemyObject {
	if v, ok := pddm.alchemyListMap[kindId]; ok {
		return v
	}
	return nil
}

//开始炼丹
func (pddm *PlayerDanDataManager) SetAlchemyStart(kindId int, num int32) bool {
	if num <= 0 {
		return false
	}
	aObj := pddm.existKindId(kindId)
	if aObj != nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	if v, ok := pddm.alchemyUsedListMap[kindId]; ok {
		v.Num = int(num)
		v.StartTime = now
		v.State = dantypes.AlchemyStateStart
		v.SetModified()
		delete(pddm.alchemyUsedListMap, kindId)
		pddm.alchemyListMap[kindId] = v
		return true
	}
	pao := NewPlayerAlchemyObject(pddm.p)
	id, _ := idutil.GetId()
	pao.Id = id
	//生成id
	pao.PlayerId = pddm.p.GetId()
	pao.KindId = kindId
	pao.Num = int(num)
	pao.StartTime = now
	pao.State = dantypes.AlchemyStateStart
	pao.CreateTime = now
	pddm.alchemyListMap[kindId] = pao
	pao.SetModified()
	return true
}

//领取丹药清数据
func (pddm *PlayerDanDataManager) ClearAlchemyReceive(kindId int) bool {
	alchemyItem := pddm.existKindId(kindId)
	if alchemyItem == nil {
		return false
	}
	if alchemyItem.State == dantypes.AlchemyStateReceive {
		return true
	}
	if !pddm.IsAlchemyFinish(kindId) {
		return false
	}
	pddm.updateAlchemyState(kindId, dantypes.AlchemyStateReceive)
	delete(pddm.alchemyListMap, kindId)
	return true
}

//计算加速炼丹消耗元宝
func (pddm *PlayerDanDataManager) GetAccelerateNeedGold(kindId int) (needGold int32, synthetiseId int32, leftNum int32) {
	leftNum = 0
	needGold = 0
	synthetiseId = 0
	alchemyObj := pddm.GetAlchemy(kindId)
	if alchemyObj == nil {
		return
	}
	alchemyTemplate := dan.GetDanService().GetAlchemy(kindId)
	alchemyMoney := alchemyTemplate.GetAchemyaMoney()
	synthetiseId = alchemyTemplate.SynthetiseId
	now := global.GetGame().GetTimeService().Now()
	num := int32(alchemyObj.Num)
	start := alchemyObj.StartTime
	passTime := now - start
	finishedNum := int32(math.Floor(float64(passTime) / float64(alchemyTemplate.GetAchemyTime())))
	leftNum = num - finishedNum
	needGold = leftNum * alchemyMoney
	return
}

//判断炼丹能否炼丹
func (pddm *PlayerDanDataManager) IsCanAlchemy(kindId int) bool {
	aObj := pddm.existKindId(kindId)
	if aObj != nil {
		return false
	}
	return true
}

//更新加速炼丹状态
func (pddm *PlayerDanDataManager) AlchemyAccelerateState(kindId int) {
	pddm.updateAlchemyState(kindId, dantypes.AlchemyStateEnd)
}

//是否完成炼丹
func (pddm *PlayerDanDataManager) IsAlchemyFinish(kindId int) bool {
	achemyItem := pddm.existKindId(kindId)
	if achemyItem == nil {
		return false
	}
	if achemyItem.State == dantypes.AlchemyStateEnd {
		return true
	}
	if achemyItem.State == dantypes.AlchemyStateStart {
		now := global.GetGame().GetTimeService().Now()
		achemyTemplate := dan.GetDanService().GetAlchemy(kindId)
		totalTime := achemyTemplate.GetAchemyTime() * int32(achemyItem.Num)
		//处理心跳处理器时间误差
		diffTime := now - achemyItem.StartTime
		if diffTime < int64(totalTime) {
			return false
		}
	}
	return true
}

//更新炼丹状态
func (pddm *PlayerDanDataManager) updateAlchemyState(kindId int, state dantypes.AlchemyState) bool {
	achemyItem := pddm.alchemyListMap[kindId]
	achemyItem.State = state
	now := global.GetGame().GetTimeService().Now()
	achemyItem.UpdateTime = now
	achemyItem.SetModified()
	return true
}

//map是否存在kindId
func (pddm *PlayerDanDataManager) existKindId(kindId int) *PlayerAlchemyObject {
	alchemyTemplate := dan.GetDanService().GetAlchemy(kindId)
	if alchemyTemplate == nil {
		return nil
	}
	if alchemyItem, ok := pddm.alchemyListMap[kindId]; ok {
		return alchemyItem
	}
	return nil
}

//gm使用  设置食丹等级
func (pddm *PlayerDanDataManager) GmSetDanLevel(danLevel int32) {
	pddm.playerDanObject.LevelId = danLevel
	pddm.clearData()
	now := global.GetGame().GetTimeService().Now()
	pddm.playerDanObject.UpdateTime = now
	pddm.playerDanObject.SetModified()
	return
}

func CreatePlayerDanDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerDanDataManager{}
	pddm.p = p
	pddm.alchemyListMap = make(map[int]*PlayerAlchemyObject)
	pddm.alchemyUsedListMap = make(map[int]*PlayerAlchemyObject)
	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerDanDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerDanDataManager))
}
