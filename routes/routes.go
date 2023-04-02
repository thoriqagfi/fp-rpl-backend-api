package routes

import (
	"FP-RPL-ECommerce/controller"
	"FP-RPL-ECommerce/middleware"
	"FP-RPL-ECommerce/services"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, UserController controller.UserController, CustContoller controller.CustController, SellerContoller controller.SellerController, ProductController controller.ProductController, WishlistController controller.WishlistController, ReviewController controller.ReviewController, jwtSvc services.JWTService) {
	router := route.Group("")
	{
		router.POST("/register", UserController.Register)
		router.GET("", ProductController.GetAllProduct)
	}

	custRouter := route.Group("/customer")
	{
		custRouter.POST("/login", CustContoller.LoginCust)
		custRouter.PUT("/update", middleware.Authenticate(jwtSvc, "customer"), CustContoller.UpdateProfileCust)
		custRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "customer"), CustContoller.DeleteCust)
		custRouter.GET("/profile", middleware.Authenticate(jwtSvc, "customer"), CustContoller.ShowCustByID)
		custRouter.GET("/wishlist/:product_id", middleware.Authenticate(jwtSvc, "customer"), WishlistController.AddWishlist)
		custRouter.POST("/review/create", middleware.Authenticate(jwtSvc, "customer"), ReviewController.CreateReview)
	}

	sellerRouter := route.Group("/seller")
	{
		sellerRouter.POST("/login", SellerContoller.LoginCust)
		sellerRouter.PUT("/update", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.UpdateProfileSeller)
		sellerRouter.DELETE("/delete", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.DeleteSeller)
		sellerRouter.GET("/profile", middleware.Authenticate(jwtSvc, "seller"), SellerContoller.ShowSellerByID)
		sellerRouter.GET("/:first_name/:last_name", UserController.GetSellerByName)
	}

	adminRouter := route.Group("/admin")
	{
		// adminRouter.GET("/customer/:id", CustContoller.GetCustByID)
		adminRouter.GET("/customer/all", CustContoller.GetAllCust)

		// adminRouter.GET("/seller/:id", SellerContoller.GetSellerByID)
		adminRouter.GET("/seller/all", SellerContoller.GetAllSeller)
	}

	productRouter := route.Group("/product")
	{
		productRouter.POST("/create", middleware.Authenticate(jwtSvc, "seller"), ProductController.CreateProduct)
		productRouter.GET("/:id", ProductController.GetProductByID)
		productRouter.GET("/search/:product_name", ProductController.GetProductByName)
		productRouter.PUT("/:id/update", middleware.Authenticate(jwtSvc, "seller"), ProductController.UpdateProduct)
		productRouter.DELETE("/:id/delete", middleware.Authenticate(jwtSvc, "seller"), ProductController.DeleteProduct)
	}

}
