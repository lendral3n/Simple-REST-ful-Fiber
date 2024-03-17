package internal

type Product struct {
	ID uint
	Name string
	Price int
	Stock int
}

type Service interface {
	CreateProduct(newProduct Product) error
	GetProduct()([]Product, error)
	UpdateProduct(updateProduct Product)error
	DeleteProduct(productId uint)error
}

type Repository interface {
	InsertProduct(newProduct Product) error
	ReadProduct()([]Product, error)
	PutProduct(updateProduct Product)error
	DeleteProduct(productId uint)error
}