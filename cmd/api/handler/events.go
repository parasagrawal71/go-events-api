package apiHandler

import (
	appContext "go-events-api/cmd/api/context"
	"go-events-api/cmd/api/dto"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// The func names must be capital to export them

// TODO: Cleanup later
var events []dto.Event = []dto.Event{
	{
		ID:          1,
		Name:        "Event 1",
		Description: "Description 1",
		Date:        "2023-01-01",
		Location:    "Location 1",
		OwnerID:     1,
	},
	{
		ID:          2,
		Name:        "Event 2",
		Description: "Description 2",
		Date:        "2023-02-01",
		Location:    "Location 2",
		OwnerID:     1,
	},
}

// createEvent creates a new event
//
// @Summary Creates a new event
// @Description Creates a new event
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param event body dto.Event true "Event details"
// @Success 201 {object} dto.Event
// @Router /api/v1/events [post]
func CreateEvent(c *gin.Context) {
	var newEvent dto.Event

	if err := c.BindJSON(&newEvent); err != nil {
		if err == io.EOF {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	newEvent.ID = len(events) + 1
	user := appContext.GetUserFromContext(c)
	newEvent.OwnerID = user.ID

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

// getAllEvents returns all events
//
// @Summary Returns all events
// @Description Returns all events
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Event
// @Router /api/v1/events [get]
func GetAllEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

// getEvent returns a single event
//
// @Summary Returns a single event by ID
// @Description Returns a single event by ID
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} dto.Event
// @Router /api/v1/events/{id} [get]
func GetEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	for _, event := range events {
		if event.ID == id {
			c.IndentedJSON(http.StatusOK, event)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}

// updateEvent updates a single event
//
// @Summary Updates a single event by ID
// @Description Updates a single event by ID
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body dto.Event true "Event details"
// @Success 200 {object} dto.Event
// @Router /api/v1/events/{id} [put]
func UpdateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	var updatedEvent dto.Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := appContext.GetUserFromContext(c)

	for i, event := range events {
		if event.ID == id {
			if event.OwnerID != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}

			updatedEvent.ID = id
			events[i] = updatedEvent
			c.IndentedJSON(http.StatusOK, updatedEvent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}

// deleteEvent deletes a single event
//
// @Summary Deletes a single event by ID
// @Description Deletes a single event by ID
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} gin.H
// @Router /api/v1/events/{id} [delete]
func DeleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	user := appContext.GetUserFromContext(c)

	for i, event := range events {
		if event.ID == id {
			if event.OwnerID != user.ID {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}

			events = append(events[:i], events[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "event deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
}
