package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/chuangshi/chuangshi"
	"fgame/fgame/cross/chuangshi/pbutil"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_CHUANGSHI_ENTER_CITY_TYPE), dispatch.HandlerFunc(handleChuangShiEnterCity))
}

//处理进入创世城池
func handleChuangShiEnterCity(s session.Session, msg interface{}) (err error) {
	log.Debug("chuangshi:处理进入创世城池")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMsg := msg.(*crosspb.SIChuangShiEnterCity)
	cityId := siMsg.GetCityId()

	err = chuangShiEnterCity(tpl, cityId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("chuangshi:处理进入创世城池,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chuangshi:处理进入创世城池,完成")
	return nil

}

//进入创世城池
func chuangShiEnterCity(pl *player.Player, cityId int64) (err error) {
	isLineup := false

	cityData := chuangshi.GetChuangShiService().GetChuangShiCityData(cityId)
	if cityData == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"cityId":   cityId,
			}).Warnln("chuangshi:处理进入创世城池,城池不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		pl.Close(nil)
		return
	}

	campType := cityData.GetCity().GetCampType()
	cityType := cityData.GetCity().GetType()
	index := cityData.GetCity().GetIndex()

	cityTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(campType, cityType, index)
	if cityTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"campType": campType,
				"cityType": cityType,
				"index":    index,
			}).Warnln("chuangshi:玩家进入城池错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if !cityTemp.IfCanEnter(now) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"campType": campType,
				"cityType": cityType,
				"index":    index,
			}).Warnln("chuangshi:玩家进入城池错误，不是进入时间")
		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotEnterCityTime)
		return
	}

	var s scene.Scene
	switch cityType {
	case chuangshitypes.ChuangShiCityTypeMain:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiMainScene(campType)
		}
	case chuangshitypes.ChuangShiCityTypeZhongli:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiZhongLiScene()
		}
	case chuangshitypes.ChuangShiCityTypeFushu:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiFuShuScene(cityId)
		}
	}

	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"campType": campType,
			}).Warnln("login:玩家进入城池错误，场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		pl.Close(nil)
		return nil
	}

	switch cityType {
	case chuangshitypes.ChuangShiCityTypeMain:
		{
			sd, ok := s.SceneDelegate().(chuangshiscene.MainSceneData)
			if !ok {
				return nil
			}
			isLineup = sd.IfLineup()
		}
	case chuangshitypes.ChuangShiCityTypeZhongli:
		{
			sd, ok := s.SceneDelegate().(chuangshiscene.ZhongLiSceneData)
			if !ok {
				return nil
			}
			isLineup = sd.IfLineup()
		}
	case chuangshitypes.ChuangShiCityTypeFushu:
		{
			sd, ok := s.SceneDelegate().(chuangshiscene.ZhongLiSceneData)
			if !ok {
				return nil
			}
			isLineup = sd.IfLineup()
		}
	}

	isMsg := pbutil.BuildISChuangShiEnterCity(isLineup)
	pl.SendMsg(isMsg)
	return
}
