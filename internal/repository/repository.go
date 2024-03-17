package repository

import (
	"Simple-REST-ful-Fiber/internal"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
	Stock int
}

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) internal.Repository {
	return &productQuery{
		db: db,
	}
}

// InsertProduct implements internal.Repository.
func (p *productQuery) InsertProduct(newProduct internal.Product) error {
	inputProduct := new(Product)
	inputProduct.Name = newProduct.Name
	inputProduct.Price = newProduct.Price
	inputProduct.Stock = newProduct.Stock

	tx := p.db.Create(&inputProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteProduct implements internal.Repository.
func (p *productQuery) DeleteProduct(productId uint) error {
	tx := p.db.Delete(&Product{}, productId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// PutProduct implements internal.Repository.
func (p *productQuery) PutProduct(updateProduct internal.Product) error {
	inputProduct := new(Product)
	inputProduct.ID = updateProduct.ID
	inputProduct.Name = updateProduct.Name
	inputProduct.Price = updateProduct.Price
	inputProduct.Stock = updateProduct.Stock

	tx := p.db.Updates(&inputProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// ReadProduct implements internal.Repository.
func (p *productQuery) ReadProduct() ([]internal.Product, error) {
	var products []Product
	tx := p.db.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []internal.Product
	for _, product := range products {
		result = append(result, internal.Product{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}
	return result, nil
}