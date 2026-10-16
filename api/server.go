package api

import (
	db "github.com/divin3circle/sba/db/sqlc"
	"github.com/gin-gonic/gin"
	`github.com/gin-gonic/gin/binding`
	`github.com/go-playground/validator/v10`
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)

	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
