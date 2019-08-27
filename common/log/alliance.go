package log

type AllianceLogReason int32

const (
	AllianceLogReasonGM AllianceLogReason = iota + 1
	AllianceLogReasonDepotAutoRemoveSetting
	AllianceLogReasonDepotItemRemove
	AllianceLogReasonDepotTakeOutItem
	AllianceLogReasonDepotSaveItem
	AllianceLogReasonDepotMergeItem
	AllianceLogReasonPlayerDepotPointChanged
)

func (zslr AllianceLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	allianceLogReasonMap = map[AllianceLogReason]string{
		AllianceLogReasonGM:                      "gm修改",
		AllianceLogReasonDepotAutoRemoveSetting:  "仙盟仓库自动销毁设置,状态：%d, 转生：%d，品质：%d",
		AllianceLogReasonDepotItemRemove:         "仙盟仓库自动销毁设置生效,物品移除",
		AllianceLogReasonDepotTakeOutItem:        "仙盟仓库取出物品,玩家id：%d",
		AllianceLogReasonDepotSaveItem:           "仙盟仓库放入物品,玩家id：%d",
		AllianceLogReasonDepotMergeItem:          "盟主整理仙盟仓库，盟主id：%d",
		AllianceLogReasonPlayerDepotPointChanged: "玩家仙盟仓库积分变化,物品id：%d",
	}
)

func (ar AllianceLogReason) String() string {
	return allianceLogReasonMap[ar]
}
