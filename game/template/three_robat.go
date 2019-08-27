package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"math/rand"
)

//3v3竞技场机器人
type ThreeRobatTemplate struct {
	*ThreeRobatTemplateVO
}

func (t *ThreeRobatTemplate) TemplateId() int {
	return t.Id
}

func (t *ThreeRobatTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ThreeRobatTemplate) PatchAfterCheck() {

}

func (t *ThreeRobatTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证robot_min
	err = validator.MinValidate(float64(t.RobotMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("RobotMin", err)
		return
	}

	//验证robot_max
	err = validator.MinValidate(float64(t.RobotMax), float64(t.RobotMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("RobotMax", err)
		return
	}

	//验证reborn_min
	err = validator.MinValidate(float64(t.RebornMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("RebornMin", err)
		return
	}
	//验证reborn_max
	err = validator.MinValidate(float64(t.RebornMax), float64(t.RebornMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("RebornMax", err)
		return
	}

	//验证time_min
	err = validator.MinValidate(float64(t.TimeMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("TimeMin", err)
		return
	}
	//验证time_max
	err = validator.MinValidate(float64(t.TimeMax), float64(t.TimeMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("TimeMax", err)
		return
	}

	return nil
}

func (t *ThreeRobatTemplate) RandomRobot() int32 {
	if t.RobotMax-t.RobotMin == 0 {
		return t.RobotMin
	}
	return rand.Int31n(t.RobotMax-t.RobotMin) + t.RobotMin
}

func (t *ThreeRobatTemplate) RandomReborn() int32 {
	if t.RebornMax-t.RebornMin == 0 {
		return t.RebornMin
	}
	return rand.Int31n(t.RebornMax-t.RebornMin) + t.RebornMin
}

func (t *ThreeRobatTemplate) RandomTime() int32 {
	if t.TimeMax-t.TimeMin == 0 {
		return t.TimeMin
	}
	return rand.Int31n(t.TimeMax-t.TimeMin) + t.TimeMin
}

func (t *ThreeRobatTemplate) FileName() string {
	return "tb_three_robat.json"
}

func init() {
	template.Register((*ThreeRobatTemplate)(nil))
}
