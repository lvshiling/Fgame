package pbutil

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	uipb "fgame/fgame/common/codec/pb/ui"
	buffcommon "fgame/fgame/game/buff/common"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/scene/scene"
)

func BuildObjectBuffList(obj scene.BattleObject) (buffDataList []*scenepb.ObjectBuffData) {
	for _, bo := range obj.GetBuffs() {
		buffDataList = append(buffDataList, BuildObjectBuffData(obj, bo))
	}
	return buffDataList
}

func BuildObjectBuffData(so scene.SceneObject, bo buffcommon.BuffObject) *scenepb.ObjectBuffData {
	bd := &scenepb.ObjectBuffData{}
	uid := so.GetId()
	bd.Uid = &uid
	typ := int32(so.GetSceneObjectType())
	bd.ObjecType = &typ
	buffId := bo.GetGroupId()

	bd.BuffId = &buffId
	remainTime := bo.GetRemainTime()
	buffTime := float32(remainTime) / float32(common.SECOND)
	if buffTime <= 0 {
		buffTime = float32(9999999)
	}
	bd.BuffTime = &buffTime
	return bd
}

func BuildSCObjectBuff(so scene.SceneObject, bo buffcommon.BuffObject) *scenepb.SCObjectBuff {
	bd := &scenepb.SCObjectBuff{}
	bd.ObjectBuffList = append(bd.ObjectBuffList, BuildObjectBuffData(so, bo))
	return bd
}

func BuildSCObjectBuffRemove(so scene.SceneObject, bo buffcommon.BuffObject) *scenepb.SCObjectBuffRemove {
	bd := &scenepb.SCObjectBuffRemove{}
	bd.ObjectBuffList = append(bd.ObjectBuffList, BuildObjectBuffData(so, bo))
	return bd
}

func BuildSCBuffList(buffList map[int32]buffcommon.BuffObject) *uipb.SCBuffList {
	bd := &uipb.SCBuffList{}
	for _, bo := range buffList {
		bd.BuffList = append(bd.BuffList, BuildBuffData(bo))
	}
	return bd
}

func BuildSCBuffSearch(buffId int32, flag bool) *uipb.SCBuffSearch {
	bd := &uipb.SCBuffSearch{}
	bd.BuffId = &buffId
	bd.Result = &flag
	return bd
}

func BuildBuffData(bo buffcommon.BuffObject) *uipb.BuffData {
	bd := &uipb.BuffData{}
	buffId := bo.GetBuffId()
	bd.BuffId = &buffId
	remainTime := bo.GetRemainTime()
	buffTime := float32(remainTime) / float32(common.SECOND)
	bd.BuffTime = &buffTime
	return bd
}
