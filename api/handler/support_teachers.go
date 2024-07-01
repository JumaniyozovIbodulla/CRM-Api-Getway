package handler

import (
	"crmapi/genproto/user_service/support_teachers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSupportTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/support-teacher [POST]
// @Summary 	Create a support teacher
// @Description API for creating support teacher
// @Tags 		SUPPORT-TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		support-teacher body user_service.CreateSupportTeacher true "support-teacher"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateSupportTeacher(c *gin.Context) {
	supportTeacher := user_service.CreateSupportTeacher{}

	if err := c.ShouldBindJSON(&supportTeacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a supportTeacher")
		return
	}

	_, err := h.grpcClient.SupportTeacher().Create(c.Request.Context(), &supportTeacher)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a support teacher")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Support teacher created successfully")
}

// GetByIdSupportTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/support-teacher/{id} [GET]
// @Summary 	Get a student
// @Description API for getting a support teacher
// @Tags 		SUPPORT-TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdSupportTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating support teacher id", http.StatusBadRequest, err.Error())
		return
	}

	supportTeacher := user_service.SupportTeacherPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.SupportTeacher().GetById(c.Request.Context(), &supportTeacher)

	if err != nil {
		handleResponse(c, h.log, "error while getting supportTeacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Support teacher got successfully", http.StatusOK, resp)
}

// UpdateSupportTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/support-teacher [PUT]
// @Summary 	Update a support teacher
// @Description API for update a support teacher
// @Tags 		SUPPORT-TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		support-teacher body user_service.UpdateSupportTeacher true "support-teacher"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateSupportTeacher(c *gin.Context) {
	supportTeacher := user_service.UpdateSupportTeacher{}

	if err := c.ShouldBindJSON(&supportTeacher); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.SupportTeacher().Update(c.Request.Context(), &supportTeacher)

	if err != nil {
		handleResponse(c, h.log, "error while updating a support teacher", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Support teacher updated successfully", http.StatusOK, resp)
}

// DeleteSupportTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/support-teacher/{id} [DELETE]
// @Summary 	Delete a support teacher
// @Description API for delete a support teacher
// @Tags 		SUPPORT-TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteSupportTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating support teacher id", http.StatusBadRequest, err.Error())
		return
	}

	supportTeacher := user_service.SupportTeacherPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.SupportTeacher().Delete(c.Request.Context(), &supportTeacher)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a support teacher", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Support teacher deleted successfully", http.StatusOK, resp)
}

// GetAllSupportTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/support-teacher [GET]
// @Summary 	Get all support teachers
// @Description API for Get all support teachers
// @Tags 		SUPPORT-TEACHERS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllSupportTeacher(c *gin.Context) {
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

	supportTeachers := user_service.GetListSupportTeachersRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.SupportTeacher().GetAll(c.Request.Context(), &supportTeachers)

	if err != nil {
		handleResponse(c, h.log, "error while getting all support teachers", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Support teachers got successfully", http.StatusOK, resp)
}