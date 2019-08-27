package logic

import (
	"encoding/json"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emailentity "fgame/fgame/game/email/entity"
	"fgame/fgame/game/email/pbutil"
	playeremail "fgame/fgame/game/email/player"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
)

//领取附件逻辑
func HandleGetEmailAttachement(pl player.Player, emailId int64) (err error) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	//验证参数

	_, emailObj := emailManager.GetEmail(emailId)
	if emailObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
			}).Warn("email:领取附件请求参数错误，无效的emailId")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	//是否存在附件
	if emailManager.HasNotOrReceiveAttachment(emailId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"emailId":  emailId,
			}).Warn("email:领取附件请求，该邮件没有附件信息")
		playerlogic.SendSystemMessage(pl, lang.EmailNoAttachment)
		return
	}

	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(emailObj.GetAttachmentInfo()) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(emailObj.GetAttachmentInfo())
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//是否足够背包空间
	if len(newItemList) > 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"emailId":  emailId,
				}).Warn("email:领取附件请求,所需背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		//物品加入背包
		flag := inventoryManager.BatchAddOfItemLevel(newItemList, commonlog.InventoryLogReasonEmailAttachment, commonlog.InventoryLogReasonEmailAttachment.String())
		if !flag {
			panic("email: getAttachment add item should be ok")
		}
	}
	if len(resMap) > 0 {
		//增加资源
		goldReasonText := commonlog.GoldLogReasonEmailAttachment.String()
		silverReasonText := commonlog.SilverLogReasonEmailAttachment.String()
		levelReasonText := commonlog.LevelLogReasonEmailAttachment.String()
		err = droplogic.AddRes(pl, resMap, commonlog.GoldLogReasonEmailAttachment, goldReasonText, commonlog.SilverLogReasonEmailAttachment, silverReasonText, commonlog.LevelLogReasonEmailAttachment, levelReasonText)
		if err != nil {
			return
		}
	}

	//设置邮件附件已领取
	emailManager.ReceiveEmailAttachment(emailId)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scGetAttachment := pbutil.BuildSCGetAttachment(emailId, newItemList)
	pl.SendMsg(scGetAttachment)

	return
}

//处理一键领取附件逻辑
func HandleGetEmailAttachementBatch(pl player.Player) (err error) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	emailObjArr := emailManager.GetNotReceiveAttachmentEmails()
	var totalAttacheList []*droptemplate.DropItemData
	var emailIdArr []int64
	if len(emailObjArr) > 0 {
		//获取所有物品map、资源map
		for _, emailObj := range emailObjArr {
			itemList := emailObj.GetAttachmentInfo()

			totalAttacheList = append(totalAttacheList, itemList...)
			emailIdArr = append(emailIdArr, emailObj.GetEmailId())
		}

		var totalItemList []*droptemplate.DropItemData
		var totalResMap map[itemtypes.ItemAutoUseResSubType]int32
		if len(totalAttacheList) != 0 {
			totalItemList, totalResMap = droplogic.SeperateItemDatas(totalAttacheList)
		}
		//物品加入背包
		if len(totalItemList) > 0 {
			//是否足够背包空间
			if !inventoryManager.HasEnoughSlotsOfItemLevel(totalItemList) {
				log.WithFields(
					log.Fields{
						"playerId":   pl.GetId(),
						"emailIdArr": emailIdArr,
					}).Warn("email:一键领取附件请求,所需背包空间不足")
				playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
				return
			}

			flag := inventoryManager.BatchAddOfItemLevel(totalItemList, commonlog.InventoryLogReasonEmailAttachment, commonlog.InventoryLogReasonEmailAttachment.String())
			if !flag {
				panic("email: getAttachmentBatch add item should be ok")
			}

		}

		//增加资源
		if len(totalResMap) > 0 {
			goldReasonText := commonlog.GoldLogReasonEmailAttachment.String()
			silverReasonText := commonlog.SilverLogReasonEmailAttachment.String()
			levelReasonText := commonlog.LevelLogReasonEmailAttachment.String()
			err = droplogic.AddRes(pl, totalResMap, commonlog.GoldLogReasonEmailAttachment, goldReasonText, commonlog.SilverLogReasonEmailAttachment, silverReasonText, commonlog.LevelLogReasonEmailAttachment, levelReasonText)
			if err != nil {
				return
			}
		}

		//设置邮件附件已领取
		for _, emailObj := range emailObjArr {
			emailManager.ReceiveEmailAttachment(emailObj.GetEmailId())
		}
	}

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	scGetAttachmentBatch := pbutil.BuildSCGetAttachmentBatch(emailIdArr, totalAttacheList)
	pl.SendMsg(scGetAttachmentBatch)

	return
}

//向玩家推送邮件
func AddEmail(pl player.Player, title string, content string, attachmentInfo map[int32]int32) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)

	//推送新的邮件
	emailObj := emailManager.AddNewEmail(title, content, convertToDropItem(attachmentInfo))
	scAddEmail := pbutil.BuildSCAddEmail(emailObj)
	pl.SendMsg(scAddEmail)
}

//向玩家推送邮件:定义创建时间
func AddEmailDefinTime(pl player.Player, title string, content string, createTime int64, attachmentInfo map[int32]int32) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)

	//推送新的邮件
	emailObj := emailManager.AddEmail(title, content, createTime, convertToDropItem(attachmentInfo))
	scAddEmail := pbutil.BuildSCAddEmail(emailObj)
	pl.SendMsg(scAddEmail)
}

//向玩家推送离线邮件
func AddOfflineEmail(playerId int64, title string, content string, attachmentInfo map[int32]int32) (err error) {
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	emailsInfoBytes, err := json.Marshal(convertToDropItem(attachmentInfo))
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mailEntity := &emailentity.PlayerEmailEntity{
		Id:              id,
		PlayerId:        playerId,
		IsRead:          0,
		IsGetAttachment: 0,
		Title:           title,
		Content:         content,
		AttachementInfo: string(emailsInfoBytes),
		UpdateTime:      now,
		CreateTime:      now,
		DeleteTime:      0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(mailEntity)
	return err
}

//向玩家推送离线邮件:定义创建时间
func AddOfflineEmailDefinTime(playerId int64, title string, content string, createTime int64, attachmentInfo map[int32]int32) (err error) {
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	emailsInfoBytes, err := json.Marshal(convertToDropItem(attachmentInfo))
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mailEntity := &emailentity.PlayerEmailEntity{
		Id:              id,
		PlayerId:        playerId,
		IsRead:          0,
		IsGetAttachment: 0,
		Title:           title,
		Content:         content,
		AttachementInfo: string(emailsInfoBytes),
		UpdateTime:      now,
		CreateTime:      createTime,
		DeleteTime:      0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(mailEntity)
	return err
}

//向玩家推送邮件-等级物品:定义创建时间
func AddEmailItemLevel(pl player.Player, title string, content string, createTime int64, itemList []*droptemplate.DropItemData) *playeremail.PlayerEmailObject {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)

	//推送新的邮件
	emailObj := emailManager.AddEmail(title, content, createTime, itemList)
	scAddEmail := pbutil.BuildSCAddEmail(emailObj)
	pl.SendMsg(scAddEmail)

	return emailObj
}

//向玩家推送离线邮件-等级物品:定义创建时间
func AddOfflineEmailItemLevel(playerId int64, title string, content string, createTime int64, itemList []*droptemplate.DropItemData) (err error) {
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	emailsInfoBytes, err := json.Marshal(itemList)
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mailEntity := &emailentity.PlayerEmailEntity{
		Id:              id,
		PlayerId:        playerId,
		IsRead:          0,
		IsGetAttachment: 0,
		Title:           title,
		Content:         content,
		AttachementInfo: string(emailsInfoBytes),
		UpdateTime:      now,
		CreateTime:      createTime,
		DeleteTime:      0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(mailEntity)
	return err
}

func convertToDropItem(itemMap map[int32]int32) (itemList []*droptemplate.DropItemData) {
	for itemId, num := range itemMap {
		level := int32(0)
		bind := itemtypes.ItemBindTypeUnBind

		newData := droptemplate.CreateItemData(itemId, num, level, bind)
		itemList = append(itemList, newData)
	}

	return
}
