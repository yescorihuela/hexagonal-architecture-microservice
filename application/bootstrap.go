package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/infrastructure/postgresql"
	"gorm.io/gorm"
)

type Server struct {
	engine   *gin.Engine
	database *gorm.DB
	httpAddr string
}

func NewServer(host string, port uint) *Server {
	dbConnection := postgresql.InitPGClient()
	db, _ := dbConnection.GetConnection()
	server := &Server{
		engine:   gin.Default(),
		database: db,
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	return server
}

func (s *Server) Run() error {
	return s.engine.Run(s.httpAddr)
}
