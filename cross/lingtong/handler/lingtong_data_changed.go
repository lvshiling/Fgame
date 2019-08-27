package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playerpbutil "fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LING_TONG_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handleLingTongDataChanged))
}

//处理灵童数据变化
func handleLingTongDataChanged(s session.Session, msg interface{}) error {
	log.Info("lingtong:处理灵童数据变化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siLingTongDataChanged := msg.(*crosspb.SILingTongDataChanged)

	err := lingTongDataChanged(pl, siLingTongDataChanged)

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lingtong:处理灵童数据变化,失败")
		return err
	}

	log.Info("lingtong:处理灵童数据变化完成")
	return nil
}

//玩家基础属性变化
func lingTongDataChanged(pl *player.Player, siLingTongDataChanged *crosspb.SILingTongDataChanged) (err error) {
	lingTong := pl.GetLingTong()
	if lingTong == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:处理跨服灵童数据,灵童不存在")
		return
	}
	//更新模板
	if siLingTongDataChanged.LingTongId != nil {
		lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(siLingTongDataChanged.GetLingTongId())
		if lingTongTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"lingTongId": siLingTongDataChanged.GetLingTongId(),
				}).Warn("lingtong:处理跨服灵童数据,灵童模板是空")
			return
		}
		name := siLingTongDataChanged.GetName()
		lingTong.UpdateLingTongTemplate(lingTongTemplate, name)
	}
	//更新名字
	if siLingTongDataChanged.Name != nil {
		lingTong.UpdateName(siLingTongDataChanged.GetName())
	}
	//更新属性
	if siLingTongDataChanged.BattlePropertyData != nil {
		battleProperties := playerpbutil.ConvertFromBattleProperty(siLingTongDataChanged.GetBattlePropertyData())
		lingTong.UpdateSystemBattleProperty(battleProperties)
		lingTong.Calculate()
	}
	//更新外观
	if siLingTongDataChanged.LingTongShowData != nil {
		lingTongShowDataChanged(lingTong, siLingTongDataChanged.GetLingTongShowData())
	}
	return
}

//外观变化
func lingTongShowDataChanged(lingTong scene.LingTong, showData *crosspb.LingTongShowData) {
	if showData.FashionId != nil {
		lingTong.SetLingTongFashionId(showData.GetFashionId())
	}

	if showData.WeaponId != nil {
		lingTong.SetLingTongWeapon(showData.GetWeaponId(), showData.GetWeaponState())
	}
	if showData.TitleId != nil {
		lingTong.SetLingTongTitleId(showData.GetTitleId())
	}
	if showData.WingId != nil {
		lingTong.SetLingTongWingId(showData.GetWingId())
	}
	if showData.MountId != nil {
		lingTong.SetLingTongMountId(showData.GetMountId())
	}
	if showData.ShenFaId != nil {
		lingTong.SetLingTongShenFaId(showData.GetShenFaId())
	}
	if showData.LingYuId != nil {
		lingTong.SetLingTongLingYuId(showData.GetLingYuId())
	}
	if showData.FaBaoId != nil {
		lingTong.SetLingTongFaBaoId(showData.GetFaBaoId())
	}
	if showData.XianTiId != nil {
		lingTong.SetLingTongXianTiId(showData.GetXianTiId())
	}
}
