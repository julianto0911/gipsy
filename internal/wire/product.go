package wire

import (
	"app/internal/adaptor"
	"app/internal/repository"
	ucproduct "app/internal/usecase/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func wireProduct(router *gin.RouterGroup, db *gorm.DB) {
	rProduct := repository.NewProductRepo(db)
	ucProduct := ucproduct.NewProductUseCase(rProduct)
	adpProduct := adaptor.NewProductAdaptor(ucProduct)

	router.POST("/create", adpProduct.Create)
}
