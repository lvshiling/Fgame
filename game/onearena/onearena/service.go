package onearena

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/onearena/dao"
	onearenaentity "fgame/fgame/game/onearena/entity"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	playeronearena "fgame/fgame/game/onearena/player"
	onearenatemplate "fgame/fgame/game/onearena/template"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"sync"
)

//灵池争夺处理
type OneArenaService interface {
	Heartbeat()
	//获取灵池信息
	GetOneArenaList() []*OneArenaObject
	//获取玩家当前灵池
	GetPlayerOneArena(playerId int64) (level onearenatypes.OneArenaLevelType, pos int32)
	//获取灵池信息
	GetOneArena(level onearenatypes.OneArenaLevelType, pos int32) *OneArenaObject
	//抢灵池
	OneArenaRob(oneArenaObj *playeronearena.PlayerOneArenaObject, level onearenatypes.OneArenaLevelType, pos int32) (ownerId int64, ownerName string, codeResult onearenatypes.OneArenaRobCodeType)
	//改变灵池占领者
	OneArenaRobSucess(pl player.Player, level onearenatypes.OneArenaLevelType, pos int32) error
	//抢夺失败
	OneArenaRobFail(pl player.Player, level onearenatypes.OneArenaLevelType, pos int32) error
	//玩家名字变化
	PlayerNameChanged(pl player.Player)
	//仅gm 命令使用
	GmResetKunTime()
	//仅gm 使用
	GmResetLastTime()
}

type oneArenaService struct {
	//读写锁
	rwm          sync.RWMutex
	oneArenaList []*OneArenaObject
	ownerMap     map[int64]*OneArenaObject
}

//初始化
func (os *oneArenaService) init() (err error) {
	os.oneArenaList = make([]*OneArenaObject, 0, 8)
	os.ownerMap = make(map[int64]*OneArenaObject)
	//灵池信息
	oneArenaList, err := dao.GetOneArenaDao().GetOneArenaList()
	if err != nil {
		return
	}
	err = os.initOneArena(oneArenaList)
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	//合服
	isMerge := merge.GetMergeService().IsMerge()
	if isMerge {
		//归还鲲
		os.outputKun(now)
		os.mergeServer()
	}
	os.initOneArenaTemplate()
	return
}

func (os *oneArenaService) initOneArena(oneArenaList []*onearenaentity.OneArenaEntity) (err error) {
	for _, oneArenaObj := range oneArenaList {
		oao := NewOneArenaObject()
		oao.FromEntity(oneArenaObj)
		os.oneArenaList = append(os.oneArenaList, oao)
		os.ownerMap[oao.OwnerId] = oao
	}
	return
}

func (os *oneArenaService) initOneArenaTemplate() {
	now := global.GetGame().GetTimeService().Now()
	oneArenaMap := onearenatemplate.GetOneArenaTemplateService().GetOneArenaMap()
	for level, levelMap := range oneArenaMap {
		for pos, _ := range levelMap {
			oneArenaObj := os.getOneArenaObj(level, pos)
			if oneArenaObj != nil {
				continue
			}
			oao := os.initOneArenaObject(level, pos, now)
			os.oneArenaList = append(os.oneArenaList, oao)
		}
	}
	return
}

func (os *oneArenaService) mergeServer() {
	now := global.GetGame().GetTimeService().Now()
	for _, oneArenaObj := range os.oneArenaList {
		oneArenaObj.DeleteTime = now
		oneArenaObj.SetModified()
		delete(os.ownerMap, oneArenaObj.OwnerId)
		gameevent.Emit(onearenaeventtypes.EventTypeOneArenaMergeServer, oneArenaObj.OwnerId, nil)
	}
	os.oneArenaList = make([]*OneArenaObject, 0, 8)
}

func (os *oneArenaService) initOneArenaObject(level onearenatypes.OneArenaLevelType, pos int32, now int64) *OneArenaObject {
	oao := NewOneArenaObject()
	id, _ := idutil.GetId()
	oao.Id = id
	oao.ServerId = global.GetGame().GetServerIndex()
	oao.Level = level
	oao.Pos = pos
	oao.OwnerId = 0
	oao.OwnerName = ""
	oao.LastTime = now
	oao.UpdateTime = now
	oao.CreateTime = now
	oao.IsRobbing = false
	oao.SetModified()
	return oao
}

func (os *oneArenaService) GetOneArena(level onearenatypes.OneArenaLevelType, pos int32) *OneArenaObject {
	os.rwm.RLock()
	defer os.rwm.RUnlock()

	oneArenaObj := os.getOneArenaObj(level, pos)
	return oneArenaObj
}

func (os *oneArenaService) GetPlayerOneArena(playerId int64) (level onearenatypes.OneArenaLevelType, pos int32) {
	os.rwm.RLock()
	defer os.rwm.RUnlock()

	oneArena, exist := os.ownerMap[playerId]
	if !exist {
		return 0, 0
	}
	level = oneArena.Level
	pos = oneArena.Pos
	return
}

func (os *oneArenaService) GetOneArenaList() []*OneArenaObject {
	os.rwm.RLock()
	defer os.rwm.RUnlock()
	return os.oneArenaList
}

func (os *oneArenaService) OneArenaRobFail(pl player.Player, level onearenatypes.OneArenaLevelType, pos int32) (err error) {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	os.removeRobbing(level, pos)

	oneArenaObj := os.getOneArenaObj(level, pos)
	if oneArenaObj == nil {
		return
	}

	peerArenaData := onearenaeventtypes.CreateOneArenaData(oneArenaObj.OwnerId, level, pos)
	eventData := onearenaeventtypes.CreateOneArenaRobFailEventData(pl.GetId(), peerArenaData)
	gameevent.Emit(onearenaeventtypes.EventTypePlayerOneArenaFail, nil, eventData)
	return
}

func (os *oneArenaService) PlayerNameChanged(pl player.Player) {
	os.rwm.Lock()
	defer os.rwm.Unlock()
	playerId := pl.GetId()
	oneArena, exist := os.ownerMap[playerId]
	if !exist {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	oneArena.OwnerName = pl.GetName()
	oneArena.UpdateTime = now
	oneArena.SetModified()
	return
}

func (os *oneArenaService) OneArenaRobSucess(pl player.Player, level onearenatypes.OneArenaLevelType, pos int32) (err error) {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	oneArenaObj := os.getOneArenaObj(level, pos)
	if oneArenaObj == nil {
		return
	}
	ownerId := oneArenaObj.OwnerId
	now := global.GetGame().GetTimeService().Now()

	//移除抢夺
	os.removeRobbing(level, pos)
	myOneArenaObj, exist := os.ownerMap[pl.GetId()]
	if !exist {
		if ownerId != 0 {
			delete(os.ownerMap, ownerId)
		}
		oneArenaObj.OwnerId = pl.GetId()
		oneArenaObj.OwnerName = pl.GetName()
		oneArenaObj.UpdateTime = now
		oneArenaObj.SetModified()
		os.ownerMap[pl.GetId()] = oneArenaObj
	} else {
		ownerLevel := myOneArenaObj.Level
		ownerPos := myOneArenaObj.Pos

		myOneArenaObj.Level = level
		myOneArenaObj.Pos = pos
		myOneArenaObj.UpdateTime = now
		myOneArenaObj.SetModified()

		oneArenaObj.Level = ownerLevel
		oneArenaObj.Pos = ownerPos
		oneArenaObj.UpdateTime = now
		oneArenaObj.SetModified()

	}

	var eventData *onearenaeventtypes.OneArenaRobSucessEventData
	var peerArenaData *onearenaeventtypes.OneArenaData
	oneArenaData := onearenaeventtypes.CreateOneArenaData(pl.GetId(), level, pos)
	if ownerId == 0 {
		eventData = onearenaeventtypes.CreateOneArenaRobSucessEventData(oneArenaData, nil)
	} else if myOneArenaObj == nil { //玩家没灵池取抢1级  || 玩家有灵池 再抢夺过程中自己的被抢了
		peerArenaData = onearenaeventtypes.CreateOneArenaData(ownerId, 0, 0)
		eventData = onearenaeventtypes.CreateOneArenaRobSucessEventData(oneArenaData, peerArenaData)
	} else {
		peerArenaData = onearenaeventtypes.CreateOneArenaData(oneArenaObj.OwnerId, oneArenaObj.Level, oneArenaObj.Pos)
		eventData = onearenaeventtypes.CreateOneArenaRobSucessEventData(oneArenaData, peerArenaData)
	}
	gameevent.Emit(onearenaeventtypes.EventTypePlayerOneArenaSucess, nil, eventData)
	return
}

func (os *oneArenaService) removeRobbing(level onearenatypes.OneArenaLevelType, pos int32) (err error) {
	now := global.GetGame().GetTimeService().Now()
	oneArenaObj := os.getOneArenaObj(level, pos)
	if oneArenaObj == nil {
		return
	}
	if !oneArenaObj.IsRobbing {
		return
	}
	oneArenaObj.IsRobbing = false
	oneArenaObj.UpdateTime = now
	oneArenaObj.SetModified()
	return
}

func (os *oneArenaService) robbing(oneArenaObj *OneArenaObject) {
	oneArenaObj.IsRobbing = true
	now := global.GetGame().GetTimeService().Now()
	oneArenaObj.UpdateTime = now
	oneArenaObj.SetModified()
}

func (os *oneArenaService) getOneArenaObj(level onearenatypes.OneArenaLevelType, pos int32) *OneArenaObject {

	for _, oneArenaObj := range os.oneArenaList {
		if oneArenaObj.Level == level && oneArenaObj.Pos == pos {
			return oneArenaObj
		}
	}
	return nil
}

func (os *oneArenaService) OneArenaRob(plOneArenaObj *playeronearena.PlayerOneArenaObject, level onearenatypes.OneArenaLevelType, pos int32) (ownerId int64, ownerName string, codeResult onearenatypes.OneArenaRobCodeType) {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	codeResult = onearenatypes.OneArenaRobCodeTypeEnter
	robbedOneArenaObj := os.getOneArenaObj(level, pos)
	ownerId = robbedOneArenaObj.OwnerId
	ownerName = robbedOneArenaObj.OwnerName
	isRobbing := robbedOneArenaObj.IsRobbing
	if isRobbing {
		codeResult = onearenatypes.OneArenaRobCodeTypeIsRobbing
		return
	}
	os.robbing(robbedOneArenaObj)
	return
}

func (os *oneArenaService) outputKun(now int64) {
	for _, oneArenaObj := range os.oneArenaList {
		level := oneArenaObj.Level
		pos := oneArenaObj.Pos
		lastTime := oneArenaObj.LastTime

		oneArenaTemplate := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
		refreshTime := int64(oneArenaTemplate.RefreshTime)

		diffTime := now - lastTime
		ownerId := oneArenaObj.OwnerId
		if diffTime >= refreshTime {
			num := int32(diffTime / refreshTime)
			offTime := int64(num) * refreshTime
			oneArenaObj.LastTime += offTime
			oneArenaObj.SetModified()
			if ownerId != 0 {
				gameevent.Emit(onearenaeventtypes.EventTypeOneArenaOutputKun, oneArenaObj, num)
			}
		}
	}
}

//心跳
func (os *oneArenaService) Heartbeat() {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	os.outputKun(now)
}

//仅gm 命令使用
func (os *oneArenaService) GmResetKunTime() {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	for _, oneArenaObj := range os.oneArenaList {
		oneArenaObj.UpdateTime = now
		oneArenaObj.LastTime = now
		oneArenaObj.SetModified()
	}
}

//仅gm 使用
func (os *oneArenaService) GmResetLastTime() {
	os.rwm.Lock()
	defer os.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	for _, oneArenaObj := range os.oneArenaList {
		oneArenaObj.UpdateTime = now
		oneArenaObj.LastTime = now
		oneArenaObj.SetModified()
	}

}

var (
	once sync.Once
	cs   *oneArenaService
)

func Init() (err error) {
	once.Do(func() {
		cs = &oneArenaService{}
		err = cs.init()
	})
	return err
}

func GetOneArenaService() OneArenaService {
	return cs
}
