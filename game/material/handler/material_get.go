package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/material/pbutil"
	playermaterial "fgame/fgame/game/material/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MATERIAL_INFO_GET_TYPE), dispatch.HandlerFunc(handlerMaterialGet))
}

//材料副本信息处理
func handlerMaterialGet(s session.Session, msg interface{}) (err error) {
	log.Debug("material:处理材料副本信息获取请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = materialGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("material:处理材料副本信息获取请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("material：处理材料副本信息获取请求完成")

	return
}

//仙府信息处理逻辑
func materialGet(pl player.Player) (err error) {
	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	//刷新数据
	materialManager.RefreshData()
	//获取信息List
	materialList := materialManager.GetPlayerMaterialInfoList()

	scMsg := pbutil.BuildSCMaterialInfoGet(materialList)
	pl.SendMsg(scMsg)
	return
}
