package main

import (
	"fmt"

	"github.com/emilianooferreyra/MagicStreamMoviesServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupUnProtectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Fail to start server", err)
	}
}
