package handler

import (
	"crmapi/genproto/schedule_service/tasks"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTask godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/task [POST]
// @Summary 	Create a task
// @Description API for creating task
// @Tags 		TASKS
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.CreateTask true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateTask(c *gin.Context) {
	task := schedule_service.CreateTask{}

	if err := c.ShouldBindJSON(&task); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a task")
		return
	}

	_, err := h.grpcClient.Tasks().Create(c.Request.Context(), &task)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a task")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Task created successfully")
}

// GetByIdTask godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/task/{id} [GET]
// @Summary 	Get a task
// @Description API for getting a task
// @Tags 		TASKS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdTask(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating schedule id", http.StatusBadRequest, err.Error())
		return
	}

	task := schedule_service.TaskPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Tasks().GetById(c.Request.Context(), &task)

	if err != nil {
		handleResponse(c, h.log, "error while getting task", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Task got successfully", http.StatusOK, resp)
}

// UpdateTask godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/task [PUT]
// @Summary 	Update a task
// @Description API for update a task
// @Tags 		TASKS
// @Accept  	json
// @Produce  	json
// @Param		schedule body schedule_service.UpdateTask true "schedule"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateTask(c *gin.Context) {
	task := schedule_service.UpdateTask{}

	if err := c.ShouldBindJSON(&task); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Tasks().Update(c.Request.Context(), &task)

	if err != nil {
		handleResponse(c, h.log, "error while updating a task", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Task updated successfully", http.StatusOK, resp)
}

// DeleteTask godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/task/{id} [DELETE]
// @Summary 	Delete a task
// @Description API for delete a task
// @Tags 		TASKS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating task id", http.StatusBadRequest, err.Error())
		return
	}

	task := schedule_service.TaskPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.Tasks().Delete(c.Request.Context(), &task)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a task", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Task deleted successfully", http.StatusOK, resp)
}

// GetAllTasks godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/tasks [GET]
// @Summary 	Get all tasks
// @Description API for Get all tasks
// @Tags 		TASKS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllTasks(c *gin.Context) {

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

	tasks := schedule_service.GetListTaskRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.Tasks().GetAll(c.Request.Context(), &tasks)

	if err != nil {
		handleResponse(c, h.log, "error while getting all tasks", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Tasks got successfully", http.StatusOK, resp)
}
