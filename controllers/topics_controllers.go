package controllers

import (
	"net/http"
	"study_notes/config"
	"study_notes/models"

	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {
	var topics []models.Topic
	config.DB.Find(&topics)
	c.JSON(http.StatusOK, gin.H{"data": topics})
}

func CreateTopics(c *gin.Context) {
	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&topic)
	c.JSON(http.StatusCreated, gin.H{"data": topic})
}

func GetTopicsByID(c *gin.Context) {
	var topic models.Topic
	id := c.Param("id")
	if err := config.DB.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Topic tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": topic})
}

func UpdateTopics(c *gin.Context) {
	var topic models.Topic
	id := c.Param("id")

	if err := config.DB.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Topic tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&topic)
	c.JSON(http.StatusOK, gin.H{"message": "Topic berhasil diperbarui", "data": topic})
}

func DeleteTopics(c *gin.Context) {
	var topic models.Topic
	id := c.Param("id")

	if err := config.DB.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Topic tidak ditemukan"})
		return
	}

	config.DB.Delete(&topic)
	c.JSON(http.StatusOK, gin.H{"message": "Topic berhasil dihapus"})
}