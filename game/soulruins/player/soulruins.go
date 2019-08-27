package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/soulruins/dao"
	soulruinsentity "fgame/fgame/game/soulruins/entity"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
	"fgame/fgame/game/soulruins/soulruins"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	soulruinstypes "fgame/fgame/game/soulruins/types"
)

//帝陵遗迹对象
type PlayerSoulRuinsObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Chapter    int32
	Type       soulruinstypes.SoulRuinsType
	Level      int32
	Star       int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSoulRuinsObject(pl player.Player) *PlayerSoulRuinsObject {
	pso := &PlayerSoulRuinsObject{
		player: pl,
	}
	return pso
}

func (psro *PlayerSoulRuinsObject) GetPlayerId() int64 {
	return psro.PlayerId
}

func (psro *PlayerSoulRuinsObject) GetDBId() int64 {
	return psro.Id
}

func (psro *PlayerSoulRuinsObject) ToEntity() (e storage.Entity, err error) {
	e = &soulruinsentity.PlayerSoulRuinsEntity{
		Id:         psro.Id,
		PlayerId:   psro.PlayerId,
		Chapter:    psro.Chapter,
		Type:       int32(psro.Type),
		Level:      psro.Level,
		Star:       psro.Star,
		UpdateTime: psro.UpdateTime,
		CreateTime: psro.CreateTime,
		DeleteTime: psro.DeleteTime,
	}
	return e, err
}

func (psro *PlayerSoulRuinsObject) FromEntity(e storage.Entity) error {
	psre, _ := e.(*soulruinsentity.PlayerSoulRuinsEntity)
	psro.Id = psre.Id
	psro.PlayerId = psre.PlayerId
	psro.Chapter = psre.Chapter
	psro.Type = soulruinstypes.SoulRuinsType(psre.Type)
	psro.Level = psre.Level
	psro.Star = psre.Star
	psro.UpdateTime = psre.UpdateTime
	psro.CreateTime = psre.CreateTime
	psro.DeleteTime = psre.DeleteTime
	return nil
}

func (psro *PlayerSoulRuinsObject) SetModified() {
	e, err := psro.ToEntity()
	if err != nil {
		panic(fmt.Errorf("SoulRuins: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psro.player.AddChangedObject(obj)
	return
}

//帝陵遗迹挑战次数对象
type PlayerSoulRuinsNumObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	Num         int32
	ExtraBuyNum int32
	RewNum      int32
	UsedNum     int32
	UsedBuyNum  int32
	UsedRewNum  int32
	BuyNum      int32
	LastTime    int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerSoulRuinsNumObject(pl player.Player) *PlayerSoulRuinsNumObject {
	pseo := &PlayerSoulRuinsNumObject{
		player: pl,
	}
	return pseo
}

func (psrno *PlayerSoulRuinsNumObject) GetPlayerId() int64 {
	return psrno.PlayerId
}

func (psrno *PlayerSoulRuinsNumObject) GetDBId() int64 {
	return psrno.Id
}

func (psrno *PlayerSoulRuinsNumObject) ToEntity() (e storage.Entity, err error) {

	e = &soulruinsentity.PlayerSoulRuinsNumEntity{
		Id:          psrno.Id,
		PlayerId:    psrno.PlayerId,
		Num:         psrno.Num,
		ExtraBuyNum: psrno.ExtraBuyNum,
		RewNum:      psrno.RewNum,
		UsedNum:     psrno.UsedNum,
		UsedBuyNum:  psrno.UsedBuyNum,
		UsedRewNum:  psrno.UsedRewNum,
		BuyNum:      psrno.BuyNum,
		LastTime:    psrno.LastTime,
		UpdateTime:  psrno.UpdateTime,
		CreateTime:  psrno.CreateTime,
		DeleteTime:  psrno.DeleteTime,
	}
	return e, err
}

func (psrno *PlayerSoulRuinsNumObject) FromEntity(e storage.Entity) error {
	psrne, _ := e.(*soulruinsentity.PlayerSoulRuinsNumEntity)

	psrno.Id = psrne.Id
	psrno.PlayerId = psrne.PlayerId
	psrno.Num = psrne.Num
	psrno.ExtraBuyNum = psrne.ExtraBuyNum
	psrno.RewNum = psrne.RewNum
	psrno.UsedNum = psrne.UsedNum
	psrno.UsedBuyNum = psrne.UsedBuyNum
	psrno.UsedRewNum = psrne.UsedRewNum
	psrno.BuyNum = psrne.BuyNum
	psrno.LastTime = psrne.LastTime
	psrno.UpdateTime = psrne.UpdateTime
	psrno.CreateTime = psrne.CreateTime
	psrno.DeleteTime = psrne.DeleteTime
	return nil
}

func (psrno *PlayerSoulRuinsNumObject) SetModified() {
	e, err := psrno.ToEntity()
	if err != nil {
		panic(fmt.Errorf("SoulRuins: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psrno.player.AddChangedObject(obj)
	return
}

//玩家帝魂遗迹章节奖励对象
type PlayerSoulRuinsRewChapterObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Chapter    int32
	Type       soulruinstypes.SoulRuinsType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerSoulRuinsRewChapterObject(pl player.Player) *PlayerSoulRuinsRewChapterObject {
	pseo := &PlayerSoulRuinsRewChapterObject{
		player: pl,
	}
	return pseo
}

func (psrrco *PlayerSoulRuinsRewChapterObject) GetPlayerId() int64 {
	return psrrco.PlayerId
}

func (psrrco *PlayerSoulRuinsRewChapterObject) GetDBId() int64 {
	return psrrco.Id
}

func (psrrco *PlayerSoulRuinsRewChapterObject) ToEntity() (e storage.Entity, err error) {

	e = &soulruinsentity.PlayerSoulRuinsRewChapterEntity{
		Id:         psrrco.Id,
		PlayerId:   psrrco.PlayerId,
		Chapter:    psrrco.Chapter,
		Type:       int32(psrrco.Type),
		UpdateTime: psrrco.UpdateTime,
		CreateTime: psrrco.CreateTime,
		DeleteTime: psrrco.DeleteTime,
	}
	return e, err
}

func (psrrco *PlayerSoulRuinsRewChapterObject) FromEntity(e storage.Entity) error {
	psrcre, _ := e.(*soulruinsentity.PlayerSoulRuinsRewChapterEntity)

	psrrco.Id = psrcre.Id
	psrrco.PlayerId = psrcre.PlayerId
	psrrco.Chapter = psrcre.Chapter
	psrrco.Type = soulruinstypes.SoulRuinsType(psrcre.Type)
	psrrco.UpdateTime = psrcre.UpdateTime
	psrrco.CreateTime = psrcre.CreateTime
	psrrco.DeleteTime = psrcre.DeleteTime
	return nil
}

func (psrrco *PlayerSoulRuinsRewChapterObject) SetModified() {
	e, err := psrrco.ToEntity()
	if err != nil {
		panic(fmt.Errorf("SoulRuins: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	psrrco.player.AddChangedObject(obj)
	return
}

//玩家帝陵遗迹管理器
type PlayerSoulRuinsDataManager struct {
	p player.Player
	//玩家帝陵遗迹对象
	playerSoulRuinsMap map[int32]map[soulruinstypes.SoulRuinsType]map[int32]*PlayerSoulRuinsObject
	//玩家帝陵遗迹挑战次数
	playerSoulRuinsNumObject *PlayerSoulRuinsNumObject
	//玩家帝陵遗迹奖励章节
	playerSoulRuinsRewChapterMap map[int32]map[soulruinstypes.SoulRuinsType]*PlayerSoulRuinsRewChapterObject
	//玩家帝陵遗迹章节星数
	playerSoulRuinsChapterStarMap map[int32]map[soulruinstypes.SoulRuinsType]int32
}

func (psrdm *PlayerSoulRuinsDataManager) Player() player.Player {
	return psrdm.p
}

//加载
func (psrdm *PlayerSoulRuinsDataManager) Load() (err error) {
	psrdm.playerSoulRuinsMap = make(map[int32]map[soulruinstypes.SoulRuinsType]map[int32]*PlayerSoulRuinsObject)
	psrdm.playerSoulRuinsRewChapterMap = make(map[int32]map[soulruinstypes.SoulRuinsType]*PlayerSoulRuinsRewChapterObject)
	psrdm.playerSoulRuinsChapterStarMap = make(map[int32]map[soulruinstypes.SoulRuinsType]int32)
	//加载玩家帝陵遗迹
	soulRuinsList, err := dao.GetSoulRuinsDao().GetSoulRuinsList(psrdm.p.GetId())
	if err != nil {
		return
	}
	//帝陵遗迹信息
	for _, soulRuins := range soulRuinsList {
		psro := NewPlayerSoulRuinsObject(psrdm.p)
		psro.FromEntity(soulRuins)
		psrdm.addSoulRuins(psro)
	}

	//加载玩家帝陵遗迹章节奖励列表
	soulRuinsRewChapterList, err := dao.GetSoulRuinsDao().GetSoulRuinsRewChapterList(psrdm.p.GetId())
	if err != nil {
		return
	}
	for _, chapterRew := range soulRuinsRewChapterList {
		psrrco := NewPlayerSoulRuinsRewChapterObject(psrdm.p)
		psrrco.FromEntity(chapterRew)
		psrdm.addSoulRuinsRewChapter(psrrco)
	}

	//加载玩家帝陵遗迹挑战次数
	soulRuinsNumEntity, err := dao.GetSoulRuinsDao().GetSoulRuinsNumEntity(psrdm.p.GetId())
	if err != nil {
		return
	}
	if soulRuinsNumEntity == nil {
		psrdm.initPlayerSoulRuinsNumObject()
	} else {
		psrdm.playerSoulRuinsNumObject = NewPlayerSoulRuinsNumObject(psrdm.p)
		psrdm.playerSoulRuinsNumObject.FromEntity(soulRuinsNumEntity)
	}

	return nil
}

//第一次初始化
func (psrdm *PlayerSoulRuinsDataManager) initPlayerSoulRuinsNumObject() {
	psrneo := NewPlayerSoulRuinsNumObject(psrdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	psrneo.Id = id
	//生成id
	psrneo.PlayerId = psrdm.p.GetId()
	psrneo.Num = 0
	psrneo.BuyNum = 0
	psrneo.LastTime = 0
	psrneo.CreateTime = now
	psrdm.playerSoulRuinsNumObject = psrneo
	psrneo.SetModified()
}

//加载后
func (psrdm *PlayerSoulRuinsDataManager) AfterLoad() (err error) {
	err = psrdm.refreshSoulRuinsNum()
	return
}

//心跳
func (psrdm *PlayerSoulRuinsDataManager) Heartbeat() {

}

//刷新帝陵遗迹次数
func (psrdm *PlayerSoulRuinsDataManager) refreshSoulRuinsNum() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := psrdm.playerSoulRuinsNumObject.LastTime
	if lastTime != 0 {
		//flag, err := timeutils.IsSameDay(lastTime, now)
		flag, err := timeutils.IsSameFive(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			psrdm.playerSoulRuinsNumObject.Num = 0
			psrdm.playerSoulRuinsNumObject.BuyNum = 0
			psrdm.playerSoulRuinsNumObject.UsedNum = 0
			psrdm.playerSoulRuinsNumObject.RewNum -= psrdm.playerSoulRuinsNumObject.UsedRewNum
			psrdm.playerSoulRuinsNumObject.UsedRewNum = 0
			psrdm.playerSoulRuinsNumObject.ExtraBuyNum -= psrdm.playerSoulRuinsNumObject.UsedBuyNum
			psrdm.playerSoulRuinsNumObject.UsedBuyNum = 0
			psrdm.playerSoulRuinsNumObject.LastTime = 0
			psrdm.playerSoulRuinsNumObject.UpdateTime = now
			psrdm.playerSoulRuinsNumObject.SetModified()
		}
	}
	return nil
}

//获取挑战次数和购买次数(外部调用)
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsNum() *PlayerSoulRuinsNumObject {
	//刷新屠魔次数
	psrdm.refreshSoulRuinsNum()
	return psrdm.getSoulRuinsNum()
}

//获取挑战次数和购买次数
func (psrdm *PlayerSoulRuinsDataManager) getSoulRuinsNum() *PlayerSoulRuinsNumObject {
	return psrdm.playerSoulRuinsNumObject
}

//获取帝陵遗迹奖励挑战次数未使用
func (psrdm *PlayerSoulRuinsDataManager) getSoulRuinsLeftRewNum() int32 {
	numObj := psrdm.getSoulRuinsNum()
	return numObj.RewNum - numObj.UsedRewNum
}

//获取帝陵遗迹额外购买挑战次数未使用
func (psrdm *PlayerSoulRuinsDataManager) getSoulRuinsLeftBuyNum() int32 {
	numObj := psrdm.getSoulRuinsNum()
	return numObj.ExtraBuyNum - numObj.UsedBuyNum
}

//获取领取星级奖励记录map
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsRewChapterMap() map[int32]map[soulruinstypes.SoulRuinsType]*PlayerSoulRuinsRewChapterObject {
	return psrdm.playerSoulRuinsRewChapterMap
}

//获取领取星级奖励记录
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsRewChapter(chapter int32, typ soulruinstypes.SoulRuinsType) *PlayerSoulRuinsRewChapterObject {
	rewChapterTypeMap, ok := psrdm.playerSoulRuinsRewChapterMap[chapter]
	if !ok {
		return nil
	}
	rewChapter, ok := rewChapterTypeMap[typ]
	if !ok {
		return nil
	}
	return rewChapter
}

//获取通过过帝陵遗迹
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsMap() map[int32]map[soulruinstypes.SoulRuinsType]map[int32]*PlayerSoulRuinsObject {
	return psrdm.playerSoulRuinsMap
}

//获取章节的总星数
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsChapterStar(chapter int32, typ soulruinstypes.SoulRuinsType) int32 {
	chapterStarTypeMap, ok := psrdm.playerSoulRuinsChapterStarMap[chapter]
	if !ok {
		return 0
	}
	chapterStar, ok := chapterStarTypeMap[typ]
	if !ok {
		return 0
	}
	return chapterStar
}

//获取通关关卡对象
func (psrdm *PlayerSoulRuinsDataManager) getSoulRuinsLevelInfo(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) *PlayerSoulRuinsObject {
	soulRuinsMap := psrdm.GetSoulRuinsMap()
	soulRuinsTypeMap, ok := soulRuinsMap[chapter]
	if !ok {
		return nil
	}
	soulRuinsLevelMap, ok := soulRuinsTypeMap[typ]
	if !ok {
		return nil
	}
	soulRuinsObj, ok := soulRuinsLevelMap[level]
	if !ok {
		return nil
	}
	return soulRuinsObj
}

//获取通关关卡获取星数
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsLevelStar(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) int32 {
	soulRuinsObj := psrdm.getSoulRuinsLevelInfo(chapter, typ, level)
	if soulRuinsObj == nil {
		return 0
	}
	return soulRuinsObj.Star
}

//参数有效性
func (psrdm *PlayerSoulRuinsDataManager) IsValid(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) bool {
	to := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	if to == nil {
		return false
	}
	return true
}

//参数有效性
func (psrdm *PlayerSoulRuinsDataManager) IsChapterAndTypValid(chapter int32, typ soulruinstypes.SoulRuinsType) bool {
	soulRuinsTypeMap, ok := psrdm.playerSoulRuinsMap[chapter]
	if !ok {
		return false
	}
	_, ok = soulRuinsTypeMap[typ]
	if !ok {
		return false
	}
	return true
}

//挑战次数是否足够
func (psrdm *PlayerSoulRuinsDataManager) HasEnoughChallengeNum(needNum int32) bool {
	defaultNum := soulruins.GetSoulRuinsService().GetSoulRuinsChallengeNum()
	numObj := psrdm.GetSoulRuinsNum()
	leftNum := defaultNum + numObj.ExtraBuyNum + numObj.RewNum - numObj.Num
	if leftNum >= needNum {
		return true
	}
	return false
}

//前置是否通关
func (psrdm *PlayerSoulRuinsDataManager) IfPreSoulRuinsPassed(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) bool {
	to := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	if to == nil {
		return false
	}
	frontId := to.FrontId
	if frontId == 0 {
		return true
	}
	preTo := soulruins.GetSoulRuinsService().GetSoulRuinsTemplateById(frontId)
	bChapter := preTo.Chapter
	bTyp := preTo.GetType()
	bLevel := preTo.Level
	flag := psrdm.IfSoulRuinsExist(bChapter, bTyp, bLevel)
	if !flag {
		return false
	}
	return true
}

//是否存在帝陵遗迹
func (psrdm *PlayerSoulRuinsDataManager) IfSoulRuinsExist(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) bool {
	chapterMap, ok := psrdm.playerSoulRuinsMap[chapter]
	if !ok {
		return false
	}
	levelMap, ok := chapterMap[typ]
	if !ok {
		return false
	}
	_, ok = levelMap[level]
	if !ok {
		return false
	}
	return true
}

//星级奖励是否领取过
func (psrdm *PlayerSoulRuinsDataManager) IfSoulRuinsRewReceived(chapter int32, typ soulruinstypes.SoulRuinsType) bool {
	rewTypMap, ok := psrdm.playerSoulRuinsRewChapterMap[chapter]
	if !ok {
		return false
	}
	_, ok = rewTypMap[typ]
	if !ok {
		return false
	}
	return true
}

//章节星级数是否达标
func (psrdm *PlayerSoulRuinsDataManager) IfStarsReachStandard(chapter int32, typ soulruinstypes.SoulRuinsType) bool {
	flag := psrdm.IsChapterAndTypValid(chapter, typ)
	if !flag {
		return false
	}
	soulRuinsStarTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsStarTemplate(chapter, typ)
	needStar := soulRuinsStarTemplate.NeedStar

	curStar := psrdm.GetSoulRuinsChapterStar(chapter, typ)
	if curStar < needStar {
		return false
	}
	return true
}

//购买挑战次数是否达上限
func (psrdm *PlayerSoulRuinsDataManager) IfBuyNumReachLimit(buyNum int32) bool {
	numObj := psrdm.GetSoulRuinsNum()
	curBuyNum := numObj.BuyNum
	maxBuyNum := soulruins.GetSoulRuinsService().GetSoulRuinsBuyChallengeNum()
	num := maxBuyNum - curBuyNum - buyNum
	if num < 0 {
		return true
	}
	return false
}

//刷新帝陵遗迹
func (psrdm *PlayerSoulRuinsDataManager) RefreshSoulRuins(chapter int32, typ soulruinstypes.SoulRuinsType, level int32, star int32, addNumFlag bool) {
	if star < soulruinstypes.MinStar || star > soulruinstypes.MaxStar {
		return
	}
	to := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	if to == nil {
		return
	}
	rewTime := to.RewTime
	soulRuinsObj := psrdm.getSoulRuinsLevelInfo(chapter, typ, level)

	//首次通关
	if soulRuinsObj == nil {
		psro := NewPlayerSoulRuinsObject(psrdm.p)
		now := global.GetGame().GetTimeService().Now()
		id, _ := idutil.GetId()
		psro.Id = id
		//生成id
		psro.PlayerId = psrdm.p.GetId()
		psro.Chapter = chapter
		psro.Type = typ
		psro.Level = level
		psro.Star = star
		psro.CreateTime = now
		psro.SetModified()
		psrdm.addSoulRuins(psro)
		//奖励挑战次数
		if addNumFlag {
			psrdm.addSoulRuinsRewNum(rewTime)
		}
		return
	}
	//更新星数
	oldStar := soulRuinsObj.Star
	if star < oldStar {
		return
	}
	addStar := star - oldStar
	psrdm.updateSoulRuinsStar(soulRuinsObj, addStar)
	return
}

//章节星级奖励领取记录
func (psrdm *PlayerSoulRuinsDataManager) RewReceiveChapter(chapter int32, typ soulruinstypes.SoulRuinsType) bool {
	flag := psrdm.IsChapterAndTypValid(chapter, typ)
	if !flag {
		return false
	}
	flag = psrdm.IfSoulRuinsRewReceived(chapter, typ)
	if flag {
		return false
	}
	flag = psrdm.IfStarsReachStandard(chapter, typ)
	if !flag {
		return false
	}

	psrrco := NewPlayerSoulRuinsRewChapterObject(psrdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	psrrco.Id = id
	//生成id
	psrrco.PlayerId = psrdm.p.GetId()
	psrrco.Chapter = chapter
	psrrco.Type = typ
	psrrco.CreateTime = now
	psrrco.SetModified()

	psrdm.addSoulRuinsRewChapter(psrrco)
	return true
}

//添加帝魂遗迹
func (psrdm *PlayerSoulRuinsDataManager) addSoulRuins(psro *PlayerSoulRuinsObject) {
	chapter := psro.Chapter
	typ := psro.Type
	level := psro.Level
	star := psro.Star
	soulRuinsTypeMap, ok := psrdm.playerSoulRuinsMap[chapter]
	if !ok {
		soulRuinsTypeMap = make(map[soulruinstypes.SoulRuinsType]map[int32]*PlayerSoulRuinsObject)
		psrdm.playerSoulRuinsMap[chapter] = soulRuinsTypeMap
	}
	soulRuinsLevelMap, ok := soulRuinsTypeMap[typ]
	if !ok {
		soulRuinsLevelMap = make(map[int32]*PlayerSoulRuinsObject)
		soulRuinsTypeMap[typ] = soulRuinsLevelMap
	}
	soulRuinsLevelMap[level] = psro
	psrdm.addSoulRuinsChapterStar(chapter, typ, star)
}

//添加帝陵遗迹章节奖励
func (psrdm *PlayerSoulRuinsDataManager) addSoulRuinsRewChapter(psrrco *PlayerSoulRuinsRewChapterObject) {
	chapter := psrrco.Chapter
	typ := psrrco.Type
	soulRuinsChapterRewTypeMap, ok := psrdm.playerSoulRuinsRewChapterMap[chapter]
	if !ok {
		soulRuinsChapterRewTypeMap = make(map[soulruinstypes.SoulRuinsType]*PlayerSoulRuinsRewChapterObject)
		psrdm.playerSoulRuinsRewChapterMap[chapter] = soulRuinsChapterRewTypeMap
	}
	soulRuinsChapterRewTypeMap[typ] = psrrco
}

//增加章节的星数
func (psrdm *PlayerSoulRuinsDataManager) addSoulRuinsChapterStar(chapter int32, typ soulruinstypes.SoulRuinsType, addStar int32) {
	if addStar < soulruinstypes.MinStar || addStar > soulruinstypes.MaxStar {
		return
	}
	chapterStarTypeMap, ok := psrdm.playerSoulRuinsChapterStarMap[chapter]
	if !ok {
		chapterStarTypeMap = make(map[soulruinstypes.SoulRuinsType]int32)
		psrdm.playerSoulRuinsChapterStarMap[chapter] = chapterStarTypeMap
		chapterStarTypeMap[typ] = addStar
	} else {
		_, ok := chapterStarTypeMap[typ]
		if !ok {
			chapterStarTypeMap[typ] = addStar
		} else {
			chapterStarTypeMap[typ] += addStar
		}

	}
}

//更新帝魂遗迹关卡星数
func (psrdm *PlayerSoulRuinsDataManager) updateSoulRuinsStar(psro *PlayerSoulRuinsObject, addStar int32) {
	now := global.GetGame().GetTimeService().Now()
	psro.Star += addStar
	psro.UpdateTime = now
	psro.SetModified()

	chapter := psro.Chapter
	typ := psro.Type
	psrdm.addSoulRuinsChapterStar(chapter, typ, addStar)
	return
}

//首次挑战成功奖励挑战次数
func (psrdm *PlayerSoulRuinsDataManager) addSoulRuinsRewNum(rewTime int32) {
	now := global.GetGame().GetTimeService().Now()
	numObj := psrdm.getSoulRuinsNum()
	numObj.RewNum += rewTime
	numObj.UpdateTime = now
	numObj.SetModified()
	return
}

//购买挑战次数
func (psrdm *PlayerSoulRuinsDataManager) AddSoulRuinsBuyNum(buyNum int32) {
	if buyNum <= 0 {
		return
	}

	flag := psrdm.IfBuyNumReachLimit(buyNum)
	if flag {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	numObj := psrdm.getSoulRuinsNum()
	numObj.BuyNum += buyNum
	numObj.ExtraBuyNum += buyNum
	numObj.UpdateTime = now
	numObj.SetModified()
	return
}

//获取挑战次数使用顺序大小
func (psrdm *PlayerSoulRuinsDataManager) getUsedNum(num int32) (usedBuyNum int32, usedRewNum int32) {
	usedBuyNum = 0
	usedRewNum = 0

	leftBuyNum := psrdm.getSoulRuinsLeftBuyNum()
	leftRewNum := psrdm.getSoulRuinsLeftRewNum()

	if leftBuyNum > 0 {
		if leftBuyNum >= num {
			usedBuyNum = num
			num = 0
		} else {
			usedBuyNum = leftBuyNum
			num -= leftBuyNum
		}
	}

	if num > 0 && leftRewNum > 0 {
		if leftRewNum >= num {
			usedRewNum = num
		} else {
			usedRewNum = leftRewNum
		}
	}

	return
}

//更新挑战次数使用
func (psrdm *PlayerSoulRuinsDataManager) updateUsendNum(num int32) {
	usedBuyNum, usedRewNum := psrdm.getUsedNum(num)
	numObj := psrdm.getSoulRuinsNum()

	now := global.GetGame().GetTimeService().Now()
	//优先使用购买次数
	if usedBuyNum > 0 {
		numObj.UsedBuyNum += usedBuyNum
	}
	//其次使用奖励次数
	if usedRewNum > 0 {
		numObj.UsedRewNum += usedRewNum
	}

	usedNum := num - usedBuyNum - usedRewNum
	//最后使用默认次数
	if usedNum > 0 {
		numObj.UsedNum += usedNum
	}
	psrdm.playerSoulRuinsNumObject.Num += num
	psrdm.playerSoulRuinsNumObject.UpdateTime = now
	psrdm.playerSoulRuinsNumObject.LastTime = now
	psrdm.playerSoulRuinsNumObject.SetModified()
	return
}

//使用挑战次数
func (psrdm *PlayerSoulRuinsDataManager) UseChallengeNum(num int32) {
	if num <= 0 {
		return
	}
	flag := psrdm.HasEnoughChallengeNum(num)
	if !flag {
		return
	}
	psrdm.updateUsendNum(num)
	gameevent.Emit(soulruinseventtypes.EventTypeSoulruinsChallenge, psrdm.p, num)
	return
}

//获取当前剩余挑战次数
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsLeftNum() (leftNum int32) {
	psrdm.refreshSoulRuinsNum()
	leftBuyNum := psrdm.getSoulRuinsLeftBuyNum()
	leftRewNum := psrdm.getSoulRuinsLeftRewNum()
	defaultNum := soulruins.GetSoulRuinsService().GetSoulRuinsChallengeNum()
	defaultLeftNum := defaultNum - psrdm.playerSoulRuinsNumObject.UsedNum
	leftNum = leftBuyNum + leftRewNum + defaultLeftNum
	return
}

//获取剩余的默认免费次数
func (psrdm *PlayerSoulRuinsDataManager) GetSoulRuinsDefaultLeftNum() (leftNum int32) {
	psrdm.refreshSoulRuinsNum()
	defaultNum := soulruins.GetSoulRuinsService().GetSoulRuinsChallengeNum()
	leftNum = defaultNum - psrdm.playerSoulRuinsNumObject.UsedNum
	return
}

//获取当前已通关最高关卡
func (psrdm *PlayerSoulRuinsDataManager) GetCurMaxLevel() (chapter int32, typ soulruinstypes.SoulRuinsType, level int32) {
	curMaxChapter := int32(1)
	level = 1
	for curChapter, _ := range psrdm.playerSoulRuinsMap {
		if curChapter > curMaxChapter {
			curMaxChapter = curChapter
		}
	}

	soulRuinsTypMap, exist := psrdm.playerSoulRuinsMap[curMaxChapter]
	if !exist {
		to := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(curMaxChapter, soulruinstypes.SoulRuinsTypeEasy, 1)
		if to != nil {
			chapter = to.Chapter
			typ = to.GetType()
			level = to.Level
			return
		}
		to = soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(curMaxChapter, soulruinstypes.SoulRuinsTypeHard, 1)
		if to != nil {
			chapter = to.Chapter
			typ = to.GetType()
			level = to.Level
		}
		return
	}
	chapter = curMaxChapter
	typ = soulruinstypes.SoulRuinsTypeHard
	levelMap, exist := soulRuinsTypMap[soulruinstypes.SoulRuinsTypeHard]
	if !exist {
		typ = soulruinstypes.SoulRuinsTypeEasy
		levelMap, _ = soulRuinsTypMap[soulruinstypes.SoulRuinsTypeEasy]
	}
	for soulLevel, _ := range levelMap {
		if soulLevel > level {
			level = soulLevel
		}
	}
	return
}

//仅gm使用 清空挑战次数
func (psrdm *PlayerSoulRuinsDataManager) GMClearSoulRuinsNum() {
	now := global.GetGame().GetTimeService().Now()
	psrdm.playerSoulRuinsNumObject.Num = 0
	psrdm.playerSoulRuinsNumObject.UpdateTime = now
	psrdm.playerSoulRuinsNumObject.LastTime = 0
	psrdm.playerSoulRuinsNumObject.SetModified()
}

//仅gm使用 清空帝陵遗迹
func (psrdm *PlayerSoulRuinsDataManager) GMClearSoulRuins() {
	soulRuinsMap := psrdm.GetSoulRuinsMap()
	rewChapterMap := psrdm.GetSoulRuinsRewChapterMap()
	now := global.GetGame().GetTimeService().Now()

	for _, soulRuinsTypMap := range soulRuinsMap {
		for _, soulRuinsLevelMap := range soulRuinsTypMap {
			for level, levelObj := range soulRuinsLevelMap {
				delete(soulRuinsLevelMap, level)
				levelObj.DeleteTime = now
				levelObj.SetModified()
			}
		}
	}

	for _, rewChapterTypMap := range rewChapterMap {
		for typ, rewChapterObj := range rewChapterTypMap {
			delete(rewChapterTypMap, typ)
			rewChapterObj.DeleteTime = now
			rewChapterObj.SetModified()
		}
	}

	for _, chapterStarTypMap := range psrdm.playerSoulRuinsChapterStarMap {
		for typ, _ := range chapterStarTypMap {
			delete(chapterStarTypMap, typ)
		}
	}
}

//gm使用 设置等级
func (psrdm *PlayerSoulRuinsDataManager) GMSetLevel(chapter int32, typ soulruinstypes.SoulRuinsType, level int32) {
	now := global.GetGame().GetTimeService().Now()

	soulRuinsTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	rewTime := soulRuinsTemplate.RewTime

	soulRuinsObj := psrdm.getSoulRuinsLevelInfo(chapter, typ, level)
	if soulRuinsObj == nil {
		for curLevel := level; curLevel > 0; {
			obj := psrdm.getSoulRuinsLevelInfo(chapter, typ, curLevel)
			if obj == nil {
				psrdm.gmAddSoulRuins(chapter, typ, curLevel, rewTime, now)
			} else {
				break
			}
			tempLevel := curLevel
			curLevel -= 1
			if curLevel == 0 {
				to := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, tempLevel)
				frontId := to.FrontId
				if frontId != 0 {
					bTo := soulruins.GetSoulRuinsService().GetSoulRuinsTemplateById(frontId)
					chapter = bTo.Chapter
					typ = bTo.GetType()
					curLevel = bTo.Level
				}
			}
		}
	} else {
		soulRuinsObj.UpdateTime = now
		soulRuinsObj.Star = soulruinstypes.MaxStar
		soulRuinsObj.SetModified()
	}
	return
}

//gm使用
func (psrdm *PlayerSoulRuinsDataManager) gmAddSoulRuins(chapter int32, typ soulruinstypes.SoulRuinsType, level int32, rewTime int32, now int64) {
	psro := NewPlayerSoulRuinsObject(psrdm.p)
	id, _ := idutil.GetId()
	psro.Id = id
	//生成id
	psro.PlayerId = psrdm.p.GetId()
	psro.Chapter = chapter
	psro.Type = typ
	psro.Level = level
	psro.Star = soulruinstypes.MaxStar
	psro.CreateTime = now
	psro.SetModified()

	psrdm.addSoulRuins(psro)
	//奖励挑战次数
	psrdm.addSoulRuinsRewNum(rewTime)
	return
}

func CreatePlayerSoulRuinsDataManager(p player.Player) player.PlayerDataManager {
	psrdm := &PlayerSoulRuinsDataManager{}
	psrdm.p = p
	return psrdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSoulRuinsDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSoulRuinsDataManager))
}
