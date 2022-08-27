package connection

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

type PostgresqlConnection struct {
	options *PostgresqlOptions
	url     string
}

func NewPostgreSQLConnection(opts ...*PostgresqlOptions) *PostgresqlConnection {
	databaseOptions := MergeOptions(opts...)
	url := databaseOptions.GetURLConnection()
	return &PostgresqlConnection{
		options: databaseOptions,
		url:     url,
	}
}

func (p *PostgresqlConnection) GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil {
		connection, err = gorm.Open(postgres.Open(p.url), &gorm.Config{})
	}
	if err != nil {
		log.WithError(err).Fatalf("error to trying to open connection in DB")
	} else {
		_, err := connection.DB()
		if err != nil {
			log.WithError(err).Fatalf("error to trying to connect DB")
		}
	}
	return connection, nil
}

func InitPGClient() *PostgresqlConnection {
	databaseName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.Atoi(os.Getenv("PGPORT"))
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	connection := NewPostgreSQLConnection(Config().Server(host).Port(port).DatabaseName(databaseName).User(user).Password(password))
	return connection
}
