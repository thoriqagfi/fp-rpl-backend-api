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
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
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
