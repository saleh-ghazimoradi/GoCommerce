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

}
