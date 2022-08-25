package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/infrastructure/postgresql"
	"github.com/yescorihuela/agrak/infrastructure/product"
	"github.com/yescorihuela/agrak/usecase"
)

type Server struct {
	engine   *gin.Engine
	dbClient *postgresql.PostgresqlConnection
	httpAddr string
}

func NewServer(host string, port uint) *Server {
	dbClient := postgresql.InitPGClient()
	postgresql.AutoMigrateEntities(dbClient)
	server := &Server{
		engine:   gin.Default(),
		dbClient: dbClient,
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	productRepository := product.NewPersistenceProductRepository(s.dbClient)
	productService := usecase.NewProductService(productRepository)

	ph := NewProductHandlers(productService)

	v1 := s.engine.Group("/v1")
	v1.POST("/products", ph.CreateProduct)
	v1.GET("/products/:sku", ph.GetProductBySku)
}
