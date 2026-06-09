package validators

import (
	"errors"
	"fmt"
	"gohub/pkg/database"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

func init() {
	govalidator.AddCustomRule("not_exists", func(field, rule, msg string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]
		dbField := rng[1]

		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		reqValue := value.(string)

		query := database.DB.Table(tableName).Where(dbField+" = ?", reqValue)

		if len(exceptID) > 0 {
			query = query.Where("except_id = ?", exceptID)
		}

		var count int64
		query.Count(&count)

		if count != 0 {
			if len(msg) > 0 {
				return errors.New(msg)
			}
			return fmt.Errorf("%v already exists", reqValue)
		}

		return nil
	})
}
