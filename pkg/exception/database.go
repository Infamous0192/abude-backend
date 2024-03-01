package exception

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func DB(err error, field ...string) error {
	if err == nil {
		return nil
	}

	_, ok := err.(HttpError)
	if ok {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if len(field) > 0 {
			return NotFound(field[0])
		}

		return NotFound("Record")
	}

	return Http(500, fmt.Sprintf("unexpected database error: %s", err.Error()))
}

func CatchDB(err error, field ...string) {
	if err == nil {
		return
	}

	panic(DB(err))
}
