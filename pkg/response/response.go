package response

import (
	"bumimedika-backend/constants"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	Detail      interface{} `json:"detail,omitempty"`
	Filter      interface{} `json:"filter,omitempty"`
	Token       interface{} `json:"token,omitempty"`
	NextURL     string      `json:"next_url,omitempty"`
	Institution string      `json:"institution,omitempty"`
	Title       interface{} `json:"title,omitempty"`
	TotalPages  interface{} `json:"total_page,omitempty"`
	TotalData   interface{} `json:"total_data,omitempty"`
}

type ResponseSingleData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   interface{} `json:"token,omitempty"`
}
type ResponseError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type,omitempty"`
}

type IResponse interface {
	Success(c *gin.Context)
	Error(c *gin.Context, message ...string)
}

func (res Response) Success(c *gin.Context) {
	if res.Status == 0 {
		res.Status = http.StatusOK
	}
	res.Message = constants.OK
	c.JSON(http.StatusOK, res)
	c.Set("response_status", res.Status)
}

func (res ResponseSingleData) Success(c *gin.Context) {
	res.Status = http.StatusOK
	res.Message = constants.OK
	c.JSON(http.StatusOK, res)
	c.Set("response_status", res.Status)
}

// Error is response with json format
func (res Response) Error(c *gin.Context, message ...string) {
	s, m := generateError(message)
	if res.Status == 0 {
		res.Status = s
	}
	res.Message = m
	c.JSON(http.StatusOK, res)
	c.Set("response_status", res.Status)
}

func (res Response) ErrorTransact(c *gin.Context, message ...string) {
	res.Error(c, message...)
	c.Set("trx_error_status", res.Status)
}

func (res ResponseError) Error(c *gin.Context, message ...string) {
	s, m := generateError(message)
	if res.Status == 0 {
		res.Status = s
	}
	res.Message = m

	if res.Status >= 500 {
		c.JSON(http.StatusInternalServerError, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
	c.Set("response_status", res.Status)
}

// generateError is used to map error message
func generateError(messages []string) (int, string) {
	var status int
	if messages == nil {
		messages = append(messages, constants.RequestErrorMessage)
	}
	switch messages[0] {
	case constants.InternalErrorMessage:
		status = http.StatusInternalServerError
	case constants.NotAuthorizedMessage:
		status = http.StatusUnauthorized
	case constants.RequestErrorMessage, constants.EmailNotValidMessage:
		status = http.StatusBadRequest
	case constants.NoRecordFound:
		status = http.StatusNotFound
	default:
		status = http.StatusBadRequest

	}
	message := fmt.Sprintf(strings.Join(messages, "\n"))

	return status, message
}

//endregion functions
