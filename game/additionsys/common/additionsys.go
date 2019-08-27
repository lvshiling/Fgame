package common

//附加系统类装备信息
type AdditionSysSlotInfo struct {
	SlotId     int32 `json:"slotId"`
	Level      int32 `json:"level"`
	ItemId     int32 `json:"itemId"`
	ShenZhuLev int32 `json:"shenZhuLev"`
	ShenZhuPro int32 `json:"shenZhuPro"`
}

//附加系统类通灵信息
type AdditionSysTongLingInfo struct {
	TongLingLev int32 `json:"tongLingLev"`
	TongLingPro int32 `json:"tongLingPro"`
}

//附加系统类信息
type AdditionSysInfo struct {
	SysType         int32                    `json:"sysType"`
	Level           int32                    `json:"level"`
	UpPro           int32                    `json:"upPro"`
	LingLevel       int32                    `json:"lingLevel"`
	LingPro         int32                    `json:"lingPro"`
	IsAwake         int32                    `json:"isAwake"`
	TongLingInfo    *AdditionSysTongLingInfo `json:"tongLingInfo"`
	SysTypeSlotList []*AdditionSysSlotInfo   `json:"sysTypeSlotList"`
}

type AllAdditionSysInfo struct {
	AdditionSysList []*AdditionSysInfo `json:"additionSysInfo"`
}
