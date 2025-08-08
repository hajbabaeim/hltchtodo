package app

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"time"
)

func (a *App) initValidator() {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := v.RegisterValidation("due_date_validator", a.ValidateDueDate); err != nil {
		a.panicOnError(err)
	}
	a.validator = v
}

func (a *App) ValidateDueDate(fl validator.FieldLevel) bool {
	dueDate := fl.Field().Interface().(time.Time)
	now := time.Now()
	if dueDate.Before(now) {
		return false
	}
	return true
}
