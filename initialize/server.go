package initialize

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paiz4ever/go-web-framework/api"
	"github.com/paiz4ever/go-web-framework/global"
)

func InitServer() {
	r := gin.Default()
	api.InitRoutes(r)
	port := global.CONFIG.Server.Port
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	log.Printf("starting server on port %d...\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
