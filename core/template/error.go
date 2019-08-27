package template

import (
	"fmt"
)

//模板错误
type TemplateError struct {
	err      error
	fileName string
	id       int
}

func (te *TemplateError) Error() string {
	return fmt.Sprintf("template: file[%s] id[%d]  error[%s]", te.fileName, te.id, te.err.Error())
}

func NewTemplateError(fileName string, id int, err error) *TemplateError {
	te := &TemplateError{
		fileName: fileName,
		id:       id,

		err: err,
	}
	return te
}

type TemplateFieldError struct {
	err       error
	fieldName string
}

func (te *TemplateFieldError) Error() string {
	if te.err == nil {
		return fmt.Sprintf("template: field[%s]", te.fieldName)
	}
	return fmt.Sprintf("template: field[%s]  error[%s]", te.fieldName, te.err.Error())
}
func NewTemplateFieldError(fieldName string, err error) *TemplateFieldError {
	te := &TemplateFieldError{
		fieldName: fieldName,
		err:       err,
	}
	return te
}
