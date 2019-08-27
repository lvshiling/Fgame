/*此类自动生成,请勿修改*/
package template

/*灵童时装配置*/
type LingTongFashionTemplateVO struct {

	//id
	Id int `json:"id"`

	//灵童时装名称
	Name string `json:"name"`

	//时装类型
	Type int32 `json:"type"`

	//激活灵童使用的物品id
	NeedItemId int32 `json:"need_item_id"`

	//时效性
	Time int32 `json:"time"`

	//激活灵童使用的物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//灵童独立攻击力
	LingTongAttack int64 `json:"lingtong_attack"`

	//灵童独立暴击
	LingTongCritical int64 `json:"lingtong_critical"`

	//灵童独立命中值
	LingTongHit int64 `json:"lingtong_hit"`

	//灵童独立破格
	LingTongAbnormality int64 `json:"lingtong_abnormality"`

	//玩家增加的生命上限
	Hp int64 `json:"hp"`

	//玩家增加的攻击值
	Attack int64 `json:"attack"`

	//玩家增加的防御值
	Defence int64 `json:"defence"`

	//灵童升级起始id
	LingTongUpstarId int32 `json:"lingtong_upstar_id"`
}
