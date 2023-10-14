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

	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks" : tasks})
}

func GetTaskById(c *gin.Context) {
	db := database.GetDB()
	var task models.Task
	taskId := c.Param("taskId")

	if err := db.First(&task, taskId).Error; err != nil {
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

	if Task.Title == "" || Task.Description == "" || Task.CategoryID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "title, description, category_id value required"})
		return
	}

	Task.UserID = userID
	var category models.Category

	if err := db.First(&category, Task.CategoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
		return
	} else {
		Task.Status = false
	}

	err := db.Debug().Create(&Task).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	TaskResponse := models.TaskResponse{
		ID: Task.ID,
		Title: Task.Title,
		Status: Task.Status,
		Description: Task.Description,
		UserID: Task.UserID,
		CategoryID: Task.CategoryID,
		CreatedAt: Task.CreatedAt,
		UpdatedAt: Task.UpdatedAt,
	}

	c.JSON(http.StatusCreated, TaskResponse)
}

func UpdateTask(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	if Task.Title == "" || Task.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "title, description value required"})
		return
	}

	Task.UserID = userID
	Task.ID = uint(taskId)

	err := db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{Title: Task.Title, Description: Task.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	TaskResponse := models.TaskResponse{
		ID: Task.ID,
		Title: Task.Title,
		Status: Task.Status,
		Description: Task.Description,
		UserID: Task.UserID,
		CategoryID: Task.CategoryID,
		CreatedAt: Task.CreatedAt,
		UpdatedAt: Task.UpdatedAt,
	}

	c.JSON(http.StatusOK, TaskResponse)
}

func UpdateTaskStatus(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))
	db.First(&Task, taskId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.UserID = userID
	Task.ID = uint(taskId)

	err := db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{Status: Task.Status}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	TaskResponse := models.TaskResponse{
		ID: Task.ID,
		Title: Task.Title,
		Status: Task.Status,
		Description: Task.Description,
		UserID: Task.UserID,
		CategoryID: Task.CategoryID,
		CreatedAt: Task.CreatedAt,
		UpdatedAt: Task.UpdatedAt,
	}

	c.JSON(http.StatusOK, TaskResponse)
}

func UpdateTaskCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Task := models.Task{}

	taskId, _ := strconv.Atoi(c.Param("taskId"))
	userID := uint(userData["id"].(float64))
	db.First(&Task, taskId)
	Task.CategoryID = 0

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	
	if Task.CategoryID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "category_id value required"})
		return
	}

	Task.UserID = userID
	Task.ID = uint(taskId)
	var category models.Category

	if err := db.First(&category, Task.CategoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
		return
	}

	err := db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{CategoryID: Task.CategoryID}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	TaskResponse := models.TaskResponse{
		ID: Task.ID,
		Title: Task.Title,
		Status: Task.Status,
		Description: Task.Description,
		UserID: Task.UserID,
		CategoryID: Task.CategoryID,
		CreatedAt: Task.CreatedAt,
		UpdatedAt: Task.UpdatedAt,
	}

	c.JSON(http.StatusOK, TaskResponse)
}

func DeleteTask(c *gin.Context) {
	db := database.GetDB()
	taskId := c.Param("taskId")
	var task models.Task

	if db.Delete(&task, taskId).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Task has been successfully deleted"})
}