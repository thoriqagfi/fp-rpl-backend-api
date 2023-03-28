package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type custController struct {
	custSvc services.CustSvc
	jwtSvc  services.JWTService
}

type CustController interface {
	LoginCust(ctx *gin.Context)
}

func NewCustController(cs services.CustSvc, jwt services.JWTService) CustController {
	return &custController{
		custSvc: cs,
		jwtSvc:  jwt,
	}
}

func (c *custController) LoginCust(ctx *gin.Context) {
	var custParam dto.UserLogin
	errParam := ctx.ShouldBindJSON(&custParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	verify, _ := c.custSvc.VerifyCust(ctx.Request.Context(), custParam.Email, custParam.Password)
	if !verify {
		response := utils.BuildErrorResponse("Gagal, email/password salah", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.FindCustByEmail(ctx.Request.Context(), custParam.Email)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtSvc.GenerateToken(tx.ID, tx.Role)
	custResponse := dto.UserResponse{
		Token: token,
		Role:  tx.Role,
	}

	response := utils.BuildResponse("Login", http.StatusOK, custResponse)
	ctx.JSON(http.StatusCreated, response)
}
