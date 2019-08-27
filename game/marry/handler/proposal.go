package handler

func init() {
	// processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_PROPOSAL_TYPE), dispatch.HandlerFunc(handleMarryProposal))
}

// //处理求婚信息，旧求婚，已停用
// func handleMarryProposal(s session.Session, msg interface{}) (err error) {
// 	log.Debug("marry:处理求婚消息")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)
// 	csMarryProposal := msg.(*uipb.CSMarryProposal)
// 	ring := csMarryProposal.GetRingType()
// 	proposalId := csMarryProposal.GetPlayerId()
// 	proposalName := csMarryProposal.GetPlayerName()

// 	err = marryProposal(tpl, marrytypes.MarryRingType(ring), proposalId, proposalName)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"ring":       ring,
// 				"proposalId": proposalId,
// 				"error":      err,
// 			}).Error("marry:处理求婚消息,错误")
// 		return
// 	}
// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("marry:处理求婚消息完成")
// 	return nil
// }

// //处理求婚信息逻辑
// func marryProposal(pl player.Player, ringType marrytypes.MarryRingType, proposalId int64, proposalName string) (err error) {
// 	//参数校验
// 	if !ringType.Valid() {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:参数错误")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
// 	flag := friendManager.IsFriend(proposalId)
// 	if !flag {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:未添加对方为好友,无法同意对方的求婚请求")
// 		playerlogic.SendSystemMessage(pl, lang.MarryProposalNoFriend, proposalName)
// 		return
// 	}

// 	spl := player.GetOnlinePlayerManager().GetPlayerById(proposalId)
// 	if spl == nil {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:玩家不在线,无法求婚")
// 		playerlogic.SendSystemMessage(pl, lang.MarrySpouseNoOnline)
// 		return
// 	}

// 	mySex := pl.GetSex()
// 	spouseSex := spl.GetSex()
// 	if mySex == spouseSex {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:只能向异性求婚")
// 		playerlogic.SendSystemMessage(pl, lang.MarrySpouseSameSex)
// 		return
// 	}

// 	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
// 	marryInfo := manager.GetMarryInfo()
// 	marryObj := marry.GetMarryService().GetMarry(proposalId)
// 	if marryInfo.SpouseId != 0 || marryObj != nil {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:您或对方已拥有伴侣,无法进行求婚")
// 		playerlogic.SendSystemMessage(pl, lang.MarryHasedSpouse)
// 		return
// 	}

// 	friendObj := friendManager.GetFriend(proposalId)
// 	intimacy := friendObj.Point
// 	needIntimacy := marrytemplate.GetMarryTemplateService().GetMarryConstIntimacy()
// 	if intimacy < needIntimacy {
// 		log.WithFields(log.Fields{
// 			"playerId":   pl.GetId(),
// 			"ringType":   ringType,
// 			"proposalId": proposalId,
// 		}).Warn("marry:双方亲密度过低")
// 		intimacyStr := fmt.Sprintf("%d", needIntimacy)
// 		playerlogic.SendSystemMessage(pl, lang.MarryIntimacyNoEnough, intimacyStr)
// 		return
// 	}

// 	//是否已向其它人求过婚
// 	marryRingObj := marry.GetMarryService().GetMarryProposalRing(pl.GetId())
// 	if marryRingObj != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"ringType":   ringType,
// 				"proposalId": proposalId,
// 			}).Warn("marry:同时只能向一个人求婚,您已经向其它求婚")
// 		playerlogic.SendSystemMessage(pl, lang.MarryProposalIsExist, marryRingObj.PeerName)
// 		return
// 	}

// 	//获取婚烟对戒
// 	itemTempalte := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())
// 	ringItem := int32(itemTempalte.TemplateId())
// 	//判断物品
// 	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	flag = inventoryManager.HasEnoughItem(ringItem, 1)
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"ringType":   ringType,
// 				"proposalId": proposalId,
// 			}).Warn("marry:婚戒不足")
// 		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 		return
// 	}

// 	//记录求婚婚戒
// 	marry.GetMarryService().MarryProposalRing(pl.GetId(), proposalId, spl.GetName(), ringType)
// 	//扣除婚戒
// 	inventoryReason := commonlog.InventoryLogReasonMarryProposal
// 	flag = inventoryManager.UseItem(ringItem, 1, inventoryReason, inventoryReason.String())
// 	if !flag {
// 		panic(fmt.Errorf("marry: marryProposal UseItem should be ok"))
// 	}
// 	//同步物品
// 	inventorylogic.SnapInventoryChanged(pl)

// 	//发送事件
// 	eventData := marryeventtypes.CreateMarryProposalEventData(pl.GetId(), proposalId, ringType)
// 	gameevent.Emit(marryeventtypes.EventTypeMarryProposal, nil, eventData)

// 	scMarryProposal := pbuitl.BuildSCMarryProposal(0)
// 	pl.SendMsg(scMarryProposal)
// 	return
// }
