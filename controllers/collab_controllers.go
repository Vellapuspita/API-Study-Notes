package controllers

import (
	"study_notes/config"
	"study_notes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCollab(c *gin.Context) {
	var collab []models.Collab
	config.DB.Find(&collab)
	c.JSON(http.StatusOK, gin.H{"data": collab})
}

func CreateCollab(c *gin.Context) {
	var collab models.Collab
	if err := c.ShouldBindJSON(&collab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&collab)
	c.JSON(http.StatusCreated, gin.H{"data": collab})
}

func GetCollabByID(c *gin.Context) {
	var collab models.Collab
	id := c.Param("id")
	if err := config.DB.First(&collab, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collab tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": collab})
}

func UpdateCollab(c *gin.Context) {
	var collab models.Collab
	id := c.Param("id")

	if err := config.DB.First(&collab, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collab tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&collab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&collab)
	c.JSON(http.StatusOK, gin.H{"message": "collab berhasil diperbarui", "data": collab})
}

func DeleteCollab(c *gin.Context) {
	var collab models.Collab
	id := c.Param("id")

	if err := config.DB.First(&collab, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collab tidak ditemukan"})
		return
	}

	config.DB.Delete(&collab)
	c.JSON(http.StatusOK, gin.H{"message": "collab berhasil dihapus"})
}