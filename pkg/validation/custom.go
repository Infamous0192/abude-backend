package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Validates if a field value exists in a specified database table and column.
func exist(db *gorm.DB) func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ".")
		if len(params) == 0 {
			return false
		}

		table, column := params[0], "id"
		if len(params) > 1 {
			column = params[1]
		}

		var count int64
		db.Table(table).
			Where(fmt.Sprintf("`%s` = ?", column), fl.Field().Interface()).
			Count(&count)

		return count > 0
	}
}

// Validates if a field value is unique in a specified database table and column.
func unique(db *gorm.DB) func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ".")
		if len(params) == 0 {
			return false
		}

		table, column := params[0], "id"
		if len(params) > 1 {
			column = params[1]
		}

		var count int64
		db.Table(table).
			Where(fmt.Sprintf("`%s` = ?", column), fl.Field().Interface()).
			Count(&count)

		return count == 0
	}
}
