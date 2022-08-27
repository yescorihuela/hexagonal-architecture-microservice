package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/agrak/infrastructure/postgresql/connection"
	"github.com/yescorihuela/agrak/infrastructure/postgresql/product"
	"github.com/yescorihuela/agrak/usecase"
)

type Server struct {
	engine   *gin.Engine
	dbClient *connection.PostgresqlConnection
	httpAddr string
}

func NewServer(host string, port uint) *Server {
	dbClient := connection.InitPGClient()
	connection.AutoMigrateEntities(dbClient)
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

	v1 := s.engine.Group("api/v1")
	v1.GET("/products/", ph.GetAllProducts)
	v1.GET("/products/:sku", ph.GetProductBySku)
	v1.POST("/products", ph.CreateProduct)
	v1.PUT("/products/:sku", ph.UpdateProduct)
	v1.DELETE("/products/:sku", ph.Delete)
}
