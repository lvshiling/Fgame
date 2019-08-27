package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/types"
	itemtypes "fgame/fgame/game/item/types"
	playertypes "fgame/fgame/game/player/types"
	treasureboxtypes "fgame/fgame/game/treasurebox/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

const (
	maxYuanBaoKa = 10000
)

type ItemTemplate struct {
	*ItemTemplateVO
	//背包类型
	bagType inventorytypes.BagType
	//类型
	itemType types.ItemType
	//子类型
	itemSubType types.ItemSubType
	//cd组
	cdGroup *CdGroupTemplate
	//过期类型
	limitTimeType types.ItemLimitTimeType
	//使用标识
	itemUseFlag types.ItemUseFlag
	//过期时间
	expiredTime int64
	//角色
	role playertypes.RoleType
	//性别
	sex playertypes.SexType
	//标签
	tag inventorytypes.InventoryTag
	//绑定属性
	bindType types.ItemBindType
	//品质
	quality types.ItemQualityType
	//丹药对应的属性
	danAttrTemplate *AttrTemplate
	//装备
	equipmentTemplate *EquipTemplate
	//技能模板
	skillTemplate *SkillTemplate
	//宝石属性
	gemAttrTemplate *AttrTemplate
	//神龙道具属性
	dragonAttrTemplate *AttrTemplate
	//宝箱模板
	boxTemplate *BoxTemplate
	//宝箱模板2
	boxTemplateSecond *BoxTemplate
	//元神金装
	goldEquipTemplate *GoldEquipTemplate
	//附加系统装备
	systemEquipTemplate AdditionSystemEquip
	//神器器灵
	shenQiQiLingTemplate *ShenQiQiLingTemplate
	//无双神器基础模板
	wushuangWeaponBaseTemplate *WushuangWeaponBaseTemplate
	//关联运营活动groupId
	groupList []int32
	//生效时间
	effectiveTime int64
	//屠龙装备
	tuLongEquipTemplate *TuLongEquipTemplate
	//宝宝玩具
	babyToyTemplate *BabyToyTemplate
	//拆解获得物品
	chaiJieItemMap map[int32]int32
}

func (it *ItemTemplate) GetBagType() inventorytypes.BagType {
	return it.bagType
}

func (it *ItemTemplate) IsLingTongCanEquip(curLevel int32) bool {
	if curLevel >= it.TypeFlag2 {
		return true
	}
	return false
}

func (it *ItemTemplate) GetItemType() types.ItemType {
	return it.itemType
}

func (it *ItemTemplate) IsYunYingItem() bool {
	if len(it.groupList) > 0 {
		return true
	}
	return false
}

func (it *ItemTemplate) IsRelationToGroup(groupId int32) bool {
	return coreutils.ContainInt32(it.groupList, groupId)
}

func (it *ItemTemplate) GetItemSubType() types.ItemSubType {
	return it.itemSubType
}

func (it *ItemTemplate) GetItemUseFlag() types.ItemUseFlag {
	return it.itemUseFlag
}

func (it *ItemTemplate) GetEquipmentTemplate() *EquipTemplate {
	return it.equipmentTemplate
}

func (it *ItemTemplate) GetSkillTemplate() *SkillTemplate {
	return it.skillTemplate
}

func (it *ItemTemplate) GetBoxTemplate() *BoxTemplate {
	return it.boxTemplate
}

func (it *ItemTemplate) GetChaiJieItemMap() map[int32]int32 {
	return it.chaiJieItemMap
}

func (it *ItemTemplate) GetBoxTemplateByCostType(costType treasureboxtypes.BoxCostType) *BoxTemplate {
	switch costType {
	case treasureboxtypes.BoxCostTypeFree:
		return it.boxTemplateSecond
	case treasureboxtypes.BoxCostTypeGold:
		return it.boxTemplate
	}

	return nil
}

func (it *ItemTemplate) GetGoldEquipTemplate() *GoldEquipTemplate {
	return it.goldEquipTemplate
}

func (it *ItemTemplate) GetWushuangBaseTemplate() *WushuangWeaponBaseTemplate {
	return it.wushuangWeaponBaseTemplate
}

func (it *ItemTemplate) GetTuLongEquipTemplate() *TuLongEquipTemplate {
	return it.tuLongEquipTemplate
}

func (it *ItemTemplate) GetBabyToyTemplate() *BabyToyTemplate {
	return it.babyToyTemplate
}

func (it *ItemTemplate) GetShenQiQiLingTemplate() *ShenQiQiLingTemplate {
	return it.shenQiQiLingTemplate
}

func (it *ItemTemplate) GetSystemEquipTemplate() AdditionSystemEquip {
	return it.systemEquipTemplate
}

func (it *ItemTemplate) IsEquipment() bool {
	return it.itemType == types.ItemTypeEquipment
}

func (it *ItemTemplate) IsCanSaveInAllianceDepot() bool {
	if it.bindType == itemtypes.ItemBindTypeBind {
		return false
	}

	if it.UnionGet == 0 {
		return false
	}

	return true
}

func (it *ItemTemplate) IsDelAfterUse() bool {
	return it.itemUseFlag&types.ItemUseFlagDelete == types.ItemUseFlagDelete
}

func (it *ItemTemplate) IsGoldEquip() bool {
	return it.itemType == types.ItemTypeGoldEquip
}

func (it *ItemTemplate) IsTeRing() bool {
	return it.itemSubType == types.ItemTeRingSubTypeRing
}

func (it *ItemTemplate) IsCanChaiJie() bool {
	return len(it.chaiJieItemMap) != 0
}

//元神金装/系统装备分解
func (it *ItemTemplate) IfCanFenJie() bool {
	return it.IsGoldEquip() || it.systemEquipTemplate != nil
}

func (it *ItemTemplate) IsBaoBaoCard() bool {
	return it.itemType == types.ItemTypeBaoBaoCard
}

func (it *ItemTemplate) IsTianGongChui() bool {
	return it.itemType == itemtypes.ItemTypeGoldEquipStrengthen && it.itemSubType == itemtypes.ItemGoldEquipStrengthenSubTypeChuiZi
}

func (it *ItemTemplate) IsKaiGuangZuan() bool {
	return it.itemType == itemtypes.ItemTypeGoldEquipStrengthen && it.itemSubType == itemtypes.ItemGoldEquipStrengthenSubTypeKaiGuangZuan
}

func (it *ItemTemplate) IsQiangHuaShengShi() bool {
	return it.itemType == itemtypes.ItemTypeGoldEquipStrengthen && it.itemSubType == itemtypes.ItemGoldEquipStrengthenSubTypeQiangHuaShengShi
}

func (it *ItemTemplate) IsTuLongEquip() bool {
	return it.itemType == types.ItemTypeTuLongEquip
}

func (it *ItemTemplate) IsGem() bool {
	return it.itemType == types.ItemTypeGem
}

func (it *ItemTemplate) IsNotice() bool {
	if it.GetItemType() == itemtypes.ItemTypeAutoUseRes {
		return false
	}

	if it.Quality < int32(itemtypes.ItemQualityTypeOrange) {
		return false
	}

	return true
}

func (it *ItemTemplate) GetRole() playertypes.RoleType {
	return it.role
}

func (it *ItemTemplate) GetSex() playertypes.SexType {
	return it.sex
}

func (it *ItemTemplate) IsLimitDayUseTimes() bool {
	return it.LimitTimeDay != 0
}

func (it *ItemTemplate) GetEffectiveTime() int64 {
	return it.effectiveTime
}

func (it *ItemTemplate) IsLimitTotalUseTimes() bool {
	return it.LimitTimeAll != 0
}

func (it *ItemTemplate) IsCDItem() bool {
	return it.CdTime != 0
}

func (it *ItemTemplate) GetLimitTimeType() inventorytypes.NewItemLimitTimeType {
	return it.limitTimeType.ConvertToNewItemLimitTimeType()
}

func (it *ItemTemplate) GetTrueExpireTime(itemGetTime int64) int64 {
	if it.limitTimeType == itemtypes.ItemLimitTimeTypeExpired {
		expireTime := it.expiredTime + itemGetTime
		return expireTime
	} else if it.limitTimeType == itemtypes.ItemLimitTimeTypeTodayExpired {
		expireTime, _ := timeutils.BeginOfNow(itemGetTime)
		expireTime += it.expiredTime
		return expireTime
	}
	return 0
}

func (it *ItemTemplate) IsExpire(itemGetTime, now int64) bool {
	if it.limitTimeType == itemtypes.ItemLimitTimeTypeExpired {
		expireTime := it.expiredTime + itemGetTime
		if now > expireTime {
			return true
		}
	} else if it.limitTimeType == itemtypes.ItemLimitTimeTypeTodayExpired {
		expireTime, err := timeutils.BeginOfNow(itemGetTime)
		if err != nil {
			return true
		}
		expireTime += it.expiredTime
		if now > expireTime {
			return true
		}
	}
	return false
}

func (it *ItemTemplate) GetExpireTime() int64 {
	return it.expiredTime
}

func (it *ItemTemplate) FormateItemNameOfNum(num int32) string {
	itemName := it.Name
	if it.GetTag() != inventorytypes.InventoryTagEquipment {
		if num > 0 {
			numString := fmt.Sprintf("x%d", num)
			itemName = it.Name + numString
		}
	}

	return itemName
}

//是否可以叠加
func (it *ItemTemplate) CanOverlap() bool {
	//过期时间可变的
	if it.limitTimeType == types.ItemLimitTimeTypeExpired {
		return false
	}

	//最多叠加1
	if it.MaxOverlap == 1 {
		return false
	}
	return true
}

func (it *ItemTemplate) GetDanAttrTemplate() *AttrTemplate {
	return it.danAttrTemplate
}

func (it *ItemTemplate) GetGemAttrTemplate() *AttrTemplate {
	return it.gemAttrTemplate
}

func (it *ItemTemplate) GetDragonAttrTemplate() *AttrTemplate {
	return it.dragonAttrTemplate
}

func (it *ItemTemplate) GetTag() inventorytypes.InventoryTag {
	return it.tag
}

func (it *ItemTemplate) GetBindType() types.ItemBindType {
	return it.bindType
}

func (it *ItemTemplate) GetQualityType() types.ItemQualityType {
	return it.quality
}

func (it *ItemTemplate) CanSell() bool {
	return it.SaleRate != 0
}
func (it *ItemTemplate) CanTrade() bool {
	return it.MarketId != 0
}

func (it *ItemTemplate) TemplateId() int {
	return it.Id
}

func (it *ItemTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(it.FileName(), it.TemplateId(), err)
			return
		}
	}()
	//检查背包类型
	it.bagType = inventorytypes.BagType(it.BagType)
	if !it.bagType.Valid() {
		err = fmt.Errorf("[%d] invalid", it.bagType)
		return template.NewTemplateFieldError("bagType", err)
	}

	//检查类型
	it.itemType = types.ItemType(it.Type)
	if !it.itemType.Valid() {
		err = fmt.Errorf("[%d] invalid", it.Type)
		return template.NewTemplateFieldError("type", err)
	}
	//检查子类型
	it.itemSubType = types.CreateItemSubType(it.itemType, it.SubType)
	if it.itemSubType == nil || !it.itemSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", it.SubType)
		return template.NewTemplateFieldError("subType", err)
	}
	//检查使用标识
	it.itemUseFlag = types.ItemUseFlag(it.UseFlag)
	//验证cd组
	if it.CdGroup != 0 {
		to := template.GetTemplateService().Get(int(it.CdGroup), (*CdGroupTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", it.CdGroup)
			return template.NewTemplateFieldError("cdGroup", err)
		}
	}

	//过期时间类型
	it.limitTimeType = types.ItemLimitTimeType(it.LimitTimeType)
	if !it.limitTimeType.Valid() {
		err = fmt.Errorf("[%d] invalid", it.LimitTimeType)
		return template.NewTemplateFieldError("limitTimeType", err)
	}
	//获取过期时间
	switch it.limitTimeType {
	case types.ItemLimitTimeTypeNone:
		break
	case types.ItemLimitTimeTypeExpired:
		{
			expireInt, err := strconv.ParseInt(it.LimitTime, 10, 64)
			if err != nil {
				return err
			}
			it.expiredTime = expireInt
			break
		}
	case types.ItemLimitTimeTypeTodayExpired:
		{
			if it.LimitTime == "240000" {
				it.expiredTime = int64(common.DAY)
			} else {
				expireTime, err := timeutils.ParseDayOfHHMMSS(it.LimitTime)
				if err != nil {
					return err
				}
				it.expiredTime = expireTime
			}
			break
		}
	}

	//验证角色
	if it.NeedProfession != 0 {
		it.role = playertypes.RoleType(it.NeedProfession)
		if !it.role.Valid() {
			err = fmt.Errorf("[%d] invalid", it.NeedProfession)
			return template.NewTemplateFieldError("needProfession", err)
		}
	}

	//性别
	if it.NeedGender != 0 {
		it.sex = playertypes.SexType(it.NeedGender)
		if !it.sex.Valid() {
			err = fmt.Errorf("[%d] invalid", it.NeedGender)
			return template.NewTemplateFieldError("NeedGender", err)
		}
	}

	//标签
	it.tag = inventorytypes.InventoryTag(it.Tag)
	if !it.tag.Valid() {
		err = fmt.Errorf("[%d] invalid", it.Tag)
		return template.NewTemplateFieldError("tag", err)
	}

	// 绑定属性
	it.bindType = types.ItemBindType(it.BindType)
	if !it.bindType.Valid() {
		err = fmt.Errorf("[%d] invalid", it.BindType)
		return template.NewTemplateFieldError("BindType", err)
	}

	// 品质属性
	it.quality = types.ItemQualityType(it.Quality)
	if !it.quality.Valid() {
		err = fmt.Errorf("[%d] invalid", it.Quality)
		return template.NewTemplateFieldError("Quality", err)
	}

	//TODO: 封装成类型验证方法
	switch it.itemType {
	case itemtypes.ItemTypeDan:
		{
			if it.itemSubType == itemtypes.ItemDanSubTypeEat {
				tempAttrTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*AttrTemplate)(nil))
				if tempAttrTemplate == nil {
					err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				attrTemplate, ok := tempAttrTemplate.(*AttrTemplate)
				if !ok {
					err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				it.danAttrTemplate = attrTemplate
			}
			break
		}
	case itemtypes.ItemTypeEquipment:
		{
			tempEquipTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*EquipTemplate)(nil))
			if tempEquipTemplate == nil {
				err = fmt.Errorf("equipment [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.equipmentTemplate = tempEquipTemplate.(*EquipTemplate)
			break
		}
	case itemtypes.ItemTypeSkill:
		{
			tempSkillTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*SkillTemplate)(nil))
			if tempSkillTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.skillTemplate = tempSkillTemplate.(*SkillTemplate)
			break
		}
	case itemtypes.ItemTypeGem:
		{
			tempAttrTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*AttrTemplate)(nil))
			if tempAttrTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.gemAttrTemplate = tempAttrTemplate.(*AttrTemplate)
			break
		}
	case itemtypes.ItemTypeShenLong:
		{
			tempAttrTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*AttrTemplate)(nil))
			if tempAttrTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			attrTemplate, ok := tempAttrTemplate.(*AttrTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.dragonAttrTemplate = attrTemplate
		}
	case itemtypes.ItemTypeGiftBag:
		{
			tempBoxTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*BoxTemplate)(nil))
			if tempBoxTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			boxTemplate, ok := tempBoxTemplate.(*BoxTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.boxTemplate = boxTemplate

			if it.TypeFlag2 > 0 {
				tempBoxTemplate := template.GetTemplateService().Get(int(it.TypeFlag2), (*BoxTemplate)(nil))
				if tempBoxTemplate == nil {
					err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				boxTemplate, ok := tempBoxTemplate.(*BoxTemplate)
				if !ok {
					err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				it.boxTemplateSecond = boxTemplate
			}

			break
		}
	case itemtypes.ItemTypeGoldEquip:
		{
			tempGoldEquipTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*GoldEquipTemplate)(nil))
			if tempGoldEquipTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			goldEquipTemplate, ok := tempGoldEquipTemplate.(*GoldEquipTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.goldEquipTemplate = goldEquipTemplate

			//验证下一级神铸装备的性别、角色
			godcastingEquip := it.goldEquipTemplate.GetGodCastingEquipTemp()
			if godcastingEquip != nil {
				nextItemTemp := godcastingEquip.GetNextItemTemplate()
				if nextItemTemp != nil {
					if it.NeedGender != nextItemTemp.NeedGender {
						err = fmt.Errorf("godcastingequip [%d]->[%d] invalid, sex type wrong", it.TypeFlag1, nextItemTemp.TypeFlag1)
						return template.NewTemplateFieldError("typeFlag1", err)
					}
					if it.NeedProfession != nextItemTemp.NeedProfession {
						err = fmt.Errorf("godcastingequip [%d]->[%d] invalid, role type wrong", it.TypeFlag1, nextItemTemp.TypeFlag1)
						return template.NewTemplateFieldError("typeFlag1", err)
					}
					if it.SubType != nextItemTemp.SubType {
						err = fmt.Errorf("godcastingequip [%d]->[%d] invalid, subtype wrong", it.TypeFlag1, nextItemTemp.TypeFlag1)
						return template.NewTemplateFieldError("typeFlag1", err)
					}
				}
			}
			break
		}
	case itemtypes.ItemTypeShenQi:
		{
			if it.itemSubType == itemtypes.ItemShenQiSubTypeQiLing {
				temp := template.GetTemplateService().Get(int(it.TypeFlag1), (*ShenQiQiLingTemplate)(nil))
				if temp == nil {
					err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				qiLingTemplate, ok := temp.(*ShenQiQiLingTemplate)
				if !ok {
					err = fmt.Errorf("qiling [%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
				it.shenQiQiLingTemplate = qiLingTemplate
			}
			break
		}
	case itemtypes.ItemTypeMountEquip,
		itemtypes.ItemTypeWingStone,
		itemtypes.ItemTypeAnqiJiguan,
		itemtypes.ItemTypeFaBaoSuit,
		itemtypes.ItemTypeXianTiLingYu,
		itemtypes.ItemTypeLingyuEquip,
		itemtypes.ItemTypeShenfaEquip,
		itemtypes.ItemTypeShiHunFanEquip,
		itemtypes.ItemTypeTianMoTiEquip,
		itemtypes.ItemTypeLingTongMountEquip,
		itemtypes.ItemTypeLingTongWingEquip,
		itemtypes.ItemTypeLingTongShenFaEquip,
		itemtypes.ItemTypeLingTongLingYuEquip,
		itemtypes.ItemTypeLingTongFaBaoEquip,
		itemtypes.ItemTypeLingTongXianTiEquip,
		itemtypes.ItemTypeLingTongWeaponEquip,
		itemtypes.ItemTypeLingTongEquip:
		{
			tempSystemEquipTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*SystemEquipTemplate)(nil))
			if tempSystemEquipTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			systemEquipTemplate, ok := tempSystemEquipTemplate.(*SystemEquipTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.systemEquipTemplate = systemEquipTemplate
			break
		}

	case itemtypes.ItemTypeShengHenEquipBaiHu,
		itemtypes.ItemTypeShengHenEquipQingLong,
		itemtypes.ItemTypeShengHenEquipZhuQue,
		itemtypes.ItemTypeShengHenEquipXuanWu:
		{
			shengHenEquipTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*SystemEquipShengHenTemplate)(nil))
			if shengHenEquipTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			equipTemplate, ok := shengHenEquipTemplate.(*SystemEquipShengHenTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.systemEquipTemplate = equipTemplate
			break
		}
	case itemtypes.ItemTypeTuLongEquip:
		{
			tulongEquipTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*TuLongEquipTemplate)(nil))
			if tulongEquipTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			equipTemplate, ok := tulongEquipTemplate.(*TuLongEquipTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.tuLongEquipTemplate = equipTemplate
			break
		}
	case itemtypes.ItemTypeBabyToy:
		{
			to := template.GetTemplateService().Get(int(it.TypeFlag1), (*BabyToyTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			toyTemplate, ok := to.(*BabyToyTemplate)
			if !ok {
				err = fmt.Errorf("dan [%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.babyToyTemplate = toyTemplate
			break
		}
	case itemtypes.ItemTypeTitle:
		{
			if it.itemSubType == itemtypes.ItemTitleSubTypeDingZhiCard {
				titleDingZhiTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*TitleDingZhiTemplate)(nil))
				if titleDingZhiTemplate == nil {
					err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
					return template.NewTemplateFieldError("typeFlag1", err)
				}
			}
			break
		}
	case itemtypes.ItemTypeExpendBagSlotCard:
		{
			if validator.MinValidate(float64(it.TypeFlag1), 0, false); err != nil {
				return template.NewTemplateFieldError("TypeFlag1", err)
			}
			break
		}
	case itemtypes.ItemTypeYuanBaoKa:
		{
			if validator.RangeValidate(float64(it.TypeFlag1), 0, false, maxYuanBaoKa, true); err != nil {
				return template.NewTemplateFieldError("TypeFlag1", err)
			}
			break
		}
	case itemtypes.ItemTypeVigorousPill:
		{
			linShiTemplate := template.GetTemplateService().Get(int(it.TypeFlag1), (*LinShiAttrTemplate)(nil))
			if linShiTemplate == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			break
		}
	case itemtypes.ItemTypeWushuangWeapon:
		{
			wuhsuangBase := template.GetTemplateService().Get(int(it.TypeFlag1), (*WushuangWeaponBaseTemplate)(nil))
			if wuhsuangBase == nil {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			wuhsuangBaseTemp, ok := wuhsuangBase.(*WushuangWeaponBaseTemplate)
			if !ok {
				err = fmt.Errorf("[%d] invalid", it.TypeFlag1)
				return template.NewTemplateFieldError("typeFlag1", err)
			}
			it.wushuangWeaponBaseTemplate = wuhsuangBaseTemp
			break
		}
	}

	//关联运营活动group
	it.groupList, err = coreutils.SplitAsIntArray(it.RankGroup)
	if err != nil {
		err = template.NewTemplateFieldError("RankGroup", err)
		return
	}

	//拆解获得物品
	it.chaiJieItemMap = make(map[int32]int32)
	chaiJieItemArr, err := coreutils.SplitAsIntArray(it.ChaiJieItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", it.ChaiJieItemId)
		err = template.NewTemplateFieldError("ChaiJieItemId", err)
		return
	}
	chaiJieCountArr, err := coreutils.SplitAsIntArray(it.ChaiJieItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", it.ChaiJieItemCount)
		err = template.NewTemplateFieldError("ChaiJieItemCount", err)
		return
	}
	if len(chaiJieItemArr) != len(chaiJieCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", it.ChaiJieItemId, it.ChaiJieItemCount)
		err = template.NewTemplateFieldError("ChaiJieItemId or ChaiJieItemCount", err)
		return
	}
	for index, itemId := range chaiJieItemArr {
		it.chaiJieItemMap[itemId] += chaiJieCountArr[index]
	}
	return nil
}

func (it *ItemTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(it.FileName(), it.TemplateId(), err)
			return
		}
	}()

	//验证cd时间
	if err = validator.MinValidate(float64(it.CdTime), 0, true); err != nil {
		return template.NewTemplateFieldError("cdTime", err)
	}

	//验证购买元宝
	// if err = validator.MinValidate(float64(it.BuyGold), 0, true); err != nil {
	// 	return template.NewTemplateFieldError("buyGold", err)
	// }
	//验证出售比例
	if err = validator.RangeValidate(float64(it.SaleRate), 0, true, common.MAX_RATE, true); err != nil {
		return template.NewTemplateFieldError("saleRate", err)
		//验证购买银两
		if err = validator.MinValidate(float64(it.BuySilver), 0, false); err != nil {
			return template.NewTemplateFieldError("buySilver", err)
		}

	}

	//验证叠加上限
	if err = validator.MinValidate(float64(it.MaxOverlap), float64(1), true); err != nil {
		return template.NewTemplateFieldError("maxOverlap", err)
	}

	//限制使用次数
	if err = validator.MinValidate(float64(it.LimitTimeDay), float64(0), true); err != nil {
		return template.NewTemplateFieldError("limitTimeDay", err)
	}
	//限制使用总次数
	if err = validator.MinValidate(float64(it.LimitTimeAll), float64(0), true); err != nil {
		return template.NewTemplateFieldError("limitTimeAll", err)
	}

	//检查叠加属性
	if it.itemType == itemtypes.ItemTypeGoldEquip || it.itemType == itemtypes.ItemTypeEquipment {
		if it.MaxOverlap > 1 {
			err := fmt.Errorf("%d invalid, equip can`t overlap", it.MaxOverlap)
			return template.NewTemplateFieldError("MaxOverlap", err)
		}
	}

	//仓库积分获取
	err = validator.MinValidate(float64(it.UnionGet), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", it.UnionGet)
		return template.NewTemplateFieldError("UnionGet", err)
	}

	//仓库积分消耗
	err = validator.MinValidate(float64(it.UnionUse), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", it.UnionUse)
		return template.NewTemplateFieldError("UnionUse", err)
	}

	//生效时间
	if len(it.ShengXiaoTime) > 0 {
		it.effectiveTime, err = timeutils.ParseYYYYMMDD(it.ShengXiaoTime)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", it.ShengXiaoTime)
			return template.NewTemplateFieldError("ShengXiaoTime", err)
		}
	}

	if it.MarketId != 0 {
		//最低上架价格
		err = validator.MinValidate(float64(it.MarketMinPrice), float64(0), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", it.MarketMinPrice)
			return template.NewTemplateFieldError("MarketMinPrice", err)
		}

		//回购价格
		err = validator.MinValidate(float64(it.HuigouPrice), float64(it.MarketMinPrice), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", it.HuigouPrice)
			return template.NewTemplateFieldError("HuigouPrice", err)
		}
	}

	//拆解物品校验
	for itemId, num := range it.chaiJieItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", it.ChaiJieItemId)
			err = template.NewTemplateFieldError("ChaiJieItemId", err)
			return
		}
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", it.ChaiJieItemCount)
			return template.NewTemplateFieldError("ChaiJieItemCount", err)
		}
	}

	return nil
}

func (it *ItemTemplate) PatchAfterCheck() {

}

func (it *ItemTemplate) FileName() string {
	return "tb_item.json"
}
func init() {
	template.Register((*ItemTemplate)(nil))
}
