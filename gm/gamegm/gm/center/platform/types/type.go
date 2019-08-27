package types

//中心平台设置类型
type SettingType int32

const (
	SettingTypeInput SettingType = iota + 1
	SettingTypeNumber
	SettingTypeSelect
)

var (
	settingTypeMap = map[SettingType]string{
		SettingTypeInput:  "字符串配置",
		SettingTypeNumber: "数字配置",
		SettingTypeSelect: "选择框配置",
	}
)

func (m SettingType) Valid() bool {
	switch m {
	case SettingTypeInput:
	case SettingTypeNumber:
	case SettingTypeSelect:
	default:
		return false
	}
	return true
}

func (m SettingType) String() string {
	return settingTypeMap[m]
}

//配置项结构
type SettingObject struct {
	Code        string                     `json:"code"`        //设置编码
	Type        SettingType                `json:"settingType"` //设置类型
	TipName     string                     `json:"tipName"`     //设置前面提示的文字
	SelectArray []*SettingObjectSelectItem `json:"selectItem"`  //如果是下拉列表的话下拉列表项目
}

type SettingObjectSelectItem struct {
	Key   int32  `json:"key"`   //下拉列表项
	Value string `json:"value"` //下拉列表显示名称
}

func GetPlatformSetting() []*SettingObject {
	rst := make([]*SettingObject, 0)
	rst = append(rst, newMarrySetting())
	rst = append(rst, newXianMengCangKu())
	rst = append(rst, newJiaoYiHang())
	rst = append(rst, newCashTiXianFlag())
	rst = append(rst, newNeiWanJiaoYiFlag())
	rst = append(rst, newZhiZuanFlag())
	return rst
}

func newMarrySetting() *SettingObject {
	marryItem := &SettingObject{
		Code:    "marrySet",
		Type:    SettingTypeSelect,
		TipName: "结婚价格版本类型",
	}
	marryItem.SelectArray = make([]*SettingObjectSelectItem, 0)
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "当前版本",
	}
	item2 := &SettingObjectSelectItem{
		Key:   2,
		Value: "廉价版本",
	}
	item3 := &SettingObjectSelectItem{
		Key:   3,
		Value: "豪华版本",
	}
	marryItem.SelectArray = append(marryItem.SelectArray, item1)
	marryItem.SelectArray = append(marryItem.SelectArray, item2)
	marryItem.SelectArray = append(marryItem.SelectArray, item3)
	return marryItem
}

func newXianMengCangKu() *SettingObject {
	xianMengItem := &SettingObject{
		Code:    "allianceWarehouseFlag",
		Type:    SettingTypeSelect,
		TipName: "仙盟仓库开关",
	}
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "开",
	}
	xianMengItem.SelectArray = append(xianMengItem.SelectArray, item1)

	item2 := &SettingObjectSelectItem{
		Key:   0,
		Value: "关",
	}
	xianMengItem.SelectArray = append(xianMengItem.SelectArray, item2)
	return xianMengItem
}

func newJiaoYiHang() *SettingObject {
	jiaoYiHangItem := &SettingObject{
		Code:    "jiaoYiHangFlag",
		Type:    SettingTypeSelect,
		TipName: "交易行开关",
	}
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "开",
	}
	jiaoYiHangItem.SelectArray = append(jiaoYiHangItem.SelectArray, item1)

	item2 := &SettingObjectSelectItem{
		Key:   0,
		Value: "关",
	}
	jiaoYiHangItem.SelectArray = append(jiaoYiHangItem.SelectArray, item2)
	return jiaoYiHangItem
}

func newCashTiXianFlag() *SettingObject {
	cashTiXianFlag := &SettingObject{
		Code:    "cashTiXianFlag",
		Type:    SettingTypeSelect,
		TipName: "现金提现开关",
	}
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "开",
	}
	cashTiXianFlag.SelectArray = append(cashTiXianFlag.SelectArray, item1)

	item2 := &SettingObjectSelectItem{
		Key:   0,
		Value: "关",
	}
	cashTiXianFlag.SelectArray = append(cashTiXianFlag.SelectArray, item2)
	return cashTiXianFlag
}

func newNeiWanJiaoYiFlag() *SettingObject {
	neiWanJiaoYiFlag := &SettingObject{
		Code:    "neiWanJiaoYiFlag",
		Type:    SettingTypeSelect,
		TipName: "内玩交易开关",
	}
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "开",
	}
	neiWanJiaoYiFlag.SelectArray = append(neiWanJiaoYiFlag.SelectArray, item1)

	item2 := &SettingObjectSelectItem{
		Key:   0,
		Value: "关",
	}
	neiWanJiaoYiFlag.SelectArray = append(neiWanJiaoYiFlag.SelectArray, item2)
	return neiWanJiaoYiFlag
}

func newZhiZuanFlag() *SettingObject {
	newZhiZuanFlag := &SettingObject{
		Code:    "zhiZuanFlag",
		Type:    SettingTypeSelect,
		TipName: "至尊会员",
	}
	item1 := &SettingObjectSelectItem{
		Key:   1,
		Value: "当前版本",
	}
	newZhiZuanFlag.SelectArray = append(newZhiZuanFlag.SelectArray, item1)

	item2 := &SettingObjectSelectItem{
		Key:   2,
		Value: "豪华版本",
	}
	newZhiZuanFlag.SelectArray = append(newZhiZuanFlag.SelectArray, item2)
	return newZhiZuanFlag
}

func newInputSetting() *SettingObject {
	inputItem := &SettingObject{
		Code:    "textInput",
		Type:    SettingTypeInput,
		TipName: "文本配置",
	}
	return inputItem
}

func newNumberSetting() *SettingObject {
	numberItem := &SettingObject{
		Code:    "numberInput",
		Type:    SettingTypeNumber,
		TipName: "数字配置",
	}
	return numberItem
}
