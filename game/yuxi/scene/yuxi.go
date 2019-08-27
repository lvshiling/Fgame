package scene

import (
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxieventtypes "fgame/fgame/game/yuxi/event/types"
	"fgame/fgame/game/yuxi/pbutil"
	yuxitemplate "fgame/fgame/game/yuxi/template"
	yuxitypes "fgame/fgame/game/yuxi/types"
	"sort"
)

const (
	yuXiBroadcastTaskTime = 10 * int64(common.SECOND) //玉玺信息广播间隔
)

// 玉玺之战结束
func (sd *yuxiSceneData) onFinishYuXi() {
	allianceId := int64(0)
	// 持有玉玺
	if sd.yuXiOwner != nil {
		allianceId = sd.yuXiOwner.GetAllianceId()
	} else {
		allianceId = sd.getHightestForceAl()
	}

	gameevent.Emit(yuxieventtypes.EventTypeYuXiWin, sd, allianceId)
}

// 玉玺重置
func (sd *yuxiSceneData) onYuXiReset(spl scene.Player, rebornType yuxitypes.YuXiReborType) {
	if spl != sd.yuXiOwner {
		return
	}

	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	scenelogic.RemoveBuff(spl, constantTemp.BuffId)

	//重置
	sd.initYuXiCollect(rebornType)
	sd.yuXiOwner = nil
	sd.startTime = 0
}

// 初始化玉玺采集物
func (sd *yuxiSceneData) initYuXiCollect(rebornType yuxitypes.YuXiReborType) {
	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	bornPos := constantTemp.GetYuXiPos()
	newYuxi := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, constantTemp.GetYuXiBiologyTemp(), bornPos, 0, 0)
	sd.GetScene().AddSceneObject(newYuxi)
	sd.yuxiNpc = newYuxi
	sd.onBroadcastYuXiCollect(rebornType)
}

// 玉玺持有者广播
func (sd *yuxiSceneData) onBroadcastYuXiInfo() {
	if sd.yuXiOwner == nil {
		return
	}

	scMsg := pbutil.BuildSCYuXiPosBroadcast(sd.yuXiOwner, sd.startTime)
	sd.GetScene().BroadcastMsg(scMsg)
}

// 玉玺采集物广播
func (sd *yuxiSceneData) onBroadcastYuXiCollect(rebornType yuxitypes.YuXiReborType) {
	if sd.yuxiNpc == nil {
		return
	}

	scMsg := pbutil.BuildSCYuXiCollectInfoBroadcast(sd.yuxiNpc, sd.yuXiOwner, sd.startTime, int32(rebornType))
	sd.GetScene().BroadcastMsg(scMsg)
}

// 玩家进入场景
func (sd *yuxiSceneData) onPlayerEnter(spl scene.Player) {
	scMsg := pbutil.BuildSCYuXiCollectInfoBroadcast(sd.yuxiNpc, sd.yuXiOwner, sd.startTime, int32(yuxitypes.YuXiReborTypeInitNone))
	spl.SendMsg(scMsg)

	allianceId := spl.GetAllianceId()
	_, ok := sd.attendAlMap[allianceId]
	if !ok {
		al := alliance.GetAllianceService().GetAlliance(allianceId)
		alData := &allianceData{
			allianceId: allianceId,
			totalForce: al.GetAllianceObject().GetTotalForce(),
		}
		sd.attendAlMap[allianceId] = alData
	}
}

// 是否能提前结束
func (sd *yuxiSceneData) isCanEnd() bool {
	if sd.yuXiOwner == nil {
		return false
	}

	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	now := global.GetGame().GetTimeService().Now()
	if now > sd.startTime+constantTemp.WinTime {
		return true
	}

	return false
}

// 是否广播玉玺位置
func (sd *yuxiSceneData) isBroadcastYuXiInfo(now int64) bool {
	if sd.yuXiOwner == nil {
		return false
	}

	if now > sd.lastBroadcastTime+yuXiBroadcastTaskTime {
		return true
	}

	return false
}

// 参与过玉玺之战-战力最高的仙盟
func (sd *yuxiSceneData) getHightestForceAl() int64 {
	var attendAlList []*allianceData
	for _, alData := range sd.attendAlMap {
		attendAlList = append(attendAlList, alData)
	}

	if len(attendAlList) == 0 {
		return 0
	}

	sort.Sort(allianceDataSort(attendAlList))
	return attendAlList[len(attendAlList)-1].allianceId
}
