package handler

// func init() {
// 	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_FORBID_PLAYER_TYPE), cmd.CmdHandlerFunc(handleForbidPlayer))
// }

// func handleForbidPlayer(msg proto.Message) (err error) {
// 	log.Debug("cmd:请求封号")
// 	cmdForbidPlayer := msg.(*cmdpb.CmdForbidPlayer)
// 	forbidPlayerId := cmdForbidPlayer.GetPlayerId()
// 	forbidReason := cmdForbidPlayer.GetForbidReason()
// 	forbidName := cmdForbidPlayer.GetForbidName()
// 	forbidTime := cmdForbidPlayer.GetForbidTime()
// 	err = forbidPlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"err": err,
// 			}).Error("cmd:请求替人，错误")
// 		return
// 	}
// 	log.Debug("cmd:请求踢人，成功")
// 	return
// }

// func forbidPlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
// 	p := player.GetOnlinePlayerManager().GetPlayerById(forbidPlayerId)
// 	if p == nil {
// 		//加载离线玩家数据
// 		err = forbidOfflinePlayer(forbidPlayerId, forbidReason, forbidName, forbidTime)
// 		return
// 	}
// 	ctx := scene.WithPlayer(context.Background(), p)
// 	result := &forbidPlayerData{
// 		forbidReason: forbidReason,
// 		forbidName:   forbidName,
// 		forbidTime:   forbidTime,
// 	}
// 	msg := message.NewScheduleMessage(onForbidPlayer, ctx, result, nil)
// 	p.Post(msg)
// 	return
// }

// func forbidOfflinePlayer(forbidPlayerId int64, forbidReason string, forbidName string, forbidTime int64) (err error) {
// 	offlinePlayer, err := pp.CreateOfflinePlayer(forbidPlayerId)
// 	if err != nil {
// 		return err
// 	}
// 	if offlinePlayer == nil {
// 		err = cmd.ErrorCodeCommonPlayerNoExist
// 		return err
// 	}
// 	offlinePlayer.Forbid(forbidReason, forbidName, forbidTime)
// 	return nil
// }

// type forbidPlayerData struct {
// 	forbidReason string
// 	forbidName   string
// 	forbidTime   int64
// }

// func onForbidPlayer(ctx context.Context, result interface{}, err error) error {
// 	sp := scene.PlayerInContext(ctx)
// 	p, ok := sp.(player.Player)
// 	if !ok {
// 		return nil
// 	}
// 	data := result.(*forbidPlayerData)
// 	p.Forbid(data.forbidReason, data.forbidName, data.forbidTime)
// 	//强制下线
// 	playerlogic.SendExceptionContentMessage(p, data.forbidReason)
// 	p.Close(nil)
// 	return nil
// }
