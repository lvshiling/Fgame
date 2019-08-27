package logic

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/marry/marry"
	pbutil "fgame/fgame/game/marry/pbutil"
	marryplayer "fgame/fgame/game/marry/player"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
)

//定情信物改变类型
type dingQingChange struct {
	PlayerId int64
	SuitMap  map[int32]map[int32]int32
	SuitId   int32
	PosId    int32
}

func newDingQingChange(playerId int64, suitMap map[int32]map[int32]int32, suitId int32, posId int32) *dingQingChange {
	rst := &dingQingChange{
		PlayerId: playerId,
		SuitId:   suitId,
		PosId:    posId,
	}
	rst.SuitMap = make(map[int32]map[int32]int32)
	for suitId, posMap := range suitMap {
		_, exists := rst.SuitMap[suitId]
		if !exists {
			rst.SuitMap[suitId] = make(map[int32]int32)
		}
		for posId, _ := range posMap {
			rst.SuitMap[suitId][posId] = 0
		}
	}
	return rst
}

//添加玩家定情
func AddPlayerDingQing(pl player.Player, suitId int32, posId int32) {
	playerId := pl.GetId()
	wedManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)
	wedManager.AddDingQing(suitId, posId)
	playerSuitMap := wedManager.GetAllDingQingMap()
	marry.GetMarryService().SyncMarryDingQing(pl.GetId(), playerSuitMap)
	//发送定情信物变更给自己
	changeMsg := pbutil.BuildSCMarryDingQingChange(playerId, playerSuitMap)
	pl.SendMsg(changeMsg)

	//通知情侣改变内容
	spouseId := marry.GetMarryService().GetSpouseId(playerId)
	if spouseId == 0 {
		return
	}
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	MarryPlayerDingQingPropertyChanged(pl) //战斗力改变
	if spl == nil {
		return
	}

	//TODO:cjy playerSuitMap多协程使用会奔溃,建议复制一份
	// dingQingMap := marry.GetMarryService().GetMarryDingQing(playerId) //获取玩家的定情信物
	postData := newDingQingChange(playerId, playerSuitMap, suitId, posId)

	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(addPlayerDingQingPost, splCtx, postData, nil)
	spl.Post(msg)
}

func addPlayerDingQingPost(ctx context.Context, result interface{}, err error) error {
	suit := result.(*dingQingChange)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	wedManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)
	wedManager.AddSpouseSuit(suit.SuitId, suit.PosId) //改变伴侣身上的属性
	changeMsg := pbutil.BuildSCMarryDingQingChange(suit.PlayerId, suit.SuitMap)
	pl.SendMsg(changeMsg)
	MarryPlayerDingQingPropertyChanged(pl) //战斗力改变
	return nil
}

//定情战斗力改变
func MarryPlayerDingQingPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMarry.Mask())
	return
}

//定情信物战斗力改变双方通知
func MarryPlayerDingQingPropertyChangedBoth(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMarry.Mask())
	spouseId := marry.GetMarryService().GetSpouseId(pl.GetId())
	if spouseId == 0 {
		return
	}
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(onMarryPlayerDingQingPropertyChangedBoth, splCtx, nil, nil)
	spl.Post(msg)
	return
}

func onMarryPlayerDingQingPropertyChangedBoth(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	return MarryPlayerDingQingPropertyChanged(pl)
}

//更新定情信物值service,同时通知伴侣,并使伴侣更新自己的数据,结婚的时候使用
func SyncMarryPlayerDingQing(pl player.Player) {
	splCtx := scene.WithPlayer(context.Background(), pl)
	msg := message.NewScheduleMessage(onSyncMarryPlayerDingQing, splCtx, nil, nil)
	pl.Post(msg)
	return
}

func onSyncMarryPlayerDingQing(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	dingQingSuitMap := manager.GetAllDingQingMap()
	marry.GetMarryService().SyncMarryDingQing(pl.GetId(), dingQingSuitMap)
	MarryPlayerDingQingPropertyChangedBoth(pl)

	spouseId := marry.GetMarryService().GetSpouseId(pl.GetId())
	if spouseId <= 0 {
		return nil
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		return nil
	}
	splCtx := scene.WithPlayer(context.Background(), spl)
	postData := newDingQingChange(pl.GetId(), dingQingSuitMap, int32(0), int32(0))

	msg := message.NewScheduleMessage(onSyncMarryPlayerDingQingNoticSpouse, splCtx, postData, nil)
	spl.Post(msg)

	return nil
}

func onSyncMarryPlayerDingQingNoticSpouse(ctx context.Context, result interface{}, err error) error {
	suit := result.(*dingQingChange)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.UpdateSpouseSuit(suit.SuitMap)
	changeMsg := pbutil.BuildSCMarryDingQingChange(suit.PlayerId, suit.SuitMap)
	pl.SendMsg(changeMsg)
	return nil
}

//更新定情信物值service,同时通知伴侣,结束
