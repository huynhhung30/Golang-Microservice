package main

import (
	"net/http"
	"os"
	"time"
	"user_service/config"
	"user_service/controllers"
	errorsController "user_service/controllers/errors"
	"user_service/middlewares"
	"user_service/models"
	"user_service/routes"
	"user_service/utils/functions"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	initialGinConfig(router)
	functions.ShowLog("router", router)
	router.Use(middlewares.GinBodyLogMiddleware)
	router.Use(errorsController.Handler)
	routes.ApplicationV1Router(router)
	controllers.Migrate()
	go models.StartRpcServer()
	startServer(router)
}

func initialGinConfig(router *gin.Engine) {
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, authorization  "},
	}))

	var err error
	config.DB, err = config.GormOpen()

	if err != nil {
		functions.ShowLog("Connect database error", err.Error())
	}
}

func startServer(router http.Handler) {
	serverPort := os.Getenv("PORT")
	addr := ":" + serverPort
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		functions.ShowLog("Start server error", err.Error())
	} else {
		functions.ShowLog("Start server success", s)
	}
}
