package handler

import (
	schedule_service "crmapi/genproto/schedule_service/attendances"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAttendance godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/attendance [POST]
// @Summary 	Create a attendance
// @Description API for creating attendance
// @Tags 		ATTENDANCES
// @Accept  	json
// @Produce  	json
// @Param		attendance body schedule_service.CreateAttendance true "attendance"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateAttendance(c *gin.Context) {
	attendance := schedule_service.CreateAttendance{}

	if err := c.ShouldBindJSON(&attendance); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a attendance")
		return
	}

	_, err := h.grpcClient.AttendancesService().Create(c.Request.Context(), &attendance)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a attendance")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Attendance created successfully")
}

// GetByIdAttendance godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/attendance/{id} [GET]
// @Summary 	Get a attendance
// @Description API for getting a attendance
// @Tags 		ATTENDANCES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdAttendance(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating attendance id", http.StatusBadRequest, err.Error())
		return
	}

	attendance := schedule_service.AttendancePrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.AttendancesService().GetById(c.Request.Context(), &attendance)

	if err != nil {
		handleResponse(c, h.log, "error while getting attendance", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Attendance got successfully", http.StatusOK, resp)
}

// UpdateAttendance godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/attendance [PUT]
// @Summary 	Update a attendance
// @Description API for update a attendance
// @Tags 		ATTENDANCES
// @Accept  	json
// @Produce  	json
// @Param		attendance body schedule_service.UpdateAttendance true "attendance"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateAttendance(c *gin.Context) {
	attendance := schedule_service.UpdateAttendance{}

	if err := c.ShouldBindJSON(&attendance); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.AttendancesService().Update(c.Request.Context(), &attendance)

	if err != nil {
		handleResponse(c, h.log, "error while updating a attendance", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "attendance updated successfully", http.StatusOK, resp)
}

// DeleteAttendance godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/attendance/{id} [DELETE]
// @Summary 	Delete a attendance
// @Description API for delete a attendance
// @Tags 		ATTENDANCES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteAttendance(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating attendance id", http.StatusBadRequest, err.Error())
		return
	}

	attendance := schedule_service.AttendancePrimaryKey{
		Id: id,
	}
	_, err := h.grpcClient.AttendancesService().Delete(c.Request.Context(), &attendance)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a attendance", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Attendance deleted successfully", http.StatusOK, "Attendance deleted successfully")
}

// GetAllAttendances godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/attendances [GET]
// @Summary 	Get all attendances
// @Description API for Get all attendances
// @Tags 		ATTENDANCES
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllAttendances(c *gin.Context) {

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

	attendances := schedule_service.GetListAttendanceRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.AttendancesService().GetAll(c.Request.Context(), &attendances)

	if err != nil {
		handleResponse(c, h.log, "error while getting all attendances", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Attendances got successfully", http.StatusOK, resp)
}
