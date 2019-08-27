/*此类自动生成,请勿修改*/
package template

/*上古之灵升级配置*/
type ShangguzhilingLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一Id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升级所需经验
	Experience int64 `json:"experience"`

	//hp
	Hp int32 `json:"hp"`

	//攻击力
	Attack int32 `json:"attack"`

	//防御力
	Defence int32 `json:"defence"`

	//宝箱CD
	BaoxiangCd int32 `json:"baoxiang_cd"`

	//宝箱掉落关联ID
	BaoxiangDrop int32 `json:"baoxiang_drop"`
}
