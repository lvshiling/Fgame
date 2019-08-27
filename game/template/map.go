package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	pktypes "fgame/fgame/game/pk/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
	"math/rand"
)

type MapTemplate struct {
	//模板数据
	*MapTemplateVO
	//自定义数据
	mapType scenetypes.SceneType
	//副本地图专用
	sceneBiologyGroup map[int32]map[int32]*SceneTemplate
	//世界地图使用
	sceneBiologyMap map[int32]*SceneTemplate
	//地图数据
	mapData *template.Map
	//出生点
	bornPos coretypes.Position
	//复活地图
	rebornMap *MapTemplate
	//复活点
	rebornPos coretypes.Position
	//pk模式
	pkState pktypes.PkState
	//复活类型
	reliveType scenetypes.ReliveType
	//安全区
	safeAreaList         [][]coretypes.Position
	anQuanQuTemplateList []*AnquanquTemplate
	//限制区
	xianZhiAreaList         [][]coretypes.Position
	xianZhiQuYuTemplateList []*XianZhiQuYuTemplate

	standList []int32
}

func (mt *MapTemplate) GetReliveType() scenetypes.ReliveType {
	return mt.reliveType
}

func (mt *MapTemplate) IfCanEnter(enterType scenetypes.SceneEnterType) bool {
	return mt.SceneXianZhi&enterType.Mask() != 0
}

func (mt *MapTemplate) IsWorld() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeWorld
}

func (mt *MapTemplate) IsBoss() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeBoss
}

// func (mt *MapTemplate) IsWorldBoss() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeWorldBoss
// }
// func (mt *MapTemplate) IsUnrealBoss() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeUnrealBoss
// }

// func (mt *MapTemplate) IsOutlandBoss() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeOutlandBoss
// }

// func (mt *MapTemplate) IsCangJingGe() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeCangJingGe
// }

// func (mt *MapTemplate) IsZhenXi() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeZhenXiBoss
// }

// func (mt *MapTemplate) IsDingShi() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeDingShiBoss
// }
// func (mt *MapTemplate) IsCrossWorldBoss() bool {
// 	return mt.mapType.MapType() == scenetypes.MapTypeCrossWorldBoss
// }

func (mt *MapTemplate) IsActivitySub() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeActivitySub
}

func (mt *MapTemplate) IsActivity() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeActivity
}

func (mt *MapTemplate) IsActivityFuBen() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeActivityFuBen
}

func (mt *MapTemplate) IsFuBen() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeFuBen
}

func (mt *MapTemplate) IsArena() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeArena
}

func (mt *MapTemplate) IsMarry() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeMarry
}

func (mt *MapTemplate) IsTower() bool {
	return mt.mapType.MapType() == scenetypes.MapTypeTower
}

func (mt *MapTemplate) IsHiddenTitle() bool {
	return mt.IsTitle != 0
}

func (mt *MapTemplate) GetMapType() scenetypes.SceneType {
	return mt.mapType
}

func (mt *MapTemplate) GetBornPos() coretypes.Position {
	return mt.bornPos
}

func (mt *MapTemplate) GetRebornPos() coretypes.Position {
	return mt.rebornPos
}

func (mt *MapTemplate) GetSceneBiologyMap() map[int32]*SceneTemplate {
	return mt.sceneBiologyMap
}

func (mt *MapTemplate) GetSceneBiologyMapByGroup(groupId int32) map[int32]*SceneTemplate {
	return mt.sceneBiologyGroup[groupId]
}

func (mt *MapTemplate) GetNumGroup() int32 {
	return int32(len(mt.sceneBiologyGroup))
}

func (mt *MapTemplate) GetMap() *template.Map {
	return mt.mapData
}

func (mt *MapTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MapTemplate) GetPkState() pktypes.PkState {
	return mt.pkState
}

func (mt *MapTemplate) GetSafeArea() [][]coretypes.Position {
	return mt.safeAreaList
}

func (mt *MapTemplate) IsSafe(pos coretypes.Position) bool {
	if len(mt.safeAreaList) == 0 {
		return false
	}
	for _, safeArena := range mt.safeAreaList {
		if coreutils.PointInPolygon(pos, safeArena) {
			return true
		}
	}
	return false
}

func (mt *MapTemplate) IsInLimitArea(pos coretypes.Position) bool {
	if len(mt.xianZhiAreaList) == 0 {
		return true
	}
	for _, xianZhiArea := range mt.xianZhiAreaList {
		if coreutils.PointInPolygon(pos, xianZhiArea) {
			return true
		}
	}
	return false
}

func (mt *MapTemplate) CanShaqiDrop() bool {
	return mt.IsShaqiDrop != 0
}

func (mt *MapTemplate) CanShengWeiDrop() bool {
	return mt.IsShengWeiDrop != 0
}

func (mt *MapTemplate) IfCanShaLuDrop() bool {
	return mt.IsShaLuDrop != 0
}

func (mt *MapTemplate) IfChangeSceneProtect() bool {
	return mt.IsSceneProtect != 0
}

func (mt *MapTemplate) IfChangeScenePvp() bool {
	return mt.IsPkScene != 0
}

func (mt *MapTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()
	//场景类型
	mt.mapType = scenetypes.SceneType(mt.ScriptType)
	//几波怪
	mt.sceneBiologyGroup = make(map[int32]map[int32]*SceneTemplate)
	//所有怪物
	mt.sceneBiologyMap = make(map[int32]*SceneTemplate)

	toMap := template.GetTemplateService().GetAll((*SceneTemplate)(nil))
	for _, to := range toMap {
		sceneBiology := to.(*SceneTemplate)
		if sceneBiology.SceneID != int32(mt.TemplateId()) {
			continue
		}
		sceneBiologyMap, ok := mt.sceneBiologyGroup[sceneBiology.GroupID]
		if !ok {
			sceneBiologyMap = make(map[int32]*SceneTemplate)
			mt.sceneBiologyGroup[sceneBiology.GroupID] = sceneBiologyMap
		}
		sceneBiologyMap[int32(sceneBiology.TemplateId())] = sceneBiology
		mt.sceneBiologyMap[int32(sceneBiology.TemplateId())] = sceneBiology
	}

	//读取mask
	mt.mapData, err = template.GetTemplateService().ReadMap(mt.Resource)
	if err != nil {
		return nil
	}
	//出生地
	mt.bornPos = coretypes.Position{
		X: mt.BirthPosX,
		Y: mt.BirthPosY,
		Z: mt.BirthPosZ,
	}

	tempRebornMapTemplate := template.GetTemplateService().Get(int(mt.RebornId), (*MapTemplate)(nil))
	if tempRebornMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", mt.RebornId)
		return template.NewTemplateError("rebornId", mt.TemplateId(), err)
	}
	mt.rebornMap = tempRebornMapTemplate.(*MapTemplate)
	//重生地
	mt.rebornPos = coretypes.Position{
		X: mt.RebornX,
		Y: mt.RebornY,
		Z: mt.RebornZ,
	}
	//pk模式
	mt.pkState = pktypes.PkState(mt.PkMode)
	mt.reliveType = scenetypes.ReliveType(mt.ResurrectionType)

	if mt.AnquanquId != "" {
		anQuanQuIdArr, err := utils.SplitAsIntArray(mt.AnquanquId)
		if err != nil {
			return err
		}
		mt.anQuanQuTemplateList = make([]*AnquanquTemplate, 0, 4)
		for _, anQuanQuId := range anQuanQuIdArr {
			tempAnquanquTemplate := template.GetTemplateService().Get(int(anQuanQuId), (*AnquanquTemplate)(nil))
			if tempAnquanquTemplate == nil {
				err = fmt.Errorf("[%s] invalid", mt.AnquanquId)
				return template.NewTemplateError("AnquanquId", mt.TemplateId(), err)
			}
			anQuanQuTemplate := tempAnquanquTemplate.(*AnquanquTemplate)
			mt.anQuanQuTemplateList = append(mt.anQuanQuTemplateList, anQuanQuTemplate)
		}
	}
	if mt.XianzhiquyuId != "" {
		xianZhiQuYuIdArr, err := utils.SplitAsIntArray(mt.XianzhiquyuId)
		if err != nil {
			return err
		}
		mt.xianZhiQuYuTemplateList = make([]*XianZhiQuYuTemplate, 0, 4)
		for _, xianZhiQuYuId := range xianZhiQuYuIdArr {
			tempXianZhiQuYuTemplate := template.GetTemplateService().Get(int(xianZhiQuYuId), (*XianZhiQuYuTemplate)(nil))
			if tempXianZhiQuYuTemplate == nil {
				err = fmt.Errorf("[%s] invalid", mt.XianzhiquyuId)
				return template.NewTemplateError("XianzhiquyuId", mt.TemplateId(), err)
			}
			xianZhiQuYuTemplate := tempXianZhiQuYuTemplate.(*XianZhiQuYuTemplate)
			mt.xianZhiQuYuTemplateList = append(mt.xianZhiQuYuTemplateList, xianZhiQuYuTemplate)
		}
	}

	return nil
}

func (mt *MapTemplate) PatchAfterCheck() {
	maskMap := mt.GetMap().GetMaskMap().Mask
	length := mt.GetMap().GetLength()

	for z, zMask := range maskMap {
		for x, mask := range zMask {
			if mask == 0 {
				continue
			}
			pos := mt.GetMap().GetPosition(int32(x), int32(z))
			if mt.IsInLimitArea(pos) {
				mt.standList = append(mt.standList, int32(z)*int32(length)+int32(x))
			}
		}
	}
	return
}

func (mt *MapTemplate) RandomPosition() coretypes.Position {
	length := mt.GetMap().GetLength()

	for i := 0; i < 10; i++ {
		standIndex := rand.Intn(len(mt.standList))
		index := mt.standList[standIndex]
		x := index % int32(length)
		z := index / int32(length)
		pos := mt.GetMap().GetPosition(int32(x), int32(z))
		if !mt.GetMap().IsMask(pos.X, pos.Z) {
			continue
		}
		pos.Y = mt.GetMap().GetHeight(pos.X, pos.Z)
		return pos
	}
	panic("position error")

}

func (mt *MapTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//TODO 验证地图类型
	switch mt.mapType {
	case scenetypes.SceneTypeWorld:
		break
	case scenetypes.SceneTypeTianJieTa:
		//验证存活时间
		if err = validator.RangeValidate(float64(mt.LastTime), float64(common.MIN_FUBEN_TIME), true, float64(common.MAX_FUBEN_TIME), true); err != nil {
			return template.NewTemplateFieldError("LastTime", err)
		}
		//验证副本失败时间
		if err = validator.RangeValidate(float64(mt.PointsTime), float64(common.MIN_FUBEN_FAILTURE_TIME), true, float64(mt.LastTime), true); err != nil {
			return template.NewTemplateFieldError("PointsTime", err)
		}
		break
	}

	if mt.mapData == nil {
		err = fmt.Errorf("[%s] invalid", mt.Resource)
		return template.NewTemplateFieldError("Resource", err)
	}

	if !mt.mapData.IsMask(mt.bornPos.X, mt.bornPos.Z) {
		err = fmt.Errorf("[%.2f] [%2.f] invalid", mt.bornPos.X, mt.bornPos.Z)
		return template.NewTemplateFieldError("pos", err)
	}
	mt.bornPos.Y = mt.mapData.GetHeight(mt.bornPos.X, mt.bornPos.Z)

	if !mt.rebornMap.GetMap().IsMask(mt.rebornPos.X, mt.rebornPos.Z) {
		err = fmt.Errorf("[%.2f] [%2.f] invalid", mt.rebornPos.X, mt.rebornPos.Z)
		return template.NewTemplateFieldError("rebornPos", err)
	}

	mt.rebornPos.Y = mt.rebornMap.GetMap().GetHeight(mt.rebornPos.X, mt.rebornPos.Z)

	if !mt.rebornMap.IsInLimitArea(mt.rebornPos) {
		err = fmt.Errorf("重生点[%.2f] [%2.f] 不在限制区内", mt.rebornPos.X, mt.rebornPos.Z)
		return template.NewTemplateFieldError("rebornPos", err)
	}

	//验证pk模式
	if !mt.pkState.Valid() {
		err = fmt.Errorf("[%d]  invalid", mt.PkMode)
		return template.NewTemplateFieldError("PkState", err)
	}
	//验证限制pk模式是否包含pk模式
	if mt.LimitPkMode&mt.pkState.Mask() == 0 {
		err = fmt.Errorf("[%d]  invalid", mt.LimitPkMode)
		return template.NewTemplateFieldError("LimitPkMode", err)
	}

	//复活
	if !mt.reliveType.Valid() {
		err = fmt.Errorf("[%d]  invalid", mt.ResurrectionType)
		return template.NewTemplateFieldError("ResurrectionType", err)
	}
	//复活点不能是副本
	if mt.rebornMap.GetMapType().MapType() == scenetypes.MapTypeFuBen {
		err = fmt.Errorf("[%d]  invalid", mt.RebornId)
		return template.NewTemplateFieldError("RebornId", err)
	}

	for _, anQuanQuTemplate := range mt.anQuanQuTemplateList {
		safeArea := make([]coretypes.Position, 0, 1)
		safeArea = append(safeArea, anQuanQuTemplate.GetPos())
		currentAnquanquTemplate := anQuanQuTemplate
		for {
			currentAnquanquTemplate = currentAnquanquTemplate.GetNext()
			if currentAnquanquTemplate == nil {
				break
			}
			safeArea = append(safeArea, currentAnquanquTemplate.GetPos())
		}
		if len(safeArea) < 3 {
			err = fmt.Errorf("[%d]  不是多边形", mt.AnquanquId)
			return template.NewTemplateFieldError("AnquanquId", err)
		}
		mt.safeAreaList = append(mt.safeAreaList, safeArea)
	}

	for _, xianZhiQuYuTemplate := range mt.xianZhiQuYuTemplateList {
		xianZhiArea := make([]coretypes.Position, 0, 1)
		xianZhiArea = append(xianZhiArea, xianZhiQuYuTemplate.GetPos())
		currentXainZhiQuYuTemplate := xianZhiQuYuTemplate
		for {
			currentXainZhiQuYuTemplate = currentXainZhiQuYuTemplate.GetNext()
			if currentXainZhiQuYuTemplate == nil {
				break
			}
			xianZhiArea = append(xianZhiArea, currentXainZhiQuYuTemplate.GetPos())
		}
		if len(xianZhiArea) < 3 {
			err = fmt.Errorf("[%d]  不是多边形", mt.XianzhiquyuId)
			return template.NewTemplateFieldError("XianzhiquyuId", err)
		}
		mt.xianZhiAreaList = append(mt.xianZhiAreaList, xianZhiArea)
	}

	if !mt.IsInLimitArea(mt.bornPos) {
		err = fmt.Errorf("出生点 [%.2f] [%2.f] 不在限制区", mt.bornPos.X, mt.bornPos.Z)
		return template.NewTemplateFieldError("pos", err)
	}

	// pk保护等级
	if err = validator.RangeValidate(float64(mt.ProtectLevel), float64(0), true, float64(common.MAX_LEVEL), true); err != nil {
		return template.NewTemplateFieldError("ProtectLevel", err)
	}

	return nil
}
func (mt *MapTemplate) FileName() string {
	return "tb_map.json"
}

func init() {
	template.Register((*MapTemplate)(nil))
}
