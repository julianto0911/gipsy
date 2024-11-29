package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/pkg/codes"
)

type Pagination struct {
	CurrentPage  int   `json:"current_page"`
	PerPage      int   `json:"per_page"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

type MetaData struct {
	Pagination  *Pagination `json:"pagination,omitempty"`
	ProcessTime float64     `json:"process_time,omitempty"`
	ServerID    string      `json:"server_id,omitempty"`
}

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	MetaData *MetaData   `json:"meta_data,omitempty"`
}

func Success(c *gin.Context, code int, data interface{}, meta *MetaData) {
	lang := getLang(c)
	c.JSON(http.StatusOK, Response{
		Code:     code,
		Message:  codes.GetMessage(lang, code),
		Data:     data,
		MetaData: meta,
	})
}

func Error(c *gin.Context, code int) {
	lang := getLang(c)
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: codes.GetMessage(lang, code),
	})
}

func getLang(c *gin.Context) string {
	// Check header first
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		// Fallback to query param
		lang = c.Query("lang")
	}
	if lang == "" {
		// Default to English
		lang = codes.LangEn
	}
	return lang
}
