package handler

// func init() {
// 	command.Register(gmcommandtypes.CommandTypeArenaFourGodCancel, command.CommandHandlerFunc(handleArenaFourGodCancel))
// }

// func handleArenaFourGodCancel(pl player.Player, c *command.Command) (err error) {
// 	log.Debug("gm:3v3竞技场四神取消")

// 	err = arenaSelectFourGodCancel(pl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"id":    pl.GetId(),
// 				"error": err,
// 			}).Warn("gm:3v3竞技场四神取消,错误")
// 		return
// 	}
// 	log.WithFields(
// 		log.Fields{
// 			"id": pl.GetId(),
// 		}).Debug("gm:3v3竞技场选择四神,完成")
// 	return
// }

// func arenaSelectFourGodCancel(pl player.Player) (err error) {
// 	err = arenalogic.ArenaFourGodCancel(pl)
// 	if err != nil {
// 		return
// 	}
// 	return

// }
