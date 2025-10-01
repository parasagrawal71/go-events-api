package main

import (
	apiHandler "go-events-api/cmd/api/handler"
	"go-events-api/cmd/api/middleware"
	"go-events-api/cmd/api/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1OpenRoutes := g.Group("/api/v1")
	{ // Bracket is not necessary, it's just for grouping.
		v1OpenRoutes.POST("/auth/register", apiHandler.RegisterUser)
		v1OpenRoutes.POST("/auth/login", apiHandler.LoginUser)
	}

	v1PrivateRoutes := v1OpenRoutes.Group("/") // IMPORTANT: Grouping on top of v1OpenRoutes
	v1PrivateRoutes.Use(middleware.AuthMiddleware())
	{
		v1PrivateRoutes.GET("/events", apiHandler.GetAllEvents)
		v1PrivateRoutes.POST("/events", validator.ValidateCreateEvent(), apiHandler.CreateEvent)
		v1PrivateRoutes.GET("/events/:id", apiHandler.GetEvent)
		v1PrivateRoutes.PUT("/events/:id", apiHandler.UpdateEvent)
		v1PrivateRoutes.DELETE("/events/:id", apiHandler.DeleteEvent)
	}

	g.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(http.StatusFound, "/swagger/index.html")
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json"))(c)
	})

	return g
}
