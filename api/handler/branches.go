package handler

import (
	"crmapi/genproto/user_service/branches"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBranch godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/branch [POST]
// @Summary 	Create a branch
// @Description API for creating branch
// @Tags 		BRANCHES
// @Accept  	json
// @Produce  	json
// @Param		super-admin body user_service.CreateBranch true "branch"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateBranch(c *gin.Context) {
	branch := user_service.CreateBranch{}

	if err := c.ShouldBindJSON(&branch); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a branch")
		return
	}

	resp, err := h.grpcClient.Branches().Create(c.Request.Context(), &branch)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating a branch")
		return
	}

	handleResponse(c, h.log, "Branch created successfully", http.StatusCreated, resp)
}

// GetByIdBranch godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/branch/{id} [GET]
// @Summary 	Get a branch
// @Description API for getting a branch
// @Tags 		BRANCHES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdBranch(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating branch", http.StatusBadRequest, err.Error())
		return
	}

	branch := user_service.BranchePrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Branches().GetById(c.Request.Context(), &branch)

	if err != nil {
		handleResponse(c, h.log, "error while getting branch", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Branch got successfully", http.StatusOK, resp)
}

// UpdateBranch godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/branch [PUT]
// @Summary 	Update a branch
// @Description API for update a branch
// @Tags 		BRANCHES
// @Accept  	json
// @Produce  	json
// @Param		branch body user_service.UpdateBranch true "branch"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateBranch(c *gin.Context) {
	branch := user_service.UpdateBranch{}

	if err := c.ShouldBindJSON(&branch); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.Branches().Update(c.Request.Context(), &branch)

	if err != nil {
		handleResponse(c, h.log, "error while updating a branch", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Branch updated successfully", http.StatusOK, resp)
}

// DeleteBranch godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/branch/{id} [DELETE]
// @Summary 	Delete a branch
// @Description API for delete a branch
// @Tags 		BRANCHES
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating branch id", http.StatusBadRequest, err.Error())
		return
	}

	branch := user_service.BranchePrimaryKey{
		Id: id,
	}

	_, err := h.grpcClient.Branches().Delete(c.Request.Context(), &branch)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an branch", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Branch deleted successfully", http.StatusOK, "Branch deleted successfully")
}

// GetAllBranches godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/branches [GET]
// @Summary 	Get all branches
// @Description API for Get all branches
// @Tags 		BRANCHES
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllBranches(c *gin.Context) {
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

	branches := user_service.GetListBranchesRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Branches().GetAll(c.Request.Context(), &branches)

	if err != nil {
		handleResponse(c, h.log, "error while getting all branches", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Branches got successfully", http.StatusOK, resp)
}
