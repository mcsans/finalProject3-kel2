package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mcsans/finalProject3-kel2/database"
	"github.com/mcsans/finalProject3-kel2/helpers"
	"github.com/mcsans/finalProject3-kel2/models"
	"gorm.io/gorm"
)

func GetTask(c *gin.Context) {
	db := database.GetDB()
	var tasks []models.Task

	db.Preload("Tasks").Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks" : tasks})
}

func GetTaskById(c *gin.Context) {
	db := database.GetDB()
	var task models.Task
	id := c.Param("taskId")

	if err := db.First(&task, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
		}
	}

	c.JSON(http.StatusOK, gin.H{"task" : task})
}

func CreateTask(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Task := models.Task{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID

	err := db.Debug().Create(&Task).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Task)
}

func UpdateTask(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID
	Task.ID = uint(productId)

	err := db.Model(&Task).Where("id = ?", productId).Updates(models.Task{Title: Task.Title, Description: Task.Description, Status: Task.Status}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Task)
}

func DeleteTask(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var task models.Task

	// if err := db.Where("category_id = ?", id).Delete(models.Task{}).Error; err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus task"})
	// 	return
	// }

	if db.Delete(&task, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Data berhasil dihapus"})
}