package listener

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	commontypes "fgame/fgame/game/common/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fmt"
)

//10个发一次
const sendEmailInTen = 10

//装备宝库跨天事件
func playerEquipBaoKuCrossDay(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	baoKuData, ok := data.(*equipbaokutypes.BaoKuData)
	if !ok {
		return
	}
	typ := baoKuData.Typ
	luckyPoints := baoKuData.LuckyPoints

	equipBaoKuTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetEquipBaoKuByLevAndZhuanNum(pl.GetLevel(), pl.GetZhuanSheng(), typ)
	if equipBaoKuTemplate == nil {
		return
	}

	//幸运值是否足够
	num := int32(luckyPoints / equipBaoKuTemplate.NeedXingYunZhi)
	if num == 0 {
		return
	}

	//掉落
	now := global.GetGame().GetTimeService().Now()
	var rewList []*droptemplate.DropItemData
	var logRewList []*droptemplate.DropItemData
	for num > 0 {
		dropData := droptemplate.GetDropTemplateService().GetDropBaoKuItemLevel(equipBaoKuTemplate.ScriptXingYun)
		if dropData != nil {
			rewList = append(rewList, dropData)
			logRewList = append(logRewList, dropData)
		}
		num -= 1
		if len(rewList)%sendEmailInTen == 0 || num == 0 {
			title := lang.GetLangService().ReadLang(lang.EmailEquipBaoKuLuckyBoxTitle)
			econtent := lang.GetLangService().ReadLang(lang.EmailEquipBaoKuLuckyBoxContent)
			emaillogic.AddEmailItemLevel(pl, title, econtent, now, rewList)
			rewList = nil
		}
	}

	//宝库幸运值日志
	luckyReason := commonlog.EquipBaoKuLogReasonLuckyPointsChange
	luckyReasonText := fmt.Sprintf(luckyReason.String(), typ.GetBaoKuName(), commontypes.ChangeTypeRefresh.String())
	luckyData := equipbaokueventtypes.CreatePlayerEquipBaoKuLuckyPointsLogEventData(luckyPoints, 0, logRewList, luckyReason, luckyReasonText, typ)
	gameevent.Emit(equipbaokueventtypes.EventTypeEquipBaoKuLuckyPointsLog, pl, luckyData)
	return
}

func init() {
	gameevent.AddEventListener(equipbaokueventtypes.EventTypeEquipBaoKuLuckyBox, event.EventListenerFunc(playerEquipBaoKuCrossDay))
}
