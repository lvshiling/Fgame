package types

type DropItemEventType string

const (
	EventTypeDropItemAuto         DropItemEventType = "DropItemAuto"
	EventTypeDropItemRemove                         = "DropItemRemove"
	EventTypeDropItemOwnerChanged                   = "DropItemOwnerChanged"
	//掉落被捡起
	EventTypeDropItemGet = "DropItemGet"
	//副本物品全部捡起
	EventTypeFubenDropItemsGet = "FubenDropItemsGet"
)
