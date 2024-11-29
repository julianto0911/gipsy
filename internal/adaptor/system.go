package adaptor

import (
	ucsystem "app/internal/usecase/system"
	"app/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSystemAdaptor(ucSystem ucsystem.ISystemUC) ISystemAdaptor {
	return &systemAdaptor{
		ucSystem: ucSystem,
	}
}

type ISystemAdaptor interface {
	VerifySAdminLogin(c *gin.Context)
}

type systemAdaptor struct {
	ucSystem ucsystem.ISystemUC
}

func (adp *systemAdaptor) VerifySAdminLogin(c *gin.Context) {
	input := ucsystem.InputSadmin{}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := adp.ucSystem.VerifySAdminLogin(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
