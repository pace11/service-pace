package utils

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(c *gin.Context, form any) error {
	if err := c.ShouldBindJSON(form); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)

			// Ambil nama json dari tag, bukan Field Go
			refType := reflect.TypeOf(form).Elem()
			for _, fe := range ve {
				jsonField := fe.Field()
				if f, ok := refType.FieldByName(fe.Field()); ok {
					jsonTag := f.Tag.Get("json")
					parts := strings.Split(jsonTag, ",")
					if len(parts) > 0 && parts[0] != "" {
						jsonField = parts[0]
					}
				}
				out[jsonField] = fmt.Sprintf("%s is %s", jsonField, fe.Tag())
			}

			HttpResponse(c, nil, http.StatusBadRequest, "data", c.Request.Method, out)
			return err
		}
		HttpResponse(c, nil, http.StatusBadRequest, "data", c.Request.Method, nil)
		return err
	}
	return nil
}
