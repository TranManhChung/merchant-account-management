package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func NewBunDB(format, username, password, address, database, driverName string) (*bun.DB, error) {
	dataSource := fmt.Sprintf(format, username, password, address, database)
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	bunDB := bun.NewDB(db, sqlitedialect.New())

	return bunDB, nil
}
