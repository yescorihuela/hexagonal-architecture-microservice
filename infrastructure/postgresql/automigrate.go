package postgresql

import (
	log "github.com/sirupsen/logrus"
	"github.com/yescorihuela/agrak/infrastructure/product/model"
)

type migrate struct {
	connection *PostgresqlConnection
}

func NewMigrate(conn *PostgresqlConnection) *migrate {
	return &migrate{
		connection: conn,
	}
}

func (m *migrate) AutoMigrateAll(tables ...interface{}) {
	db, _ := m.connection.GetConnection()
	err := db.AutoMigrate(tables...)
	if err != nil {
		log.WithError(err).Errorln("error to try to migrate tables on DB...")
	}
}

func AutoMigrateEntities(connection *PostgresqlConnection) {
	migrate := NewMigrate(connection)
	migrate.AutoMigrateAll(
		model.ProductModel{},
	)
}
