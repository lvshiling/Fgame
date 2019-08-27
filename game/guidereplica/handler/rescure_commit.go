package handler

import (
	lang "fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	constant "fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	guidereplica "fgame/fgame/game/guidereplica/guidereplica"
	"fgame/fgame/game/guidereplica/pbutil"
	guidereplicascene "fgame/fgame/game/guidereplica/scene"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guidereplica.RegisterGuideCommonOperateHandler(guidereplicatypes.GuideReplicaTypeRescure, guidereplica.GuideCommonOperateHandlerFunc(rescureHerbsCommit))
}

//药草提交请求逻辑
func rescureHerbsCommit(pl player.Player) (err error) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaRescue {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("guidereplica:引导副本提交药草,不是救援副本")
		playerlogic.SendSystemMessage(pl, lang.GuideNotInRescureMap)
		return
	}

	sd := s.SceneDelegate()
	rescureSd, ok := sd.(guidereplicascene.GuideRescureSceneData)
	if !ok {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("guidereplica:引导副本提交药草,不是救援副本")
		playerlogic.SendSystemMessage(pl, lang.GuideNotInRescureMap)
		return
	}

	yiXianNPC := scenetemplate.GetSceneTemplateService().GetQuestNPC(rescureSd.GetGuideTemp().GetRescureGuideTemp().RescureBiologyId)
	distanceOk := false
	if yiXianNPC != nil {
		distance := coreutils.Distance(yiXianNPC.GetPos(), pl.GetPos())
		commitDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
		if distance <= float64(commitDistance) {
			distanceOk = true
		}
	}

	if !distanceOk {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("guidereplica:引导副本提交药草,不在提交范围内")
		playerlogic.SendSystemMessage(pl, lang.GuideCommitNoDistance)
		return
	}

	guideTemp := rescureSd.GetGuideTemp()
	buffTemp := guideTemp.GetRescureGuideTemp().GetHerbsBuffTemplate()
	hasHerbsBuff := pl.GetBuff(buffTemp.Group) != nil
	if !hasHerbsBuff {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("guidereplica:引导副本提交药草,身上没有草药")
		playerlogic.SendSystemMessage(pl, lang.GuideNotHerbs)
		return
	}

	//移除草药buff
	scenelogic.RemoveBuff(pl, int32(buffTemp.TemplateId()))
	// scMsg := pbutil.BuildSCGuideReplicaSceneDataChangedNoticeWithRescure(int32(rescureSd.GetGuideTemp().GetGuideType()), false)
	// pl.SendMsg(scMsg)

	//场景完成
	s.Finish(true)
	scCommonMsg := pbutil.BuildSCGuideReplicaRescureCommitHerbs(int32(rescureSd.GetGuideTemp().GetGuideType()))
	pl.SendMsg(scCommonMsg)
	return
}
