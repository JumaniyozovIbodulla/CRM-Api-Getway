package handler

import (
	user_service "crmapi/genproto/user_service/super_admins"
	"crmapi/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSuperAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/super-admin [POST]
// @Summary 	Create an super admin
// @Description API for creating super admin
// @Tags 		SUPER-ADMINS
// @Accept  	json
// @Produce  	json
// @Param		super-admin body user_service.CreateSuperAdmin true "super-admin"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateSuperAdmin(c *gin.Context) {
	superAdmin := user_service.CreateSuperAdmin{}

	if err := c.ShouldBindJSON(&superAdmin); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating an order")
		return
	}

	if err := pkg.ValidateFullName(superAdmin.FullName); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "full name is not valid")
		return
	}

	if err := pkg.ValidatePassword(superAdmin.Password); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "password is not valid")
		return
	}

	if err := pkg.ValidatePhone(superAdmin.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "phone number starts with +998...")
		return
	}

	resp, err := h.grpcClient.SuperAdmins().Create(c.Request.Context(), &superAdmin)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating superAdmin")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetByIdSuperAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/super-admin/{id} [GET]
// @Summary 	Get an super admin
// @Description API for getting superadmin
// @Tags 		SUPER-ADMINS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdSuperAdmin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating super-admin", http.StatusBadRequest, err.Error())
		return
	}

	superAdmin := user_service.SuperAdminPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.SuperAdmins().GetById(c.Request.Context(), &superAdmin)

	if err != nil {
		handleResponse(c, h.log, "error while getting superAdmin", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "superAdmin got successfully", http.StatusOK, resp)
}

// UpdateSuperAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/super-admin [PUT]
// @Summary 	Update super admin
// @Description API for update super admin
// @Tags 		SUPER-ADMINS
// @Accept  	json
// @Produce  	json
// @Param		super-admin body user_service.UpdateSuperAdmin true "super-admin"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateSuperAdmin(c *gin.Context) {
	superAdmin := user_service.UpdateSuperAdmin{}

	if err := c.ShouldBindJSON(&superAdmin); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	if err := pkg.ValidateFullName(superAdmin.FullName); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "full name is not valid")
		return
	}

	if err := pkg.ValidatePassword(superAdmin.Password); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "password is not valid")
		return
	}

	if err := pkg.ValidatePhone(superAdmin.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "phone number starts with +998...")
		return
	}

	resp, err := h.grpcClient.SuperAdmins().Update(c.Request.Context(), &superAdmin)

	if err != nil {
		handleResponse(c, h.log, "error while updating superAdmin", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "superAdmin updated successfully", http.StatusOK, resp)
}

// DeleteOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/super-admin/{id} [DELETE]
// @Summary 	Delete super admin
// @Description API for delete a super admin
// @Tags 		SUPER-ADMINS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteSuperAdmin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating super-admin id", http.StatusBadRequest, err.Error())
		return
	}

	superAdminId := user_service.SuperAdminPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.SuperAdmins().Delete(c.Request.Context(), &superAdminId)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an superAdmin", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "superAdminId deleted successfully", http.StatusOK, resp)
}

// GetAllSuperAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/super-admins [GET]
// @Summary 	Get all super admins
// @Description API for Get all super admins
// @Tags 		SUPER-ADMINS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllSuperAdmin(c *gin.Context) {
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

	superAdmins := user_service.GetListSuperAdminRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.SuperAdmins().GetAll(c.Request.Context(), &superAdmins)

	if err != nil {
		handleResponse(c, h.log, "error while getting all superAdmins", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "superAdmins got successfully", http.StatusOK, resp)
}
