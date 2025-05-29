package routes

import (
	"encoding/json"
	"gin/controllers"
	"gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(c.Writer)
		encoder.SetEscapeHTML(false)
		encoder.Encode(gin.H{
			"message": "EPITECH > EPSI",
		})
	})
	api := router.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			controllers.GetUsers(c, db)
		})
		api.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			controllers.GetUserById(c, db, id)
		})
		api.POST("/users", func(c *gin.Context) {
			controllers.CreateUser(c, db, &models.User{})
		})
		api.PUT("/users", func(c *gin.Context) {
			controllers.UpdateUser(c, db, &models.User{})
		})
		api.DELETE("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			controllers.DeleteUser(c, db, id)
		})
	}
	{
		api.POST("/resources", func(c *gin.Context) {
			controllers.CreateResource(c, db, &models.Resource{})
		})
		api.PUT("/resources/:id", func(c *gin.Context) {
			controllers.UpdateResource(c, db, &models.Resource{})
		})
		api.GET("/resources", func(c *gin.Context) {
			controllers.GetResources(c, db)
		})
		api.DELETE("/resources/:id", func(c *gin.Context) {
			id := c.Param("id")
			controllers.DeleteResource(c, db, id)
		})
	}
	{
		api.POST("/loans", func(c *gin.Context) {
			controllers.LoanResources(c, db, []*models.Loan{})
		})
		// api.POST("/restitute", func(c *gin.Context) {
		// 	controllres.RestituteResources(c, db, []*models.Loan{})
		// })
	}

	return router
}
