package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	playersecretcard "fgame/fgame/game/secretcard/player"
)

func BuildSCQuestSecretCardGet(secretCard *playersecretcard.PlayerSecretCardObject) *uipb.SCQuestSecretCardGet {
	secretCardGet := &uipb.SCQuestSecretCardGet{}
	num := secretCard.Num
	totalStar := secretCard.TotalStar

	secretCardGet.Num = &num
	secretCardGet.TotalStar = &totalStar
	for _, openBox := range secretCard.OpenBoxList {
		secretCardGet.OpenBoxList = append(secretCardGet.OpenBoxList, openBox)
	}

	if secretCard.CardId != 0 {
		cardId := secretCard.CardId
		star := secretCard.Star
		secretCardGet.CardList = append(secretCardGet.CardList, buildCard(cardId, star))
	} else {
		for cardId, star := range secretCard.CardMap {
			secretCardGet.CardList = append(secretCardGet.CardList, buildCard(cardId, star))
		}
	}

	return secretCardGet
}

func BuildSCQuestSecretPickUp(cardId int32) *uipb.SCQuestSecretPickUp {
	secretPickUp := &uipb.SCQuestSecretPickUp{}
	secretPickUp.CardId = &cardId
	return secretPickUp
}

func BuildSCQuestSecretStarRew(openBox int32) *uipb.SCQuestSecretStarRew {
	secretStarRew := &uipb.SCQuestSecretStarRew{}
	secretStarRew.OpenBox = &openBox
	return secretStarRew
}

func BuildSCQuestSecretFinish(dropMap []*droptemplate.DropItemData, result bool) *uipb.SCQuestSecretFinish {
	secretFinish := &uipb.SCQuestSecretFinish{}
	secretFinish.Result = &result

	for _, itemData := range dropMap {
		itemId := itemData.ItemId
		num := itemData.Num
		level := itemData.Level
		secretFinish.ItemList = append(secretFinish.ItemList, buildItem(itemId, num, level))
	}
	return secretFinish
}

func BuildSCQuestSecretSpy(num int32, cardMap map[int32]int32) *uipb.SCQuestSecretSpy {
	secretSpy := &uipb.SCQuestSecretSpy{}
	secretSpy.Num = &num

	for cardId, star := range cardMap {
		secretSpy.CardList = append(secretSpy.CardList, buildCard(cardId, star))
	}
	return secretSpy
}

func BuildSCQuestSecretImmediateFinish(itemDataList []*droptemplate.DropItemData) *uipb.SCQuestSecretImmediate {
	secretImmediate := &uipb.SCQuestSecretImmediate{}
	for _, itemData := range itemDataList {
		if itemData == nil {
			continue
		}
		secretImmediate.ItemInfo = append(secretImmediate.ItemInfo, buildItem(itemData.ItemId, itemData.Num, itemData.Level))
	}

	return secretImmediate
}

func buildItem(itemId int32, num int32, level int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	itemInfo.Level = &level
	return itemInfo
}

func buildCard(cardId int32, star int32) *uipb.SecretCard {
	cardInfo := &uipb.SecretCard{}
	cardInfo.CardId = &cardId
	cardInfo.Star = &star
	return cardInfo
}
