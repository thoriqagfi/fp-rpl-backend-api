package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"

	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sellerController struct {
	sellerSvc services.SellerSvc
	jwtSvc    services.JWTService
}

type SellerController interface {
	LoginCust(ctx *gin.Context)
	ShowSellerByID(ctx *gin.Context)
	GetAllSeller(ctx *gin.Context)
	UpdateSeller(ctx *gin.Context)
	DeleteSeller(ctx *gin.Context)
}

func NewSellerController(cs services.SellerSvc, jwt services.JWTService) SellerController {
	return &sellerController{
		sellerSvc: cs,
		jwtSvc:    jwt,
	}
}

func (c *sellerController) LoginCust(ctx *gin.Context) {
	var sellerParam dto.UserLogin
	errParam := ctx.ShouldBindJSON(&sellerParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process login request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	verify, _ := c.sellerSvc.VerifySeller(ctx.Request.Context(), sellerParam.Email, sellerParam.Password)
	if !verify {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.FindSellerByEmail(ctx.Request.Context(), sellerParam.Email)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to Create New Account", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtSvc.GenerateToken(tx.ID, tx.Role)
	sellerResponse := dto.UserResponse{
		Token: token,
		Role:  tx.Role,
	}

	response := utils.BuildResponse("New Account Created", http.StatusCreated, sellerResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) ShowSellerByID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.FindSellerByID(ctx.Request.Context(), id)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Berhasil dapat", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) GetAllSeller(ctx *gin.Context) {
	seller, err := c.sellerSvc.GetAllSeller(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari semua customer", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("berhasil cari", http.StatusOK, seller)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) UpdateSeller(ctx *gin.Context) {
	var sellerParam dto.UserUpdate
	errParam := ctx.ShouldBindJSON(&sellerParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process token request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.UpdateSeller(ctx.Request.Context(), sellerParam, id)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Login", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) DeleteSeller(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process token request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.DeleteSeller(ctx.Request.Context(), id)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Login", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}
