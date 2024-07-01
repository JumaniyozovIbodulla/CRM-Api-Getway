package handler

import (
	"crmapi/genproto/schedule_service/lessons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateLesson godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/lesson [POST]
// @Summary 	Create a lesson
// @Description API for creating lesson
// @Tags 		LESSONS
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.CreateLesson true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateLesson(c *gin.Context) {
	lesson := schedule_service.CreateLesson{}

	if err := c.ShouldBindJSON(&lesson); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a lesson")
		return
	}

	_, err := h.grpcClient.LessonsService().Create(c.Request.Context(), &lesson)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a lesson")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Lesson created successfully")
}

// GetByIdLesson godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/lesson/{id} [GET]
// @Summary 	Get a lesson
// @Description API for getting a lesson
// @Tags 		LESSONS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdLesson(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating lesson id", http.StatusBadRequest, err.Error())
		return
	}

	lesson := schedule_service.LessonPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.LessonsService().GetById(c.Request.Context(), &lesson)

	if err != nil {
		handleResponse(c, h.log, "error while getting lesson", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Lesson got successfully", http.StatusOK, resp)
}

// UpdateLesson godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/lesson [PUT]
// @Summary 	Update a lesson
// @Description API for update a lesson
// @Tags 		LESSONS
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.UpdateLesson true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateLesson(c *gin.Context) {
	lesson := schedule_service.UpdateLesson{}

	if err := c.ShouldBindJSON(&lesson); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.LessonsService().Update(c.Request.Context(), &lesson)

	if err != nil {
		handleResponse(c, h.log, "error while updating a lesson", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Lesson updated successfully", http.StatusOK, resp)
}

// GetAllLessons godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/lessons [GET]
// @Summary 	Get all lessons
// @Description API for Get all lessons
// @Tags 		LESSONS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllLessons(c *gin.Context) {

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

	lessons := schedule_service.GetListLessonRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.LessonsService().GetAll(c.Request.Context(), &lessons)

	if err != nil {
		handleResponse(c, h.log, "error while getting all lessons", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Lessons got successfully", http.StatusOK, resp)
}
