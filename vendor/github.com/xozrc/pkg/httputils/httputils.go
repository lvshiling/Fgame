package httputils

import (
	"encoding/json"
	"errors"

	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

import (
	"github.com/xozrc/pkg/reflectutils"
)

const (
	formKey   string = "form"
	maxMemory        = int64(1024 * 1024 * 10)
)

func Bind(req *http.Request, obj interface{}) error {
	if req == nil {
		return errors.New("request is nil")
	}

	contentType := req.Header.Get("Content-Type")

	if strings.Contains(contentType, "form-urlencoded") {
		err := req.ParseForm()
		if err != nil {
			return err
		}

		return BindForm(obj, req.Form, nil)

	} else if strings.Contains(contentType, "multipart/form-data") {
		err := req.ParseMultipartForm(maxMemory)
		if err != nil {
			return err
		}
		return BindForm(obj, req.Form, req.MultipartForm.File)

	} else if strings.Contains(contentType, "application/json") {
		err := json.NewDecoder(req.Body).Decode(obj)
		if err != nil {
			return err
		}
		return nil
	} else {
		if contentType == "" {
			return errors.New("empty Content-Type")
		} else {
			return errors.New("Unsupported Content-Type")
		}

	}

}

func BindForm(form interface{}, mapVal map[string][]string, fileMap map[string][]*multipart.FileHeader) (err error) {

	formVal := reflect.ValueOf(form)

	typ := reflect.TypeOf(form)

	if formVal.Kind() != reflect.Ptr || formVal.IsNil() {
		err = NewInvalidBindError(typ)
		return
	}

	typ = typ.Elem()
	formVal = formVal.Elem()

	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldVal := formVal.Field(i)

		if tagName := fieldType.Tag.Get(formKey); tagName != "" {
			if !fieldVal.CanSet() {
				continue
			}

			var tagVal []string
			tagVal, ok := mapVal[tagName]
			if !ok {
				continue
			}

			lenOfTagVal := len(tagVal)
			if lenOfTagVal == 0 {
				fieldVal.Set(reflect.Zero(fieldType.Type))
				continue
			}

			if fieldType.Type.Kind() == reflect.Slice {
				tmpSlice := reflect.MakeSlice(fieldType.Type, lenOfTagVal, lenOfTagVal)
				for j := 0; j < lenOfTagVal; j++ {
					tempVal, err := reflectutils.ParsePrimitive(fieldType.Type.Elem(), tagVal[j])
					if err != nil {
						return err
					}
					tmpSlice.Index(j).Set(reflect.ValueOf(tempVal))
				}
				formVal.Field(i).Set(tmpSlice)
			} else {
				tempVal, err := reflectutils.ParsePrimitive(fieldType.Type, tagVal[0])
				if err != nil {
					return err
				}
				fieldVal.Set(reflect.ValueOf(tempVal))
			}
		}
	}
	return
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
