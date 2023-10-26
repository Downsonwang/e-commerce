/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-07-08 13:44:47
 * @LastEditTime: 2023-07-08 13:58:01
 */
package routers

import (
	"ecommerce-shop/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewer())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
