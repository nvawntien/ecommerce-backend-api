package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`           // Status Code
	Message string      `json:"message"`        // Message
	Data    interface{} `json:"data,omitempty"` // Data Payload
}

// success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}

// error response
func Error(c *gin.Context, httpStatus int, internalCode int, cusMsg string) {
	message := msg[internalCode]

	if cusMsg != "" {
		message = cusMsg
	}

	c.JSON(httpStatus, ResponseData{
		Code:    internalCode,
		Message: message,
		Data:    nil,
	})
}
