package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	"fgame/fgame/game/global"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_STORAGE_TYPE), dispatch.HandlerFunc(handleEmperorStorageGet))
}

//处理领取帝王金库信息
func handleEmperorStorageGet(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理领取帝王金库信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = emperorStorageGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理领取帝王金库信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理领取帝王金库信息完成")
	return nil
}

//处理领取帝王金库界面信息逻辑
func emperorStorageGet(pl player.Player) (err error) {
	playerId := pl.GetId()
	emperorId, robTime := emperor.GetEmperorService().GetEmperorIdAndRobTime()
	now := global.GetGame().GetTimeService().Now()
	pastTime := now - robTime
	if playerId != emperorId || pastTime < int64(common.DAY) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:条件不满足,无法领取")
		playerlogic.SendSystemMessage(pl, lang.EmperorStorageGetNotReach)
		return
	}

	//领取帝王金库
	storage, success := emperor.GetEmperorService().EmperorStorageGet(playerId, now)
	if success {
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		reasonSilver := commonlog.SilverLogReasonEmperorStorageGet
		silverReasonText := reasonSilver.String()
		propertyManager.AddSilver(storage, reasonSilver, silverReasonText)
		//同步元宝
		propertylogic.SnapChangedProperty(pl)

		//公告
		playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		storageStr := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("%d", storage))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmperorGetStorageNotice), playerName, storageStr)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	scEmperorStorageGet := pbuitl.BuildSCEmperorStorageGet(emperorObj, success)
	pl.SendMsg(scEmperorStorageGet)
	return
}
