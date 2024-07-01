package handler

import (
	"crmapi/genproto/user_service/managers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateManager godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/manager [POST]
// @Summary 	Create a manager
// @Description API for creating manager
// @Tags 		MANAGERS
// @Accept  	json
// @Produce  	json
// @Param		manager body user_service.CreateManager true "manager"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateManager(c *gin.Context) {
	manager := user_service.CreateManager{}

	if err := c.ShouldBindJSON(&manager); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a manager")
		return
	}

	_, err := h.grpcClient.Managers().Create(c.Request.Context(), &manager)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a manager")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Manager created successfully")
}

// GetByIdManager godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/manager/{id} [GET]
// @Summary 	Get a group
// @Description API for getting a manager
// @Tags 		MANAGERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdManager(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating manager id", http.StatusBadRequest, err.Error())
		return
	}

	manager := user_service.ManagerPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Managers().GetById(c.Request.Context(), &manager)

	if err != nil {
		handleResponse(c, h.log, "error while getting manager", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Manager got successfully", http.StatusOK, resp)
}

// UpdateManager godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/manager [PUT]
// @Summary 	Update a manager
// @Description API for update a manager
// @Tags 		MANAGERS
// @Accept  	json
// @Produce  	json
// @Param		manager body user_service.UpdateManager true "manager"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateManager(c *gin.Context) {
	manager := user_service.UpdateManager{}

	if err := c.ShouldBindJSON(&manager); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Managers().Update(c.Request.Context(), &manager)

	if err != nil {
		handleResponse(c, h.log, "error while updating a manager", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Manager updated successfully", http.StatusOK, resp)
}

// DeleteManager godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/manager/{id} [DELETE]
// @Summary 	Delete a manager
// @Description API for delete a manager
// @Tags 		MANAGERS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteManager(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating manager id", http.StatusBadRequest, err.Error())
		return
	}

	manager := user_service.ManagerPrimaryKey{
		Id: id,
	}
	_, err := h.grpcClient.Managers().Delete(c.Request.Context(), &manager)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a manager", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Manager deleted successfully", http.StatusOK, "Manager deleted successfully")
}

// GetAllManagers godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/managers [GET]
// @Summary 	Get all managers
// @Description API for Get all managers
// @Tags 		MANAGERS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllManagers(c *gin.Context) {
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

	managers := user_service.GetListManagersRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Managers().GetAll(c.Request.Context(), &managers)

	if err != nil {
		handleResponse(c, h.log, "error while getting all managers", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Managers got successfully", http.StatusOK, resp)
}