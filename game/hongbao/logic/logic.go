package logic

import (
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/hongbao/hongbao"
	hongbaotemplate "fgame/fgame/game/hongbao/template"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//塞红包
func HongBaoPlugAward(pl player.Player, hongBaoType itemtypes.ItemHongBaoSubType, countMax int32) (id int64, err error) {
	curTemplate := hongbaotemplate.GetHongBaoTemplateService().GetHongBaoByTemplateType(hongBaoType)
	if curTemplate == nil {
		err = fmt.Errorf("hongbao: HongBaoPlugAward error")
		return
	}

	if countMax > curTemplate.CountMax || countMax < curTemplate.CountMin {
		err = fmt.Errorf("hongbao: snatch max count  error")
		return
	}

	awardArr := make([]*hongbao.AwardInfo, 0, countMax)
	//最佳号码
	goodProportionNum := mathutils.RandomRange(int(1), int(countMax))
	//随机普通
	switch hongBaoType {
	case itemtypes.ItemHongBaoSubTypeSilver:
		tempArr := curTemplate.GetAwardNumList(countMax)

		for _, tempNum := range tempArr {
			awardObj := &hongbao.AwardInfo{}
			awardObj.ItemId = constanttypes.SilverItem
			awardObj.ItemCnt = tempNum
			awardArr = append(awardArr, awardObj)
		}

		break
	case itemtypes.ItemHongBaoSubTypeGold:
		tempArr := curTemplate.GetAwardNumList(countMax)

		for _, tempNum := range tempArr {
			awardObj := &hongbao.AwardInfo{}
			awardObj.ItemId = constanttypes.BindGoldItem
			awardObj.ItemCnt = tempNum
			awardArr = append(awardArr, awardObj)
		}
		break
	case itemtypes.ItemHongBaoSubTypeZhenXi:
		for i := 1; i <= int(countMax); i++ {
			awardObj := &hongbao.AwardInfo{}
			dropId := curTemplate.ProportionMin
			if i == goodProportionNum {
				dropId = curTemplate.GoodProportionMax
			}
			dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
			if dropData == nil {
				log.WithField("dropId", dropId).Warn("塞红包:掉落包随机为空")
				err = fmt.Errorf("hongbao: HongBaoPlugAward error")
				return
			}
			awardObj.ItemId = dropData.ItemId
			awardObj.ItemCnt = dropData.Num
			awardObj.Level = dropData.Level
			awardArr = append(awardArr, awardObj)
		}
		break
	default:
		err = fmt.Errorf("hongbao: HongBaoPlugAward error")
		return
	}

	hongBaoService := hongbao.GetHongBaoService()
	hongBaoObj := hongBaoService.CreateHongBaoObj(awardArr, curTemplate.GetHongBaoType(), pl.GetId())
	id = hongBaoObj.GetDBId()
	return
}

//打乱顺序
func disorderAward(awardArr []*hongbao.AwardInfo) []*hongbao.AwardInfo {
	var weights []int64
	for i := 0; i < len(awardArr); i++ {
		weights = append(weights, 1)
	}

	var newAwardArr []*hongbao.AwardInfo
	orderIdxList := mathutils.RandomListFromWeights(weights, int32(len(awardArr)))
	for j := 0; j < len(orderIdxList); j++ {
		newAwardArr = append(newAwardArr, awardArr[orderIdxList[j]])
	}
	return newAwardArr
}
