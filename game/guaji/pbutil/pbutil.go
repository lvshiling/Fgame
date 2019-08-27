package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	guajitypes "fgame/fgame/game/guaji/types"
)

func BuildSCGuaJi(guaJiDataList []*guajitypes.GuaJiData, guaJiAdvanceSettingMap map[guajitypes.GuaJiAdvanceType]int32) *uipb.SCGuaJi {
	scGuaJi := &uipb.SCGuaJi{}
	for _, data := range guaJiDataList {
		guaJiData := BuildGuaJiData(data)
		scGuaJi.GuaJiDataList = append(scGuaJi.GuaJiDataList, guaJiData)
	}

	for typ, val := range guaJiAdvanceSettingMap {
		guaJiAdvanceSettingData := BuildGuaJiAdvanceSettingData(typ, val)
		scGuaJi.GuaJiAdvanceSettingDataList = append(scGuaJi.GuaJiAdvanceSettingDataList, guaJiAdvanceSettingData)
	}
	return scGuaJi
}

func BuildGuaJiAdvanceSettingData(typ guajitypes.GuaJiAdvanceType, val int32) *uipb.GuaJiAdvanceSettingData {
	guaJiAdvanceSettingData := &uipb.GuaJiAdvanceSettingData{}
	typInt := int32(typ)
	guaJiAdvanceSettingData.AdvanceType = &typInt
	guaJiAdvanceSettingData.AdvanceValue = &val

	return guaJiAdvanceSettingData
}

func BuildGuaJiData(data *guajitypes.GuaJiData) *uipb.GuaJiData {
	guaJiData := &uipb.GuaJiData{}
	typ := int32(data.GetType())
	guaJiData.GuaJiType = &typ
	for k, v := range data.GetOptions() {
		op := BuildGuaJiOption(k, v)
		guaJiData.Options = append(guaJiData.Options, op)
	}

	return guaJiData
}

func BuildGuaJiOption(optionType guajitypes.GuaJiOptionType, val int32) *uipb.GuaJiOption {
	guaJiOption := &uipb.GuaJiOption{}
	typ := guaJiOption.GetOptionType()
	guaJiOption.OptionType = &typ
	guaJiOption.OptionValue = &val

	return guaJiOption
}

func BuildSCCurrentGuaJi(guaJiType int32) *uipb.SCCurrentGuaJi {
	scCurrentGuaJi := &uipb.SCCurrentGuaJi{}

	scCurrentGuaJi.GuaJiType = &guaJiType
	return scCurrentGuaJi
}

func BuildSCStopGuaJi() *uipb.SCStopGuaJi {
	scStopGuaJi := &uipb.SCStopGuaJi{}
	return scStopGuaJi
}

func BuildSCGuaJiPos(mapId int32, pos coretypes.Position) *uipb.SCGuaJiPos {
	scGuaJiPos := &uipb.SCGuaJiPos{}
	scGuaJiPos.MapId = &mapId
	guaJiPos := &uipb.GuaJiPosition{}
	x := float32(pos.X)
	guaJiPos.PosX = &x
	y := float32(pos.Y)
	guaJiPos.PosY = &y
	z := float32(pos.Z)
	guaJiPos.PosZ = &z
	scGuaJiPos.Pos = guaJiPos
	return scGuaJiPos
}

func BuildSCGuaJiAdvanceList(advanceMap map[guajitypes.GuaJiAdvanceType]int32) *uipb.SCGuaJiAdvanceList {
	scGuaJiAdvanceList := &uipb.SCGuaJiAdvanceList{}
	for typ, val := range advanceMap {
		scGuaJiAdvanceList.GuaJiAdvanceDataList = append(scGuaJiAdvanceList.GuaJiAdvanceDataList, BuildGuaJiAdvanceData(typ, val))
	}

	return scGuaJiAdvanceList
}

func BuildGuaJiAdvanceData(typ guajitypes.GuaJiAdvanceType, val int32) *uipb.GuaJiAdvanceData {
	guaJiAdvanceData := &uipb.GuaJiAdvanceData{}
	typInt := int32(typ)
	guaJiAdvanceData.AdvanceType = &typInt
	guaJiAdvanceData.AdvanceValue = &val

	return guaJiAdvanceData
}

func BuildSCGuaJiAdvanceUpdateList(typ guajitypes.GuaJiAdvanceType, val int32) *uipb.SCGuaJiAdvanceUpdateList {
	scGuaJiAdvanceUpdateList := &uipb.SCGuaJiAdvanceUpdateList{}
	scGuaJiAdvanceUpdateList.GuaJiAdvanceDataList = append(scGuaJiAdvanceUpdateList.GuaJiAdvanceDataList, BuildGuaJiAdvanceData(typ, val))
	return scGuaJiAdvanceUpdateList
}
