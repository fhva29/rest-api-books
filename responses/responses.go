package responses

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error   string            `json:"error"`
	Code    int               `json:"code"`
	Details map[string]string `json:"details,omitempty"`
}

func SendError(ctx *gin.Context, code int, message string, details ...map[string]string) {
	response := gin.H{
		"message":   message,
		"errorCode": code,
	}

	if len(details) > 0 && details[0] != nil {
		response["details"] = details
	}

	ctx.JSON(code, response)
}

func SendSuccess(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"data": data,
	})
}
