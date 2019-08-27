/*此类自动生成,请勿修改*/
package template

/*仙桃大会常量配置*/
type XianTaoConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//玩家身上仙桃的最低数量
	XianTaoMin int32 `json:"xiantao_min"`

	//玩家身上仙桃的最高数量
	XianTaoMax int32 `json:"xiantao_max"`

	//玩家采集的次数限制
	CaiJiLimit int32 `json:"caiji_limit"`

	//玩家劫取限制次数
	JieQuLimit int32 `json:"jiequ_limit"`

	//千年仙桃的buffid
	XianTaoBuff int32 `json:"xiantao_buff"`

	//百年仙桃的buffid
	XianTaoBuff2 int32 `json:"xiantao_buff2"`

	//被劫取不再损失仙桃时身上的buff
	BuSunBuff int32 `json:"busun_buff"`

	//上交NPC的id
	TaoXianBiologyId int32 `json:"taoxian_biology_id"`

	//提交仙桃距离
	TiJiaoService int32 `json:"tijiao_service"`
}
