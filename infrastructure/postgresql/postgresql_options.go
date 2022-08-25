package postgresql

import (
	"fmt"
	"net/url"
)

type SQLDialect string

var (
	PostgresqlDefaultPort            = 5432
	Posgres               SQLDialect = "postgres"
)

type PostgresqlOptions struct {
	databaseName *string
	server       *string
	port         *int
	user         *string
	password     *string
}

func Config() *PostgresqlOptions {
	return &PostgresqlOptions{}
}

func (p *PostgresqlOptions) DatabaseName(name string) *PostgresqlOptions {
	p.databaseName = &name
	return p
}

func (p *PostgresqlOptions) Server(server string) *PostgresqlOptions {
	p.server = &server
	return p
}

func (p *PostgresqlOptions) Port(port int) *PostgresqlOptions {
	p.port = &port
	return p
}

func (p *PostgresqlOptions) User(user string) *PostgresqlOptions {
	p.user = &user
	return p
}

func (p *PostgresqlOptions) Password(password string) *PostgresqlOptions {
	p.password = &password
	return p
}

func MergeOptions(opts ...*PostgresqlOptions) *PostgresqlOptions {
	option := new(PostgresqlOptions)
	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.server != nil {
			option.server = opt.server
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

func (a *PostgresqlOptions) GetURLConnection() string {
	if a.port == nil {
		a.port = &PostgresqlDefaultPort
	}
	query := url.Values{}
	query.Add("sslmode", "disable")
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(*a.user, *a.password),
		Host:     fmt.Sprintf("%s:%d", *a.server, *a.port),
		Path:     *a.databaseName,
		RawQuery: query.Encode(),
	}

	return u.String()
}
