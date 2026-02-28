package main

import (
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lbodlev888/url_shortener/routes"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.SetFuncMap(template.FuncMap{
        "formatDate": func(t time.Time) string {
            return t.Format("15:04:05 02/01/2006")
        },
    })

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
