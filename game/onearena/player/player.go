package player

import (
	"encoding/json"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"sort"

	"fgame/fgame/game/onearena/dao"
	onearenaentity "fgame/fgame/game/onearena/entity"
	onearenatemplate "fgame/fgame/game/onearena/template"
	onearenatypes "fgame/fgame/game/onearena/types"

	"github.com/pkg/errors"
)

//灵池争夺对象
type PlayerOneArenaObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	Level       onearenatypes.OneArenaLevelType
	KunSilver   int64
	KunBindGold int64
	Pos         int32
	RobTime     int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerOneArenaObject(pl player.Player) *PlayerOneArenaObject {
	poao := &PlayerOneArenaObject{
		player: pl,
	}
	return poao
}

func (poao *PlayerOneArenaObject) GetPlayerId() int64 {
	return poao.PlayerId
}

func (poao *PlayerOneArenaObject) GetDBId() int64 {
	return poao.Id
}

func (poao *PlayerOneArenaObject) ToEntity() (e storage.Entity, err error) {
	e = &onearenaentity.PlayerOneArenaEntity{
		Id:          poao.Id,
		PlayerId:    poao.PlayerId,
		Level:       int32(poao.Level),
		Pos:         poao.Pos,
		RobTime:     poao.RobTime,
		KunSilver:   poao.KunSilver,
		KunBindGold: poao.KunBindGold,
		UpdateTime:  poao.UpdateTime,
		CreateTime:  poao.CreateTime,
		DeleteTime:  poao.DeleteTime,
	}
	return e, nil
}

func (poao *PlayerOneArenaObject) FromEntity(e storage.Entity) error {
	poae, _ := e.(*onearenaentity.PlayerOneArenaEntity)

	poao.Id = poae.Id
	poao.PlayerId = poae.PlayerId
	poao.Level = onearenatypes.OneArenaLevelType(poae.Level)
	poao.Pos = poae.Pos
	poao.RobTime = poae.RobTime
	poao.KunSilver = poae.KunSilver
	poao.KunBindGold = poae.KunBindGold
	poao.UpdateTime = poae.UpdateTime
	poao.CreateTime = poae.CreateTime
	poao.DeleteTime = poae.DeleteTime
	return nil
}

func (poao *PlayerOneArenaObject) SetModified() {
	e, err := poao.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OneArena"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	poao.player.AddChangedObject(obj)
	return
}

//灵池争夺记录对象
type PlayerOneArenaRecordObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Level      onearenatypes.OneArenaLevelType
	Pos        int32
	RobTime    int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerOneArenaRecordObject(pl player.Player) *PlayerOneArenaRecordObject {
	poaro := &PlayerOneArenaRecordObject{
		player: pl,
	}
	return poaro
}

func (poaro *PlayerOneArenaRecordObject) GetPlayerId() int64 {
	return poaro.PlayerId
}

func (poaro *PlayerOneArenaRecordObject) GetDBId() int64 {
	return poaro.Id
}

func (poaro *PlayerOneArenaRecordObject) ToEntity() (e storage.Entity, err error) {
	e = &onearenaentity.PlayerOneArenaRecordEntity{
		Id:         poaro.Id,
		PlayerId:   poaro.PlayerId,
		Level:      int32(poaro.Level),
		Pos:        poaro.Pos,
		RobTime:    poaro.RobTime,
		UpdateTime: poaro.UpdateTime,
		CreateTime: poaro.CreateTime,
		DeleteTime: poaro.DeleteTime,
	}
	return e, nil
}

func (poaro *PlayerOneArenaRecordObject) FromEntity(e storage.Entity) error {
	poare, _ := e.(*onearenaentity.PlayerOneArenaRecordEntity)

	poaro.Id = poare.Id
	poaro.PlayerId = poare.PlayerId
	poaro.Level = onearenatypes.OneArenaLevelType(poare.Level)
	poaro.Pos = poare.Pos
	poaro.RobTime = poare.RobTime
	poaro.UpdateTime = poare.UpdateTime
	poaro.CreateTime = poare.CreateTime
	poaro.DeleteTime = poare.DeleteTime
	return nil
}

func (poaro *PlayerOneArenaRecordObject) SetModified() {
	e, err := poaro.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OneArenaRecord"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	poaro.player.AddChangedObject(obj)
	return
}

//灵池被抢
type PlayerOneArenaRobbedObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	RobName    string
	RobTime    int64
	Status     onearenatypes.OneArenaRobbedStatusType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

//抢夺记录排序
type PlayerOneArenaRobbedObjectList []*PlayerOneArenaRobbedObject

func (oarol PlayerOneArenaRobbedObjectList) Len() int {
	return len(oarol)
}

func (oarol PlayerOneArenaRobbedObjectList) Less(i, j int) bool {
	return oarol[i].RobTime < oarol[j].RobTime
}

func (oarol PlayerOneArenaRobbedObjectList) Swap(i, j int) {
	oarol[i], oarol[j] = oarol[j], oarol[i]
}

func NewPlayerOneArenaRobbedObject(pl player.Player) *PlayerOneArenaRobbedObject {
	poaro := &PlayerOneArenaRobbedObject{
		player: pl,
	}
	return poaro
}

func (poaro *PlayerOneArenaRobbedObject) GetPlayerId() int64 {
	return poaro.PlayerId
}

func (poaro *PlayerOneArenaRobbedObject) GetDBId() int64 {
	return poaro.Id
}

func (poaro *PlayerOneArenaRobbedObject) ToEntity() (e storage.Entity, err error) {
	e = &onearenaentity.PlayerOneArenaRobbedEntity{
		Id:         poaro.Id,
		PlayerId:   poaro.PlayerId,
		RobName:    poaro.RobName,
		RobTime:    poaro.RobTime,
		Status:     int32(poaro.Status),
		UpdateTime: poaro.UpdateTime,
		CreateTime: poaro.CreateTime,
		DeleteTime: poaro.DeleteTime,
	}
	return e, nil
}

func (poaro *PlayerOneArenaRobbedObject) FromEntity(e storage.Entity) error {
	poare, _ := e.(*onearenaentity.PlayerOneArenaRobbedEntity)

	poaro.Id = poare.Id
	poaro.PlayerId = poare.PlayerId
	poaro.RobName = poare.RobName
	poaro.RobTime = poare.RobTime
	poaro.Status = onearenatypes.OneArenaRobbedStatusType(poare.Status)
	poaro.UpdateTime = poare.UpdateTime
	poaro.CreateTime = poare.CreateTime
	poaro.DeleteTime = poare.DeleteTime
	return nil
}

func (poaro *PlayerOneArenaRobbedObject) SetModified() {
	e, err := poaro.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OneArenaRobbed"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	poaro.player.AddChangedObject(obj)
	return
}

//玩家下线灵池产出的鲲
type PlayerOneArenaKunObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	ItemMap    map[int32]int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerOneArenaKunObject(pl player.Player) *PlayerOneArenaKunObject {
	poako := &PlayerOneArenaKunObject{
		player: pl,
	}
	return poako
}

func (poako *PlayerOneArenaKunObject) GetPlayerId() int64 {
	return poako.PlayerId
}

func (poako *PlayerOneArenaKunObject) GetDBId() int64 {
	return poako.Id
}

func (poako *PlayerOneArenaKunObject) ToEntity() (e storage.Entity, err error) {
	itemInfoBytes, err := json.Marshal(poako.ItemMap)
	if err != nil {
		return nil, err
	}

	e = &onearenaentity.PlayerOneArenaKunEntity{
		Id:         poako.Id,
		PlayerId:   poako.PlayerId,
		KunInfo:    string(itemInfoBytes),
		UpdateTime: poako.UpdateTime,
		CreateTime: poako.CreateTime,
		DeleteTime: poako.DeleteTime,
	}
	return e, nil
}

func (poako *PlayerOneArenaKunObject) FromEntity(e storage.Entity) error {
	poake, _ := e.(*onearenaentity.PlayerOneArenaKunEntity)

	itemInfoMap := make(map[int32]int32)
	if err := json.Unmarshal([]byte(poake.KunInfo), &itemInfoMap); err != nil {
		return err
	}
	poako.Id = poake.Id
	poako.PlayerId = poake.PlayerId
	poako.ItemMap = itemInfoMap
	poako.UpdateTime = poake.UpdateTime
	poako.CreateTime = poake.CreateTime
	poako.DeleteTime = poake.DeleteTime
	return nil
}

func (poako *PlayerOneArenaKunObject) SetModified() {
	e, err := poako.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OneArenaKun"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	poako.player.AddChangedObject(obj)
	return
}

//玩家灵池争夺管理器
type PlayerOneArenaDataManager struct {
	p player.Player
	//玩家灵池争夺对象
	playerOneArenaObject *PlayerOneArenaObject
	//玩家灵池争夺记录
	playerOneArenaRecordMap map[onearenatypes.OneArenaLevelType]map[int32]*PlayerOneArenaRecordObject
	//玩家灵池被抢记录
	playerOneArenaRobbedList []*PlayerOneArenaRobbedObject
	//玩家鲲记录
	playerOneArenaKunList []*PlayerOneArenaKunObject
	//鲲map
	kunMap map[int32]int32
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (poadm *PlayerOneArenaDataManager) Player() player.Player {
	return poadm.p
}

//加载
func (poadm *PlayerOneArenaDataManager) Load() (err error) {
	poadm.playerOneArenaRecordMap = make(map[onearenatypes.OneArenaLevelType]map[int32]*PlayerOneArenaRecordObject)
	poadm.playerOneArenaRobbedList = make([]*PlayerOneArenaRobbedObject, 0, 8)
	poadm.playerOneArenaKunList = make([]*PlayerOneArenaKunObject, 0, 8)
	poadm.kunMap = make(map[int32]int32)
	//玩家灵池对象
	oneArenaEntity, err := dao.GetOneArenaDao().GetPlayerOneArenaEntity(poadm.p.GetId())
	if err != nil {
		return
	}
	if oneArenaEntity == nil {
		poadm.initPlayerOneArenaObject()
	} else {
		poadm.playerOneArenaObject = NewPlayerOneArenaObject(poadm.p)
		poadm.playerOneArenaObject.FromEntity(oneArenaEntity)
	}

	//玩家灵池争夺记录
	oneArenaRecordList, err := dao.GetOneArenaDao().GetPlayerOneArenaRecordList(poadm.p.GetId())
	if err != nil {
		return
	}

	for _, oneArenaRecord := range oneArenaRecordList {
		poaro := NewPlayerOneArenaRecordObject(poadm.p)
		poaro.FromEntity(oneArenaRecord)

		typ := onearenatypes.OneArenaLevelType(oneArenaRecord.Level)

		oneArenaLevelMap, exist := poadm.playerOneArenaRecordMap[typ]
		if !exist {
			oneArenaLevelMap = make(map[int32]*PlayerOneArenaRecordObject)
			poadm.playerOneArenaRecordMap[typ] = oneArenaLevelMap
		}
		oneArenaLevelMap[poaro.Pos] = poaro
	}

	//玩家灵池被抢记录
	oneArenaRobbedList, err := dao.GetOneArenaDao().GetPlayerOneArenaRobbedList(poadm.p.GetId())
	if err != nil {
		return
	}

	for _, oneArenaRobbed := range oneArenaRobbedList {
		poaro := NewPlayerOneArenaRobbedObject(poadm.p)
		poaro.FromEntity(oneArenaRobbed)
		poadm.playerOneArenaRobbedList = append(poadm.playerOneArenaRobbedList, poaro)
	}

	//灵池产出的鲲
	oneArenaKunList, err := dao.GetOneArenaDao().GetPlayerOneArenaKunList(poadm.p.GetId())
	if err != nil {
		return
	}

	for _, oneArenaKun := range oneArenaKunList {
		poako := NewPlayerOneArenaKunObject(poadm.p)
		poako.FromEntity(oneArenaKun)
		poadm.playerOneArenaKunList = append(poadm.playerOneArenaKunList, poako)
		for itemId, num := range poako.ItemMap {
			curNum, exist := poadm.kunMap[itemId]
			if exist {
				num += curNum
			}
			poadm.kunMap[itemId] = num
		}
	}

	return nil
}

//第一次初始化
func (poadm *PlayerOneArenaDataManager) initPlayerOneArenaObject() {
	poao := NewPlayerOneArenaObject(poadm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	poao.Id = id
	//生成id
	poao.PlayerId = poadm.p.GetId()
	poao.Level = 0
	poao.Pos = 0
	poao.RobTime = 0
	poao.CreateTime = now
	poadm.playerOneArenaObject = poao
	poao.SetModified()
}

func (poadm *PlayerOneArenaDataManager) newPlayerOneArenaRecordObject(now int64, level onearenatypes.OneArenaLevelType, pos int32) (poaro *PlayerOneArenaRecordObject) {
	poaro = NewPlayerOneArenaRecordObject(poadm.p)

	id, _ := idutil.GetId()
	poaro.Id = id
	//生成id
	poaro.PlayerId = poadm.p.GetId()
	poaro.Level = level
	poaro.Pos = pos
	poaro.RobTime = now
	poaro.CreateTime = now
	poaro.SetModified()
	return
}

func (poadm *PlayerOneArenaDataManager) newPlayerOneArenaRobbedObject(robName string, robTime int64, sucess bool) *PlayerOneArenaRobbedObject {
	poaro := NewPlayerOneArenaRobbedObject(poadm.p)
	status := onearenatypes.OneArenaRobbedStatusTypeSucess
	if !sucess {
		status = onearenatypes.OneArenaRobbedStatusTypeFail
	}

	id, _ := idutil.GetId()
	poaro.Id = id
	poaro.PlayerId = poadm.p.GetId()
	poaro.RobName = robName
	poaro.RobTime = robTime
	poaro.Status = status
	poaro.CreateTime = robTime
	poaro.SetModified()
	return poaro
}

//加载后
func (poadm *PlayerOneArenaDataManager) AfterLoad() (err error) {

	poadm.heartbeatRunner.AddTask(CreateOneArenaTask(poadm.p))
	return nil
}

//心跳
func (poadm *PlayerOneArenaDataManager) Heartbeat() {
	poadm.heartbeatRunner.Heartbeat()
}

func (poadm *PlayerOneArenaDataManager) IsValid(level onearenatypes.OneArenaLevelType, pos int32) bool {
	oneArenaTemplate := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
	if oneArenaTemplate == nil {
		return false
	}
	return true
}

func (poadm *PlayerOneArenaDataManager) GetOneArenaKunMap() map[int32]int32 {
	return poadm.kunMap
}

func (poadm *PlayerOneArenaDataManager) GetOneArenaRecordMap() map[onearenatypes.OneArenaLevelType]map[int32]*PlayerOneArenaRecordObject {
	return poadm.playerOneArenaRecordMap
}

func (poadm *PlayerOneArenaDataManager) GetOneArenaRecord(level onearenatypes.OneArenaLevelType, pos int32) *PlayerOneArenaRecordObject {
	levelMap, exist := poadm.playerOneArenaRecordMap[level]
	if !exist {
		return nil
	}
	oneArenaRecord, exist := levelMap[pos]
	if !exist {
		return nil
	}
	return oneArenaRecord
}

func (poadm *PlayerOneArenaDataManager) GetOneArena() *PlayerOneArenaObject {
	return poadm.playerOneArenaObject
}

func (poadm *PlayerOneArenaDataManager) IsRobCoolTime(level onearenatypes.OneArenaLevelType, pos int32) bool {
	levelMap, exist := poadm.playerOneArenaRecordMap[level]
	if !exist {
		return false
	}
	oneArenaRecord, exist := levelMap[pos]
	if !exist {
		return false
	}

	oneArenaTemplate := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
	cdTime := oneArenaTemplate.CoolTime
	lastTime := oneArenaRecord.RobTime
	now := global.GetGame().GetTimeService().Now()

	if now-lastTime < int64(cdTime) {
		return true
	}
	return false
}

func (poadm *PlayerOneArenaDataManager) robOneArenaRecord(now int64, Level onearenatypes.OneArenaLevelType, pos int32) {
	oneArenaLevelMap, exist := poadm.playerOneArenaRecordMap[Level]
	if !exist {
		oneArenaLevelMap = make(map[int32]*PlayerOneArenaRecordObject)
		poadm.playerOneArenaRecordMap[Level] = oneArenaLevelMap
		poaro := poadm.newPlayerOneArenaRecordObject(now, Level, pos)
		oneArenaLevelMap[poaro.Pos] = poaro
	} else {
		oneArenaObj, exist := oneArenaLevelMap[pos]
		if !exist {
			poaro := poadm.newPlayerOneArenaRecordObject(now, Level, pos)
			oneArenaLevelMap[poaro.Pos] = poaro
		} else {
			oneArenaObj.RobTime = now
			oneArenaObj.UpdateTime = now
			oneArenaObj.SetModified()
		}
	}
	return
}

func (poadm *PlayerOneArenaDataManager) RobOneArenaRecord(level onearenatypes.OneArenaLevelType, pos int32) {
	now := global.GetGame().GetTimeService().Now()
	poadm.robOneArenaRecord(now, level, pos)
	return
}

func (poadm *PlayerOneArenaDataManager) ReplaceOneArena(level onearenatypes.OneArenaLevelType, pos int32) {
	now := global.GetGame().GetTimeService().Now()
	poadm.playerOneArenaObject.Level = level
	poadm.playerOneArenaObject.Pos = pos
	poadm.playerOneArenaObject.RobTime = now
	poadm.playerOneArenaObject.UpdateTime = now
	poadm.playerOneArenaObject.SetModified()
	return
}

func (poadm *PlayerOneArenaDataManager) ReplaceOneArenaAfter(level onearenatypes.OneArenaLevelType, pos int32, robName string, robTime int64) {
	poadm.ReplaceOneArena(level, pos)
	//容错处理 (全局表灵池写入 & 个人的被抢没写入)
	if len(poadm.playerOneArenaRobbedList) == 0 && robName != "" {
		poadm.robbedRecords(robName, robTime, true)
	}
}

func (poadm *PlayerOneArenaDataManager) robbedRecords(robName string, robTime int64, sucess bool) {
	maxLen := onearenatypes.RECORD_MAXLEN
	ero := poadm.newPlayerOneArenaRobbedObject(robName, robTime, sucess)
	poadm.playerOneArenaRobbedList = append(poadm.playerOneArenaRobbedList, ero)
	curLen := len(poadm.playerOneArenaRobbedList)

	sort.Sort(sort.Reverse(PlayerOneArenaRobbedObjectList(poadm.playerOneArenaRobbedList)))

	if curLen > maxLen {
		for index := maxLen; index < curLen; index++ {
			poadm.playerOneArenaRobbedList[index].DeleteTime = robTime
			poadm.playerOneArenaRobbedList[index].UpdateTime = robTime
			poadm.playerOneArenaRobbedList[index].SetModified()
		}
		poadm.playerOneArenaRobbedList = poadm.playerOneArenaRobbedList[:maxLen]
	}
	return
}

func (poadm *PlayerOneArenaDataManager) RobbedRecord(name string, sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	poadm.robbedRecords(name, now, sucess)
	return
}

func (poadm *PlayerOneArenaDataManager) DeleteKunRecord() {
	now := global.GetGame().GetTimeService().Now()
	for _, kunObj := range poadm.playerOneArenaKunList {
		kunObj.DeleteTime = now
		kunObj.UpdateTime = now
		kunObj.SetModified()
	}
	return
}

func (poadm *PlayerOneArenaDataManager) GetFirstRecord() *PlayerOneArenaRobbedObject {
	if len(poadm.playerOneArenaRobbedList) > 0 {
		return poadm.playerOneArenaRobbedList[0]
	}
	return nil
}

func (poadm *PlayerOneArenaDataManager) GetRecordByLogTime(logTime int64) []*PlayerOneArenaRobbedObject {
	starIndex := int(-1)
	for index, oneArenaRobbed := range poadm.playerOneArenaRobbedList {
		if oneArenaRobbed.RobTime <= logTime {
			break
		}
		starIndex = index
	}

	if starIndex >= 0 {
		return poadm.playerOneArenaRobbedList[0 : starIndex+1]

	}
	return nil
}

func (poadm *PlayerOneArenaDataManager) SellKunAddRes(kunSilver int64, kunBindGold int64) (totalKunSilver int64, totalBindGold int64) {
	if kunSilver <= 0 || kunBindGold <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	poadm.playerOneArenaObject.KunSilver += kunSilver
	poadm.playerOneArenaObject.KunBindGold += kunBindGold
	poadm.playerOneArenaObject.UpdateTime = now
	poadm.playerOneArenaObject.SetModified()

	totalKunSilver = poadm.playerOneArenaObject.KunSilver
	totalBindGold = poadm.playerOneArenaObject.KunBindGold
	return
}

// 是否满足任务
func (poadm *PlayerOneArenaDataManager) IsFullQuestCondition(level onearenatypes.OneArenaLevelType, needOccupyTime int64) (flag bool) {
	if poadm.playerOneArenaObject.Level < level {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	occupyTime := now - poadm.playerOneArenaObject.RobTime
	if occupyTime < needOccupyTime {
		return
	}

	flag = true
	return
}

func CreatePlayerOneArenaDataManager(p player.Player) player.PlayerDataManager {
	poadm := &PlayerOneArenaDataManager{}
	poadm.p = p
	poadm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return poadm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerOneArenaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerOneArenaDataManager))
}
