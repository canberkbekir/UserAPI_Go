package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
}

func createServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() {
	router := gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "test")
	})
	router.Run(os.Getenv("PORT"))
}
