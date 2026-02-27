package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register"})
}
