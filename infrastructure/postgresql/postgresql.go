package postgresql

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

type PostgresqlRepository interface {
	GetConnection() (*gorm.DB, error)
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
		log.WithError(err).Errorln("error to trying to open connection in DB")
	} else {
		_, err := connection.DB()
		if err != nil {
			log.WithError(err).Errorln("error to trying to connect DB")
		}
	}
	return connection, nil
}

func InitPGClient() *PostgresqlConnection {
	databaseName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connection := NewPostgreSQLConnection(Config().Server(host).Port(port).DatabaseName(databaseName).User(user).Password(password))
	return connection
}
