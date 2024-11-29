package repository

import "gorm.io/gorm"

// entity for product
type EProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (c *EProduct) TableName() string {
	return "products"
}

// repository for product
func NewProductRepo(db *gorm.DB) RProduct {
	repo := &rProduct{
		db: db,
	}

	return repo
}

// interface for product repository
type RProduct interface {
	Create(name string) (*EProduct, error)
}

// internal repository for product
type rProduct struct {
	db *gorm.DB
}

// create product with name
func (r *rProduct) Create(name string) (*EProduct, error) {
	product := EProduct{}

	err := r.db.Create(&product).Error

	return &product, err
}
