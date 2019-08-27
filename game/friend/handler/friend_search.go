package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_SEARCH_TYPE), dispatch.HandlerFunc(handleFriendSearch))
}

//处理好友添加
func handleFriendSearch(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友查找")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendSearch := msg.(*uipb.CSFriendSearch)
	search := csFriendSearch.GetSearch()
	err := friendSearch(tpl, search)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"search":   search,
				"error":    err,
			}).Error("friend:处理好友查找,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友查找,完成")
	return nil

}

//查找好友id
func friendSearchId(pl player.Player, searchId int64) (err error) {
	//TODO 异步获取
	info, err := player.GetPlayerService().GetPlayerInfo(searchId)
	if err != nil {
		return
	}

	scFriendSearch := pbutil.BuildSCFriendSearch(info)
	pl.SendMsg(scFriendSearch)
	return
}

//查找好友查询
func friendSearch(pl player.Player, search string) (err error) {
	serverId := global.GetGame().GetServerIndex()
	info, err := player.GetPlayerService().GetPlayerInfoByName(search, serverId)
	if err != nil {
		return
	}

	scFriendSearch := pbutil.BuildSCFriendSearch(info)
	pl.SendMsg(scFriendSearch)
	return
}
