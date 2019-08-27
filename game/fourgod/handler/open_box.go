package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	pbuitl "fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// (废弃-走通用采集)
func init() {
	// processor.Register(codec.MessageType(uipb.MessageType_CS_FOURGOD_OPEN_BOX_TYPE), dispatch.HandlerFunc(handleFourGodOpenBox))
}

//处理开宝箱信息
func handleFourGodOpenBox(s session.Session, msg interface{}) (err error) {
	log.Debug("fourgod:处理开宝箱信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFourGodOpenBox := msg.(*uipb.CSFourGodOpenBox)
	npcId := csFourGodOpenBox.GetNpcId()
	err = fourGodOpenBox(tpl, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"error":    err,
			}).Error("fourgod:处理开宝箱信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fourgod:处理开宝箱信息完成")
	return nil
}

//处理开宝箱信息逻辑
func fourGodOpenBox(pl player.Player, npcId int64) (err error) {
	//判断场景
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}
	sd := s.SceneDelegate()
	sceneData := sd.(fourgodscene.FourGodWarSceneData)
	//参数校验
	npc := sceneData.GetNpc(npcId)
	if npc == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断是否是宝箱
	if npc.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeFourGodCollect {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:不是宝箱")
		playerlogic.SendSystemMessage(pl, lang.FourGodNPCNoBox)
		return
	}

	//判断宝箱是否存在
	if npc.IsDead() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:该宝箱已被开启")
		playerlogic.SendSystemMessage(pl, lang.FourGodBoxIsOpened)
		return
	}

	distance := coreutils.Distance(npc.GetPosition(), pl.GetPos())
	collectDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
	if distance > float64(collectDistance) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:不在采集范围内")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectNoDistance)
		return
	}

	collectBoxInfo, flag := sceneData.IfCanCollectBox(npcId)
	if collectBoxInfo.GetPlayerId() == pl.GetId() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:宝箱正在开启中")
		playerlogic.SendSystemMessage(pl, lang.FourGodOpenBoxRepeat)
		return
	}
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:该宝箱有人正在开启")
		playerlogic.SendSystemMessage(pl, lang.FourGodOpenBoxIsExist)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := manager.GetKeyNum()
	// keyMax := template.GetFourGodTemplateService().GetFourGodConstTemplate().KeyMax
	// useKeyNum := keyNum
	// if keyNum > keyMax {
	// 	useKeyNum = keyMax
	// }
	//boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplate(useKeyNum)

	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	boxTemplate := template.GetFourGodTemplateService().GetFourGodBoxTemplateByBiologyId(biologyId)
	if boxTemplate == nil {
		return
	}
	if keyNum < boxTemplate.UseItemCount {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("fourgod:钥匙数不够")
		keyNumStr := fmt.Sprintf("%d", boxTemplate.UseItemCount)
		playerlogic.SendSystemMessage(pl, lang.FourGodKeyNoEnough, keyNumStr)
		return
	}

	flag = sceneData.CollectBox(npcId, pl.GetId())
	if !flag {
		panic("fourgod:四神遗迹采集宝箱应该是ok的")
	}
	scFourGodOpenBox := pbuitl.BuildSCFourGodOpenBox(npcId)
	pl.SendMsg(scFourGodOpenBox)
	return
}
