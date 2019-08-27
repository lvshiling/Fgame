/*此类自动生成,请勿修改*/
package template

/*场景部怪配置*/
type SceneTemplateVO struct {

	//idx
	Idx int `json:"idx"`

	//id
	Id int `json:"id"`

	//场景id
	SceneID int32 `json:"sceneID"`

	//模型id
	TempId int `json:"tempID"`

	//生物类型
	BiologyType int32 `json:"biologyType"`

	//索引id
	IndexID int32 `json:"indexID"`

	//组id
	GroupID int32 `json:"groupID"`

	//名字
	Name string `json:"name"`

	//posX
	PosX float64 `json:"posX"`

	//posY
	PosY float64 `json:"posY"`

	//posZ
	PosZ float64 `json:"posZ"`

	//角度
	Angle float64 `json:"angle"`

	//警戒半径
	AlertRadius float64 `json:"alertRadius"`

	//最大移动半径
	RandRadius float64 `json:"randRadius"`
}
