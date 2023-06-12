package postgresql

import (
	"vm-backend/pkg/helpers/db"

	"gorm.io/gorm"
)

func MakeQuery(db *gorm.DB, query *db.Query) *gorm.DB {
	if query == nil {
		return db
	}

	if query.LimitVal != nil && *query.LimitVal > 0 {
		db = db.Limit(*query.LimitVal)
	}

	if query.OffsetVal != nil && *query.OffsetVal > 0 {
		db = db.Offset(*query.OffsetVal)
	}

	if len(query.WhereVal) > 0 {
		for _, where := range query.WhereVal {
			db = db.Where(where.Query, where.Args...)
		}
	}

	if query.OrderVal != nil {
		if query.OrderVal.Decending {
			db = db.Order(query.OrderVal.Field + " DESC")
		} else {
			db = db.Order(query.OrderVal.Field)
		}
	}

	if len(query.PerloadVal) > 0 {
		for _, perload := range query.PerloadVal {
			db = db.Preload(perload.Query, perload.Args...)
		}
	}

	return db
}
