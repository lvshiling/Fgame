/*此类自动生成,请勿修改*/
package template

/*英灵谱配置*/
type YinglingpuTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型,0为英灵图鉴,1为至尊英灵图鉴
	Type int32 `json:"type"`

	//图鉴id
	TujianId int32 `json:"tujian_id"`

	//图鉴名称
	Name string `json:"name"`

	//图鉴升级起始id
	LevelBeginId int32 `json:"level_begin_id"`

	//碎片起始id
	SuiPianBeginId int32 `json:"suipian_begin_id"`

	//图片标志，0为无标志，1为稀有标志
	SpecialTag int32 `json:"special_tag"`

	//位置
	PosId int32 `json:"pos_id"`

	//图鉴图片资源
	Resource string `json:"resource"`
}
