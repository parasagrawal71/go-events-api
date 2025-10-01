package apiHandler

// The func names must be capital to export them
import (
	appContext "go-events-api/cmd/api/context"
	"go-events-api/cmd/api/dto"
	"go-events-api/cmd/api/models"
	"go-events-api/cmd/api/repository"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// createEvent creates a new event
//
// @Summary Creates a new event
// @Description Creates a new event
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param event body dto.CreateEvent true "Event details"
// @Success 201 {object} dto.CreateEvent
// @Router /api/v1/events [post]
func CreateEvent(c *gin.Context) {
	var newEvent dto.CreateEvent
	if err := c.BindJSON(&newEvent); err != nil {
		if err == io.EOF {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	user := appContext.GetUserFromContext(c)

	eventToCreate := models.Event{
		Name:        newEvent.Name,
		Description: newEvent.Description,
		Date:        newEvent.Date,
		Location:    newEvent.Location,
		OwnerID:     user.ID,
	}

	createdEvent, err := repository.EventRepo.Create(&eventToCreate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdEvent)
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
	// // -- 1. Other way to get all events
	// var events []models.Event
	// if err := config.DB.Find(&events).Error; err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // -- 2. Other way to get all events
	// var events []models.Event
	// config.DB.Raw("SELECT * FROM events;").Scan(&events)

	// Using EventRepo
	events, err := repository.EventRepo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	id := uint(idParam)

	event, err := repository.EventRepo.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, event)
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
// @Param event body dto.UpdateEvent true "Event details"
// @Success 200 {object} dto.Event
// @Router /api/v1/events/{id} [put]
func UpdateEvent(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	id := uint(idParam)

	var eventPayload dto.UpdateEvent
	if err := c.BindJSON(&eventPayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := appContext.GetUserFromContext(c)

	event, err := repository.EventRepo.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event.OwnerID != user.ID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventToUpdate := models.Event{
		Name:        eventPayload.Name,
		Description: eventPayload.Description,
		Date:        eventPayload.Date,
		Location:    eventPayload.Location,
	}

	updatedEvent, err := repository.EventRepo.Update(id, &eventToUpdate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEvent)
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
// @Success 200 {object} dto.Event
// @Router /api/v1/events/{id} [delete]
func DeleteEvent(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	id := uint(idParam)
	user := appContext.GetUserFromContext(c)

	event, err := repository.EventRepo.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event.OwnerID != user.ID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	deletedEvent, err := repository.EventRepo.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedEvent)
}
