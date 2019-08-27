package chess

import (
	"fgame/fgame/core/runner"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"

	"sync"

	log "github.com/Sirupsen/logrus"
)

//苍龙棋局接口处理
type ChessService interface {
	runner.Task
	//添加日志
	AddLog(plName string, itemId, num int32)
	//获取日志
	GetLogByTime(time int64) []*ChessLogObject
	//苍龙棋局前三次固定
	GetSpecialHandleChessItemId(role playertypes.RoleType, attendNum int32) int32
	//清空日志
	GMClearLog()
}

type chessService struct {
	rwm sync.RWMutex
	//棋局日志列表
	chessLogList []*ChessLogObject
	//上次系统插入日志时间
	lastAddDummyLogLime int64
}

//初始化
func (s *chessService) init() (err error) {
	// entityList, err := dao.GetChessDao().GetChessLogEntityList()
	// if err != nil {
	// 	return
	// }

	// for _, entity := range entityList {
	// 	logObj := NewChessLogObject()
	// 	logObj.FromEntity(entity)
	// 	s.chessLogList = append(s.chessLogList, logObj)
	// }

	return
}

//TODO 修改为n秒加一次
//心跳
func (s *chessService) Heartbeat() {

	err := s.addDummyLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("chess:系统生成假日志,错误")
		return
	}
}

//生成系统假日志
func (s *chessService) addDummyLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := s.lastAddDummyLogLime
	diffTime := now - lastTime
	randTime := s.getRandomLogTime()
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	chessTemp := chesstemplate.GetChessTemplateService().GetChessRandom(chesstypes.RandomChessType(), 0)
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(chessTemp.DropId)
	if dropData == nil {
		return
	}
	s.AddLog(name, dropData.ItemId, dropData.Num)

	s.lastAddDummyLogLime = now
	return
}

func (s *chessService) GetLogList() []*ChessLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.chessLogList
}

func (s *chessService) AddLog(plName string, itemId, num int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.appendLog(plName, itemId, num)
}

func (s *chessService) GetSpecialHandleChessItemId(role playertypes.RoleType, attendNum int32) (dropId int32) {
	// TODO:xzk:优化，封装成一个map[role]map[attendNum]dropId
	switch role {
	case playertypes.RoleTypeKaiTian:
		if attendNum == 1 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiTianSpecialChessDropId1))
		}
		if attendNum == 2 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiTianSpecialChessDropId2))
		}
		if attendNum == 3 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeKaiTianSpecialChessDropId3))
		}
	case playertypes.RoleTypeYiJian:
		if attendNum == 1 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeYiJianSpecialChessDropId1))
		}
		if attendNum == 2 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeYiJianSpecialChessDropId2))
		}
		if attendNum == 3 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeYiJianSpecialChessDropId3))
		}
	case playertypes.RoleTypePoYue:
		if attendNum == 1 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePoYueSpecialChessDropId1))
		}
		if attendNum == 2 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePoYueSpecialChessDropId2))
		}
		if attendNum == 3 {
			dropId = int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePoYueSpecialChessDropId3))
		}
	}

	return dropId
}

func (s *chessService) GetLogByTime(time int64) []*ChessLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for index, log := range s.chessLogList {
		if time < log.updateTime {
			return s.chessLogList[index:]
		}
	}

	return nil
}

func (s *chessService) GMClearLog() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	var empty []*ChessLogObject
	s.chessLogList = empty
}

func (s *chessService) appendLog(playerName string, itemId, itemNum int32) {
	obj := s.createLogObj(playerName, itemId, itemNum)
	s.chessLogList = append(s.chessLogList, obj)
}

func (s *chessService) createLogObj(playerName string, itemId, itemNum int32) *ChessLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChessLogMaxNum)
	var obj *ChessLogObject
	if len(s.chessLogList) >= int(maxLogLen) {
		obj = s.chessLogList[0]
		s.chessLogList = s.chessLogList[1:]
	} else {
		obj = NewChessLogObject()
		id, _ := idutil.GetId()
		obj.id = id
		obj.serverId = global.GetGame().GetServerIndex()
		obj.createTime = now
	}

	obj.playerName = playerName
	obj.itemId = itemId
	obj.itemNum = itemNum
	obj.updateTime = now
	// obj.SetModified()

	return obj
}

//系统假日志生成间隔
func (s *chessService) getRandomLogTime() int64 {
	min := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChessLogAddTimeMin))
	max := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChessLogAddTimeMax))
	randTime := int64(mathutils.RandomRange(min, max))
	return randTime
}

var (
	once sync.Once
	cs   *chessService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chessService{}
		err = cs.init()
	})
	return err
}

func GetChessService() ChessService {
	return cs
}
