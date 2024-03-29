package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mcsans/finalProject3-kel2/controllers"
	"github.com/mcsans/finalProject3-kel2/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/update-account/:userId", controllers.UserUpdate)
		userRouter.DELETE("/delete-account/:userId", controllers.UserDelete)
	}

	categoryRouter := r.Group("/categories")
	{
		categoryRouter.Use(middlewares.Authentication())
		categoryRouter.GET("/", controllers.GetCategory)
		categoryRouter.GET("/:categoryId", controllers.GetCategoryById)
		categoryRouter.POST("/", controllers.CreateCategory)
		categoryRouter.PATCH("/:categoryId", middlewares.CategoryAuthorization(), controllers.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", controllers.DeleteCategory)
	}

	taskRouter := r.Group("/tasks")
	{
		taskRouter.Use(middlewares.Authentication())
		taskRouter.GET("/", controllers.GetTask)
		taskRouter.GET("/:taskId", controllers.GetTaskById)
		taskRouter.POST("/", controllers.CreateTask)
		taskRouter.PUT("/:taskId", middlewares.TaskAuthorization(), controllers.UpdateTask)
		taskRouter.PATCH("/update-status/:taskId", middlewares.TaskAuthorization(), controllers.UpdateTaskStatus)
		taskRouter.PATCH("/update-category/:taskId", middlewares.TaskAuthorization(), controllers.UpdateTaskCategory)
		taskRouter.DELETE("/:taskId", controllers.DeleteTask)
	}

	return r
}