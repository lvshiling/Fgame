package template

type WorldBossTemplateInterface interface {
	GetBiologyId() int32
	GetMapId() int32
	GetRecForce() int64
	GetBiologyTemplate() *BiologyTemplate
}
