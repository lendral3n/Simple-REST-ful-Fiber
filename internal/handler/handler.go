package handler

import (
	"Simple-REST-ful-Fiber/internal"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	srv internal.Service
}

func New(s internal.Service) *ProductHandler {
	return &ProductHandler{
		srv: s,
	}
}

func (p *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	req := new(ProducctRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot Parse JSON"})
	}

	err := p.srv.CreateProduct(internal.Product{
		Name: req.Name,
		Price: req.Price,
		Stock: req.Stock,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (p *ProductHandler) GetProduct(c *fiber.Ctx) error {
	products, err := p.srv.GetProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var productResponss []ProductResponse
	for _, product := range products {
		productResponss = append(productResponss, ProductResponse{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}

	return c.JSON(productResponss)
}

func (p *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	req := new(ProducctRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON"})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse id"})
	}

	err = p.srv.UpdateProduct(internal.Product{
		ID: uint(id),
		Name: req.Name,
		Price: req.Price,
		Stock: req.Stock,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (p *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse id"})
	}

	err = p.srv.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
