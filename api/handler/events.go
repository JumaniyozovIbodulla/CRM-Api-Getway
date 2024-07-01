package handler

import (
	"crmapi/genproto/user_service/events"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateEvent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/event [POST]
// @Summary 	Create an event
// @Description API for creating event
// @Tags 		EVENTS
// @Accept  	json
// @Produce  	json
// @Param		event body user_service.CreateEvent true "event"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateEvent(c *gin.Context) {
	event := user_service.CreateEvent{}

	if err := c.ShouldBindJSON(&event); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating an event")
		return
	}

	_, err := h.grpcClient.Events().Create(c.Request.Context(), &event)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating an event")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Event created successfully")
}

// GetByIdEvent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/event/{id} [GET]
// @Summary 	Get an event
// @Description API for getting an event
// @Tags 		EVENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdEvent(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating admin", http.StatusBadRequest, err.Error())
		return
	}

	event := user_service.EventPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Events().GetById(c.Request.Context(), &event)

	if err != nil {
		handleResponse(c, h.log, "error while getting event", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Event got successfully", http.StatusOK, resp)
}

// UpdateEvent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/event [PUT]
// @Summary 	Update an event
// @Description API for update an event
// @Tags 		EVENTS
// @Accept  	json
// @Produce  	json
// @Param		admin body user_service.UpdateEvent true "admin"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateEvent(c *gin.Context) {
	event := user_service.UpdateEvent{}

	if err := c.ShouldBindJSON(&event); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Events().Update(c.Request.Context(), &event)

	if err != nil {
		handleResponse(c, h.log, "error while updating an event", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Event updated successfully", http.StatusOK, resp)
}

// DeleteEvent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/event/{id} [DELETE]
// @Summary 	Delete an event
// @Description API for delete an event
// @Tags 		EVENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating event id", http.StatusBadRequest, err.Error())
		return
	}

	event := user_service.EventPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Events().Delete(c.Request.Context(), &event)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an event", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Event deleted successfully", http.StatusOK, resp)
}

// GetAllEvents godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/events [GET]
// @Summary 	Get all events
// @Description API for Get all events
// @Tags 		EVENTS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllEvents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)

	if err != nil {
		handleResponse(c, h.log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	events := user_service.GetListEventsRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Events().GetAll(c.Request.Context(), &events)

	if err != nil {
		handleResponse(c, h.log, "error while getting all events", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Events got successfully", http.StatusOK, resp)
}