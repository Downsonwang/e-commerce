/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-07-08 13:35:36
 * @LastEditTime: 2023-07-08 14:10:01
 */
package main

import (
	"ecommerce-shop/controllers"
	"ecommerce-shop/database"
	"ecommerce-shop/middleware"
	"ecommerce-shop/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRoutes(router)

	router.Use(middleware.Auth())

	router.GET("addtocart", app.AddToCart())
	router.GET("removeItemCart", app.RemoveItem())
	router.GET("/checkout", app.BuyFromCart())
	router.GET("/instantBuy", app.InstantBuy)

	log.Fatal(router.Run(":" + port))
}
