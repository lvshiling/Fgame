package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_INVITE_TYPE), dispatch.HandlerFunc(handleArenaInvite))
}

//处理3v3吆喝
func handleArenaInvite(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理3v3吆喝")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaInvite(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3吆喝,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3吆喝,完成")
	return nil

}

//3v3匹配
func arenaInvite(pl player.Player) (err error) {
	t := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if t == nil {
		return
	}
	if t.GetCaptain().GetPlayerId() != pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:处理3v3吆喝,不是队长")
		return
	}
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	if !arenaManager.IfInviteCD() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:处理3v3吆喝,吆喝cd中")
		playerlogic.SendSystemMessage(pl, lang.TransportationDistressCD)
		return
	}
	arenaManager.Invite()
	args := []int64{int64(chattypes.ChatLinkTypeTeamApply), t.GetTeamId()}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToParticipant, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenaInvite), t.GetTeamName(), joinLink)
	//发送系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	//TODO: 发送系统公告 记录cd
	scArenaInvite := pbutil.BuildSCArenaInvite()
	pl.SendMsg(scArenaInvite)
	return
}
