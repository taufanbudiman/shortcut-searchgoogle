package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"shorturl/models"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindUrl(c *gin.Context) {
	var dataUrl []models.DataUrl
	models.DB.Find(&dataUrl)

	c.JSON(http.StatusOK, gin.H{"data": dataUrl})
}

func GetAlias(c *gin.Context) {
	var dataUrl models.DataUrl

	if err := models.DB.Where("alias = ?", c.Param("alias")).First(&dataUrl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alias Record not found!"})
		return
	}
	q := url.Values{}
	q.Set("q", dataUrl.Query)
	location := url.URL{Path: "http://www.google.com/search", RawQuery: q.Encode()}

	c.Redirect(http.StatusFound, location.RequestURI())
}

func CreateUrl(c *gin.Context) {
	// Validate input
	var input models.CreateDataUrl
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	dataurl := models.DataUrl{
		Title:    input.Title,
		Alias:    input.Alias,
		ShortUrl: fmt.Sprint("http://localhost:8090/", input.Alias),
		Query:    input.Query,
	}
	models.DB.Create(&dataurl)

	c.JSON(http.StatusCreated, gin.H{"data": dataurl})
}

func UpdateUrl(c *gin.Context) {
	// Get model if exist
	var dataUrl models.DataUrl

	if err := models.DB.Where("id = ?", c.Param("id")).First(&dataUrl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateDataUrl
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fields = struct {
		Title    string
		ShortUrl string
		Alias    string
		Query    string
	}{input.Title, fmt.Sprint("http://localhost:8090/", input.Alias), input.Alias, input.Query}

	models.DB.Model(&dataUrl).Updates(fields)

	c.JSON(http.StatusOK, gin.H{"data": dataUrl})
}

func DeleteUrl(c *gin.Context) {
	// Get model if exist
	var dataUrl models.DataUrl
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dataUrl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&dataUrl)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
