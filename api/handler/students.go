package handler

import (
	"crmapi/genproto/user_service/students"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStudent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/student [POST]
// @Summary 	Create a student
// @Description API for creating student
// @Tags 		STUDENTS
// @Accept  	json
// @Produce  	json
// @Param		student body user_service.CreateStudent true "student"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateStudent(c *gin.Context) {
	student := user_service.CreateStudent{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a student")
		return
	}

	_, err := h.grpcClient.Students().Create(c.Request.Context(), &student)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a student")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Student created successfully")
}

// GetByIdStudent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/student/{id} [GET]
// @Summary 	Get a student
// @Description API for getting a student
// @Tags 		STUDENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdStudent(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating student id", http.StatusBadRequest, err.Error())
		return
	}

	student := user_service.StudentPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Students().GetById(c.Request.Context(), &student)

	if err != nil {
		handleResponse(c, h.log, "error while getting student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Student got successfully", http.StatusOK, resp)
}

// UpdateStudent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/student [PUT]
// @Summary 	Update a student
// @Description API for update a student
// @Tags 		STUDENTS
// @Accept  	json
// @Produce  	json
// @Param		student body user_service.UpdateStudent true "student"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateStudent(c *gin.Context) {
	student := user_service.UpdateStudent{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Students().Update(c.Request.Context(), &student)

	if err != nil {
		handleResponse(c, h.log, "error while updating a student", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Student updated successfully", http.StatusOK, resp)
}

// DeleteStudent godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/student/{id} [DELETE]
// @Summary 	Delete a student
// @Description API for delete a student
// @Tags 		STUDENTS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating manager id", http.StatusBadRequest, err.Error())
		return
	}

	student := user_service.StudentPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.Students().Delete(c.Request.Context(), &student)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a student", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Student deleted successfully", http.StatusOK, resp)
}

// GetAllStudents godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/students [GET]
// @Summary 	Get all students
// @Description API for Get all students
// @Tags 		STUDENTS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllStudents(c *gin.Context) {
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

	students := user_service.GetListStudentsRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Students().GetAll(c.Request.Context(), &students)

	if err != nil {
		handleResponse(c, h.log, "error while getting all students", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Students got successfully", http.StatusOK, resp)
}