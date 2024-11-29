package entity

func NewProductEntity() IProduct {
	return &eProduct{}
}

type IProduct interface {
	TableName() string
}

type eProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (c *eProduct) TableName() string {
	return "products"
}
