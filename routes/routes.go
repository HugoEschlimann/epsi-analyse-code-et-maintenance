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
		api.POST("/users", func(c *gin.Context) {
			controllers.CreateUser(c, db, &models.User{})
		})
		api.PUT("/users/:id", func(c *gin.Context) {
			controllers.UpdateUser(c, db, &models.User{})
		})
		api.DELETE("/users/:id", func(c *gin.Context) {
			controllers.DeleteUser(c, db)
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
			controllers.DeleteResource(c, db)
		})
	}
	{
		api.POST("/loans", func(c *gin.Context) {
			controllers.LoanResources(c, db, []*models.Loan{})
		})
		api.GET("/loans", func(c *gin.Context) {
			controllers.GetLoans(c, db)
		})
		api.POST("/restitute", func(c *gin.Context) {
			controllers.Restitute(c, db, []*models.Loan{})
		})
	}

	return router
}
