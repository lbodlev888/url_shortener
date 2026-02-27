package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lbodlev888/url_shortener/routes"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
