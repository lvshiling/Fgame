package template

import (
	"fgame/fgame/core/template"
	droptemplate "fgame/fgame/game/drop/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//元神金装模板处理
type GoldEquipTemplateService interface {
	GetGoldEquipTemplate(id int) *gametemplate.GoldEquipTemplate
	//重铸模板
	GetGoldEquipChongzhuTemplat(role playertypes.RoleType, sex playertypes.SexType, quality, level int32) *gametemplate.GoldEquipChongzhuTemplate
	//金装模板
	GetGoldEquipTemplateByGroup(groupId int32) *gametemplate.GoldEquipTemplate
	//转生模板
	GetZhuanShengTemplate(zhuanShu int32) *gametemplate.ZhuanShengTemplate
	//元神等级模板
	GetGoldYuanTemplate(level int32) *gametemplate.GoldYuanTemplate
	//获取强化概率模板
	GetGoldEquipStrengthenRateTemplate(id int32) *gametemplate.GoldEquipStrengthenRateTemplate
	//附加属性模板
	GetFuJiaAttrTemplate(id int32) *gametemplate.GoldEquipFuJiaAttrTemplate
	// 套装目标模板
	GetMuBiaoTaoZhuangTemplate(mubiaoType goldequiptypes.TaoZhuangMuBiaoType, curLevel int32) *gametemplate.TaozhuangMuBiaoTemplate
	//装备继承模板
	GetGoldEquipJiChengTemplate(level int32) *gametemplate.GoldEquipJiChengTemplate
	//获取强化部位模板
	GetGoldEquipStrengthenBuWeiTemplate(pos inventorytypes.BodyPositionType, level int32) *gametemplate.GoldEquipStrengthenBuWeiTemplate
	//获取铸灵模板
	GetCastingSpiritTemplate(bodyPos inventorytypes.BodyPositionType, spiritType goldequiptypes.SpiritType) *gametemplate.GodCastingCastingSpiritTemplate
	//获取锻魂模板
	GetForgeSoulTemplate(bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType) *gametemplate.GodCastingForgeSoulTemplate
	//获取神铸继承模板
	GetGodcastingJiChengTemplate(level int32) *gametemplate.GodCastingJichengTemplate
	//获得神铸物品目标阶级的进阶物品ID
	GetGodCastingInheritTargetLevelItemId(goldequipTemp *gametemplate.GoldEquipTemplate, targetLevel int32) (itemId int32, flag bool)
}

type goldEquipTemplateService struct {
	//元神金装配置
	goldEquipMap map[int]*gametemplate.GoldEquipTemplate
	//元神金装重铸配置
	goldEquipChongzhuMap map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.GoldEquipChongzhuTemplate
	//元神金装模板按套装
	goldEquipByGroupMap map[int32]*gametemplate.GoldEquipTemplate
	//转生配置
	zhuanShengMap map[int32]*gametemplate.ZhuanShengTemplate
	//元神等级配置
	goldYuanMap map[int32]*gametemplate.GoldYuanTemplate
	//元神金装强化概率配置
	goldEquipStrengthenRateMap map[int32]*gametemplate.GoldEquipStrengthenRateTemplate
	//金装附加属性配置
	goldequipAttrMap map[int32]*gametemplate.GoldEquipFuJiaAttrTemplate
	//套装目标配置
	mubiaoTaoZhuangMap map[goldequiptypes.TaoZhuangMuBiaoType]map[int32]*gametemplate.TaozhuangMuBiaoTemplate
	//强化继承
	goldEquipJiChengMap map[int32]*gametemplate.GoldEquipJiChengTemplate
	//强化
	goldEquipStrengthenBuWeiMap map[inventorytypes.BodyPositionType]map[int32]*gametemplate.GoldEquipStrengthenBuWeiTemplate
	//铸灵
	castingSpiritMap map[inventorytypes.BodyPositionType]map[goldequiptypes.SpiritType]*gametemplate.GodCastingCastingSpiritTemplate
	//锻魂
	forgeSoulMap map[inventorytypes.BodyPositionType]map[goldequiptypes.ForgeSoulType]*gametemplate.GodCastingForgeSoulTemplate
	//神铸继承
	godCastingJichengMap map[int32]*gametemplate.GodCastingJichengTemplate
}

func (s *goldEquipTemplateService) GetGoldEquipTemplate(id int) *gametemplate.GoldEquipTemplate {
	return s.goldEquipMap[id]
}

//初始化
func (s *goldEquipTemplateService) init() error {
	s.goldEquipMap = make(map[int]*gametemplate.GoldEquipTemplate)
	//元神金装
	templateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipTemplate)(nil))
	for _, templateObject := range templateMap {
		goldequipTemplate, _ := templateObject.(*gametemplate.GoldEquipTemplate)
		s.goldEquipMap[goldequipTemplate.TemplateId()] = goldequipTemplate
	}

	s.goldEquipChongzhuMap = make(map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.GoldEquipChongzhuTemplate)
	//金装重铸
	chongzhuTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipChongzhuTemplate)(nil))
	for _, templateObject := range chongzhuTemplateMap {
		chongzhuTemplate, _ := templateObject.(*gametemplate.GoldEquipChongzhuTemplate)

		role := playertypes.RoleType(chongzhuTemplate.Profession)
		sex := playertypes.SexType(chongzhuTemplate.Gender)

		sexMap, ok := s.goldEquipChongzhuMap[role]
		if !ok {
			sexMap = make(map[playertypes.SexType][]*gametemplate.GoldEquipChongzhuTemplate)
			s.goldEquipChongzhuMap[role] = sexMap
		}
		sexMap[sex] = append(sexMap[sex], chongzhuTemplate)

		//验证掉落
		dropId := chongzhuTemplate.DropId
		flag := droptemplate.GetDropTemplateService().CheckSureDrop(dropId)
		if !flag {
			return fmt.Errorf("goldequip: 元神金装重铸配置掉落应该是必定掉落的")
		}
	}

	//金装模板
	s.goldEquipByGroupMap = make(map[int32]*gametemplate.GoldEquipTemplate)
	groupTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipTemplate)(nil))
	for _, templateObject := range groupTemplateMap {
		groupTemplate, _ := templateObject.(*gametemplate.GoldEquipTemplate)
		s.goldEquipByGroupMap[groupTemplate.SuitGroup] = groupTemplate
	}

	s.zhuanShengMap = make(map[int32]*gametemplate.ZhuanShengTemplate)
	zhuanShengTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ZhuanShengTemplate)(nil))
	for _, templateObject := range zhuanShengTemplateMap {
		zhuanShengTemplate, _ := templateObject.(*gametemplate.ZhuanShengTemplate)
		s.zhuanShengMap[int32(zhuanShengTemplate.TemplateId())] = zhuanShengTemplate
	}

	// 元神等级配置
	s.goldYuanMap = make(map[int32]*gametemplate.GoldYuanTemplate)
	goldYuanTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldYuanTemplate)(nil))
	for _, temp := range goldYuanTemplateMap {
		goldYuanTemplate, _ := temp.(*gametemplate.GoldYuanTemplate)
		s.goldYuanMap[goldYuanTemplate.Level] = goldYuanTemplate
	}

	// 强化概率
	s.goldEquipStrengthenRateMap = make(map[int32]*gametemplate.GoldEquipStrengthenRateTemplate)
	rateTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipStrengthenRateTemplate)(nil))
	for _, temp := range rateTemplateMap {
		rateTemp, _ := temp.(*gametemplate.GoldEquipStrengthenRateTemplate)
		s.goldEquipStrengthenRateMap[int32(rateTemp.Id)] = rateTemp
	}

	// 强化概率
	s.goldequipAttrMap = make(map[int32]*gametemplate.GoldEquipFuJiaAttrTemplate)
	attrTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipFuJiaAttrTemplate)(nil))
	for _, temp := range attrTemplateMap {
		attrTemp, _ := temp.(*gametemplate.GoldEquipFuJiaAttrTemplate)
		s.goldequipAttrMap[int32(attrTemp.Id)] = attrTemp
	}

	//宝石套装
	s.mubiaoTaoZhuangMap = make(map[goldequiptypes.TaoZhuangMuBiaoType]map[int32]*gametemplate.TaozhuangMuBiaoTemplate)
	mubiaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.TaozhuangMuBiaoTemplate)(nil))
	for _, temp := range mubiaoTemplateMap {
		mubiaoTemp, _ := temp.(*gametemplate.TaozhuangMuBiaoTemplate)

		subMap, ok := s.mubiaoTaoZhuangMap[mubiaoTemp.GetMuBiaoType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.TaozhuangMuBiaoTemplate)
			s.mubiaoTaoZhuangMap[mubiaoTemp.GetMuBiaoType()] = subMap
		}

		subMap[mubiaoTemp.NeedLevel] = mubiaoTemp
	}

	// 强化概率
	s.goldEquipJiChengMap = make(map[int32]*gametemplate.GoldEquipJiChengTemplate)
	jichengTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipJiChengTemplate)(nil))
	for _, temp := range jichengTemplateMap {
		jichengTemp, _ := temp.(*gametemplate.GoldEquipJiChengTemplate)
		s.goldEquipJiChengMap[jichengTemp.Level] = jichengTemp
	}

	s.goldEquipStrengthenBuWeiMap = make(map[inventorytypes.BodyPositionType]map[int32]*gametemplate.GoldEquipStrengthenBuWeiTemplate)
	strengthenBuWeiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GoldEquipStrengthenBuWeiTemplate)(nil))
	for _, templateObject := range strengthenBuWeiTemplateMap {
		strengthenBuWeiTemp, _ := templateObject.(*gametemplate.GoldEquipStrengthenBuWeiTemplate)
		levelMap, ok := s.goldEquipStrengthenBuWeiMap[strengthenBuWeiTemp.GetPosition()]
		if !ok {
			levelMap = make(map[int32]*gametemplate.GoldEquipStrengthenBuWeiTemplate)
			s.goldEquipStrengthenBuWeiMap[strengthenBuWeiTemp.GetPosition()] = levelMap
		}
		levelMap[strengthenBuWeiTemp.Level] = strengthenBuWeiTemp
	}

	//锻魂铸灵
	s.castingSpiritMap = make(map[inventorytypes.BodyPositionType]map[goldequiptypes.SpiritType]*gametemplate.GodCastingCastingSpiritTemplate)
	s.forgeSoulMap = make(map[inventorytypes.BodyPositionType]map[goldequiptypes.ForgeSoulType]*gametemplate.GodCastingForgeSoulTemplate)

	//初始化铸灵
	spiritTempMap := template.GetTemplateService().GetAll((*gametemplate.GodCastingCastingSpiritTemplate)(nil))
	for _, temp := range spiritTempMap {
		spiritTemp, _ := temp.(*gametemplate.GodCastingCastingSpiritTemplate)
		_, ok := s.castingSpiritMap[spiritTemp.GetBodyPos()]
		if !ok {
			s.castingSpiritMap[spiritTemp.GetBodyPos()] = make(map[goldequiptypes.SpiritType]*gametemplate.GodCastingCastingSpiritTemplate)
		}
		s.castingSpiritMap[spiritTemp.GetBodyPos()][spiritTemp.GetSpiritType()] = spiritTemp
	}

	//初始化锻魂
	soulTempMap := template.GetTemplateService().GetAll((*gametemplate.GodCastingForgeSoulTemplate)(nil))
	for _, temp := range soulTempMap {
		soulTemp, _ := temp.(*gametemplate.GodCastingForgeSoulTemplate)
		_, ok := s.forgeSoulMap[soulTemp.GetBodyPos()]
		if !ok {
			s.forgeSoulMap[soulTemp.GetBodyPos()] = make(map[goldequiptypes.ForgeSoulType]*gametemplate.GodCastingForgeSoulTemplate)
		}
		s.forgeSoulMap[soulTemp.GetBodyPos()][soulTemp.GetSoulType()] = soulTemp
	}

	//继承模板
	s.godCastingJichengMap = make(map[int32]*gametemplate.GodCastingJichengTemplate)
	gcjichengTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GodCastingJichengTemplate)(nil))
	for _, temp := range gcjichengTemplateMap {
		jichengTemp, _ := temp.(*gametemplate.GodCastingJichengTemplate)
		s.godCastingJichengMap[jichengTemp.Level] = jichengTemp
	}

	return nil
}

func (st *goldEquipTemplateService) GetCastingSpiritTemplate(bodyPos inventorytypes.BodyPositionType, spiritType goldequiptypes.SpiritType) *gametemplate.GodCastingCastingSpiritTemplate {
	return st.castingSpiritMap[bodyPos][spiritType]
}

func (st *goldEquipTemplateService) GetForgeSoulTemplate(bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType) *gametemplate.GodCastingForgeSoulTemplate {
	return st.forgeSoulMap[bodyPos][soulType]
}

func (st *goldEquipTemplateService) GetGodcastingJiChengTemplate(level int32) *gametemplate.GodCastingJichengTemplate {
	return st.godCastingJichengMap[level]
}

func (st *goldEquipTemplateService) GetGodCastingInheritTargetLevelItemId(goldequipTemp *gametemplate.GoldEquipTemplate, targetLevel int32) (itemId int32, flag bool) {
	goldTemp := goldequipTemp
	godcastingTemp := goldTemp.GetGodCastingEquipTemp()
	if godcastingTemp == nil {
		return 0, false
	}
	itemTemp := godcastingTemp.GetNextItemTemplate()
	if itemTemp == nil {
		return 0, false
	}
	for goldTemp.GetGodCastingEquipLevel() != targetLevel {
		godcastingTemp = goldTemp.GetGodCastingEquipTemp()
		if godcastingTemp == nil {
			return 0, false
		}
		itemTemp = godcastingTemp.GetNextItemTemplate()
		if itemTemp == nil {
			return 0, false
		}
		goldTemp = itemTemp.GetGoldEquipTemplate()
		if goldTemp == nil {
			return 0, false
		}
	}
	return int32(itemTemp.Id), true
}

func (s goldEquipTemplateService) GetGoldEquipChongzhuTemplat(role playertypes.RoleType, sex playertypes.SexType, quality, level int32) *gametemplate.GoldEquipChongzhuTemplate {
	chongzhuList := s.goldEquipChongzhuMap[role][sex]
	for _, temp := range chongzhuList {
		if temp.Quality == quality && temp.Level == level {
			return temp
		}
	}
	return nil
}

// 元神等级
func (s goldEquipTemplateService) GetGoldYuanTemplate(level int32) *gametemplate.GoldYuanTemplate {
	temp, ok := s.goldYuanMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (s goldEquipTemplateService) GetGoldEquipStrengthenRateTemplate(id int32) *gametemplate.GoldEquipStrengthenRateTemplate {
	temp, ok := s.goldEquipStrengthenRateMap[id]
	if !ok {
		return nil
	}

	return temp
}

func (s goldEquipTemplateService) GetFuJiaAttrTemplate(id int32) *gametemplate.GoldEquipFuJiaAttrTemplate {
	temp, ok := s.goldequipAttrMap[id]
	if !ok {
		return nil
	}

	return temp
}

func (s goldEquipTemplateService) GetGoldEquipJiChengTemplate(level int32) *gametemplate.GoldEquipJiChengTemplate {
	temp, ok := s.goldEquipJiChengMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (s goldEquipTemplateService) GetMuBiaoTaoZhuangTemplate(mubiaoType goldequiptypes.TaoZhuangMuBiaoType, curLevel int32) (temp *gametemplate.TaozhuangMuBiaoTemplate) {
	subMap, ok := s.mubiaoTaoZhuangMap[mubiaoType]
	if !ok {
		return nil
	}

	maxLevel := int32(0)
	for needLevel, _ := range subMap {
		if curLevel < needLevel {
			continue
		}
		if maxLevel >= needLevel {
			continue
		}
		maxLevel = needLevel
	}

	temp, ok = subMap[maxLevel]
	if !ok {
		return nil
	}
	return
}

//获取强化部位模板
func (s goldEquipTemplateService) GetGoldEquipStrengthenBuWeiTemplate(pos inventorytypes.BodyPositionType, level int32) *gametemplate.GoldEquipStrengthenBuWeiTemplate {
	to, ok := s.goldEquipStrengthenBuWeiMap[pos][level]
	if !ok {
		return nil
	}
	return to
}

//套装模板
func (s goldEquipTemplateService) GetGoldEquipTemplateByGroup(groupId int32) *gametemplate.GoldEquipTemplate {
	return s.goldEquipByGroupMap[groupId]
}

func (s goldEquipTemplateService) GetZhuanShengTemplate(zhuanShu int32) *gametemplate.ZhuanShengTemplate {
	zhuanShengTemplate, exist := s.zhuanShengMap[zhuanShu]
	if !exist {
		return nil
	}
	return zhuanShengTemplate
}

var (
	once sync.Once
	cs   *goldEquipTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &goldEquipTemplateService{}
		err = cs.init()
	})
	return err
}

func GetGoldEquipTemplateService() GoldEquipTemplateService {
	return cs
}
