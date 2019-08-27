package types

type ShopConsumeType int32

const (
	//银两
	ShopConsumeTypeSliver ShopConsumeType = 1 + iota
	//绑元
	ShopConsumeTypeBindGold
	//元宝
	ShopConsumeTypeGold
	//功勋
	ShopConsumeTypeGongXun
	//物品
	ShopConsumeTypeItem
	//3v3积分
	ShopConsumeTypeArenaJiFen
	//pvp积分
	ShopConsumeTypeArenapvpJiFen
	//创世之战积分
	ShopConsumeTypeChuangShiJiFen
	//特戒寻宝积分
	ShopConsumeTypeRingXunBaoJiFen
)

func (sct ShopConsumeType) Valid() bool {
	switch sct {
	case ShopConsumeTypeSliver,
		ShopConsumeTypeBindGold,
		ShopConsumeTypeGold,
		ShopConsumeTypeGongXun,
		ShopConsumeTypeItem,
		ShopConsumeTypeArenaJiFen,
		ShopConsumeTypeArenapvpJiFen,
		ShopConsumeTypeChuangShiJiFen,
		ShopConsumeTypeRingXunBaoJiFen:
		return true
	}
	return false
}

func (sct ShopConsumeType) Priority(oldType ShopConsumeType) bool {
	if sct == ShopConsumeTypeSliver {
		return false
	}
	return sct < oldType
}
