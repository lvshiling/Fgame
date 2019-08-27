package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIChuangShiEnterCity(cityId int64) *crosspb.SIChuangShiEnterCity {
	siMsg := &crosspb.SIChuangShiEnterCity{}
	siMsg.CityId = &cityId
	return siMsg
}
