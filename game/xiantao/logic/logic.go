package logic

import (
	"fgame/fgame/common/lang"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/xiantao/pbutil"
	playerxiantao "fgame/fgame/game/xiantao/player"
	xiantaotemplate "fgame/fgame/game/xiantao/template"
	xiantaotypes "fgame/fgame/game/xiantao/types"
	"math"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterXianTaoScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	ns := pl.GetScene()
	mapType := ns.MapTemplate().GetMapType()
	if mapType == scenetypes.SceneTypeXianTaoDaHui {
		return
	}
	s := getXianTaoScene(activityTemplate)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xiantao:仙桃大会场景不存在")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}

	bornPos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		return
	}
	flag = true
	return
}

// 仙桃大会结束
func onFinishXianTao(spl scene.Player) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCXiantaoResult()
	pl.SendMsg(scMsg)
}

// 退出仙桃大会
func onExistXianTao(spl scene.Player, sd XianTaoSceneData) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	xianTaoManager := pl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoManager.ExitXianTao()
}

// 进入仙桃大会
func onEnterXianTao(spl scene.Player, sd XianTaoSceneData) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	xianTaoManager := pl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObj := xianTaoManager.GetXianTaoObject()

	endTime := sd.GetScene().GetEndTime()
	pCollectCount := sd.GetPlayerCollectCount(spl.GetId())
	xianTaoManager.EnterXianTao(endTime)

	npcMap := sd.GetScene().GetNPCS(scenetypes.BiologyScriptTypeXianTaoQianNianCollect)
	var cpnList []*collectnpc.CollectPointNPC
	for _, npc := range npcMap {
		cpn, ok := npc.(*collectnpc.CollectPointNPC)
		if !ok {
			continue
		}
		cpnList = append(cpnList, cpn)
	}
	scMsg := pbutil.BuildSCXiantaoGet(xianTaoObj, cpnList, pCollectCount)
	pl.SendMsg(scMsg)

	//buff变化
	PlayerXianTaoChangedBuff(pl)
}

func XianTaoDrop(pl player.Player, spl player.Player) (dropNumMap, gainNumMap map[xiantaotypes.XianTaoType]int32) {
	dropNumMap = make(map[xiantaotypes.XianTaoType]int32)
	gainNumMap = make(map[xiantaotypes.XianTaoType]int32)
	manager := pl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	sManager := spl.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObject := manager.GetXianTaoObject()
	sXianTaoObject := sManager.GetXianTaoObject()

	xianTaoService := xiantaotemplate.GetXianTaoTemplateService()
	constTemp := xianTaoService.GetXianTaoConstTemplate()
	timesTemp := manager.GetNextBeRobTimesTemplate()
	sTimesTemp := sManager.GetNextRobTimesTemplate()

	//劫取者
	if sTimesTemp != nil {
		//百年仙桃
		if sXianTaoObject.JuniorPeachCount < constTemp.XianTaoMax {
			gainNumMap[xiantaotypes.XianTaoTypeBaiNian] = int32(math.Ceil(float64(sTimesTemp.JieQuPercent) / float64(common.MAX_RATE) * float64(xianTaoObject.JuniorPeachCount)))
		}
		//千年仙桃
		if sXianTaoObject.HighPeachCount < constTemp.XianTaoMax {
			gainNumMap[xiantaotypes.XianTaoTypeQianNian] = int32(math.Ceil(float64(sTimesTemp.JieQuPercent) / float64(common.MAX_RATE) * float64(xianTaoObject.HighPeachCount)))
		}
	}
	//被劫取者
	if xianTaoObject.JuniorPeachCount > constTemp.XianTaoMin {
		dropNumMap[xiantaotypes.XianTaoTypeBaiNian] = int32(math.Ceil(float64(timesTemp.SunShiPercent) / float64(common.MAX_RATE) * float64(xianTaoObject.JuniorPeachCount)))
	}
	if xianTaoObject.HighPeachCount > constTemp.XianTaoMin {
		dropNumMap[xiantaotypes.XianTaoTypeQianNian] = int32(math.Ceil(float64(timesTemp.SunShiPercent) / float64(common.MAX_RATE) * float64(xianTaoObject.HighPeachCount)))
	}

	//仙桃数量 改变
	isDrop, isGain := false, false
	for typ, count := range dropNumMap {
		if typ == xiantaotypes.XianTaoTypeQianNian {
			_, subCount := manager.SubHighPeachCount(count)
			dropNumMap[typ] = subCount
			if subCount > 0 {
				isDrop = true
			}
		}
		if typ == xiantaotypes.XianTaoTypeBaiNian {
			_, subCount := manager.SubJuniorPeachCount(count)
			dropNumMap[typ] = subCount
			if subCount > 0 {
				isDrop = true
			}
		}
	}
	for typ, count := range gainNumMap {
		if typ == xiantaotypes.XianTaoTypeQianNian {
			_, addCount := sManager.AddHighPeachCount(count)
			gainNumMap[typ] = addCount
			if addCount > 0 {
				isGain = true
			}
		}
		if typ == xiantaotypes.XianTaoTypeBaiNian {
			_, addCount := sManager.AddJuniorPeachCount(count)
			gainNumMap[typ] = addCount
			if addCount > 0 {
				isGain = true
			}
		}
	}

	//劫取次数
	if isDrop {
		manager.SetBeRobCount(timesTemp.Times)
		PlayerXianTaoInfoChanged(pl)
	}
	if isGain {
		sManager.AddRobCount(1)
		PlayerXianTaoInfoChanged(spl)
	}

	return
}

//仙桃信息变化
func PlayerXianTaoInfoChanged(p player.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}

	sd := s.SceneDelegate()
	xiantaoSd, ok := sd.(XianTaoSceneData)
	if !ok {
		return
	}
	xianTaoManager := p.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObject := xianTaoManager.GetXianTaoObject()
	pCollectCount := xiantaoSd.GetPlayerCollectCount(p.GetId())
	scMsg := pbutil.BuildSCXiantaoPlayerAttendChange(xianTaoObject, pCollectCount)
	p.SendMsg(scMsg)
}

//玩家仙桃变化
func PlayerXianTaoChangedBuff(p player.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObject := manager.GetXianTaoObject()
	juniorPeachCount := xianTaoObject.JuniorPeachCount
	highPeachCount := xianTaoObject.HighPeachCount
	xianTaoConstTemplate := xiantaotemplate.GetXianTaoTemplateService().GetXianTaoConstTemplate()
	hasJuniorBuff := p.GetBuff(xianTaoConstTemplate.GetHundredPeachBuffTemplate().Group) != nil
	hasHighBuff := p.GetBuff(xianTaoConstTemplate.GetThousandPeachBuffTemplate().Group) != nil
	hasBuSunBuff := p.GetBuff(xianTaoConstTemplate.GetBuSunPeachBuffTemplate().Group) != nil

	if hasHighBuff {
		//移除buff
		if highPeachCount <= 0 {
			scenelogic.RemoveBuff(p, xianTaoConstTemplate.XianTaoBuff)
		}
	} else {
		if highPeachCount > 0 {
			scenelogic.AddBuff(p, xianTaoConstTemplate.XianTaoBuff, p.GetId(), common.MAX_RATE)
		}
	}

	if hasJuniorBuff {
		//移除buff
		if juniorPeachCount <= 0 || highPeachCount > 0 {
			scenelogic.RemoveBuff(p, xianTaoConstTemplate.XianTaoBuff2)
		}
	} else {
		if juniorPeachCount > 0 && highPeachCount <= 0 {
			scenelogic.AddBuff(p, xianTaoConstTemplate.XianTaoBuff2, p.GetId(), common.MAX_RATE)
		}
	}

	//不损buff
	if hasBuSunBuff {
		isInit := juniorPeachCount == 0 && highPeachCount == 0
		if isInit || juniorPeachCount > xianTaoConstTemplate.XianTaoMin || highPeachCount > xianTaoConstTemplate.XianTaoMin {
			scenelogic.RemoveBuff(p, xianTaoConstTemplate.BuSunBuff)
		}
	} else if juniorPeachCount > 0 || highPeachCount > 0 {
		if juniorPeachCount <= xianTaoConstTemplate.XianTaoMin && highPeachCount <= xianTaoConstTemplate.XianTaoMin {
			scenelogic.AddBuff(p, xianTaoConstTemplate.BuSunBuff, p.GetId(), common.MAX_RATE)
		}
	}

}
