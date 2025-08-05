package main

import (
	apiHandler "go-events-api/cmd/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{ // Bracket is not necessary, it's just for grouping.
		v1.GET("/events", apiHandler.GetAllEvents)
		v1.POST("/events", apiHandler.CreateEvent)
		v1.GET("/events/:id", apiHandler.GetEvent)
		v1.PUT("/events/:id", apiHandler.UpdateEvent)
		v1.DELETE("/events/:id", apiHandler.DeleteEvent)
	}

	return g
}
