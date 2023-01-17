package entity

import "database/sql"

type DeletedAt struct {
	sql.NullTime
}
