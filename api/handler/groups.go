package handler

import (
	"crmapi/genproto/user_service/groups"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateGroup godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/group [POST]
// @Summary 	Create a group
// @Description API for creating group
// @Tags 		GROUPS
// @Accept  	json
// @Produce  	json
// @Param		group body user_service.CreateGroup true "group"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateGroup(c *gin.Context) {
	group := user_service.CreateGroup{}

	if err := c.ShouldBindJSON(&group); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a group")
		return
	}

	_, err := h.grpcClient.Groups().Create(c.Request.Context(), &group)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a group")
		return
	}
	handleGrpcErrWithDescription(c, h.log, err, "Group created successfully")
}

// GetByIdGroup godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/group/{id} [GET]
// @Summary 	Get a group
// @Description API for getting a group
// @Tags 		GROUPS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdGroup(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating group id", http.StatusBadRequest, err.Error())
		return
	}

	group := user_service.GroupPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Groups().GetById(c.Request.Context(), &group)

	if err != nil {
		handleResponse(c, h.log, "error while getting group", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Group got successfully", http.StatusOK, resp)
}

// UpdateGroup godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/group [PUT]
// @Summary 	Update a group
// @Description API for update a group
// @Tags 		GROUPS
// @Accept  	json
// @Produce  	json
// @Param		group body user_service.UpdateGroup true "group"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateGroup(c *gin.Context) {
	group := user_service.UpdateGroup{}

	if err := c.ShouldBindJSON(&group); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Groups().Update(c.Request.Context(), &group)

	if err != nil {
		handleResponse(c, h.log, "error while updating a group", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Group updated successfully", http.StatusOK, resp)
}

// DeleteGroup godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/group/{id} [DELETE]
// @Summary 	Delete a group
// @Description API for delete a group
// @Tags 		GROUPS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating group id", http.StatusBadRequest, err.Error())
		return
	}

	group := user_service.GroupPrimaryKey{
		Id: id,
	}
	resp, err := h.grpcClient.Groups().Delete(c.Request.Context(), &group)

	if err != nil {
		handleResponse(c, h.log, "error while deleting a group", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Group deleted successfully", http.StatusOK, resp)
}

// GetAllGroups godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/groups [GET]
// @Summary 	Get all groups
// @Description API for Get all groups
// @Tags 		GROUPS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllGroups(c *gin.Context) {
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

	groups := user_service.GetListGroupsRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Groups().GetAll(c.Request.Context(), &groups)

	if err != nil {
		handleResponse(c, h.log, "error while getting all groups", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Groups got successfully", http.StatusOK, resp)
}