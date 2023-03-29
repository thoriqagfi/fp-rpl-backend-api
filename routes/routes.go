package routes

import (
	"FP-RPL-ECommerce/controller"
	"FP-RPL-ECommerce/middleware"
	"FP-RPL-ECommerce/services"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, UserController controller.UserController, CustContoller controller.CustController, SellerContoller controller.SellerController, jwtSvc services.JWTService) {
	router := route.Group("")
	{
		router.POST("/register", UserController.Register)
	}

	custRouter := route.Group("/customer")
	{
		custRouter.POST("/login", CustContoller.LoginCust)
		custRouter.PUT("/update", middleware.Authenticate(jwtSvc, "customer"), CustContoller.UpdateCust)
		custRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "customer"), CustContoller.DeleteCust)
	}

	sellerRouter := route.Group("/seller")
	{
		sellerRouter.POST("/login", SellerContoller.LoginCust)
		sellerRouter.PUT("/update", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.UpdateSeller)
		sellerRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.DeleteSeller)

	}

	adminRouter := route.Group("/admin")
	{
		adminRouter.GET("/customer/all", CustContoller.GetAllCust)

		adminRouter.GET("/seller/all", SellerContoller.GetAllSeller)
	}
}
