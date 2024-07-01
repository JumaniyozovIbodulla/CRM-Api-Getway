package handler

import (
	"crmapi/genproto/schedule_service/schedules"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSchedule godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/schedule [POST]
// @Summary 	Create a schedule
// @Description API for creating schedule
// @Tags 		SCHEDULES
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.CreateSchedule true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateSchedule(c *gin.Context) {
	schedule := schedule_service.CreateSchedule{}

	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a schedule")
		return
	}

	_, err := h.grpcClient.Schedules().Create(c.Request.Context(), &schedule)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a schedule")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Schedule created successfully")
}

// GetByIdSchedule godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/schedule/{id} [GET]
// @Summary 	Get a schedule
// @Description API for getting a schedule
// @Tags 		SCHEDULES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdSchedule(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating schedule id", http.StatusBadRequest, err.Error())
		return
	}

	schedule := schedule_service.SchedulePrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Schedules().GetById(c.Request.Context(), &schedule)

	if err != nil {
		handleResponse(c, h.log, "error while getting schedule", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Schedule got successfully", http.StatusOK, resp)
}

// UpdateSchedule godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/schedule [PUT]
// @Summary 	Update a schedule
// @Description API for update a schedule
// @Tags 		SCHEDULES
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.UpdateSchedule true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateSchedule(c *gin.Context) {
	schedule := schedule_service.UpdateSchedule{}

	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Schedules().Update(c.Request.Context(), &schedule)

	if err != nil {
		handleResponse(c, h.log, "error while updating a schedule", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Schedule updated successfully", http.StatusOK, resp)
}

// DeleteSchedule godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/schedule/{id} [DELETE]
// @Summary 	Delete a schedule
// @Description API for delete a schedule
// @Tags 		SCHEDULES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating schedule id", http.StatusBadRequest, err.Error())
		return
	}

	schedule := schedule_service.SchedulePrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.Schedules().Delete(c.Request.Context(), &schedule)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a schedule", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Schedule deleted successfully", http.StatusOK, resp)
}

// GetAllSchedules godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/schedules [GET]
// @Summary 	Get all schedules
// @Description API for Get all schedules
// @Tags 		SCHEDULES
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllSchedules(c *gin.Context) {

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

	schedules := schedule_service.GetListScheduleRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.Schedules().GetAll(c.Request.Context(), &schedules)

	if err != nil {
		handleResponse(c, h.log, "error while getting all schedules", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Schedules got successfully", http.StatusOK, resp)
}
