package repository

import "gorm.io/gorm"

// repository for product
func NewProductRepo(db *gorm.DB) IProduct {
	repo := &rProduct{
		db: db,
	}

	return repo
}

// interface for product repository
type IProduct interface {
	Create(name string) (*EProduct, error)
}

// entity for product
type EProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (c *EProduct) TableName() string {
	return "products"
}

// internal repository for product
type rProduct struct {
	db *gorm.DB
}

func (r *rProduct) Create(name string) (*EProduct, error) {
	product := EProduct{}

	err := r.db.Create(&product).Error

	return &product, err
}
