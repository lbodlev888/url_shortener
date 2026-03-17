package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lbodlev888/url_shortener/services"
)

func IndexPage(c *gin.Context) {
	raw_token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	token, status := services.ValidateToken(raw_token)
	if !status {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	shorts, err := services.GetAllShorts(token)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home", "shorts": shorts})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register"})
}
