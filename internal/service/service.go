package service

import (
	"Simple-REST-ful-Fiber/internal"
	"errors"
)

type productService struct {
	repo internal.Repository
}

func New(r internal.Repository) internal.Service {
	return &productService{
		repo: r,
	}
}

// CreateProduct implements internal.Service.
func (p *productService) CreateProduct(newProduct internal.Product) error {
	if newProduct.Name == "" {
		return errors.New("Product name is required")
	}
	if newProduct.Price <= 0 {
		return errors.New("Product price must be greater than 0")
	}
	if newProduct.Stock < 0 {
		return errors.New("Product stock cannot be negative")
	}

	err := p.repo.InsertProduct(newProduct)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements internal.Service.
func (p *productService) DeleteProduct(productId uint) error {
	if productId == 0 {
		return errors.New("product ID is required")
	}
	err := p.repo.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}

// GetProduct implements internal.Service.
func (p *productService) GetProduct() ([]internal.Product, error) {
	products, err := p.repo.ReadProduct()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("No products found")
	}
	return products, nil
}

// UpdateProduct implements internal.Service.
func (p *productService) UpdateProduct(updateProduct internal.Product) error {
	err := p.repo.PutProduct(updateProduct)
	if err != nil {
		return err
	}
	return nil
}