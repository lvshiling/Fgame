package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_SHOW_DATA_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerShowDataChanged))
}

//处理外观属性推送
func handlePlayerShowDataChanged(s session.Session, msg interface{}) error {
	log.Debug("login:处理跨服系统属性推送消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerShowDataChanged := msg.(*crosspb.SIPlayerShowDataChanged)

	err := playerShowChanged(pl, siPlayerShowDataChanged.GetPlayerShowData())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家跨服登陆,失败")
		return err
	}

	log.Debug("login:处理跨服登陆消息完成")
	return nil
}

//玩家显示变化
func playerShowChanged(pl *player.Player, playerShowData *crosspb.PlayerShowData) (err error) {
	if pl.PlayerShowManager == nil {
		return
	}
	if playerShowData.TitleId != nil {
		pl.SetTitleId(playerShowData.GetTitleId())
	}
	if playerShowData.WeaponId != nil {
		pl.SetWeapon(playerShowData.GetWeaponId(), playerShowData.GetWeaponState())
	}

	if playerShowData.ClothesId != nil {
		pl.SetFashionId(playerShowData.GetClothesId())
	}
	if playerShowData.RideId != nil {
		pl.SetMountId(playerShowData.GetRideId(), playerShowData.GetMountAdvanceId())
	}
	if playerShowData.WingId != nil {
		pl.SetWingId(playerShowData.GetWingId())
	}
	if playerShowData.ShenFaId != nil {
		pl.SetShenFaId(playerShowData.GetShenFaId())
	}
	if playerShowData.LingYuId != nil {
		pl.SetLingYuId(playerShowData.GetLingYuId())
	}
	if playerShowData.FourGodKey != nil {
		pl.SetFourGodKey(playerShowData.GetFourGodKey())
	}
	if playerShowData.Realm != nil {
		pl.SetRealm(playerShowData.GetRealm())
	}
	if playerShowData.Spouse != nil {
		pl.SetSpouse(playerShowData.GetSpouse())
	}
	if playerShowData.WeddingStatus != nil {
		pl.SetWeddingStatus(playerShowData.GetWeddingStatus())
	}
	if playerShowData.Model != nil {
		pl.SetModel(playerShowData.GetModel())
	}

	if playerShowData.RingType != nil {
		pl.SetRingType(playerShowData.GetRingType())
	}
	if playerShowData.PetId != nil {
		pl.SetPetId(playerShowData.GetPetId())
	}
	if playerShowData.FaBaoId != nil {
		pl.SetFaBaoId(playerShowData.GetFaBaoId())
	}
	if playerShowData.XianTiId != nil {
		pl.SetXianTiId(playerShowData.GetXianTiId())
	}
	if playerShowData.BaGua != nil {
		pl.SetBaGua(playerShowData.GetBaGua())
	}
	return nil
}
