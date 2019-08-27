package friend

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/dao"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/global"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"math"
	"sync"
)

type FriendData struct {
	friendMap    map[int64]map[int64]*FriendObject
	revFriendMap map[int64]map[int64]*FriendObject
}

func newFriendData() *FriendData {
	d := &FriendData{
		friendMap:    make(map[int64]map[int64]*FriendObject),
		revFriendMap: make(map[int64]map[int64]*FriendObject),
	}
	return d
}

//好友接口处理
type FriendService interface {
	GetFriendList(playerId int64) (friendMap map[int64]*FriendObject)
	AddFriend(pl player.Player, friendId int64) (fo *FriendObject, err error)
	DeleteFriend(pl player.Player, friendId int64) (err error)
	AddPoint(pl player.Player, friendId int64, pointNum int32)
	DivorceSubPoint(pl player.Player, spouseId int64, divorceType marrytypes.MarryDivorceType, percent float64) (point int32)
	//添加表白日志
	AddMarryDevelopLog(logData *friendtypes.MarryDevelopLogData)
	//获取表白日志
	GetMarryDevelopLogByTime(time int64) []*FriendMarryDevelopLogObject
}

type friendService struct {
	//读写锁
	rwm sync.RWMutex

	friendData *FriendData
	//所有表白日志列表
	marryDevelopLogList []*FriendMarryDevelopLogObject
}

//初始化
func (fs *friendService) init() (err error) {

	fs.friendData = newFriendData()

	firendAllList, err := dao.GetFriendDao().GetFriendAllList()
	if err != nil {
		return
	}
	for _, friendEntity := range firendAllList {
		fo := NewFriendObject()
		fo.FromEntity(friendEntity)
		fs.initFriendObj(fo)
	}
	//表白日志
	logEntityList, err := dao.GetFriendDao().GetMarryDevelopLogEntityList()
	if err != nil {
		return
	}
	for _, logEntity := range logEntityList {
		logObj := NewFriendMarryDevelopLogObject()
		logObj.FromEntity(logEntity)
		fs.marryDevelopLogList = append(fs.marryDevelopLogList, logObj)
	}

	return
}

func (fs *friendService) newFriendObj(playerId int64, friendId int64) (fo *FriendObject) {
	fo = NewFriendObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	fo.Id = id
	fo.ServerId = global.GetGame().GetServerIndex()
	fo.PlayerId = playerId
	fo.FriendId = friendId
	fo.Point = 0
	fo.CreateTime = now
	fo.UpdateTime = now
	fo.SetModified()

	fs.initFriendObj(fo)
	return
}

func (fs *friendService) initFriendObj(fo *FriendObject) {
	friendMap, exist := fs.friendData.friendMap[fo.PlayerId]
	if !exist {
		friendMap = make(map[int64]*FriendObject)
		fs.friendData.friendMap[fo.PlayerId] = friendMap
	}
	friendMap[fo.FriendId] = fo

	revFriendMap, exist := fs.friendData.revFriendMap[fo.FriendId]
	if !exist {
		revFriendMap = make(map[int64]*FriendObject)
		fs.friendData.revFriendMap[fo.FriendId] = revFriendMap
	}
	revFriendMap[fo.PlayerId] = fo
}

func (fs *friendService) getFriendList(playerId int64) (friendMap map[int64]*FriendObject) {
	friendMap = make(map[int64]*FriendObject)
	for friendId, revFriendMap := range fs.friendData.friendMap {
		for revFriendId, fo := range revFriendMap {
			if playerId == friendId {
				friendMap[revFriendId] = fo
			}
			if playerId == revFriendId {
				friendMap[friendId] = fo
			}
		}
	}
	return
}

func (fs *friendService) getFriendFromMap(friendMap map[int64]*FriendObject, friendId int64) (fo *FriendObject) {
	if len(friendMap) == 0 {
		return
	}
	for peerId, obj := range friendMap {
		if peerId == friendId {
			fo = obj
			return
		}
	}
	return
}

func (fs *friendService) getFriendObj(playerId int64, friendId int64) (fo *FriendObject) {
	friendMap := fs.friendData.friendMap[playerId]
	revFriendMap := fs.friendData.revFriendMap[playerId]
	if friendMap == nil && revFriendMap == nil {
		return
	}
	if revFriendMap != nil {
		fo = fs.getFriendFromMap(revFriendMap, friendId)
	}
	if fo != nil {
		return
	}
	if friendMap != nil {
		fo = fs.getFriendFromMap(friendMap, friendId)
		return
	}
	return
}

func (fs *friendService) deleteFriendFromMap(friendMap map[int64]*FriendObject, friendId int64) {
	if len(friendMap) == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	fo, exist := friendMap[friendId]
	if !exist {
		return
	}
	fo.UpdateTime = now
	fo.DeleteTime = now
	fo.SetModified()
	delete(friendMap, friendId)
}

func (fs *friendService) deleteFriendObj(playerId int64, friendId int64) {
	friendMap := fs.friendData.friendMap[playerId]
	revFriendMap := fs.friendData.revFriendMap[playerId]
	if friendMap == nil && revFriendMap == nil {
		return
	}

	if friendMap != nil {
		fs.deleteFriendFromMap(friendMap, friendId)
	}
	if revFriendMap != nil {
		fs.deleteFriendFromMap(revFriendMap, friendId)
	}
}

func (fs *friendService) GetFriendList(playerId int64) (friendMap map[int64]*FriendObject) {
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	friendMap = fs.getFriendList(playerId)
	return
}

func (fs *friendService) getFriendNum(friendId int64) (num int32) {
	num = 0
	friendMap := fs.friendData.friendMap[friendId]
	if friendMap != nil {
		num += int32(len(friendMap))
	}
	revFriendMap := fs.friendData.revFriendMap[friendId]
	if revFriendMap != nil {
		num += int32(len(revFriendMap))
	}
	return
}

func (fs *friendService) AddFriend(pl player.Player, friendId int64) (fo *FriendObject, err error) {
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	obj := fs.getFriendObj(pl.GetId(), friendId)
	if obj != nil {
		return obj, nil
	}

	//判断对方好友数好友是否达上限
	numFriends := fs.getFriendNum(friendId)
	maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
	if numFriends >= maxFriends {
		err = ErrorFriendPeerAlreadFull
		return
	}

	fo = fs.newFriendObj(pl.GetId(), friendId)
	gameevent.Emit(friendeventtypes.EventTypeFriendAdd, pl, fo)
	return
}

func (fs *friendService) DeleteFriend(pl player.Player, friendId int64) (err error) {
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	obj := fs.getFriendObj(pl.GetId(), friendId)
	if obj == nil {
		err = ErrorFriendIsNotFriend
		return
	}
	gameevent.Emit(friendeventtypes.EventTypeFriendDelete, pl, obj)
	fs.deleteFriendObj(pl.GetId(), friendId)
	fs.deleteFriendObj(friendId, pl.GetId())
	return
}

func (fs *friendService) AddPoint(pl player.Player, friendId int64, pointNum int32) {
	if pointNum <= 0 {
		return
	}
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	obj := fs.getFriendObj(pl.GetId(), friendId)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.Point += pointNum
	obj.UpdateTime = now
	obj.SetModified()
	gameevent.Emit(friendeventtypes.EventTypeFriendPointChanged, pl, obj)
}

func (fs *friendService) DivorceSubPoint(pl player.Player, spouseId int64, divorceType marrytypes.MarryDivorceType, percent float64) (point int32) {
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	obj := fs.getFriendObj(pl.GetId(), spouseId)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if divorceType == marrytypes.MarryDivorceTypeForce {
		obj.Point = 0
	} else {
		obj.Point = int32(math.Ceil(float64(obj.Point) * percent))
	}
	obj.UpdateTime = now
	obj.SetModified()
	gameevent.Emit(friendeventtypes.EventTypeFriendPointChanged, pl, obj)
	return obj.Point
}

func (fs *friendService) GetMarryDevelopLogByTime(time int64) []*FriendMarryDevelopLogObject {
	fs.rwm.RLock()
	defer fs.rwm.RUnlock()

	for index, log := range fs.marryDevelopLogList {
		if time < log.UpdateTime {
			return fs.marryDevelopLogList[index:]
		}
	}

	return nil
}

func (fs *friendService) AddMarryDevelopLog(logData *friendtypes.MarryDevelopLogData) {
	fs.rwm.Lock()
	defer fs.rwm.Unlock()

	fs.appendMarryDevelopLog(logData)
}

func (fs *friendService) appendMarryDevelopLog(logData *friendtypes.MarryDevelopLogData) {
	obj := fs.createMarryDevelopLogObj(logData)
	fs.marryDevelopLogList = append(fs.marryDevelopLogList, obj)
}

func (fs *friendService) createMarryDevelopLogObj(logData *friendtypes.MarryDevelopLogData) *FriendMarryDevelopLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMarryDevelopLogMaxNum)
	var obj *FriendMarryDevelopLogObject
	if len(fs.marryDevelopLogList) >= int(maxLogLen) {
		obj = fs.marryDevelopLogList[0]
		fs.marryDevelopLogList = fs.marryDevelopLogList[1:]
	} else {
		obj = NewFriendMarryDevelopLogObject()
		id, _ := idutil.GetId()
		obj.Id = id
		obj.ServerId = global.GetGame().GetServerIndex()
		obj.CreateTime = now
	}

	obj.SendId = logData.SendId
	obj.RecvId = logData.RecvId
	obj.SendName = logData.SendName
	obj.RecvName = logData.RecvName
	obj.ItemId = logData.ItemId
	obj.ItemNum = logData.ItemNum
	obj.CharmNum = logData.CharmNum
	obj.DevelopExp = logData.DevelopExp
	obj.ContextStr = logData.ContextStr
	obj.UpdateTime = now
	obj.SetModified()

	return obj
}

var (
	once sync.Once
	cs   *friendService
)

func Init() (err error) {
	once.Do(func() {
		cs = &friendService{}
		err = cs.init()
	})
	return err
}

func GetFriendService() FriendService {
	return cs
}
