package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/cross/lingtong/pbutil"
	playerpbutil "fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/lingtong/lingtong"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_INIT_TYPE), dispatch.HandlerFunc(handleLingTongDataInit))
}

//战斗数据变化
func handleLingTongDataInit(s session.Session, msg interface{}) error {
	log.Info("lingtong:处理跨服灵童数据")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siLingTongDataInit := msg.(*crosspb.SILingTongDataInit)

	err := lingTongDataInit(pl, siLingTongDataInit.GetLingTongData())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
				"err":      err,
			}).Error("lingtong:处理跨服灵童数据,失败")
		return err
	}

	log.Info("lingtong:处理跨服灵童数据,完成")
	return nil
}

//灵童数据初始化
func lingTongDataInit(pl *player.Player, lingTongData *crosspb.LingTongData) (err error) {
	if pl.GetLingTong() != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:处理跨服灵童数据,灵童已经存在")
		return
	}
	id, _ := idutil.GetId()
	pos := coretypes.Position{}
	angle := float64(0)
	if pl.GetScene() != nil {
		pos = pl.GetPosition()
		angle = pl.GetAngle()
	}
	name := lingTongData.GetName()
	lingTongId := lingTongData.GetLingTongId()
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
			}).Warn("lingtong:处理跨服灵童数据,灵童模板不存在")
		return
	}
	//更新灵童
	lingTongShowData := pbutil.ConvertFromLingTongShowData(lingTongData.GetLingTongShowData())
	battleProperties := playerpbutil.ConvertFromBattleProperty(lingTongData.GetBattlePropertyData())
	lingTong := lingtong.CreateLingTong(pl, id, name, pos, angle, lingTongTemplate, lingTongShowData, battleProperties)
	pl.UpdateLingTong(lingTong)
	return nil
}
