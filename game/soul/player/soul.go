package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/soul/dao"
	soulentity "fgame/fgame/game/soul/entity"
	souleventtypes "fgame/fgame/game/soul/event/types"
	"fgame/fgame/game/soul/soul"
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	soultypes "fgame/fgame/game/soul/types"
)

//帝魂对象
type PlayerSoulObject struct {
	player          player.Player
	Id              int64
	PlayerId        int64
	SoulTag         int32
	Level           int32 //吞噬等级
	Experience      int32 //吞噬进度值
	IsAwaken        int32
	AwakenOrder     int32 //技能等级
	StrengthenLevel int32 //强化等级
	StrengthenNum   int32 //当前强化等级强化次数
	StrengthenPro   int32 //强化等级进度值
	UpdateTime      int64
	CreateTime      int64
	DeleteTime      int64
}

func NewPlayerSoulObject(pl player.Player) *PlayerSoulObject {
	pso := &PlayerSoulObject{
		player: pl,
	}
	return pso
}

func (pso *PlayerSoulObject) GetPlayerId() int64 {
	return pso.PlayerId
}

func (pso *PlayerSoulObject) GetDBId() int64 {
	return pso.Id
}

func (pso *PlayerSoulObject) ToEntity() (e storage.Entity, err error) {
	e = &soulentity.PlayerSoulEntity{
		Id:              pso.Id,
		PlayerId:        pso.PlayerId,
		SoulTag:         pso.SoulTag,
		Level:           pso.Level,
		Experience:      pso.Experience,
		IsAwaken:        pso.IsAwaken,
		AwakenOrder:     pso.AwakenOrder,
		StrengthenLevel: pso.StrengthenLevel,
		StrengthenNum:   pso.StrengthenNum,
		StrengthenPro:   pso.StrengthenPro,
		UpdateTime:      pso.UpdateTime,
		CreateTime:      pso.CreateTime,
		DeleteTime:      pso.DeleteTime,
	}
	return e, err
}

func (pso *PlayerSoulObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*soulentity.PlayerSoulEntity)
	pso.Id = pse.Id
	pso.PlayerId = pse.PlayerId
	pso.SoulTag = pse.SoulTag
	pso.Level = pse.Level
	pso.Experience = pse.Experience
	pso.IsAwaken = pse.IsAwaken
	pso.AwakenOrder = pse.AwakenOrder
	pso.StrengthenLevel = pse.StrengthenLevel
	pso.StrengthenNum = pse.StrengthenNum
	pso.StrengthenPro = pse.StrengthenPro
	pso.UpdateTime = pse.UpdateTime
	pso.CreateTime = pse.CreateTime
	pso.DeleteTime = pse.DeleteTime
	return nil
}

func (pso *PlayerSoulObject) SetModified() {
	e, err := pso.ToEntity()
	if err != nil {
		panic(fmt.Errorf("Soul: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pso.player.AddChangedObject(obj)
	return
}

//帝魂镶嵌对象
type PlayerSoulEmbedObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	EmbedList  []soultypes.SoulType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSoulEmbedObject(pl player.Player) *PlayerSoulEmbedObject {
	pseo := &PlayerSoulEmbedObject{
		player: pl,
	}
	return pseo
}

func (pseo *PlayerSoulEmbedObject) GetPlayerId() int64 {
	return pseo.PlayerId
}

func (pseo *PlayerSoulEmbedObject) GetDBId() int64 {
	return pseo.Id
}

func convertNewPlayerSoulEmbedObjectToEntity(pso *PlayerSoulEmbedObject) (*soulentity.PlayerSoulEmbedEntity, error) {
	embedInfoBytes, err := json.Marshal(pso.EmbedList)
	if err != nil {
		return nil, err
	}

	e := &soulentity.PlayerSoulEmbedEntity{
		Id:         pso.Id,
		PlayerId:   pso.PlayerId,
		EmbedInfo:  string(embedInfoBytes),
		UpdateTime: pso.UpdateTime,
		CreateTime: pso.CreateTime,
		DeleteTime: pso.DeleteTime,
	}
	return e, nil
}

func (pseo *PlayerSoulEmbedObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerSoulEmbedObjectToEntity(pseo)
	return e, err
}

func (pseo *PlayerSoulEmbedObject) FromEntity(e storage.Entity) error {
	psee, _ := e.(*soulentity.PlayerSoulEmbedEntity)
	var embedList = make([]soultypes.SoulType, 0, 8)
	if err := json.Unmarshal([]byte(psee.EmbedInfo), &embedList); err != nil {
		return err
	}

	pseo.Id = psee.Id
	pseo.PlayerId = psee.PlayerId
	pseo.EmbedList = embedList
	pseo.UpdateTime = psee.UpdateTime
	pseo.CreateTime = psee.CreateTime
	pseo.DeleteTime = psee.DeleteTime
	return nil
}

func (pseo *PlayerSoulEmbedObject) SetModified() {
	e, err := pseo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("SoulEmbed: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pseo.player.AddChangedObject(obj)
	return
}

//玩家帝魂管理器
type PlayerSoulDataManager struct {
	p player.Player
	//玩家帝魂镶嵌对象
	playerSoulEmbedObject *PlayerSoulEmbedObject
	//玩家帝魂
	playerSoulObjectMap map[soultypes.SoulType]*PlayerSoulObject
}

func (psdm *PlayerSoulDataManager) Player() player.Player {
	return psdm.p
}

func (psdm *PlayerSoulDataManager) GetAllSoulSkillLevel() int32 {
	totalLevel := int32(0)
	for _, obj := range psdm.playerSoulObjectMap {
		totalLevel += obj.AwakenOrder
	}
	return totalLevel
}

//加载
func (psdm *PlayerSoulDataManager) Load() (err error) {
	//加载玩家帝魂
	souls, err := dao.GetSoulDao().GetSoulList(psdm.p.GetId())
	if err != nil {
		return
	}
	//帝魂信息
	for _, soul := range souls {
		pso := NewPlayerSoulObject(psdm.p)
		pso.FromEntity(soul)
		psdm.playerSoulObjectMap[soultypes.SoulType(pso.SoulTag)] = pso
	}

	//加载玩家镶嵌帝魂
	soulEmbedEntity, err := dao.GetSoulDao().GetSoulEmbedEntity(psdm.p.GetId())
	if err != nil {
		return
	}
	if soulEmbedEntity == nil {
		psdm.initPlayerSoulEmbedObject()
	} else {
		psdm.playerSoulEmbedObject = NewPlayerSoulEmbedObject(psdm.p)
		psdm.playerSoulEmbedObject.FromEntity(soulEmbedEntity)
	}

	return nil
}

//第一次初始化
func (psdm *PlayerSoulDataManager) initPlayerSoulEmbedObject() {
	pseo := NewPlayerSoulEmbedObject(psdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pseo.Id = id
	//生成id
	pseo.PlayerId = psdm.p.GetId()
	pseo.EmbedList = make([]soultypes.SoulType, 0, 8)
	pseo.CreateTime = now
	psdm.playerSoulEmbedObject = pseo
	pseo.SetModified()
}

//加载后
func (psdm *PlayerSoulDataManager) AfterLoad() (err error) {

	return nil
}

//心跳
func (ppdm *PlayerSoulDataManager) Heartbeat() {

}

func (psdm *PlayerSoulDataManager) GetSoulEmbedTagList() []soultypes.SoulType {
	return psdm.playerSoulEmbedObject.EmbedList
}

func (psdm *PlayerSoulDataManager) hasedEmbed(kindType soultypes.SoulKindType) (flag bool) {
	for _, soulTag := range psdm.playerSoulEmbedObject.EmbedList {
		to := soul.GetSoulService().GetSoulTemplateByLevel(soulTag, 1)
		if to == nil {
			continue
		}
		if to.Type != int32(kindType) {
			continue
		}
		flag = true
		return
	}
	return
}

//获取玩家帝魂镶嵌
func (psdm *PlayerSoulDataManager) GetSoulEmbed() (embedList []int32) {
	for _, soulTag := range psdm.playerSoulEmbedObject.EmbedList {
		soulId, flag := psdm.GetSoulIdByOrder(soulTag)
		if flag {
			embedList = append(embedList, soulId)
			continue
		}
		soulId, flag = psdm.GetSoulIdByOrder(soulTag)
		if flag {
			embedList = append(embedList, soulId)
		}
	}
	return
}

//是否穿戴指定帝魂
func (psdm *PlayerSoulDataManager) HasedEmbedSoulTag(soulTag soultypes.SoulType) (flag bool) {
	for _, embedTag := range psdm.playerSoulEmbedObject.EmbedList {
		if embedTag == soulTag {
			flag = true
			break
		}
	}
	return flag
}

//获取帝魂信息
func (psdm *PlayerSoulDataManager) GetSoulInfoAll() map[soultypes.SoulType]*PlayerSoulObject {
	return psdm.playerSoulObjectMap
}

//获取帝魂信息通过帝魂标签
func (psdm *PlayerSoulDataManager) GetSoulInfoByTag(soulTag soultypes.SoulType) *PlayerSoulObject {
	obj, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return nil
	}
	return obj
}

//获取帝魂等级通过帝魂标签
func (psdm *PlayerSoulDataManager) GetSoulLevelByTag(soulTag soultypes.SoulType) int32 {
	obj, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return 0
	}
	return obj.Level
}

//获取帝魂经验通过帝魂标签
func (psdm *PlayerSoulDataManager) GetSoulExpByTag(soulTag soultypes.SoulType) int32 {
	obj, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return 0
	}
	return obj.Experience
}

//获取帝魂阶别通过帝魂标签
func (psdm *PlayerSoulDataManager) GetSoulOrderByTag(soulTag soultypes.SoulType) int32 {
	obj, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return 0
	}
	return obj.AwakenOrder
}

//获取觉醒满阶的帝魂数
func (psdm *PlayerSoulDataManager) GetSoulNumFullAwaken() int32 {
	num := int32(0)
	for typ, soulObj := range psdm.playerSoulObjectMap {
		to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(typ, soulObj.AwakenOrder)
		if to.NextId == 0 {
			num++
		}
	}
	return num
}

//获取帝魂id
func (psdm *PlayerSoulDataManager) GetSoulIdByOrder(soulTag soultypes.SoulType) (int32, bool) {
	obj, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return 0, false
	}
	to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, obj.AwakenOrder)
	if to == nil {
		return 0, false
	}
	return int32(to.TemplateId()), true
}

//是否已激活
func (psdm *PlayerSoulDataManager) IfSoulTagExist(soulTag soultypes.SoulType) bool {
	_, ok := psdm.playerSoulObjectMap[soulTag]
	if !ok {
		return false
	}
	return true
}

//激活的前置帝魂条件
func (psdm *PlayerSoulDataManager) IfPreSoul(soulTag soultypes.SoulType, level int32) bool {
	if level <= 0 {
		return false
	}
	if !soulTag.Valid() {
		return false
	}
	if !psdm.IfSoulTagExist(soulTag) {
		return false
	}
	preSoul := psdm.playerSoulObjectMap[soulTag]
	if preSoul.Level < level {
		return false
	}
	return true
}

//是否已镶嵌
func (psdm *PlayerSoulDataManager) IfSoulTagEmemded(soulTag soultypes.SoulType) bool {
	embeds := psdm.playerSoulEmbedObject.EmbedList
	for _, tag := range embeds {
		if tag == soulTag {
			return true
		}
	}
	return false
}

//是否能喂养
func (psdm *PlayerSoulDataManager) IfCanFeed(soulTag soultypes.SoulType) bool {
	flag := psdm.IfSoulTagExist(soulTag)
	if !flag {
		return false
	}
	level := psdm.playerSoulObjectMap[soultypes.SoulType(soulTag)].Level
	levelTemplateMap := soul.GetSoulService().GetSoulTemplateByLevel(soultypes.SoulType(soulTag), level)
	if levelTemplateMap.NextId == 0 {
		return false
	}
	return true
}

//能否觉醒
func (psdm *PlayerSoulDataManager) IfIsAwaken(soulTag soultypes.SoulType) bool {
	obj := psdm.GetSoulInfoByTag(soulTag)
	if obj == nil {
		return false
	}
	if obj.IsAwaken == 1 {
		return false
	}
	return true
}

//能否升级
func (psdm *PlayerSoulDataManager) IfCanUpgrade(soulTag soultypes.SoulType) bool {
	obj := psdm.GetSoulInfoByTag(soulTag)
	if obj == nil {
		return false
	}
	if obj.AwakenOrder == 0 {
		return true
	}
	to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, obj.AwakenOrder)
	if to != nil {
		if to.NextId != 0 {
			return true
		}
	}
	return false
}

//能否强化
func (psdm *PlayerSoulDataManager) IfCanStrengthen(soulTag soultypes.SoulType) bool {
	obj := psdm.GetSoulInfoByTag(soulTag)
	if obj == nil {
		return false
	}
	pLevel := psdm.p.GetLevel()
	limitLevel := soul.GetSoulService().GetSoulStrengthenLevelLimit(pLevel)
	if obj.StrengthenLevel >= limitLevel {
		return false
	}
	to := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, obj.StrengthenLevel)
	if to != nil && to.NextId != 0 {
		return true
	}
	return false
}

//帝魂激活
func (psdm *PlayerSoulDataManager) SoulActive(soulTag soultypes.SoulType) (*PlayerSoulObject, bool, bool) {
	autoEmbed := false
	flag := soulTag.Valid()
	if !flag {
		return nil, autoEmbed, false
	}
	flag = psdm.IfSoulTagExist(soulTag)
	if flag {
		return nil, autoEmbed, false
	}
	id, err := idutil.GetId()
	if err != nil {
		return nil, autoEmbed, false
	}

	now := global.GetGame().GetTimeService().Now()
	pso := NewPlayerSoulObject(psdm.p)
	pso.Id = id
	pso.PlayerId = psdm.p.GetId()
	pso.SoulTag = int32(soulTag)
	pso.Level = int32(1)
	pso.Experience = int32(0)
	pso.IsAwaken = int32(0)
	pso.AwakenOrder = int32(1)
	pso.StrengthenLevel = int32(1)
	pso.StrengthenNum = int32(0)
	pso.StrengthenPro = int32(0)
	pso.CreateTime = now
	pso.SetModified()
	psdm.playerSoulObjectMap[soulTag] = pso

	//激活自动穿戴
	to := soul.GetSoulService().GetSoulTemplateByLevel(soulTag, pso.Level)
	if to != nil && to.Type == int32(soultypes.SoulKindTypeAttack) {
		if !psdm.hasedEmbed(soultypes.SoulKindType(to.Type)) {
			autoEmbed = true
			psdm.playerSoulEmbedObject.EmbedList = append(psdm.playerSoulEmbedObject.EmbedList, soulTag)
			psdm.playerSoulEmbedObject.UpdateTime = now
			psdm.playerSoulEmbedObject.SetModified()

			oldTag := soultypes.SoulType(-1)
			eventData := souleventtypes.CreateSoulEmbedEventData(oldTag, soulTag)
			gameevent.Emit(souleventtypes.EventTypeSoulEmbed, psdm.p, eventData)
		}
	}
	gameevent.Emit(souleventtypes.EventTypeSoulActive, psdm.p, soulTag)
	return pso, autoEmbed, true
}

//帝魂镶嵌
func (psdm *PlayerSoulDataManager) Embed(soulTag soultypes.SoulType) bool {
	flag := soulTag.Valid()
	if !flag {
		return false
	}
	flag = psdm.IfSoulTagExist(soulTag)
	if !flag {
		return false
	}
	flag = psdm.IfSoulTagEmemded(soulTag)
	if flag {
		return false
	}

	oldTag := soultypes.SoulType(-1)
	//获取帝魂种类
	soulKindType := soul.GetSoulService().GetSoulKindTemplate(soulTag)
	//更换镶嵌
	addFlag := true
	for index, tag := range psdm.playerSoulEmbedObject.EmbedList {
		kindType := soul.GetSoulService().GetSoulKindTemplate(tag)
		if kindType == soulKindType {
			oldTag = tag
			psdm.playerSoulEmbedObject.EmbedList[index] = soulTag
			addFlag = false
			break
		}
	}
	//添加镶嵌
	if addFlag {
		psdm.playerSoulEmbedObject.EmbedList = append(psdm.playerSoulEmbedObject.EmbedList, soulTag)
	}

	now := global.GetGame().GetTimeService().Now()
	psdm.playerSoulEmbedObject.UpdateTime = now
	psdm.playerSoulEmbedObject.SetModified()

	eventData := souleventtypes.CreateSoulEmbedEventData(oldTag, soulTag)
	gameevent.Emit(souleventtypes.EventTypeSoulEmbed, psdm.p, eventData)
	return true
}

//帝魂升级+1
func (psdm *PlayerSoulDataManager) Upgrade(soulTag soultypes.SoulType) bool {
	obj := psdm.GetSoulInfoByTag(soulTag)
	if obj == nil {
		return false
	}
	flag := psdm.IfCanUpgrade(soulTag)
	if !flag {
		return false
	}
	oldOrder := obj.AwakenOrder
	now := global.GetGame().GetTimeService().Now()
	obj.AwakenOrder += 1
	obj.UpdateTime = now
	obj.SetModified()
	eventData := souleventtypes.CreateSoulUpgradeEventData(soulTag, oldOrder, obj.AwakenOrder)
	gameevent.Emit(souleventtypes.EventTypeSoulUpgrade, psdm.p, eventData)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, psdm.p, nil)
	return true
}

//帝魂觉醒
func (psdm *PlayerSoulDataManager) Awaken(soulTag soultypes.SoulType) bool {
	obj := psdm.GetSoulInfoByTag(soulTag)
	if obj == nil {
		return false
	}
	flag := psdm.IfIsAwaken(soulTag)
	if !flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj.IsAwaken = 1
	obj.UpdateTime = now
	obj.SetModified()
	gameevent.Emit(souleventtypes.EventTypeSoulAwaken, psdm.p, soulTag)
	return true
}

//喂养能升几级
func (psdm *PlayerSoulDataManager) UpgradesByFeed(soulTag soultypes.SoulType, totalExp int32) bool {
	if totalExp <= 0 {
		return false
	}
	flag := psdm.IfCanFeed(soulTag)
	if !flag {
		return false
	}

	obj := psdm.GetSoulInfoByTag(soulTag)
	curExp := obj.Experience
	curLevel := obj.Level
	totalExp += curExp
	now := global.GetGame().GetTimeService().Now()
	to, exp := soul.GetSoulService().GetSoulTemplateByExp(soulTag, curLevel, totalExp)
	if to != nil {
		obj.Level = to.Level
		obj.Experience = exp
		obj.UpdateTime = now
		obj.SetModified()
	}
	return true
}

//获取技能
func (psdm *PlayerSoulDataManager) GetSkillId(soulTag soultypes.SoulType, order int32) (skillId int32) {
	skillId = int32(0)
	soulInfo := psdm.GetSoulInfoByTag(soulTag)
	to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, order)
	if to != nil {
		//觉醒技能
		if soulInfo.IsAwaken == 1 {
			skillId = to.SkillId
		} else { //升级技能
			skillId = to.UplevelSkillId
		}
	}

	return skillId
}

//帝魂强化
func (psdm *PlayerSoulDataManager) SoulStrengthen(soulTag soultypes.SoulType, pro int32, sucess bool) {
	if pro < 0 {
		return
	}
	soulInfo := psdm.GetSoulInfoByTag(soulTag)
	if soulInfo == nil {
		return
	}

	if sucess {
		nextStrengthenLevel := soulInfo.StrengthenLevel + 1
		strengthenTemplate := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, nextStrengthenLevel)
		if strengthenTemplate == nil {
			return
		}
		soulInfo.StrengthenLevel += 1
		soulInfo.StrengthenNum = 0
		soulInfo.StrengthenPro = pro
		gameevent.Emit(souleventtypes.EventTypeSoulStrengthen, psdm.p, soulTag)
	} else {
		soulInfo.StrengthenNum += 1
		soulInfo.StrengthenPro += pro
	}

	now := global.GetGame().GetTimeService().Now()
	soulInfo.UpdateTime = now
	soulInfo.SetModified()
	return
}

func (pwdm *PlayerSoulDataManager) GetSoulNum() int32 {
	return int32(len(pwdm.playerSoulObjectMap))
}

//获取帝魂觉醒数量
func (pwdm *PlayerSoulDataManager) GetAwakenNum() int32 {
	num := int32(0)
	for _, tempSoul := range pwdm.playerSoulObjectMap {
		if tempSoul.IsAwaken != 0 {
			num++
		}
	}
	return num
}

//帝魂信息
func (pwdm *PlayerSoulDataManager) ToAllSoulInfo() (allSoulInfo *soultypes.AllSoulInfo) {
	allSoulInfo = &soultypes.AllSoulInfo{}
	soulEmbedList := pwdm.GetSoulEmbed()
	allSoulInfo.EmbedIdList = append(allSoulInfo.EmbedIdList, soulEmbedList...)
	for _, tempSoul := range pwdm.playerSoulObjectMap {
		soulInfo := &soultypes.SoulInfo{
			SoulTag:        tempSoul.SoulTag,
			Level:          tempSoul.Level,
			Experience:     tempSoul.Experience,
			AwakenOrder:    tempSoul.AwakenOrder,
			IsAwaken:       tempSoul.IsAwaken,
			StrengthenLevl: tempSoul.StrengthenLevel,
			StrengthenPro:  tempSoul.StrengthenPro,
		}
		allSoulInfo.SoulList = append(allSoulInfo.SoulList, soulInfo)
	}

	return
}

//所有的帝魂强化是否全满
func (pwdm *PlayerSoulDataManager) IfAllStrengthenFull() (flag bool) {
	for soulTag := soultypes.SoulTypeMin; soulTag <= soultypes.SoulTypeMax; soulTag++ {
		soulObj, ok := pwdm.playerSoulObjectMap[soulTag]
		if !ok {
			return
		}
		nextLevel := soulObj.StrengthenLevel + 1
		strengthenTemplate := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, nextLevel)
		if strengthenTemplate != nil {
			return
		}
	}
	flag = true
	return
}

//所有的帝魂升级是否全满
func (pwdm *PlayerSoulDataManager) IfAllUpgradeFull() (flag bool) {
	for soulTag := soultypes.SoulTypeMin; soulTag <= soultypes.SoulTypeMax; soulTag++ {
		soulObj, ok := pwdm.playerSoulObjectMap[soulTag]
		if !ok {
			return
		}
		nextOrder := soulObj.AwakenOrder + 1
		to := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, nextOrder)
		if to != nil {
			return
		}
	}
	flag = true
	return
}

//仅Gm使用 帝魂激活 &&镶嵌
func (psdm *PlayerSoulDataManager) GmSoulEmbed(soulTag soultypes.SoulType) bool {
	now := global.GetGame().GetTimeService().Now()
	flag := psdm.IfSoulTagEmemded(soulTag)
	if flag {
		return true
	}

	flag = psdm.IfSoulTagExist(soulTag)
	if !flag {
		id, err := idutil.GetId()
		if err != nil {
			return false
		}

		pso := NewPlayerSoulObject(psdm.p)
		pso.Id = id
		pso.PlayerId = psdm.p.GetId()
		pso.SoulTag = int32(soulTag)
		pso.Level = int32(1)
		pso.Experience = int32(0)
		pso.IsAwaken = int32(0)
		pso.AwakenOrder = int32(1)
		pso.StrengthenLevel = int32(1)
		pso.StrengthenNum = int32(0)
		pso.StrengthenPro = int32(0)
		pso.CreateTime = now
		pso.SetModified()
		psdm.playerSoulObjectMap[soulTag] = pso
	}

	oldTag := soultypes.SoulType(-1)
	//获取帝魂种类
	soulKindType := soul.GetSoulService().GetSoulKindTemplate(soulTag)
	//更换镶嵌
	addFlag := true
	for index, tag := range psdm.playerSoulEmbedObject.EmbedList {
		kindType := soul.GetSoulService().GetSoulKindTemplate(tag)
		if kindType == soulKindType {
			oldTag = tag
			psdm.playerSoulEmbedObject.EmbedList[index] = soulTag
			addFlag = false
			break
		}
	}
	//添加镶嵌
	if addFlag {
		psdm.playerSoulEmbedObject.EmbedList = append(psdm.playerSoulEmbedObject.EmbedList, soulTag)
	}
	psdm.playerSoulEmbedObject.UpdateTime = now
	psdm.playerSoulEmbedObject.SetModified()
	eventData := souleventtypes.CreateSoulEmbedEventData(oldTag, soulTag)
	gameevent.Emit(souleventtypes.EventTypeSoulEmbed, psdm.p, eventData)

	return true
}

//仅Gm使用 帝魂魂技升级
func (psdm *PlayerSoulDataManager) GmSoulAwakenOrder(soulTag soultypes.SoulType, order int32) {
	soulInfo := psdm.GetSoulInfoByTag(soulTag)
	if soulInfo == nil {
		return
	}
	oldOrder := soulInfo.AwakenOrder
	now := global.GetGame().GetTimeService().Now()
	soulInfo.AwakenOrder = order
	soulInfo.UpdateTime = now
	soulInfo.SetModified()
	eventData := souleventtypes.CreateSoulUpgradeEventData(soulTag, oldOrder, soulInfo.AwakenOrder)
	gameevent.Emit(souleventtypes.EventTypeSoulUpgrade, psdm.p, eventData)
	return
}

//仅Gm使用 帝魂觉醒
func (psdm *PlayerSoulDataManager) GmSoulAwaken(soulTag soultypes.SoulType) {
	soulInfo := psdm.GetSoulInfoByTag(soulTag)
	if soulInfo == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	soulInfo.IsAwaken = 1
	soulInfo.UpdateTime = now
	soulInfo.SetModified()

	gameevent.Emit(souleventtypes.EventTypeSoulAwaken, psdm.p, soulTag)
	return
}

//仅Gm使用 帝魂强化
func (psdm *PlayerSoulDataManager) GmSoulStrengthenLevel(soulTag soultypes.SoulType, strenghthenLevel int32) {
	soulInfo := psdm.GetSoulInfoByTag(soulTag)
	if soulInfo == nil {
		return
	}

	soulInfo.StrengthenLevel = strenghthenLevel
	soulInfo.StrengthenNum = 0
	soulInfo.StrengthenPro = 0

	now := global.GetGame().GetTimeService().Now()
	soulInfo.UpdateTime = now
	soulInfo.SetModified()
	gameevent.Emit(souleventtypes.EventTypeSoulStrengthen, psdm.p, soulTag)
}

func CreatePlayerSoulDataManager(p player.Player) player.PlayerDataManager {
	psdm := &PlayerSoulDataManager{}
	psdm.p = p
	psdm.playerSoulObjectMap = make(map[soultypes.SoulType]*PlayerSoulObject)
	return psdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSoulDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSoulDataManager))
}
