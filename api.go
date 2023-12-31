package main

import (
	"net/http"

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

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "mainpage")
	})
	router.Run(":" + s.listenAddr)
}
