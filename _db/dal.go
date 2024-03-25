package db_dal

import (
	"log"
	inital "server/_init"
	t "server/_sys/type"
)

func SelectById[T comparable](orderId string, object *T) {
	if err := inital.O_DB_PG.Where("id = ?", orderId).First(&object).Error; err != nil {
		log.Println("failed to query o", orderId, err)
	}
}
func SelectAllByField[T comparable](objects *[]T, orderId string, field string) {
	if field == "" {
		field = "id"
	}
	if err := inital.O_DB_PG.Where(field+" = ?", orderId).Scan(&objects).Error; err != nil {
		log.Println("failed to query", orderId, err)
	}
}

func Raw[T comparable](objects *[]T, raw string, variables ...interface{}) {
	// db.Raw("SELECT * FROM products WHERE price > ? ORDER BY price DESC LIMIT ? OFFSET ?", 100, 10, 20).Scan(&products)
	inital.O_DB_PG.Raw(raw, variables...).Scan(&objects)
}

func Query[T comparable](object *[]T, strSelect string, where t.M, strOrder string, limit int, offset int) {
	resultOrm := inital.O_DB_PG.Model(nil)
	if strSelect != "" {
		resultOrm.Select(strSelect)
	}
	if where != nil {
		if err := resultOrm.Where(where["condition"], where["value"]).Error; err != nil {
			log.Println("failed to query ", where, err)
		}
	}
	if strOrder != "" {
		resultOrm.Order(strOrder)
	}
	if offset == 0 {
		offset = -1
	}
	resultOrm.Offset(offset)
	if limit == 0 {
		limit = 1
	}
	resultOrm.Limit(limit)

	resultOrm.Find(&object)
}

func QueryAll[T comparable](object *[]T, strSelect string, where t.M, strOrder string) {
	resultOrm := inital.O_DB_PG.Model(nil)
	if strSelect != "" {
		resultOrm.Select(strSelect)
	}
	if where != nil {
		if err := resultOrm.Where(where["condition"], where["value"]).Error; err != nil {
			log.Println("failed to query ", where, err)
		}
	}
	if strOrder != "" {
		resultOrm.Order(strOrder)
	}

	resultOrm.Find(&object)
}
