package handler

import (
	"crmapi/genproto/user_service/teachers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/teacher [POST]
// @Summary 	Create a teacher
// @Description API for creating teacher
// @Tags 		TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		teacher body user_service.CreateTeacher true "teacher"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateTeacher(c *gin.Context) {
	teacher := user_service.CreateTeacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a teacher")
		return
	}

	_, err := h.grpcClient.Teachers().Create(c.Request.Context(), &teacher)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a teacher")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Teacher created successfully")
}

// GetByIdTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/teacher/{id} [GET]
// @Summary 	Get a teacher
// @Description API for getting a teacher
// @Tags 		TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating teacher id", http.StatusBadRequest, err.Error())
		return
	}

	teacher := user_service.TeacherPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Teachers().GetById(c.Request.Context(), &teacher)

	if err != nil {
		handleResponse(c, h.log, "error while getting teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Teacher got successfully", http.StatusOK, resp)
}

// UpdateTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/teacher [PUT]
// @Summary 	Update a teacher
// @Description API for update a teacher
// @Tags 		TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		teacher body user_service.UpdateTeacher true "teacher"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateTeacher(c *gin.Context) {
	teacher := user_service.UpdateTeacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Teachers().Update(c.Request.Context(), &teacher)

	if err != nil {
		handleResponse(c, h.log, "error while updating a teacher", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Teacher updated successfully", http.StatusOK, resp)
}

// DeleteTeacher godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/teacher/{id} [DELETE]
// @Summary 	Delete a teacher
// @Description API for delete a teacher
// @Tags 		TEACHERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating teacher id", http.StatusBadRequest, err.Error())
		return
	}

	teacher := user_service.TeacherPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.Teachers().Delete(c.Request.Context(), &teacher)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a teacher", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Teacher deleted successfully", http.StatusOK, resp)
}

// GetAllTeachers godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/teachers [GET]
// @Summary 	Get all teachers
// @Description API for Get all teachers
// @Tags 		TEACHERS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllTeachers(c *gin.Context) {
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

	teachers := user_service.GetListTeachersRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Teachers().GetAll(c.Request.Context(), &teachers)

	if err != nil {
		handleResponse(c, h.log, "error while getting all teachers", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Teachers got successfully", http.StatusOK, resp)
}