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

	func GetCategory(c *gin.Context) {
		db := database.GetDB()
		var categories []models.Category

		db.Preload("Tasks").Find(&categories)
		c.JSON(http.StatusOK, gin.H{"categories" : categories})
	}

func GetCategoryById(c *gin.Context) {
	db := database.GetDB()
	var category models.Category
	id := c.Param("categoryId")

	if err := db.Preload("Tasks").First(&category, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
		}
	}

	c.JSON(http.StatusOK, gin.H{"category" : category})
}

func CreateCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Category := models.Category{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}

	Category.UserID = userID

	err := db.Debug().Create(&Category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CategoryResponse := models.CategoryResponse{
		ID: Category.ID,
		Type: Category.Type,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
	}

	c.JSON(http.StatusCreated, CategoryResponse)
}

func UpdateCategory(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Category := models.Category{}

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}

	Category.UserID = userID
	Category.ID = uint(categoryId)

	err := db.Model(&Category).Where("id = ?", categoryId).Updates(models.Category{Type: Category.Type}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CategoryResponse := models.CategoryResponse{
		ID: Category.ID,
		Type: Category.Type,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
	}

	c.JSON(http.StatusOK, CategoryResponse)
}

func DeleteCategory(c *gin.Context) {
	db := database.GetDB()
	categoryId := c.Param("categoryId")
	var category models.Category

	if err := db.Where("category_id = ?", categoryId).Delete(models.Task{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}

	if db.Delete(&category, categoryId).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Category has been successfully deleted"})
}