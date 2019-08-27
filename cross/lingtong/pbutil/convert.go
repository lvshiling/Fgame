package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/game/lingtong/lingtong"
)

//基础数据
func ConvertFromLingTongShowData(data *crosspb.LingTongShowData) *lingtong.LingTongShowObject {
	fashionId := data.GetFashionId()
	weaponId := data.GetWeaponId()
	weaponState := int32(0)
	titleId := data.GetTitleId()
	wingId := data.GetWingId()
	mountId := data.GetMountId()
	shenFaId := data.GetShenFaId()
	lingYuId := data.GetLingYuId()
	faBaoId := data.GetFaBaoId()
	xianTiId := data.GetXianTiId()
	obj := lingtong.CreateLingTongShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		true,
		shenFaId,
		lingYuId,
		faBaoId,
		xianTiId,
	)
	return obj
}
