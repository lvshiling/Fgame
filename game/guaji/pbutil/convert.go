package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	guajitypes "fgame/fgame/game/guaji/types"
)

func ConvertFromGuaJiDataList(dataList []*uipb.GuaJiData) (guaJiDataList []*guajitypes.GuaJiData, flag bool) {
	for _, data := range dataList {
		guaJiData, flag := ConvertFromGuaJiData(data)
		if !flag {
			return nil, false
		}
		guaJiDataList = append(guaJiDataList, guaJiData)
	}
	flag = true
	return
}

func ConvertFromGuaJiData(data *uipb.GuaJiData) (guaJiData *guajitypes.GuaJiData, flag bool) {

	guaJiType := guajitypes.GuaJiType(data.GetGuaJiType())
	if !guaJiType.Valid() {
		return
	}
	guaJiOptionMap := make(map[guajitypes.GuaJiOptionType]int32)
	optionList := data.GetOptions()
	if len(optionList) == 0 {
		flag = true
		guaJiData = guajitypes.CreateGuaJiData(guaJiType, guaJiOptionMap)
		return
	}
	f := guajitypes.GetGuaJiOptionTypeFactory(guaJiType)
	if f == nil {
		flag = true
		guaJiData = guajitypes.CreateGuaJiData(guaJiType, guaJiOptionMap)
		return
	}
	for _, v := range optionList {
		guaJiOptionType := f.CreateGuaJiOptionType(v.GetOptionType())
		if !guaJiOptionType.Valid() {
			return
		}
		guaJiOptionMap[guaJiOptionType] = v.GetOptionValue()
	}
	flag = true
	guaJiData = guajitypes.CreateGuaJiData(guaJiType, guaJiOptionMap)
	return
}

func ConvertFromGuaJiAdvnaceSettingDataList(dataList []*uipb.GuaJiAdvanceSettingData) (advanceTypeMap map[guajitypes.GuaJiAdvanceType]int32, flag bool) {
	advanceTypeMap = make(map[guajitypes.GuaJiAdvanceType]int32)
	for _, data := range dataList {
		advanceType := guajitypes.GuaJiAdvanceType(data.GetAdvanceType())
		if !advanceType.Valid() {
			return
		}
		advanceTypeMap[advanceType] = data.GetAdvanceValue()
	}
	flag = true
	return
}
