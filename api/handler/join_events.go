package handler

import (
	"crmapi/genproto/user_service/join_events"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateJoinEvent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/join-event [POST]
// @Summary 	Create a join event
// @Description API for creating join event
// @Tags 		JOIN-EVENTS
// @Accept  	json
// @Produce  	json
// @Param		join-event body user_service.CreateJoinEvent true "join-event"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateJoinEvent(c *gin.Context) {
	evenJoin := user_service.CreateJoinEvent{}

	if err := c.ShouldBindJSON(&evenJoin); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a evenJoin")
		return
	}

	_, err := h.grpcClient.JoinEvens().Create(c.Request.Context(), &evenJoin)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a evenJoin")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "EvenJoin created successfully")
}

// GetByIdEventJoin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/join-event/{id} [GET]
// @Summary 	Get a join event
// @Description API for getting a join event
// @Tags 		JOIN-EVENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdEventJoin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating join event id", http.StatusBadRequest, err.Error())
		return
	}

	joinevent := user_service.JoinEventPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.JoinEvens().GetById(c.Request.Context(), &joinevent)

	if err != nil {
		handleResponse(c, h.log, "error while getting joinevent", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Join event got successfully", http.StatusOK, resp)
}

// DeleteEventJoin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/join-event/{id} [DELETE]
// @Summary 	Delete a join event
// @Description API for delete a join event
// @Tags 		JOIN-EVENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteEventJoin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating join event id", http.StatusBadRequest, err.Error())
		return
	}

	joinEvent := user_service.JoinEventPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.JoinEvens().Delete(c.Request.Context(), &joinEvent)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a join event", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Join event deleted successfully", http.StatusOK, resp)
}