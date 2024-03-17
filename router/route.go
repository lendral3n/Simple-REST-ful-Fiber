package router

import (
	"Simple-REST-ful-Fiber/internal/handler"
	"Simple-REST-ful-Fiber/internal/repository"
	"Simple-REST-ful-Fiber/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, db *gorm.DB) {
	repo := repository.New(db)
	srv := service.New(repo)
	handler := handler.New(srv)

	api := app.Group("/product")

	api.Post("", handler.CreateProduct)
	api.Get("", handler.GetProduct)
	api.Delete("/:id", handler.DeleteProduct)
	api.Put("/:id", handler.UpdateProduct)
}