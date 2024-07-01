package handler

import (
	"crmapi/genproto/user_service/administrators"
	"crmapi/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/admin [POST]
// @Summary 	Create a admin
// @Description API for creating admin
// @Tags 		ADMINISTRATORS
// @Accept  	json
// @Produce  	json
// @Param		admin body user_service.CreateAdminstrator true "admin"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateAdmin(c *gin.Context) {
	admin := user_service.CreateAdminstrator{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating an admin")
		return
	}

	if err := pkg.ValidateFullName(admin.FullName); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "full name is not valid")
		return
	}

	if err := pkg.ValidatePassword(admin.Password); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "password is not valid")
		return
	}

	if err := pkg.ValidatePhone(admin.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "phone number starts with +998...")
		return
	}

	resp, err := h.grpcClient.Admins().Create(c.Request.Context(), &admin)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating an admin")
		return
	}

	handleResponse(c, h.log, "Adminstrator is created", http.StatusCreated, resp)
}

// GetByIdAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/admin/{id} [GET]
// @Summary 	Get an admin
// @Description API for getting an admin
// @Tags 		ADMINISTRATORS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdAdmin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating admin", http.StatusBadRequest, err.Error())
		return
	}

	admin := user_service.AdminstratorPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Admins().GetById(c.Request.Context(), &admin)

	if err != nil {
		handleResponse(c, h.log, "error while getting admin", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Admin got successfully", http.StatusOK, resp)
}

// UpdateAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/admin [PUT]
// @Summary 	Update an admin
// @Description API for update an admin
// @Tags 		ADMINISTRATORS
// @Accept  	json
// @Produce  	json
// @Param		admin body user_service.UpdateAdminstrator true "admin"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateAdmin(c *gin.Context) {
	admin := user_service.UpdateAdminstrator{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	if err := pkg.ValidateFullName(admin.FullName); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "full name is not valid")
		return
	}

	if err := pkg.ValidatePassword(admin.Password); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "password is not valid")
		return
	}

	if err := pkg.ValidatePhone(admin.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "phone number starts with +998...")
		return
	}

	resp, err := h.grpcClient.Admins().Update(c.Request.Context(), &admin)

	if err != nil {
		handleResponse(c, h.log, "error while updating an admin", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Admin updated successfully", http.StatusOK, resp)
}

// DeleteAdmin godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/admin/{id} [DELETE]
// @Summary 	Delete an admin
// @Description API for delete an admin
// @Tags 		ADMINISTRATORS
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating admin id", http.StatusBadRequest, err.Error())
		return
	}

	admin := user_service.AdminstratorPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.Admins().Delete(c.Request.Context(), &admin)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an admin", http.StatusBadRequest, err.Error())
		return
	}
	
	handleResponse(c, h.log, "Admin deleted successfully", http.StatusOK, resp)
}

// GetAllAdmins godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/admins [GET]
// @Summary 	Get all admins
// @Description API for Get all admins
// @Tags 		ADMINISTRATORS
// @Accept  	json
// @Produce  	json
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllAdmins(c *gin.Context) {
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

	admins := user_service.GetListAdminstratorsRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Search: search,
	}

	resp, err := h.grpcClient.Admins().GetAll(c.Request.Context(), &admins)

	if err != nil {
		handleResponse(c, h.log, "error while getting all admins", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Admins got successfully", http.StatusOK, resp)
}