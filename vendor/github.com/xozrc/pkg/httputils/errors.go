package httputils

import (
	"reflect"
)

func NewInvalidBindError(t reflect.Type) error {
	return &InvalidBindError{t}
}

type InvalidBindError struct {
	Type reflect.Type
}

func (ibe *InvalidBindError) Error() string {
	if ibe.Type == nil {
		return "bind(nil)"
	}
	if ibe.Type.Kind() != reflect.Ptr {
		return "bind(non-pointer " + ibe.Type.String() + ")"
	}
	return "bind(nil " + ibe.Type.String() + ")"
}

func NewInvalidBindFieldError(t reflect.Type) error {
	return &InvalidBindFieldError{t}
}

type InvalidBindFieldError struct {
	Type reflect.Type
}

func (ibfe *InvalidBindFieldError) Error() string {
	if ibfe.Type == nil {
		return "bind field(nil)"
	}
	if ibfe.Type.Kind() != reflect.Ptr {
		return "bind field(non-pointer " + ibfe.Type.String() + ")"
	}
	return "bind field(nil " + ibfe.Type.String() + ")"
}
