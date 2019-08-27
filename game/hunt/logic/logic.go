package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	hunttemplate "fgame/fgame/game/hunt/template"
	hunttypes "fgame/fgame/game/hunt/types"
)

func CountHuntDropItemList(huntType hunttypes.HuntType, curHuntCount int32, addTimes int32) (rewList []*droptemplate.DropItemData) {
	huntTemplate := hunttemplate.GetHuntTemplateService().GetHuntTemplat(huntType)
	for i := int32(0); i < addTimes; i++ {
		curHuntCount += 1
		dropId := huntTemplate.DropId
		for _, times := range huntTemplate.GetDropTimesDescList() {
			mustGet := int32(times)
			ret := curHuntCount % mustGet
			if ret == 0 {
				dropId = huntTemplate.GetRewDropMap()[mustGet]
				break
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData != nil {
			rewList = append(rewList, dropData)
		}
	}

	return
}
