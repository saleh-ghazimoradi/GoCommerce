package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/GoCommerce/config"
	"github.com/saleh-ghazimoradi/GoCommerce/utils"
)

func registerRoutes(app *fiber.App) {
	dbClient, err := utils.ConnectionMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	db := dbClient.Database(config.AppConfig.Databases.MongoDB.DatabaseName)

	fmt.Println(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GoCommerce Gateway")
	})

	v1 := app.Group("/v1")

	/*---------------------- User Routes -------------------*/
	v1.Post("/users/signup")
	v1.Post("/users/login")
	v1.Post("/admin/addproduct")
	v1.Get("/users/productview")
	v1.Get("/users/search")

	/*----------------------- Other Routes -------------------*/
	v1.Get("/addtocart")
	v1.Delete("/removeitem")
	v1.Get("/cartcheckout")
	v1.Get("instantbuy")

}
