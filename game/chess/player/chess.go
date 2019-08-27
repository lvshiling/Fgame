package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/chess/chess"
	chesseventtypes "fgame/fgame/game/chess/event/types"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	"fgame/fgame/game/chess/dao"
)

//玩家棋局管理器
type PlayerChessDataManager struct {
	p player.Player
	//玩家棋局对象
	playerChessMap map[chesstypes.ChessType]*PlayerChessObject
}

func (m *PlayerChessDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerChessDataManager) Load() (err error) {
	//加载玩家棋局信息
	chessEntityList, err := dao.GetChessDao().GetPlayerChessEntityList(m.p.GetId())
	if err != nil {
		return
	}

	if len(chessEntityList) > 0 {
		for _, chessEntity := range chessEntityList {
			obj := NewPlayerChessObject(m.p)
			obj.FromEntity(chessEntity)
			m.playerChessMap[obj.chessType] = obj
		}
	} else {
		m.initPlayerChessObject()
	}

	return nil
}

//第一次初始化
func (m *PlayerChessDataManager) initPlayerChessObject() {
	for typ := chesstypes.MinChessType; typ <= chesstypes.MaxChessType; typ++ {
		var chessTemplate *gametemplate.ChessTemplate
		if typ == chesstypes.ChessTypeGold {
			chessTemplate = chesstemplate.GetChessTemplateService().GetChessByRole(m.p.GetRole())
		} else {
			chessTemplate = chesstemplate.GetChessTemplateService().GetChessRandom(typ, 0)
		}
		pwo := NewPlayerChessObject(m.p)
		now := global.GetGame().GetTimeService().Now()
		id, _ := idutil.GetId()
		pwo.id = id
		pwo.chessId = chessTemplate.ChessId
		pwo.chessType = typ
		pwo.attendTimes = int32(0)
		pwo.totalAttendTimes = int32(0)
		pwo.createTime = now
		pwo.lastSystemRefreshTime = now
		pwo.SetModified()

		m.playerChessMap[typ] = pwo
	}
}

//加载后
func (m *PlayerChessDataManager) AfterLoad() (err error) {
	err = m.refreshTimes()
	if err != nil {
		return
	}

	err = m.refreshChess()
	if err != nil {
		return
	}
	return
}

//心跳
func (m *PlayerChessDataManager) Heartbeat() {

}

//获取棋局信息
func (m *PlayerChessDataManager) getChessObj(typ chesstypes.ChessType) *PlayerChessObject {
	obj, ok := m.playerChessMap[typ]
	if ok {
		return obj
	}

	return nil
}

//刷新棋局参与次数
func (m *PlayerChessDataManager) refreshTimes() (err error) {
	for _, chessObj := range m.playerChessMap {
		now := global.GetGame().GetTimeService().Now()
		isSame, err := timeutils.IsSameFive(chessObj.updateTime, now)
		if err != nil {
			return err
		}

		if !isSame {
			chessObj.attendTimes = 0
			chessObj.updateTime = now
			chessObj.SetModified()
		}
	}

	return
}

// 刷新棋局
func (m *PlayerChessDataManager) refreshChess() (err error) {
	for _, chessObj := range m.playerChessMap {
		//特殊处理
		if chessObj.chessType == chesstypes.ChessTypeGold && chessObj.totalAttendTimes < specialHandleChessTimes {
			continue
		}

		now := global.GetGame().GetTimeService().Now()
		lastSystemRefreshTime := chessObj.lastSystemRefreshTime
		isSame, err := timeutils.IsSameFive(lastSystemRefreshTime, now)
		if err != nil {
			return err
		}

		if !isSame {
			chessObj.randomChess()
			continue
		}

		//间隔刷新:2的倍数
		refreshCycleHour := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChessRefreshTime))
		curHour := timeutils.MillisecondToTime(now).Add(-5 * common.HOUR).Hour()
		lastHour := timeutils.MillisecondToTime(lastSystemRefreshTime).Add(-5 * common.HOUR).Hour()
		isRefresh := curHour/refreshCycleHour != lastHour/refreshCycleHour
		if isRefresh {
			chessObj.randomChess()
		}
	}

	return nil
}

//获取棋局列表
func (m *PlayerChessDataManager) GetChessMap() map[chesstypes.ChessType]*PlayerChessObject {
	m.refreshTimes()
	m.refreshChess()

	return m.playerChessMap
}

//次数是否足够
func (m *PlayerChessDataManager) IsEnoughTimes(typ chesstypes.ChessType, attendNum int32) bool {
	if typ == chesstypes.ChessTypeGold {
		return true
	}

	// 子福：2019.01.25特殊处理
	maxTimes := constant.GetConstantService().GetConstant(typ.GetAttendLimitConstantType())
	if maxTimes == 0 {
		return true
	}

	m.refreshTimes()

	obj := m.getChessObj(typ)
	if obj == nil {
		return false
	}

	silverTimes := obj.attendTimes + attendNum
	return silverTimes <= maxTimes
}

//破解棋局
func (m *PlayerChessDataManager) AttendChess(typ chesstypes.ChessType) (flag bool) {
	m.refreshTimes()

	obj := m.getChessObj(typ)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.attendTimes += 1
	obj.totalAttendTimes += 1
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(chesseventtypes.EventTypeAttendChess, m.p, typ)

	flag = true
	return
}

//获取棋局掉落包
func (m *PlayerChessDataManager) GetChessDrop(typ chesstypes.ChessType, times int32) (rewList []*droptemplate.DropItemData) {
	m.refreshTimes()

	obj := m.getChessObj(typ)
	if obj == nil {
		return
	}
	chessTemplate := chesstemplate.GetChessTemplateService().GetChessByTypAndChessId(typ, obj.chessId)
	curTotal := obj.totalAttendTimes
	var logRewItemIdList []int32
	for index := int32(0); index < times; index++ {
		dropId := int32(0)
		curTotal += 1
		if typ == chesstypes.ChessTypeGold && curTotal <= specialHandleChessTimes {
			dropId = chess.GetChessService().GetSpecialHandleChessItemId(m.p.GetRole(), curTotal)
		} else {
			dropId = chessTemplate.DropId

			rewDropByTimesMap := chessTemplate.GetRewDropMap()
			timesDescList := chessTemplate.GetDropTimesDescList()
			ruleTimesMap := viplogic.CountDropTimesWithCostLevel(m.p, viptypes.CostLevelRuleTypeChess, timesDescList)
			for _, times := range timesDescList {
				ruleTimes := ruleTimesMap[times]
				ret := curTotal % ruleTimes
				if ret == 0 {
					dropId = rewDropByTimesMap[int32(times)]
					break
				}
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData != nil {
			rewList = append(rewList, dropData)
			logRewItemIdList = append(logRewItemIdList, dropData.ItemId)
		}
	}

	//棋局日志
	chessReason := commonlog.ChessLogReasonAttendLog
	reasonText := fmt.Sprintf(chessReason.String(), typ, logRewItemIdList)
	eventData := chesseventtypes.CreatePlayerAttendChessLogEventData(times, chessReason, reasonText)
	gameevent.Emit(chesseventtypes.EventTypeAttendChessLog, m.p, eventData)

	return
}

//换一批
func (m *PlayerChessDataManager) ChangedChess(typ chesstypes.ChessType) int32 {
	m.refreshChess()

	obj := m.getChessObj(typ)
	if obj == nil {
		return 0
	}
	newChessDropId := obj.randomChess()
	return newChessDropId
}

//获取当前棋局id
func (m *PlayerChessDataManager) GetChessId(typ chesstypes.ChessType) int32 {
	obj := m.getChessObj(typ)
	if obj == nil {
		return 0
	}
	return obj.chessId
}

//获取剩余次数
func (m *PlayerChessDataManager) GetAllLeftTimesExcludeGold() (leftNum int32) {
	m.refreshTimes()
	for typ := chesstypes.MinChessType; typ <= chesstypes.MaxChessType; typ++ {
		if typ == chesstypes.ChessTypeGold {
			continue
		}
		chessObj := m.getChessObj(typ)
		maxTimes := constant.GetConstantService().GetConstant(typ.GetAttendLimitConstantType())
		leftNum += (maxTimes - chessObj.attendTimes)
	}
	return
}

//GM重置次数
func (m *PlayerChessDataManager) GMResetTimes(typ chesstypes.ChessType) (err error) {
	obj := m.getChessObj(typ)
	if obj == nil {
		return
	}

	obj.attendTimes = 0
	obj.totalAttendTimes = 0
	obj.SetModified()
	return
}

func CreatePlayerChessDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerChessDataManager{}
	m.p = p
	m.playerChessMap = make(map[chesstypes.ChessType]*PlayerChessObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerChessDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerChessDataManager))
}
