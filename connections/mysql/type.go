package mysql

import (
	"database/sql"

	"gorm.io/gorm"
)

type DB struct {
	Client     *sql.DB
	GORMClient *gorm.DB
}
