package data

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	chuangshipb "fgame/fgame/cross/chuangshi/pb"
	alliancetypes "fgame/fgame/game/alliance/types"
)

//阵营工资分配参数
type CamPayScheduleParam struct {
	CityId int64 //城池id
	Ratio  int32 //分配比例
}

func ConvertToCampPayScheduleParamList(paramList []*uipb.CampPaySchedule) (infoList []*CamPayScheduleParam) {
	for _, param := range paramList {
		d := &CamPayScheduleParam{}
		d.CityId = param.GetCityId()
		d.Ratio = param.GetRatio()

		infoList = append(infoList, d)
	}
	return
}

func CrossConvertToCampPayScheduleParamList(infoList []*chuangshipb.CampPaySchedule) (paramList []*CamPayScheduleParam) {
	for _, info := range infoList {
		d := &CamPayScheduleParam{}
		d.CityId = info.GetCityId()
		d.Ratio = info.GetRatio()

		paramList = append(paramList, d)
	}

	return
}

//城池工资分配参数
type CityPayScheduleParam struct {
	AlPos alliancetypes.AlliancePosition //仙盟职位
	Ratio int32                          //分配比例

}

func ConvertToCityPayScheduleParamList(paramList []*uipb.CityPaySchedule) (infoList []*CityPayScheduleParam) {
	for _, param := range paramList {
		d := &CityPayScheduleParam{}
		d.AlPos = alliancetypes.AlliancePosition(param.GetAlPos())
		d.Ratio = param.GetRatio()

		infoList = append(infoList, d)
	}
	return
}

func CrossConvertToCityPayScheduleParamList(infoList []*chuangshipb.CityPaySchedule) (paramList []*CityPayScheduleParam) {
	for _, info := range infoList {
		d := &CityPayScheduleParam{}
		d.AlPos = alliancetypes.AlliancePosition(info.GetAlPos())
		d.Ratio = info.GetRatio()

		paramList = append(paramList, d)
	}

	return
}

type CityPayScheduleList []*CityPayScheduleParam

func (t CityPayScheduleList) GetPayRatio(alPos alliancetypes.AlliancePosition) int32 {
	for _, param := range t {
		if param.AlPos != alPos {
			continue
		}

		return param.Ratio
	}

	return 0
}
