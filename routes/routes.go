package routes

import (
	"FP-RPL-ECommerce/controller"
	// "FP-RPL-ECommerce/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, UserController controller.UserController, CustContoller controller.CustController, SellerContoller controller.SellerController) {
	router := route.Group("")
	{
		router.POST("/register", UserController.Register)
		// router.GET("", middleware.Authenticate(""))

	}

	custRouter := route.Group("/customer")
	{
		custRouter.POST("/login", CustContoller.LoginCust)
	}

	sellerRouter := route.Group("/seller")
	{
		sellerRouter.POST("/login", SellerContoller.LoginCust)
	}
}
