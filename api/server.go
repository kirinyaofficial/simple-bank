package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kirinyaofficial/simple-bank/db/sqlc"
)

// Serves HTTP requests for our banking services
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewSerevr(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
