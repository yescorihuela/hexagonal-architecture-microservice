package database

import "gorm.io/gorm"

type GenericDatabaseRepository interface {
	GetConnection() (*gorm.DB, error)
}
