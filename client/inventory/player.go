package inventory

import (
	"fgame/fgame/client/player"
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ItemObject struct {
	ItemId int32
	Index  int32
	Num    int32
}

func (io *ItemObject) String() string {
	return fmt.Sprintf("item:id:%d,index:%d,num:%d", io.ItemId, io.Index, io.Num)
}

type PlayerInventoryDataManager struct {
	p           *player.Player
	itemList    []*ItemObject
	currentPage int32
	wg          *sync.WaitGroup
}

func (psdm *PlayerInventoryDataManager) GetPlayer() *player.Player {
	return psdm.p
}

func (psdm *PlayerInventoryDataManager) Load(wg *sync.WaitGroup) {
	psdm.wg = wg
	psdm.loadItems()
}

//获取物品数据
func (psdm *PlayerInventoryDataManager) loadItems() {
	InventoryGet(psdm.p, psdm.currentPage+1)
}

func (psdm *PlayerInventoryDataManager) NumOfItems(itemId int32) int32 {
	num := int32(0)
	for _, item := range psdm.itemList {
		if item.ItemId == itemId {
			num += item.Num
		}
	}
	return num
}

func (psdm *PlayerInventoryDataManager) FirstIndex(itemId int32) int32 {
	index := int32(-1)
	for _, item := range psdm.itemList {
		if item.ItemId == itemId && item.Num > 0 {
			index = item.Index
			break
		}
	}
	return index
}

func (psdm *PlayerInventoryDataManager) findItemByIndex(index int32) *ItemObject {
	for _, item := range psdm.itemList {
		if item.Index == index {
			return item
		}
	}
	return nil
}

func (psdm *PlayerInventoryDataManager) OnLoadItems(page int32, items []*ItemObject) {
	if psdm.currentPage+1 != page {
		log.WithFields(
			log.Fields{
				"playerId":    psdm.p.Id(),
				"currentPage": psdm.currentPage,
				"page":        page,
			}).Warn("inventory:获取错的页数")
		return
	}
	psdm.currentPage = page
	for _, item := range items {
		tItem := psdm.findItemByIndex(item.Index)
		if tItem != nil {
			log.WithFields(
				log.Fields{
					"playerId": psdm.p.Id(),
					"index":    item.Index,
					"page":     page,
				}).Warn("inventory:获取重复的索引")
			continue
		}
		psdm.itemList = append(psdm.itemList, item)
	}

	if len(items) != 0 {
		psdm.loadItems()
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": psdm.p.Id(),
			"items":    psdm.itemList,
		}).Debug("inventory:加载背包完成")
	psdm.wg.Done()
}

func (psdm *PlayerInventoryDataManager) OnItemsChange(items []*ItemObject) {
	for _, item := range items {
		tItem := psdm.findItemByIndex(item.Index)
		if tItem == nil {
			psdm.itemList = append(psdm.itemList, item)
			continue
		}
		tItem.ItemId = item.ItemId
		tItem.Index = item.Index
		tItem.Num = item.Num
	}
}

func CreatePlayerInventoryDataManager(pl *player.Player) player.PlayerDataManager {
	psdm := &PlayerInventoryDataManager{
		p:           pl,
		itemList:    make([]*ItemObject, 0, 16),
		currentPage: 0,
	}

	return psdm
}

func init() {
	player.RegisterPlayerDataManager(player.PlayerDataKeyInventory, player.PlayerDataManagerFactoryFunc(CreatePlayerInventoryDataManager))
}
