package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	guajitypes "fgame/fgame/game/guaji/types"
	playertypes "fgame/fgame/game/player/types"
)

func BuildCSSelectJob(role playertypes.RoleType, sex playertypes.SexType, name string) *uipb.CSSelectJob {
	csSelectJob := &uipb.CSSelectJob{}
	csSelectJob.Name = &name
	jobInt := int32(role)
	csSelectJob.Job = &jobInt
	sexInt := int32(sex)
	csSelectJob.Sex = &sexInt
	return csSelectJob
}

func BuildMainGuaJi() *uipb.CSGuaJi {
	csGuaJi := &uipb.CSGuaJi{}
	guaJiData := &uipb.GuaJiData{}
	guaJiType := int32(guajitypes.GuaJiTypeMainQuest)
	guaJiData.GuaJiType = &guaJiType
	csGuaJi.GuaJiDataList = append(csGuaJi.GuaJiDataList, guaJiData)
	return csGuaJi
}
