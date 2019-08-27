package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/gem/dao"
	gementity "fgame/fgame/game/gem/entity"
	gemeventtypes "fgame/fgame/game/gem/event/types"
	"fgame/fgame/game/gem/gem"
	gemtypes "fgame/fgame/game/gem/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//玩家原石对象
type PlayerMiningObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Level      int32
	Storage    int32
	Stone      int64
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMiningObject(pl player.Player) *PlayerMiningObject {
	pso := &PlayerMiningObject{
		player: pl,
	}
	return pso
}

func (pmo *PlayerMiningObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerMiningObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerMiningObject) ToEntity() (e storage.Entity, err error) {
	e = &gementity.PlayerMiningEntity{
		Id:         pmo.Id,
		PlayerId:   pmo.PlayerId,
		Level:      pmo.Level,
		Storage:    pmo.Storage,
		Stone:      pmo.Stone,
		LastTime:   pmo.LastTime,
		UpdateTime: pmo.UpdateTime,
		CreateTime: pmo.CreateTime,
		DeleteTime: pmo.DeleteTime,
	}
	return e, err
}

func (pmo *PlayerMiningObject) FromEntity(e storage.Entity) error {
	pme, _ := e.(*gementity.PlayerMiningEntity)
	pmo.Id = pme.Id
	pmo.PlayerId = pme.PlayerId
	pmo.Level = pme.Level
	pmo.Storage = pme.Storage
	pmo.Stone = pme.Stone
	pmo.LastTime = pme.LastTime
	pmo.UpdateTime = pme.UpdateTime
	pmo.CreateTime = pme.CreateTime
	pmo.DeleteTime = pme.DeleteTime
	return nil
}

func (pmo *PlayerMiningObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Mining"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//玩家赌石对象
type PlayerGambleObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Type       gemtypes.GemGambleType
	Num        int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerGambleObject(pl player.Player) *PlayerGambleObject {
	pso := &PlayerGambleObject{
		player: pl,
	}
	return pso
}

func (pgo *PlayerGambleObject) GetPlayerId() int64 {
	return pgo.PlayerId
}

func (pgo *PlayerGambleObject) GetDBId() int64 {
	return pgo.Id
}

func (pgo *PlayerGambleObject) ToEntity() (e storage.Entity, err error) {
	e = &gementity.PlayerGambleEntity{
		Id:         pgo.Id,
		PlayerId:   pgo.PlayerId,
		Type:       int32(pgo.Type),
		Num:        pgo.Num,
		UpdateTime: pgo.UpdateTime,
		CreateTime: pgo.CreateTime,
		DeleteTime: pgo.DeleteTime,
	}
	return e, err
}

func (pgo *PlayerGambleObject) FromEntity(e storage.Entity) error {
	pge, _ := e.(*gementity.PlayerGambleEntity)
	pgo.Id = pge.Id
	pgo.PlayerId = pge.PlayerId
	pgo.Type = gemtypes.GemGambleType(pge.Type)
	pgo.Num = pge.Num
	pgo.UpdateTime = pge.UpdateTime
	pgo.CreateTime = pge.CreateTime
	pgo.DeleteTime = pge.DeleteTime
	return nil
}

func (pgo *PlayerGambleObject) SetModified() {
	e, err := pgo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Gamble"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pgo.player.AddChangedObject(obj)
	return
}

//玩家宝石管理器
type PlayerGemDataManager struct {
	p player.Player
	//玩家原石对象
	playerMiningObject *PlayerMiningObject
	//玩家赌石map
	playerGambleMap map[gemtypes.GemGambleType]*PlayerGambleObject
}

func (pgdm *PlayerGemDataManager) Player() player.Player {
	return pgdm.p
}

//加载
func (pgdm *PlayerGemDataManager) Load() (err error) {
	pgdm.playerGambleMap = make(map[gemtypes.GemGambleType]*PlayerGambleObject)
	//加载玩家原石信息
	miningEntity, err := dao.GetGemDao().GetMineEntity(pgdm.p.GetId())
	if err != nil {
		return
	}
	if miningEntity == nil {
		pgdm.initPlayerMiningObject()
	} else {
		pgdm.playerMiningObject = NewPlayerMiningObject(pgdm.p)
		pgdm.playerMiningObject.FromEntity(miningEntity)
	}

	//加载玩家赌石信息
	gambleList, err := dao.GetGemDao().GetGambleList(pgdm.p.GetId())
	if err != nil {
		return
	}
	//赌石信息
	for _, gamble := range gambleList {
		pgo := NewPlayerGambleObject(pgdm.p)
		pgo.FromEntity(gamble)
		pgdm.playerGambleMap[pgo.Type] = pgo
	}
	return nil
}

//第一次初始化
func (pgdm *PlayerGemDataManager) initPlayerMiningObject() {
	pmo := NewPlayerMiningObject(pgdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.Id = id
	//生成id
	giveStone := gem.GetGemService().GetFuncOpenGiveStone()
	pmo.PlayerId = pgdm.p.GetId()
	pmo.Level = int32(1)
	pmo.Stone = int64(giveStone)
	pmo.CreateTime = now
	pmo.LastTime = now
	pmo.Storage = 0
	pgdm.playerMiningObject = pmo
	pmo.SetModified()
}

//加载后
func (pgdm *PlayerGemDataManager) AfterLoad() (err error) {
	//刷新
	pgdm.refreshStorage()
	return
}

//心跳
func (pgdm *PlayerGemDataManager) Heartbeat() {

}

//获取玩家原石对象
func (pgdm *PlayerGemDataManager) GetMine() *PlayerMiningObject {
	//刷新
	pgdm.refreshStorage()
	return pgdm.playerMiningObject
}

//获取赌石对象
func (pgdm *PlayerGemDataManager) getGambleTyp(typ gemtypes.GemGambleType) *PlayerGambleObject {
	gambleObj, exist := pgdm.playerGambleMap[typ]
	if !exist {
		return nil
	}
	return gambleObj
}

//获取矿工等级
func (pgdm *PlayerGemDataManager) GetMineLevel() int32 {
	return pgdm.playerMiningObject.Level
}

//获取我的原石
func (pgdm *PlayerGemDataManager) GetMineStone() int64 {
	return pgdm.playerMiningObject.Stone
}

//获取当前库存(外部调用)
func (pgdm *PlayerGemDataManager) GetMineStorage() int32 {
	//刷新
	pgdm.refreshStorage()
	return pgdm.getMineStorage()
}

//获取当前库存
func (pgdm *PlayerGemDataManager) getMineStorage() int32 {
	return pgdm.playerMiningObject.Storage
}

//获取记录时间
func (pgdm *PlayerGemDataManager) GetMineLastTime() int64 {
	return pgdm.playerMiningObject.LastTime
}

//获取赌石次数
func (pgdm *PlayerGemDataManager) GetGambleNum(typ gemtypes.GemGambleType) (num int32) {
	num = 0
	gambleObj, exist := pgdm.playerGambleMap[typ]
	if exist {
		num = gambleObj.Num
	}
	return
}

//是否有足够原石
func (pgdm *PlayerGemDataManager) HasEnoughYuanShi(needNum int32) bool {
	return int64(needNum) <= pgdm.GetMineStone()
}

//参数有效性
func (pgdm *PlayerGemDataManager) IsMineValid(level int32) bool {
	to := gem.GetGemService().GetMineTemplateByLevel(level)
	if to == nil {
		return false
	}
	curLevel := pgdm.GetMineLevel()
	//级数小于当前等级
	if level <= curLevel {
		return false
	}
	return true
}

//统计lastTime后产出的原石
func (pgdm *PlayerGemDataManager) outStone(now int64) (curStorage int32, curLastTime int64) {
	curLevel := pgdm.GetMineLevel()
	curStorge := int64(pgdm.getMineStorage())
	lastTime := pgdm.GetMineLastTime()

	to := gem.GetGemService().GetMineTemplateByLevel(curLevel)
	limitMax := int64(to.LimitMax)
	interval := int64(to.IntervalTime * common.TIME_RATE)
	if now > lastTime {
		count := (now - lastTime) / interval
		outCount := count * int64(to.Revenue)
		curStorge += outCount
		if curStorge > limitMax {
			curStorge = limitMax
		}
		lastTime += int64(count * interval)
	}
	curStorage = int32(curStorge)
	return curStorage, lastTime
}

//矿工激活
func (pgdm *PlayerGemDataManager) MineActive(level int32) bool {
	flag := pgdm.IsMineValid(level)
	if !flag {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.Level = level
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()

	gameevent.Emit(gemeventtypes.EventTypeMineActivate, pgdm.p, nil)
	return true
}

//刷新库存
func (pgdm *PlayerGemDataManager) refreshStorage() {
	now := global.GetGame().GetTimeService().Now()
	curStorge, lastTime := pgdm.outStone(now)
	pgdm.playerMiningObject.Storage = curStorge
	pgdm.playerMiningObject.LastTime = lastTime
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return
}

//领取收益
func (pgdm *PlayerGemDataManager) Receive() *PlayerMiningObject {
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.Stone += int64(pgdm.playerMiningObject.Storage)
	pgdm.playerMiningObject.Storage = int32(0)
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return pgdm.playerMiningObject
}

func (pgdm *PlayerGemDataManager) DropYuanShi(yuanShi int32) (mine *PlayerMiningObject) {
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.Stone += int64(yuanShi)
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	mine = pgdm.playerMiningObject
	return
}

//赌石消耗原石
func (pgdm *PlayerGemDataManager) GambleSubStone(needStone int32) bool {
	if needStone <= 0 {
		return false
	}
	curStone := pgdm.GetMineStone()
	if curStone < int64(needStone) {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.Stone -= int64(needStone)
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return true
}

//增加赌石次数
func (pgdm *PlayerGemDataManager) AddGambleNum(typ gemtypes.GemGambleType, batchNum int32) bool {
	if batchNum <= 0 {
		return false
	}
	flag := typ.Valid()
	if !flag {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	obj := pgdm.getGambleTyp(typ)
	if obj == nil {
		pgo := NewPlayerGambleObject(pgdm.p)
		id, _ := idutil.GetId()
		pgo.Id = id
		//生成id
		pgo.PlayerId = pgdm.p.GetId()
		pgo.Num = batchNum
		pgo.Type = typ
		pgo.CreateTime = now
		pgdm.playerGambleMap[typ] = pgo
		pgo.SetModified()
	} else {
		obj.Num += batchNum
		obj.UpdateTime = now
		obj.SetModified()
	}

	gameevent.Emit(gemeventtypes.EventTypeGemGambleFinish, pgdm.p, batchNum)
	return true
}

//gm 设置矿工等级 仅gm使用
func (pgdm *PlayerGemDataManager) GmSetMineLevel(level int32) {
	now := global.GetGame().GetTimeService().Now()
	curStorge, lastTime := pgdm.outStone(now)
	pgdm.playerMiningObject.Storage = 0
	pgdm.playerMiningObject.Stone += int64(curStorge)
	pgdm.playerMiningObject.LastTime = lastTime
	pgdm.playerMiningObject.Level = level
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return
}

//gm 设置矿山原石 仅gm使用
func (pgdm *PlayerGemDataManager) GmSetMineStone(stone int32) {
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.Stone = int64(stone)
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return
}

//gm 设置矿山当前库存 仅gm使用
func (pgdm *PlayerGemDataManager) GmSetMineStorage(storage int32) {
	now := global.GetGame().GetTimeService().Now()
	pgdm.playerMiningObject.LastTime = now
	pgdm.playerMiningObject.Storage = storage
	pgdm.playerMiningObject.UpdateTime = now
	pgdm.playerMiningObject.SetModified()
	return
}

func CreatePlayerGemMineDataManager(p player.Player) player.PlayerDataManager {
	pgdm := &PlayerGemDataManager{}
	pgdm.p = p
	return pgdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerGemDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerGemMineDataManager))
}
