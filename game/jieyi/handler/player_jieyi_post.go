package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	coredirty "fgame/fgame/core/dirty"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fmt"
	"strings"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_POST_TYPE), dispatch.HandlerFunc(handlePlayerJieYiPost))
}

func handlePlayerJieYiPost(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家发布结义请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	p := gcs.Player()
	pl := p.(player.Player)

	csMsg := msg.(*uipb.CSJieYiPost)
	leaveWord := csMsg.GetLeaveWord()

	err = playerJieYiPost(pl, leaveWord)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家发布结义请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi: 处理玩家发布结义请求消息,成功")

	return
}

const (
	minLeaveWordLen = 0
	maxLeaveWordLen = 30
)

func playerJieYiPost(pl player.Player, leaveWord string) (err error) {
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if jieYiManager.IsJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家已经结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiAlreadyJieYi)
		return
	}

	obj := jieYiManager.GetPlayerJieYiObj()
	constantTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate()

	now := global.GetGame().GetTimeService().Now()
	if obj.GetLastPostTime() != 0 {
		jianGeTime := now - obj.GetLastPostTime()
		if jianGeTime < constantTemp.FaBuCD {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("jieyi: 发布结义消息CD中")
			playerlogic.SendSystemMessage(pl, lang.JieYiPostMessageCD, fmt.Sprintf("%d", constantTemp.FaBuCD/int64(common.MINUTE)))
			return
		}
	}

	// 判断留言是否合法
	leaveWord = strings.TrimSpace(leaveWord)
	leaveWordLen := len([]rune(leaveWord))
	if !jieyi.GetJieYiService().IsAlreadyJieYi(pl.GetId()) {
		if leaveWordLen < minLeaveWordLen || leaveWordLen > maxLeaveWordLen {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"leaveWord": leaveWord,
				}).Warn("alliance:留言不合法")
			playerlogic.SendSystemMessage(pl, lang.JieYiLeaveWordIllegal)
			return
		}

		flag := coredirty.GetDirtyService().IsLegal(leaveWord)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"leaveWord": leaveWord,
				}).Warn("alliance:留言含有脏字")
			playerlogic.SendSystemMessage(pl, lang.JieYiLeaveWordDirty)
			return
		}
	}

	// 公告
	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeJieYi}
	link := coreutils.FormatLink(chattypes.ButtonTypeJieYi, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.JieYiLeaveWordGongGao), plName, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	flag := jieyi.GetJieYiService().AddJieYiLeaveWord(pl.GetId(), leaveWord)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"leaveWord": leaveWord,
			}).Warn("alliance:留言失败")
		playerlogic.SendSystemMessage(pl, lang.JieYiLiuYanFail)
		return
	}

	// 留言成功，刷新数据
	jieYiManager.LeaveWordSuccess()
	scMsg := pbutil.BuildSCJieYiPost(obj.GetLastPostTime(), leaveWord)
	pl.SendMsg(scMsg)
	return
}
