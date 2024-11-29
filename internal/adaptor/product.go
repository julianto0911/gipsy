package adaptor

import (
	usecase "app/internal/usecase/product"
	"app/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewProductAdaptor(ucProduct usecase.ProductUC) ProductAdaptor {
	return &productAdaptor{
		ucProduct: ucProduct,
	}
}

type ProductAdaptor interface {
	Create(c *gin.Context)
}

type productAdaptor struct {
	ucProduct usecase.ProductUC
}

func (adp *productAdaptor) Create(c *gin.Context) {
	input := usecase.InputProduct{}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := adp.ucProduct.Create(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
