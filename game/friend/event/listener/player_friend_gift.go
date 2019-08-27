package listener

import (
	"context"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itmetypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyentity "fgame/fgame/game/property/entity"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//送花
func friendGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*friendeventtypes.FriendGiftEventData)
	if !ok {
		return
	}
	friendId := eventData.GetFriendId()
	addPoint := eventData.GetNum()
	itemId := eventData.GetItemId()
	itemCount := eventData.GetItemCount()
	//增加亲密度
	friend.GetFriendService().AddPoint(pl, friendId, addPoint)

	//增加自身魅力值
	if eventData.GetCharmNum() != 0 {
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		charmAddReason := commonlog.CharmLogReasonGiftReward
		propertyManager.AddCharm(eventData.GetCharmNum()*2, charmAddReason, charmAddReason.String())
		propertylogic.SnapChangedProperty(pl)
	}

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		//离线增加魅力值
		charm := eventData.GetCharmNum()
		if charm != 0 {
			offineAddCharmLog(charm, friendId, pl.GetId())
		}
		return
	}
	frCtx := scene.WithPlayer(context.Background(), fr)
	friendGiftMsg := message.NewScheduleMessage(onFriendGift, frCtx, eventData, nil)
	fr.Post(friendGiftMsg)

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(itemCount)))
	linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	itemNameLink := coreutils.FormatLink(itemName, linkArgs)
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(fr.GetName()))

	playerIdList := make([]int64, 0, 2)
	playerNameList := make([]string, 0, 2)
	playerIdList = append(playerIdList, pl.GetId(), fr.GetId())
	playerNameList = append(playerNameList, pl.GetName(), fr.GetName())
	scFriendGiveFlowerLight := pbutil.BuildSCFriendGiveFlowerLight(itemId, itemCount, playerIdList, playerNameList)

	itemSubType := itemTemplate.GetItemSubType()
	switch itemSubType {
	case itmetypes.ItemXueHuaSubTypeSeniorBouquet:
		{

			numStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", addPoint))
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FriendGiveFlowersNotice), playerName, peerName, itemNameLink, numStr)
			//跑马灯
			noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
			//系统公告
			chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			player.GetOnlinePlayerManager().BroadcastMsg(scFriendGiveFlowerLight)
			break
		}
	case itmetypes.ItemBiaoBaiSubTypeShip:
		{
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FriendGiveBiaoBaiGiftNotice), playerName, peerName, itemNameLink)
			//跑马灯
			noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
			//系统公告
			chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			player.GetOnlinePlayerManager().BroadcastMsg(scFriendGiveFlowerLight)
			break
		}
	case itmetypes.ItemXueHuaSubTypeBouquet,
		itmetypes.ItemXueHuaSubTypeMidBouquet,
		itmetypes.ItemBiaoBaiSubTypeCash,
		itmetypes.ItemBiaoBaiSubTypeCar:
		{
			pl.SendMsg(scFriendGiveFlowerLight)
			fr.SendMsg(scFriendGiveFlowerLight)
			break
		}
	}

	return
}

func onFriendGift(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	if pl == nil {
		return nil
	}
	friendData := result.(*friendeventtypes.FriendGiftEventData)
	num := friendData.GetCharmNum()

	//魅力值
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	charmAddReason := commonlog.CharmLogReasonGiftReward
	propertyManager.AddCharm(num, charmAddReason, charmAddReason.String())
	propertylogic.SnapChangedProperty(pl)

	return nil
}

func offineAddCharmLog(charm int32, friendId int64, sendId int64) {
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()

	charmLogEntity := &propertyentity.PlayerCharmAddLogEntity{
		Id:         id,
		PlayerId:   friendId,
		Charm:      charm,
		SendId:     sendId,
		UpdateTime: now,
		CreateTime: now,
		DeleteTime: 0,
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(charmLogEntity)
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendGift, event.EventListenerFunc(friendGift))
}
