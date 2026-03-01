package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lbodlev888/url_shortener/services"
)

func NewShortUrl(c *gin.Context) {
	raw_token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	if !services.ValidateToken(raw_token) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	type reqBody struct {
		Url string `json:"url"`
	}
	var req reqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	err = services.NewUrl(raw_token, req.Url, c.ClientIP())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetLongUrl(c *gin.Context) {
	url := c.Param("url")

	link, err := services.GetLongUrl(url)
	if err != nil {
		c.HTML(404, "wrong.html", gin.H{"title": "not found"})
		return
	}
	services.Increment(url)

	c.Redirect(http.StatusTemporaryRedirect, link)
}

func DeleteShort(c *gin.Context) {
	raw_token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	if !services.ValidateToken(raw_token) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	url := c.Param("url")

	err = services.DeleteShort(raw_token, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
